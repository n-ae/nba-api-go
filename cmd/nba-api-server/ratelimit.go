package main

import (
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
		ip := r.RemoteAddr
		limiter := rl.getLimiter(ip)

		if !limiter.Allow() {
			writeError(w, http.StatusTooManyRequests, "rate_limit_exceeded", "Too many requests, please slow down")
			return
		}

		next.ServeHTTP(w, r)
	})
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
