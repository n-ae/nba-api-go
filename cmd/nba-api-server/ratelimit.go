package main

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*rate.Limiter
	rate     rate.Limit
	burst    int
}

func NewRateLimiter(requestsPerSecond int, burst int) *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     rate.Limit(requestsPerSecond),
		burst:    burst,
	}
}

func (rl *RateLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.limiters[ip]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.limiters[ip] = limiter
	}

	return limiter
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		limiter := rl.getLimiter(ip)

		if !limiter.Allow() {
			writeError(w, http.StatusTooManyRequests, "rate_limit_exceeded", "Too many requests, please slow down")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// clientIP extracts the connecting IP from r.RemoteAddr, stripping the
// ephemeral source port so that a client reconnecting on a new port keys
// into the same limiter bucket. Proxy headers (X-Forwarded-For,
// X-Real-IP) are deliberately not trusted here: honoring them from an
// arbitrary client would let any request forge its own rate-limit key.
// If this server is deployed behind a trusted reverse proxy, that proxy
// must be the one to enforce (or safely relay) per-client identity.
func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// RemoteAddr had no port (e.g. a unix socket or a test request
		// built without one) - fall back to the raw value.
		return r.RemoteAddr
	}
	return host
}

func (rl *RateLimiter) CleanupOldLimiters(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			rl.mu.Lock()
			for ip, limiter := range rl.limiters {
				if limiter.Tokens() == float64(rl.burst) {
					delete(rl.limiters, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()
}
