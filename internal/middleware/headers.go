package middleware

import (
	"context"
	"net/http"
)

func WithHeaders(headers http.Header) Middleware {
	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			for key, values := range headers {
				for _, value := range values {
					req.Header.Add(key, value)
				}
			}
			return next.RoundTrip(ctx, req)
		})
	}
}

func WithUserAgent(userAgent string) Middleware {
	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			if req.Header.Get("User-Agent") == "" {
				req.Header.Set("User-Agent", userAgent)
			}
			return next.RoundTrip(ctx, req)
		})
	}
}

func WithReferer(referer string) Middleware {
	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			if req.Header.Get("Referer") == "" {
				req.Header.Set("Referer", referer)
			}
			return next.RoundTrip(ctx, req)
		})
	}
}

func WithAccept(accept string) Middleware {
	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			if req.Header.Get("Accept") == "" {
				req.Header.Set("Accept", accept)
			}
			return next.RoundTrip(ctx, req)
		})
	}
}
