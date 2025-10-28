package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type defaultLogger struct{}

func (l *defaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func WithLogging(logger Logger) Middleware {
	if logger == nil {
		logger = &defaultLogger{}
	}

	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			start := time.Now()
			logger.Printf("Request: %s %s", req.Method, req.URL.String())

			resp, err := next.RoundTrip(ctx, req)

			duration := time.Since(start)
			if err != nil {
				logger.Printf("Request failed: %s %s (%v) - %v", req.Method, req.URL.String(), duration, err)
			} else {
				logger.Printf("Request completed: %s %s (%v) - Status: %d", req.Method, req.URL.String(), duration, resp.StatusCode)
			}

			return resp, err
		})
	}
}

func WithDebugLogging(logger Logger) Middleware {
	if logger == nil {
		logger = &defaultLogger{}
	}

	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			start := time.Now()
			logger.Printf("=== Request Start ===")
			logger.Printf("Method: %s", req.Method)
			logger.Printf("URL: %s", req.URL.String())
			logger.Printf("Headers:")
			for key, values := range req.Header {
				for _, value := range values {
					logger.Printf("  %s: %s", key, value)
				}
			}

			resp, err := next.RoundTrip(ctx, req)

			duration := time.Since(start)
			logger.Printf("Duration: %v", duration)

			if err != nil {
				logger.Printf("Error: %v", err)
			} else {
				logger.Printf("Status: %d %s", resp.StatusCode, resp.Status)
				logger.Printf("Response Headers:")
				for key, values := range resp.Header {
					for _, value := range values {
						logger.Printf("  %s: %s", key, value)
					}
				}
			}
			logger.Printf("=== Request End ===")

			return resp, err
		})
	}
}

func WithRequestIDLogging(logger Logger) Middleware {
	if logger == nil {
		logger = &defaultLogger{}
	}

	return func(next RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
			requestID := req.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano())
				req.Header.Set("X-Request-ID", requestID)
			}

			start := time.Now()
			logger.Printf("[%s] Request: %s %s", requestID, req.Method, req.URL.String())

			resp, err := next.RoundTrip(ctx, req)

			duration := time.Since(start)
			if err != nil {
				logger.Printf("[%s] Failed (%v): %v", requestID, duration, err)
			} else {
				logger.Printf("[%s] Completed (%v): Status %d", requestID, duration, resp.StatusCode)
			}

			return resp, err
		})
	}
}
