package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func fakeResponse(status int, header http.Header) *http.Response {
	if header == nil {
		header = make(http.Header)
	}
	return &http.Response{
		StatusCode: status,
		Header:     header,
		Body:       http.NoBody,
	}
}

func TestWithRetry_RetriesRetryableStatusUpToMaxRetries(t *testing.T) {
	var calls int
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		calls++
		return fakeResponse(http.StatusServiceUnavailable, nil), nil
	})

	config := RetryConfig{
		MaxRetries:      2,
		InitialBackoff:  1 * time.Millisecond,
		MaxBackoff:      5 * time.Millisecond,
		BackoffMultiple: 2.0,
		RetryableStatus: []int{http.StatusServiceUnavailable},
	}

	wrapped := WithRetry(config)(rt)
	resp, err := wrapped.RoundTrip(context.Background(), httptest.NewRequest(http.MethodGet, "http://example.com", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected final status 503, got %d", resp.StatusCode)
	}
	if want := config.MaxRetries + 1; calls != want {
		t.Errorf("expected %d calls (1 initial + %d retries), got %d", want, config.MaxRetries, calls)
	}
}

func TestWithRetry_NonRetryableStatusReturnsImmediately(t *testing.T) {
	var calls int
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		calls++
		return fakeResponse(http.StatusNotFound, nil), nil
	})

	wrapped := WithRetry(DefaultRetryConfig())(rt)
	resp, err := wrapped.RoundTrip(context.Background(), httptest.NewRequest(http.MethodGet, "http://example.com", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
	if calls != 1 {
		t.Errorf("expected exactly 1 call for a non-retryable status, got %d", calls)
	}
}

func TestWithRetry_NegativeMaxRetriesStillMakesInitialRequest(t *testing.T) {
	var calls int
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		calls++
		return fakeResponse(http.StatusOK, nil), nil
	})

	config := DefaultRetryConfig()
	config.MaxRetries = -1

	wrapped := WithRetry(config)(rt)
	resp, err := wrapped.RoundTrip(context.Background(), httptest.NewRequest(http.MethodGet, "http://example.com", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("expected a successful response, got %#v", resp)
	}
	if calls != 1 {
		t.Errorf("expected one initial request with negative MaxRetries, got %d", calls)
	}
}

func TestWithRetry_PermanentTransportErrorExitsImmediately(t *testing.T) {
	for _, permErr := range []error{context.Canceled, context.DeadlineExceeded} {
		t.Run(permErr.Error(), func(t *testing.T) {
			var calls int
			rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
				calls++
				return nil, permErr
			})

			wrapped := WithRetry(DefaultRetryConfig())(rt)
			_, err := wrapped.RoundTrip(context.Background(), httptest.NewRequest(http.MethodGet, "http://example.com", nil))
			if !errors.Is(err, permErr) {
				t.Fatalf("expected %v, got %v", permErr, err)
			}
			if calls != 1 {
				t.Errorf("expected exactly 1 call (no retries on a permanent transport error), got %d", calls)
			}
		})
	}
}

func TestWithRetry_TransientTransportErrorIsRetried(t *testing.T) {
	var calls int
	sentinelErr := errors.New("connection reset")
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		calls++
		return nil, sentinelErr
	})

	config := DefaultRetryConfig()
	config.MaxRetries = 2
	config.InitialBackoff = 1 * time.Millisecond
	config.MaxBackoff = 5 * time.Millisecond

	wrapped := WithRetry(config)(rt)
	_, err := wrapped.RoundTrip(context.Background(), httptest.NewRequest(http.MethodGet, "http://example.com", nil))
	if !errors.Is(err, sentinelErr) {
		t.Fatalf("expected sentinel error, got %v", err)
	}
	if want := config.MaxRetries + 1; calls != want {
		t.Errorf("expected %d calls, got %d", want, calls)
	}
}

func TestWithRetry_ContextCancelDuringBackoffExitsPromptly(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusServiceUnavailable, nil), nil
	})

	config := DefaultRetryConfig() // InitialBackoff is 1s - would be slow if not interrupted
	config.MaxRetries = 5

	wrapped := WithRetry(config)(rt)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := wrapped.RoundTrip(ctx, httptest.NewRequest(http.MethodGet, "http://example.com", nil))
	elapsed := time.Since(start)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded, got %v", err)
	}
	if elapsed >= 1*time.Second {
		t.Errorf("expected to exit during the backoff wait well under 1s, took %v", elapsed)
	}
}

