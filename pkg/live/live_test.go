package live

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/n-ae/nba-api-go/v2/pkg/client"
	"github.com/n-ae/nba-api-go/v2/pkg/client/middleware"
)

// withHeader is a small test middleware that sets a fixed header on every
// request, used to prove AdditionalMiddlewares actually runs.
func withHeader(key, value string) client.Middleware {
	return func(next client.RoundTripper) client.RoundTripper {
		return client.RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			req.Header.Set(key, value)
			return next.RoundTrip(ctx, req)
		})
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
// underlying http.Client's Timeout (a client-wide deadline). The server
// delay (1.5s) exceeds the configured Timeout (1s) but the context passed
// to Get has no deadline of its own, so only the forwarded Timeout can
// abort the request. Retry middleware is disabled to isolate a single
// RoundTrip attempt (the default chain has no retry, but rate limiting
// would otherwise still apply harmlessly; this keeps the test focused).
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

// noRetry is a passthrough middleware used to isolate a single request
// attempt from the default middleware chain when a test cares about the
// outcome of exactly one RoundTrip (e.g. a client-side timeout).
func noRetry() middleware.Middleware {
	return func(next middleware.RoundTripper) middleware.RoundTripper {
		return next
	}
}

// TestNewClient_DefaultHeadersReachTheWire guards against regressing the
// User-Agent shadowing bug: the core client used to inject
// DefaultUserAgent ("nba-api-go/2") into every request's headers before
// the middleware chain ran, and WithUserAgent only sets the header when
// absent - so the browser-style default this facade installs never won.
// This asserts the actual bytes received by the server, not just that
// some middleware ran.
func TestNewClient_DefaultHeadersReachTheWire(t *testing.T) {
	var gotUserAgent string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUserAgent = r.Header.Get("User-Agent")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := NewClient(Config{BaseURL: srv.URL})

	if _, err := c.client.Get(context.Background(), "test", url.Values{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const wantUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
	if gotUserAgent != wantUserAgent {
		t.Errorf("User-Agent = %q, want %q", gotUserAgent, wantUserAgent)
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

// TestNewClient_AdditionalMiddlewaresAppendToDefaults proves
// AdditionalMiddlewares runs alongside DefaultMiddlewares (Config.Middlewares
// left empty) rather than requiring the caller to manually
// append(DefaultMiddlewares(), ...) to get both.
func TestNewClient_AdditionalMiddlewaresAppendToDefaults(t *testing.T) {
	var gotUserAgent, gotCustomHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUserAgent = r.Header.Get("User-Agent")
		gotCustomHeader = r.Header.Get("X-Custom")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := NewClient(Config{
		BaseURL:               srv.URL,
		AdditionalMiddlewares: []client.Middleware{withHeader("X-Custom", "added")},
	})

	if _, err := c.client.Get(context.Background(), "test", url.Values{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const wantUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
	if gotUserAgent != wantUserAgent {
		t.Errorf("User-Agent = %q, want %q (defaults should still apply)", gotUserAgent, wantUserAgent)
	}
	if gotCustomHeader != "added" {
		t.Errorf("X-Custom = %q, want %q (AdditionalMiddlewares should have run)", gotCustomHeader, "added")
	}
}

// TestNewClient_AdditionalMiddlewaresLayerOnExplicitOverride proves
// AdditionalMiddlewares still runs when Config.Middlewares is also set
// (an explicit override, which on its own replaces the defaults
// entirely) - AdditionalMiddlewares layers on top of whichever chain
// Middlewares resolves to, not only the default one.
func TestNewClient_AdditionalMiddlewaresLayerOnExplicitOverride(t *testing.T) {
	var gotOverrideHeader, gotCustomHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotOverrideHeader = r.Header.Get("X-Override")
		gotCustomHeader = r.Header.Get("X-Custom")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := NewClient(Config{
		BaseURL:               srv.URL,
		Middlewares:           []client.Middleware{withHeader("X-Override", "explicit")},
		AdditionalMiddlewares: []client.Middleware{withHeader("X-Custom", "added")},
	})

	if _, err := c.client.Get(context.Background(), "test", url.Values{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotOverrideHeader != "explicit" {
		t.Errorf("X-Override = %q, want %q (explicit Middlewares should still apply)", gotOverrideHeader, "explicit")
	}
	if gotCustomHeader != "added" {
		t.Errorf("X-Custom = %q, want %q (AdditionalMiddlewares should layer on top of an explicit override too)", gotCustomHeader, "added")
	}
}
