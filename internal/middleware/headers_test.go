package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestWithHeaders_IdempotentAcrossRepeatedApplication is a regression test
// for the Header.Add -> Header.Set fix: retry middleware reuses the same
// *http.Request across attempts, so if every configured value were
// applied with Add, each retry would pile another copy of the same
// headers onto what the previous attempt already added.
func TestWithHeaders_IdempotentAcrossRepeatedApplication(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})

	headers := http.Header{}
	headers.Set("X-Custom", "value")

	wrapped := WithHeaders(headers)(rt)
	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)

	for i := 0; i < 3; i++ {
		if _, err := wrapped.RoundTrip(context.Background(), req); err != nil {
			t.Fatalf("unexpected error on application %d: %v", i, err)
		}
	}

	got := req.Header.Values("X-Custom")
	if len(got) != 1 {
		t.Fatalf("expected exactly 1 value for X-Custom after 3 applications, got %d: %v", len(got), got)
	}
	if got[0] != "value" {
		t.Errorf("expected value %q, got %q", "value", got[0])
	}
}

func TestWithHeaders_MultipleValuesForSameKeyPreserved(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})

	headers := http.Header{"Accept": {"application/json", "text/plain"}}
	wrapped := WithHeaders(headers)(rt)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	if _, err := wrapped.RoundTrip(context.Background(), req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := req.Header.Values("Accept")
	if len(got) != 2 || got[0] != "application/json" || got[1] != "text/plain" {
		t.Errorf("expected both configured values preserved in order, got %v", got)
	}
}

func TestWithUserAgent_DoesNotOverrideExisting(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})

	wrapped := WithUserAgent("default-agent")(rt)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.Header.Set("User-Agent", "custom-agent")

	if _, err := wrapped.RoundTrip(context.Background(), req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.Header.Get("User-Agent"); got != "custom-agent" {
		t.Errorf("expected existing User-Agent preserved, got %q", got)
	}
}

func TestWithUserAgent_SetsWhenAbsent(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})

	wrapped := WithUserAgent("default-agent")(rt)
	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)

	if _, err := wrapped.RoundTrip(context.Background(), req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.Header.Get("User-Agent"); got != "default-agent" {
		t.Errorf("expected default-agent, got %q", got)
	}
}

func TestWithReferer_SetsOnlyWhenAbsent(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})
	wrapped := WithReferer("https://www.nba.com/")(rt)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.Header.Set("Referer", "https://custom.example.com/")

	if _, err := wrapped.RoundTrip(context.Background(), req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.Header.Get("Referer"); got != "https://custom.example.com/" {
		t.Errorf("expected existing Referer preserved, got %q", got)
	}
}

func TestWithAccept_SetsOnlyWhenAbsent(t *testing.T) {
	rt := RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return fakeResponse(http.StatusOK, nil), nil
	})
	wrapped := WithAccept("application/json")(rt)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)

	if _, err := wrapped.RoundTrip(context.Background(), req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.Header.Get("Accept"); got != "application/json" {
		t.Errorf("expected Accept set when absent, got %q", got)
	}
}
