package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	server := NewServer(logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	server.handleHealth()(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if status, ok := response["status"].(string); !ok || status != "healthy" {
		t.Errorf("expected status=healthy, got %v", response["status"])
	}

	if version, ok := response["version"].(string); !ok || version == "" {
		t.Errorf("expected version to be set, got %v", response["version"])
	}

	if buildInfo, ok := response["build_info"].(map[string]interface{}); !ok {
		t.Error("expected build_info to be present")
	} else {
		if goVersion, ok := buildInfo["go_version"].(string); !ok || goVersion == "" {
			t.Error("expected go_version in build_info")
		}
	}

	if nbaAPIStatus, ok := response["nba_api_status"].(string); !ok {
		t.Error("expected nba_api_status to be present")
	} else if nbaAPIStatus != "operational" && nbaAPIStatus != "degraded" {
		t.Errorf("expected nba_api_status to be 'operational' or 'degraded', got %s", nbaAPIStatus)
	}

	if timestamp, ok := response["timestamp"].(float64); !ok || timestamp == 0 {
		t.Error("expected timestamp to be set")
	}
}

func TestCORSMiddleware(t *testing.T) {
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	server := NewServer(logger)

	handler := server.corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if origin := w.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Errorf("expected CORS origin *, got %s", origin)
	}

	if methods := w.Header().Get("Access-Control-Allow-Methods"); methods != "GET, OPTIONS" {
		t.Errorf("expected methods 'GET, OPTIONS', got %s", methods)
	}
}

func TestOPTIONSRequest(t *testing.T) {
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	server := NewServer(logger)

	handler := server.corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("handler should not be called for OPTIONS request")
	}))

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200 for OPTIONS, got %d", w.Code)
	}
}

func TestServerRoutes(t *testing.T) {
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	server := NewServer(logger)
	routes := server.Routes()

	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "health endpoint",
			method:         http.MethodGet,
			path:           "/health",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "unknown endpoint",
			method:         http.MethodGet,
			path:           "/unknown",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()

			routes.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}
		})
	}
}

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	writeError(w, http.StatusBadRequest, "invalid_param", "PlayerID is required")

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}

	if success, ok := response["success"].(bool); !ok || success {
		t.Error("expected success=false")
	}

	errorObj, ok := response["error"].(map[string]interface{})
	if !ok {
		t.Fatal("expected error object")
	}

	if code := errorObj["code"].(string); code != "invalid_param" {
		t.Errorf("expected error code 'invalid_param', got %s", code)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	server := NewServer(logger)

	handlerCalled := false
	handler := server.loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if !handlerCalled {
		t.Error("expected handler to be called")
	}

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestStatsHandlerMethodNotAllowed(t *testing.T) {
	handler := NewStatsHandler()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stats/playergamelog", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if success, ok := response["success"].(bool); !ok || success {
		t.Error("expected success=false for method not allowed")
	}
}

func TestStatsHandlerInvalidEndpoint(t *testing.T) {
	handler := NewStatsHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/nonexistent", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if success, ok := response["success"].(bool); !ok || success {
		t.Error("expected success=false for invalid endpoint")
	}
}

func TestMetricsEndpoint(t *testing.T) {
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	server := NewServer(logger)

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	w := httptest.NewRecorder()

	server.handleMetrics()(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode metrics response: %v", err)
	}

	if _, ok := response["uptime_seconds"]; !ok {
		t.Error("expected uptime_seconds in metrics")
	}

	if _, ok := response["total_requests"]; !ok {
		t.Error("expected total_requests in metrics")
	}
}

func TestRateLimiting(t *testing.T) {
	rl := NewRateLimiter(2, 2)

	handler := rl.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 3; i++ {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		if i < 2 {
			if w.Code != http.StatusOK {
				t.Errorf("request %d: expected status 200, got %d", i, w.Code)
			}
		} else {
			if w.Code != http.StatusTooManyRequests {
				t.Errorf("request %d: expected status 429, got %d", i, w.Code)
			}
		}
	}
}

func TestMetricsTracking(t *testing.T) {
	metrics := NewMetrics()

	metrics.RecordRequest("/api/test", 200, 10000000)
	metrics.RecordRequest("/api/test", 404, 5000000)
	metrics.RecordRequest("/health", 200, 1000000)

	snapshot := metrics.GetSnapshot()

	if snapshot.TotalRequests != 3 {
		t.Errorf("expected 3 total requests, got %d", snapshot.TotalRequests)
	}

	if snapshot.TotalErrors != 1 {
		t.Errorf("expected 1 error, got %d", snapshot.TotalErrors)
	}

	if snapshot.RequestsByStatus[200] != 2 {
		t.Errorf("expected 2 requests with status 200, got %d", snapshot.RequestsByStatus[200])
	}

	if snapshot.RequestsByPath["/api/test"] != 2 {
		t.Errorf("expected 2 requests to /api/test, got %d", snapshot.RequestsByPath["/api/test"])
	}
}
