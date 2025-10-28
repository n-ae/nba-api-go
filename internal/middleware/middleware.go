package middleware

import (
	"context"
	"net/http"
)

type Middleware func(RoundTripper) RoundTripper

type RoundTripper interface {
	RoundTrip(ctx context.Context, req *http.Request) (*http.Response, error)
}

type RoundTripperFunc func(ctx context.Context, req *http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(ctx context.Context, req *http.Request) (*http.Response, error) {
	return f(ctx, req)
}

func Chain(middlewares ...Middleware) Middleware {
	return func(next RoundTripper) RoundTripper {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
