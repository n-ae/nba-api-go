package client

import (
	"context"
	"net/http"
)

// Middleware wraps a RoundTripper to add cross-cutting behavior (retries,
// rate limiting, logging, custom headers, ...). A Middleware receives the
// next RoundTripper in the chain and returns a new one that wraps it.
type Middleware func(RoundTripper) RoundTripper

// RoundTripper is the context-aware transport interface that middleware
// and the base HTTP transport both implement. It mirrors http.RoundTripper
// but threads a context.Context through explicitly instead of relying on
// the one attached to the request.
type RoundTripper interface {
	RoundTrip(ctx context.Context, req *http.Request) (*http.Response, error)
}

// RoundTripperFunc adapts a plain function to RoundTripper, the way
// http.HandlerFunc adapts a function to http.Handler.
type RoundTripperFunc func(ctx context.Context, req *http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(ctx context.Context, req *http.Request) (*http.Response, error) {
	return f(ctx, req)
}

// Chain composes middlewares into a single Middleware. Middlewares run in
// the order given: the first one in the list is outermost and runs first
// on the way in (and last on the way out).
func Chain(middlewares ...Middleware) Middleware {
	return func(next RoundTripper) RoundTripper {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
