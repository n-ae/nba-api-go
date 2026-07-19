package stats

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/n-ae/nba-api-go/internal/middleware"
)

// noRetry is a passthrough middleware used to isolate a single request
// attempt from the default retry chain when a test cares about the
// outcome of exactly one RoundTrip (e.g. a client-side timeout).
func noRetry() middleware.Middleware {
	return func(next middleware.RoundTripper) middleware.RoundTripper {
		return next
	}
}

func TestNewClient_ForwardsHeaders(t *testing.T) {
	var gotHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotHeader = r.Header.Get("X-Test-Header")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := NewClient(Config{
		BaseURL: srv.URL,
		Headers: http.Header{"X-Test-Header": {"hello"}},
	})

	if _, err := c.client.Get(context.Background(), "test", url.Values{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotHeader != "hello" {
		t.Errorf("expected X-Test-Header=hello, got %q", gotHeader)
	}
}

// TestNewClient_ForwardsTimeout proves Config.Timeout reaches the
// underlying http.Client's Timeout (a client-wide deadline), not merely
// that request-scoped context deadlines still work regardless of Config.
// The server delay (1.5s) exceeds the configured Timeout (1s) but the
// context passed to Get has no deadline of its own, so the only thing
// that can abort the request is the forwarded Timeout.
func TestNewClient_ForwardsTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1500 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := NewClient(Config{
		BaseURL:     srv.URL,
		Timeout:     time.Second,
		Middlewares: []middleware.Middleware{noRetry()},
	})

	start := time.Now()
	_, err := c.client.Get(context.Background(), "slow", url.Values{})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected client-side timeout error, got nil")
	}
	if elapsed >= 1500*time.Millisecond {
		t.Errorf("expected request to abort around the 1s Timeout, took %v (server needed 1.5s)", elapsed)
	}
}

func TestNewClient_ForwardsMaxResponseBytes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(make([]byte, 100))
	}))
	defer srv.Close()

	c := NewClient(Config{
		BaseURL:          srv.URL,
		MaxResponseBytes: 10,
	})

	if _, err := c.client.Get(context.Background(), "test", url.Values{}); err == nil {
		t.Fatal("expected an error for a 100-byte body against a forwarded 10-byte MaxResponseBytes, got nil")
	}
}
