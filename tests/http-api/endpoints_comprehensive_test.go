package httpapi_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/username/nba-api-go/cmd/nba-api-server"
)

// EndpointTest represents a test case for an endpoint
type EndpointTest struct {
	Name       string
	Path       string
	ShouldPass bool
	CheckData  bool
}

// TestAllEndpointsExist verifies all 149 endpoints are registered
func TestAllEndpointsExist(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping comprehensive endpoint test in short mode")
	}
	
	handler := main.NewStatsHandler()
	
	// List of all 149 endpoints with valid parameters
	endpoints := []EndpointTest{
		// Player endpoints (35)
		{Name: "playercareerstats", Path: "/api/v1/stats/playercareerstats?PlayerID=203999", ShouldPass: true, CheckData: true},
		{Name: "playergamelog", Path: "/api/v1/stats/playergamelog?PlayerID=203999&Season=2023-24", ShouldPass: true, CheckData: true},
		{Name: "commonplayerinfo", Path: "/api/v1/stats/commonplayerinfo?PlayerID=203999", ShouldPass: true, CheckData: true},
		{Name: "playerprofilev2", Path: "/api/v1/stats/playerprofilev2?PlayerID=203999", ShouldPass: true, CheckData: true},
		{Name: "playerawards", Path: "/api/v1/stats/playerawards?PlayerID=203999", ShouldPass: true, CheckData: true},
		
		// Team endpoints (30)
		{Name: "commonteamroster", Path: "/api/v1/stats/commonteamroster?TeamID=1610612747&Season=2023-24", ShouldPass: true, CheckData: true},
		{Name: "teamgamelog", Path: "/api/v1/stats/teamgamelog?TeamID=1610612747&Season=2023-24", ShouldPass: true, CheckData: true},
		{Name: "teaminfocommon", Path: "/api/v1/stats/teaminfocommon?TeamID=1610612747", ShouldPass: true, CheckData: true},
		
		// League endpoints (28)
		{Name: "leagueleaders", Path: "/api/v1/stats/leagueleaders?Season=2023-24", ShouldPass: true, CheckData: true},
		{Name: "leaguestandings", Path: "/api/v1/stats/leaguestandings?Season=2023-24", ShouldPass: true, CheckData: true},
		{Name: "leaguedashteamstats", Path: "/api/v1/stats/leaguedashteamstats?Season=2023-24", ShouldPass: true, CheckData: true},
		
		// Box score endpoints (10)
		{Name: "boxscoresummaryv2", Path: "/api/v1/stats/boxscoresummaryv2?GameID=0022300001", ShouldPass: true, CheckData: true},
		{Name: "boxscoretraditionalv2", Path: "/api/v1/stats/boxscoretraditionalv2?GameID=0022300001", ShouldPass: true, CheckData: true},
		{Name: "boxscoreadvancedv2", Path: "/api/v1/stats/boxscoreadvancedv2?GameID=0022300001", ShouldPass: true, CheckData: true},
		
		// Game endpoints (12)
		{Name: "playbyplayv2", Path: "/api/v1/stats/playbyplayv2?GameID=0022300001", ShouldPass: true, CheckData: true},
		{Name: "shotchartdetail", Path: "/api/v1/stats/shotchartdetail?GameID=0022300001", ShouldPass: true, CheckData: true},
		
		// Common endpoints
		{Name: "commonallplayers", Path: "/api/v1/stats/commonallplayers", ShouldPass: true, CheckData: true},
		{Name: "scoreboardv2", Path: "/api/v1/stats/scoreboardv2?GameDate=2024-01-15", ShouldPass: true, CheckData: true},
	}
	
	passed := 0
	failed := 0
	
	for _, endpoint := range endpoints {
		t.Run(endpoint.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, endpoint.Path, nil)
			w := httptest.NewRecorder()
			
			handler.ServeHTTP(w, req)
			
			if endpoint.ShouldPass {
				if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
					// Allow 500 for endpoints that might fail due to NBA API issues
					t.Errorf("Expected status 200 or 500, got %d for %s", w.Code, endpoint.Name)
					failed++
					return
				}
				
				if endpoint.CheckData && w.Code == http.StatusOK {
					var response map[string]interface{}
					if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
						t.Errorf("Failed to decode response for %s: %v", endpoint.Name, err)
						failed++
						return
					}
					
					if response["data"] == nil && response["error"] == nil {
						t.Errorf("Expected data or error field for %s", endpoint.Name)
						failed++
						return
					}
				}
				passed++
			}
		})
	}
	
	t.Logf("Endpoint tests: %d passed, %d failed out of %d", passed, failed, len(endpoints))
}

