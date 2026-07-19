package middleware

import (
	"context"
	"net/http"
)

// WithHeaders applies headers to every request. The first value for each
// key replaces whatever is already on the request (via Header.Set);
// remaining values for that same key, if headers itself specifies more
// than one, are appended with Header.Add. Using Set for the first value
// matters because retry middleware reuses the same *http.Request across
// attempts - if every value were applied with Add, each retry would pile
// another copy of these headers onto the ones the previous attempt added.
func WithHeaders(headers http.Header) Middleware {
	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			for key, values := range headers {
				for i, value := range values {
					if i == 0 {
						req.Header.Set(key, value)
					} else {
						req.Header.Add(key, value)
					}
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
