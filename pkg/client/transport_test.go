package client

import (
	"net/http"
	"testing"
)

// TestNewClient_DefaultTransport is a regression test for ADR 003: the
// default transport must be cloned from http.DefaultTransport (so proxy
// support and HTTP/2 negotiation survive) with keep-alives left enabled,
// not built from a zero-valued struct with DisableKeepAlives set.
func TestNewClient_DefaultTransport(t *testing.T) {
	c := NewClient(Config{BaseURL: "http://example.com"})

	httpClient, ok := c.httpClient.(*http.Client)
	if !ok {
		t.Fatalf("expected the default HTTPClient to be *http.Client, got %T", c.httpClient)
	}

	transport, ok := httpClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("expected the default Transport to be *http.Transport, got %T", httpClient.Transport)
	}

	if transport.DisableKeepAlives {
		t.Error("expected keep-alives to be enabled (DisableKeepAlives=false); see ADR 003")
	}
	if transport.Proxy == nil {
		t.Error("expected Proxy support to be inherited from http.DefaultTransport, got nil")
	}
	if !transport.ForceAttemptHTTP2 {
		t.Error("expected ForceAttemptHTTP2 to be inherited from http.DefaultTransport")
	}
	if transport.TLSHandshakeTimeout <= 0 {
		t.Error("expected a positive TLSHandshakeTimeout")
	}
	if transport.ResponseHeaderTimeout <= 0 {
		t.Error("expected a positive ResponseHeaderTimeout")
	}
}

// TestNewClient_CustomHTTPClientBypassesDefaultTransport proves a
// caller-supplied HTTPClient is used as-is, with no transport
// construction happening at all - the escape hatch for anyone who does
// need different transport behavior than ADR 003's default.
func TestNewClient_CustomHTTPClientBypassesDefaultTransport(t *testing.T) {
	custom := &http.Client{}
	c := NewClient(Config{BaseURL: "http://example.com", HTTPClient: custom})

	if c.httpClient != custom {
		t.Error("expected the caller-supplied HTTPClient to be used as-is")
	}
}
