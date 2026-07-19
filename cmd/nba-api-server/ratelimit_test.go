package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimiter_SamePortStripped(t *testing.T) {
	rl := NewRateLimiter(1, 1) // 1 req/s, burst of 1

	handler := rl.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Same IP, two different ephemeral source ports - this is exactly what
	// happens when a client opens a fresh connection for each request.
	// Before stripping the port, each request got its own limiter and the
	// per-IP limit was never enforced.
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	req1.RemoteAddr = "203.0.113.5:51000"
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.RemoteAddr = "203.0.113.5:51001"

	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)
	if w1.Code != http.StatusOK {
		t.Fatalf("first request: expected 200, got %d", w1.Code)
	}

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)
	if w2.Code != http.StatusTooManyRequests {
		t.Errorf("second request from same IP on a different port: expected 429 (shared bucket exhausted), got %d", w2.Code)
	}
}

func TestRateLimiter_DifferentIPsIndependent(t *testing.T) {
	rl := NewRateLimiter(1, 1)

	handler := rl.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	req1.RemoteAddr = "203.0.113.5:51000"
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.RemoteAddr = "203.0.113.6:51000"

	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)
	if w1.Code != http.StatusOK {
		t.Fatalf("first IP: expected 200, got %d", w1.Code)
	}

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Errorf("second, distinct IP: expected 200 (independent bucket), got %d", w2.Code)
	}
}

func TestClientIP(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		want       string
	}{
		{"host and port", "203.0.113.5:51000", "203.0.113.5"},
		{"ipv6 with port", "[2001:db8::1]:51000", "2001:db8::1"},
		{"no port falls back to raw value", "203.0.113.5", "203.0.113.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.RemoteAddr = tt.remoteAddr

			if got := clientIP(req); got != tt.want {
				t.Errorf("clientIP(%q) = %q, want %q", tt.remoteAddr, got, tt.want)
			}
		})
	}
}
