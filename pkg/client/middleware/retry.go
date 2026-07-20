package middleware

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// RetryConfig controls retry behavior. MaxRetries values below zero are
// treated as zero, so a request is always attempted at least once.
type RetryConfig struct {
	MaxRetries      int
	InitialBackoff  time.Duration
	MaxBackoff      time.Duration
	BackoffMultiple float64
	RetryableStatus []int
}

func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:      3,
		InitialBackoff:  1 * time.Second,
		MaxBackoff:      30 * time.Second,
		BackoffMultiple: 2.0,
		RetryableStatus: []int{
			http.StatusTooManyRequests,
			http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		},
	}
}

func WithRetry(config RetryConfig) Middleware {
	if config.MaxRetries < 0 {
		config.MaxRetries = 0
	}

	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			var resp *http.Response
			var err error
			// retryAfter carries a server-specified delay from the previous
			// attempt's response into the next attempt's wait, taking
			// priority over the calculated exponential backoff.
			var retryAfter time.Duration

			for attempt := 0; attempt <= config.MaxRetries; attempt++ {
				if attempt > 0 {
					backoff := retryAfter
					if backoff <= 0 {
						backoff = calculateBackoff(attempt, config)
					}
					select {
					case <-time.After(backoff):
					case <-ctx.Done():
						return nil, ctx.Err()
					}
				}
				retryAfter = 0

				resp, err = next.RoundTrip(ctx, req)

				if err != nil {
					// A canceled or expired context means every further
					// attempt will fail the same way immediately - retrying
					// just burns through MaxRetries for nothing.
					if isPermanentTransportError(err) {
						return nil, err
					}
					if attempt < config.MaxRetries {
						continue
					}
					return nil, err
				}

				if !isRetryableStatus(resp.StatusCode, config.RetryableStatus) {
					return resp, nil
				}

				retryAfter = parseRetryAfter(resp.Header.Get("Retry-After"), config.MaxBackoff)

				if attempt < config.MaxRetries {
					//nolint:errcheck
					resp.Body.Close()
					continue
				}

				return resp, nil
			}

			return resp, err
		})
	}
}

func isPermanentTransportError(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

// parseRetryAfter parses an HTTP Retry-After header value, which per RFC
// 9110 §10.2.3 is either a non-negative integer number of seconds or an
// HTTP-date. It returns 0 (meaning "no preference, use the calculated
// backoff") when header is empty, unparseable, or in the past, and caps
// the result at maxBackoff so a misbehaving upstream can't stall a caller
// indefinitely.
func parseRetryAfter(header string, maxBackoff time.Duration) time.Duration {
	if header == "" {
		return 0
	}

	if seconds, err := strconv.Atoi(header); err == nil {
		if seconds <= 0 {
			return 0
		}
		delay := time.Duration(seconds) * time.Second
		if delay > maxBackoff {
			return maxBackoff
		}
		return delay
	}

	if when, err := http.ParseTime(header); err == nil {
		delay := time.Until(when)
		if delay <= 0 {
			return 0
		}
		if delay > maxBackoff {
			return maxBackoff
		}
		return delay
	}

	return 0
}

func calculateBackoff(attempt int, config RetryConfig) time.Duration {
	backoff := float64(config.InitialBackoff) * math.Pow(config.BackoffMultiple, float64(attempt-1))
	backoff = backoff + (backoff * 0.1 * (rand.Float64()*2 - 1))

	if backoff > float64(config.MaxBackoff) {
		backoff = float64(config.MaxBackoff)
	}

	return time.Duration(backoff)
}

func isRetryableStatus(statusCode int, retryable []int) bool {
	for _, code := range retryable {
		if statusCode == code {
			return true
		}
	}
	return false
}
