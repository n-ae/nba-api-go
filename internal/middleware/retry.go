package middleware

import (
	"context"
	"math"
	"math/rand"
	"net/http"
	"time"
)

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
	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			var resp *http.Response
			var err error

			for attempt := 0; attempt <= config.MaxRetries; attempt++ {
				if attempt > 0 {
					backoff := calculateBackoff(attempt, config)
					select {
					case <-time.After(backoff):
					case <-ctx.Done():
						return nil, ctx.Err()
					}
				}

				resp, err = next.RoundTrip(ctx, req)

				if err != nil {
					if attempt < config.MaxRetries {
						continue
					}
					return nil, err
				}

				if !isRetryableStatus(resp.StatusCode, config.RetryableStatus) {
					return resp, nil
				}

				if attempt < config.MaxRetries {
					resp.Body.Close()
					continue
				}

				return resp, nil
			}

			return resp, err
		})
	}
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