func TestCalculateBackoff_NeverExceedsMaxBackoff(t *testing.T) {
	config := RetryConfig{
		InitialBackoff:  1 * time.Second,
		MaxBackoff:      3 * time.Second,
		BackoffMultiple: 2.0,
	}

	for attempt := 1; attempt <= 10; attempt++ {
		backoff := calculateBackoff(attempt, config)
		if backoff > config.MaxBackoff {
			t.Errorf("attempt %d: backoff %v exceeds MaxBackoff %v", attempt, backoff, config.MaxBackoff)
		}
		if backoff < 0 {
			t.Errorf("attempt %d: backoff %v is negative", attempt, backoff)
		}
	}
}

func TestWithRetry_HonorsRetryAfterCappedAtMaxBackoff(t *testing.T) {
	var timestamps []time.Time
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		timestamps = append(timestamps, time.Now())
		if len(timestamps) == 1 {
			header := http.Header{}
			header.Set("Retry-After", "5") // far longer than MaxBackoff below
			return fakeResponse(http.StatusServiceUnavailable, header), nil
		}
		return fakeResponse(http.StatusOK, nil), nil
	})

	config := RetryConfig{
		MaxRetries:      1,
		InitialBackoff:  1 * time.Millisecond, // would be used if Retry-After weren't honored
		MaxBackoff:      20 * time.Millisecond,
		BackoffMultiple: 2.0,
		RetryableStatus: []int{http.StatusServiceUnavailable},
	}

	wrapped := WithRetry(config)(rt)
	resp, err := wrapped.RoundTrip(context.Background(), httptest.NewRequest(http.MethodGet, "http://example.com", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected eventual 200, got %d", resp.StatusCode)
	}
	if len(timestamps) != 2 {
		t.Fatalf("expected 2 calls, got %d", len(timestamps))
	}

	gap := timestamps[1].Sub(timestamps[0])
	// Retry-After asked for 5s but MaxBackoff caps it at 20ms. If the
	// server-specified delay weren't honored at all, we'd see something
	// close to InitialBackoff (1ms) instead of the ~20ms cap.
	if gap < 15*time.Millisecond {
		t.Errorf("expected the capped Retry-After (~20ms) to be honored, gap was only %v", gap)
	}
	if gap > 500*time.Millisecond {
		t.Errorf("expected the delay capped near MaxBackoff (20ms), gap was %v (looks like the full 5s Retry-After was used)", gap)
	}
}

func TestParseRetryAfter(t *testing.T) {
	maxBackoff := 10 * time.Second

	t.Run("seconds form", func(t *testing.T) {
		if got := parseRetryAfter("5", maxBackoff); got != 5*time.Second {
			t.Errorf("expected 5s, got %v", got)
		}
	})

	t.Run("http-date form", func(t *testing.T) {
		future := time.Now().Add(2 * time.Second).UTC().Format(http.TimeFormat)
		got := parseRetryAfter(future, maxBackoff)
		if got <= 0 || got > 3*time.Second {
			t.Errorf("expected ~2s delay, got %v", got)
		}
	})

	t.Run("caps at maxBackoff", func(t *testing.T) {
		if got := parseRetryAfter("100", 2*time.Second); got != 2*time.Second {
			t.Errorf("expected cap at 2s, got %v", got)
		}
	})

	t.Run("empty or invalid returns zero", func(t *testing.T) {
		for _, v := range []string{"", "not-a-number-or-date", "-5"} {
			if got := parseRetryAfter(v, maxBackoff); got != 0 {
				t.Errorf("parseRetryAfter(%q) = %v, want 0", v, got)
			}
		}
	})

	t.Run("past date returns zero", func(t *testing.T) {
		past := time.Now().Add(-1 * time.Hour).UTC().Format(http.TimeFormat)
		if got := parseRetryAfter(past, maxBackoff); got != 0 {
			t.Errorf("expected 0 for a past date, got %v", got)
		}
	})
}
