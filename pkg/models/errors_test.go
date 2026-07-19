package models

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestNewAPIError_TruncatesBody(t *testing.T) {
	t.Run("short body kept in full", func(t *testing.T) {
		body := []byte("short body")
		apiErr := NewAPIError(http.StatusInternalServerError, "http://example.com", "server error", body, nil)
		if apiErr.Body != string(body) {
			t.Errorf("expected body %q, got %q", body, apiErr.Body)
		}
	})

	t.Run("long body truncated to maxErrorBodyBytes", func(t *testing.T) {
		body := []byte(strings.Repeat("a", maxErrorBodyBytes+500))
		apiErr := NewAPIError(http.StatusInternalServerError, "http://example.com", "server error", body, nil)
		if len(apiErr.Body) != maxErrorBodyBytes {
			t.Errorf("expected truncated body length %d, got %d", maxErrorBodyBytes, len(apiErr.Body))
		}
	})

	t.Run("nil body produces empty string, not a panic", func(t *testing.T) {
		apiErr := NewAPIError(http.StatusInternalServerError, "http://example.com", "server error", nil, nil)
		if apiErr.Body != "" {
			t.Errorf("expected empty body, got %q", apiErr.Body)
		}
	})
}

func TestHTTPStatusToError_MapsStatusAndCarriesBody(t *testing.T) {
	body := []byte(`{"error": "player not found"}`)

	err := HTTPStatusToError(http.StatusNotFound, "http://example.com/player/1", body)

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected an *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, apiErr.StatusCode)
	}
	if apiErr.Body != string(body) {
		t.Errorf("expected body %q, got %q", body, apiErr.Body)
	}
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected errors.Is(err, ErrNotFound) to hold")
	}
}

func TestHTTPStatusToError_UnmappedSuccessStatusReturnsNil(t *testing.T) {
	if err := HTTPStatusToError(http.StatusOK, "http://example.com", nil); err != nil {
		t.Errorf("expected nil for a 200 status, got %v", err)
	}
}
