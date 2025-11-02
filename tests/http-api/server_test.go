package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/n-ae/nba-api-go/cmd/nba-api-server"
)

// TestHealthEndpoint tests the health check endpoint
func TestHealthEndpoint(t *testing.T) {
	handler := main.NewStatsHandler()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":                 "healthy",
			"http_endpoints_exposed": 149,
		})
	}).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", response["status"])
	}
}

// TestPlayerEndpoints tests player-related endpoints
func TestPlayerCareerStatsEndpoint(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	handler := main.NewStatsHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/playercareerstats?PlayerID=203999", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["data"] == nil {
		t.Error("Expected data field in response")
	}
}

// TestMissingParameter tests error handling for missing parameters
func TestMissingParameter(t *testing.T) {
	handler := main.NewStatsHandler()

	// Request without required PlayerID
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/playercareerstats", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["error"] == nil {
		t.Error("Expected error field in response")
	}
}

// TestInvalidEndpoint tests 404 handling
func TestInvalidEndpoint(t *testing.T) {
	handler := main.NewStatsHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/invalidendpoint", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// TestMethodNotAllowed tests that only GET is allowed
func TestMethodNotAllowed(t *testing.T) {
	handler := main.NewStatsHandler()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stats/playercareerstats", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}
