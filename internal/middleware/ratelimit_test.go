package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestWithPerHostRateLimit_HostsAreIndependent(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})

	wrapped := WithPerHostRateLimit(1, 1)(rt) // 1 req/s, burst 1

	reqA1 := httptest.NewRequest(http.MethodGet, "http://host-a.example.com", nil)
	reqA2 := httptest.NewRequest(http.MethodGet, "http://host-a.example.com", nil)
	reqB1 := httptest.NewRequest(http.MethodGet, "http://host-b.example.com", nil)

	ctx := context.Background()

	start := time.Now()
	if _, err := wrapped.RoundTrip(ctx, reqA1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if elapsed := time.Since(start); elapsed > 50*time.Millisecond {
		t.Errorf("first request to host A should not wait, took %v", elapsed)
	}

	// A different host's first request must also be immediate - this is
	// what proves hosts don't share a limiter.
	start = time.Now()
	if _, err := wrapped.RoundTrip(ctx, reqB1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if elapsed := time.Since(start); elapsed > 50*time.Millisecond {
		t.Errorf("first request to a different host should not wait for host A's limiter, took %v", elapsed)
	}

	// A second request to host A has to wait for its own bucket to refill.
	start = time.Now()
	if _, err := wrapped.RoundTrip(ctx, reqA2); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if elapsed := time.Since(start); elapsed < 500*time.Millisecond {
		t.Errorf("second request to host A should have waited for its own rate limit, took only %v", elapsed)
	}
}

// TestRateLimiter_ConcurrentWaitIsSafe is a regression test for the
// removed sync.Mutex wrapper: rate.Limiter is already safe for concurrent
// use, and this proves concurrent callers no longer need (or get) any
// extra serialization around Wait. Run with -race to catch a real
// concurrency bug, not just a logical one.
func TestRateLimiter_ConcurrentWaitIsSafe(t *testing.T) {
	rl := NewRateLimiter(1000, 1000) // generous limits: this test is about concurrency safety, not timing

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := rl.Wait(context.Background()); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestRateLimiter_AllowRespectsBurst(t *testing.T) {
	rl := NewRateLimiter(1, 2) // 1 req/s, burst of 2

	if !rl.Allow() {
		t.Error("expected first call to be allowed (within burst)")
	}
	if !rl.Allow() {
		t.Error("expected second call to be allowed (within burst)")
	}
	if rl.Allow() {
		t.Error("expected third call to be denied (burst exhausted)")
	}
}
