package middleware

import (
	"context"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter *rate.Limiter
	mu      sync.Mutex
}

func NewRateLimiter(requestsPerSecond float64, burst int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(requestsPerSecond), burst),
	}
}

func WithRateLimit(requestsPerSecond float64, burst int) Middleware {
	limiter := NewRateLimiter(requestsPerSecond, burst)

	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			if err := limiter.Wait(ctx); err != nil {
				return nil, err
			}
			return next.RoundTrip(ctx, req)
		})
	}
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.limiter.Wait(ctx)
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.limiter.Allow()
}

type PerHostRateLimiter struct {
	limiters map[string]*RateLimiter
	mu       sync.RWMutex
	rps      float64
	burst    int
}

func NewPerHostRateLimiter(requestsPerSecond float64, burst int) *PerHostRateLimiter {
	return &PerHostRateLimiter{
		limiters: make(map[string]*RateLimiter),
		rps:      requestsPerSecond,
		burst:    burst,
	}
}

func (p *PerHostRateLimiter) getLimiter(host string) *RateLimiter {
	p.mu.RLock()
	limiter, exists := p.limiters[host]
	p.mu.RUnlock()

	if exists {
		return limiter
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	limiter, exists = p.limiters[host]
	if exists {
		return limiter
	}

	limiter = NewRateLimiter(p.rps, p.burst)
	p.limiters[host] = limiter
	return limiter
}

func WithPerHostRateLimit(requestsPerSecond float64, burst int) Middleware {
	perHostLimiter := NewPerHostRateLimiter(requestsPerSecond, burst)

	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			limiter := perHostLimiter.getLimiter(req.URL.Host)
			if err := limiter.Wait(ctx); err != nil {
				return nil, err
			}
			return next.RoundTrip(ctx, req)
		})
	}
}
