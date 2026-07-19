// Package middleware holds the concrete middleware constructors used to
// build this SDK's default request pipeline (retry, rate limiting,
// headers, logging). The composable seam itself - Middleware,
// RoundTripper, RoundTripperFunc, and Chain - is defined in pkg/client so
// external consumers can reference it without importing this internal
// package; the names below are aliases for those same types, kept so the
// constructors in this package didn't need to change.
package middleware

import (
	"github.com/n-ae/nba-api-go/pkg/client"
)

type Middleware = client.Middleware

type RoundTripper = client.RoundTripper

type RoundTripperFunc = client.RoundTripperFunc

var Chain = client.Chain
