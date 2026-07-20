package client

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
)

func TestClient_Get_MaxResponseBytes(t *testing.T) {
	const limit = 10 // bytes

	body := func(n int) []byte {
		return bytes.Repeat([]byte("a"), n)
	}

	tests := []struct {
		name       string
		bodySize   int
		wantErr    bool
		wantTooBig bool
	}{
		{name: "under the limit", bodySize: limit - 1, wantErr: false},
		{name: "exactly at the limit", bodySize: limit, wantErr: false},
		{name: "one byte over the limit", bodySize: limit + 1, wantErr: true, wantTooBig: true},
		{name: "well over the limit", bodySize: limit * 100, wantErr: true, wantTooBig: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write(body(tt.bodySize))
			}))
			defer srv.Close()

			c := NewClient(Config{
				BaseURL:          srv.URL,
				MaxResponseBytes: limit,
			})

			resp, err := c.Get(context.Background(), "test", nil)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected an error for a %d-byte body against a %d-byte limit, got nil", tt.bodySize, limit)
				}
				if tt.wantTooBig && !errors.Is(err, models.ErrResponseTooLarge) {
					t.Errorf("expected errors.Is(err, models.ErrResponseTooLarge), got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(resp.Body) != tt.bodySize {
				t.Errorf("expected body length %d, got %d", tt.bodySize, len(resp.Body))
			}
		})
	}
}

func TestClient_Get_DefaultMaxResponseBytes(t *testing.T) {
	c := NewClient(Config{BaseURL: "http://example.com"})
	if c.maxResponseBytes != DefaultMaxResponseBytes {
		t.Errorf("expected maxResponseBytes to default to %d, got %d", DefaultMaxResponseBytes, c.maxResponseBytes)
	}
}

func TestClientGetOversizedErrorResponsePreservesStatusError(t *testing.T) {
	client := NewClient(Config{
		BaseURL:          "https://api.example.com",
		MaxResponseBytes: 10,
		HTTPClient: httpClientFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(bytes.Repeat([]byte("x"), 11))),
			}, nil
		}),
	})

	_, err := client.Get(context.Background(), "missing", nil)
	if !errors.Is(err, models.ErrNotFound) {
		t.Fatalf("expected the upstream 404 to remain visible, got %v", err)
	}
	if errors.Is(err, models.ErrResponseTooLarge) {
		t.Errorf("expected the upstream status error, not ErrResponseTooLarge")
	}
}
