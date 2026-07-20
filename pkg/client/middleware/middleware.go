// Package middleware holds the concrete middleware constructors used to
// build this SDK's default request pipeline (retry, rate limiting,
// headers, logging) - WithRetry, RetryConfig/DefaultRetryConfig,
// WithPerHostRateLimit, WithHeaders, WithUserAgent, WithReferer,
// WithAccept, and the WithLogging family. It was previously
// internal/middleware, inaccessible outside this module; moved here so
// external consumers can import it directly to tune retry/backoff, rate
// limits, or headers instead of only being able to add middleware
// alongside these defaults without being able to reconfigure them (see
// docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md finding
// #14). The composable seam itself - Middleware, RoundTripper,
// RoundTripperFunc, and Chain - is defined in the parent pkg/client
// package; the names below are aliases for those same types, kept so the
// constructors in this package didn't need to change.
package middleware

import (
	"github.com/n-ae/nba-api-go/v2/pkg/client"
)

type Middleware = client.Middleware

type RoundTripper = client.RoundTripper

type RoundTripperFunc = client.RoundTripperFunc

var Chain = client.Chain
