package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
)

func TestWriteEndpointError_UsesAPIErrorStatusCode(t *testing.T) {
	apiErr := models.NewAPIError(http.StatusNotFound, "http://stats.nba.com/whatever", "resource not found", nil, models.ErrNotFound)

	w := httptest.NewRecorder()
	writeEndpointError(w, apiErr)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected the upstream 404 to be reused, got %d", w.Code)
	}
}

func TestWriteEndpointError_APIErrorRateLimited(t *testing.T) {
	apiErr := models.NewAPIError(http.StatusTooManyRequests, "http://stats.nba.com/whatever", "rate limited", nil, models.ErrRateLimited)

	w := httptest.NewRecorder()
	writeEndpointError(w, apiErr)

	if w.Code != http.StatusTooManyRequests {
		t.Errorf("expected the upstream 429 to be reused, got %d", w.Code)
	}
}

func TestWriteEndpointError_NonAPIErrorFallsBackTo500(t *testing.T) {
	w := httptest.NewRecorder()
	writeEndpointError(w, errors.New("connection reset by peer"))

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 for a non-APIError, got %d", w.Code)
	}
}

func TestWriteEndpointError_WrappedAPIErrorStillUnwraps(t *testing.T) {
	apiErr := models.NewAPIError(http.StatusServiceUnavailable, "http://stats.nba.com/whatever", "server error", nil, nil)
	wrapped := errors.Join(errors.New("context"), apiErr)

	w := httptest.NewRecorder()
	writeEndpointError(w, wrapped)

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("expected the wrapped APIError's 503 to be reused, got %d", w.Code)
	}
}