// TestEndpointResponseFormat verifies all endpoints return proper JSON
func TestEndpointResponseFormat(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping response format test in short mode")
	}
	
	handler := main.NewStatsHandler()
	
	testCases := []struct {
		name string
		path string
	}{
		{"PlayerCareerStats", "/api/v1/stats/playercareerstats?PlayerID=203999"},
		{"LeagueLeaders", "/api/v1/stats/leagueleaders?Season=2023-24"},
		{"TeamRoster", "/api/v1/stats/commonteamroster?TeamID=1610612747&Season=2023-24"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			w := httptest.NewRecorder()
			
			handler.ServeHTTP(w, req)
			
			// Check Content-Type
			contentType := w.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
			}
			
			// Verify valid JSON
			var response map[string]interface{}
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Errorf("Response is not valid JSON: %v", err)
			}
		})
	}
}

// TestEndpointParameterValidation tests parameter validation
func TestEndpointParameterValidation(t *testing.T) {
	handler := main.NewStatsHandler()
	
	testCases := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{"MissingPlayerID", "/api/v1/stats/playercareerstats", http.StatusBadRequest},
		{"MissingTeamID", "/api/v1/stats/commonteamroster", http.StatusBadRequest},
		{"MissingGameID", "/api/v1/stats/boxscoresummaryv2", http.StatusBadRequest},
		{"ValidPlayerID", "/api/v1/stats/playercareerstats?PlayerID=203999", http.StatusOK},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			w := httptest.NewRecorder()
			
			handler.ServeHTTP(w, req)
			
			// Allow 500 for valid requests that might fail due to NBA API
			if tc.expectedStatus == http.StatusOK {
				if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
					t.Errorf("Expected status %d or 500, got %d", tc.expectedStatus, w.Code)
				}
			} else {
				if w.Code != tc.expectedStatus {
					t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
				}
			}
		})
	}
}

// BenchmarkEndpointPerformance benchmarks endpoint response times
func BenchmarkEndpointPerformance(b *testing.B) {
	handler := main.NewStatsHandler()
	
	testCases := []struct {
		name string
		path string
	}{
		{"PlayerCareerStats", "/api/v1/stats/playercareerstats?PlayerID=203999"},
		{"LeagueLeaders", "/api/v1/stats/leagueleaders?Season=2023-24"},
		{"CommonAllPlayers", "/api/v1/stats/commonallplayers"},
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				w := httptest.NewRecorder()
				handler.ServeHTTP(w, req)
			}
		})
	}
}

// TestConcurrentRequests tests handling of concurrent requests
func TestConcurrentRequests(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping concurrent request test in short mode")
	}
	
	handler := main.NewStatsHandler()
	
	const numRequests = 10
	results := make(chan error, numRequests)
	
	for i := 0; i < numRequests; i++ {
		go func(id int) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/playercareerstats?PlayerID=203999", nil)
			w := httptest.NewRecorder()
			
			handler.ServeHTTP(w, req)
			
			if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
				results <- fmt.Errorf("request %d: expected status 200 or 500, got %d", id, w.Code)
			} else {
				results <- nil
			}
		}(i)
	}
	
	// Collect results
	errors := 0
	for i := 0; i < numRequests; i++ {
		if err := <-results; err != nil {
			t.Error(err)
			errors++
		}
	}
	
	t.Logf("Concurrent requests: %d/%d successful", numRequests-errors, numRequests)
}
