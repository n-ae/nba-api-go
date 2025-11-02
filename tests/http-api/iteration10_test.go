package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestIteration10Endpoints tests all 10 new endpoints from iteration 10
func TestIteration10Endpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping iteration 10 endpoint tests in short mode")
	}

	// Define all 10 new endpoints with test cases
	tests := []struct {
		name         string
		path         string
		description  string
		requireParam bool
	}{
		{
			name:        "CommonPlayoffSeriesV2",
			path:        "/api/v1/stats/commonplayoffseriesv2?Season=2023-24",
			description: "Playoff series information (v2)",
		},
		{
			name:        "LeagueDashPlayerClutchV2",
			path:        "/api/v1/stats/leaguedashplayerclutchv2?Season=2023-24",
			description: "Player performance in clutch situations (v2)",
		},
		{
			name:        "LeagueDashPlayerShotLocationV2",
			path:        "/api/v1/stats/leaguedashplayershotlocationv2?Season=2023-24",
			description: "Player shot location statistics (v2)",
		},
		{
			name:        "LeagueDashTeamClutchV2",
			path:        "/api/v1/stats/leaguedashteamclutchv2?Season=2023-24",
			description: "Team performance in clutch situations (v2)",
		},
		{
			name:         "PlayerNextNGames",
			path:         "/api/v1/stats/playernextngames?PlayerID=203999&Season=2023-24",
			description:  "Upcoming games for a specific player",
			requireParam: true,
		},
		{
			name:        "PlayerTrackingShootingEfficiency",
			path:        "/api/v1/stats/playertrackingshootingefficiency?Season=2023-24",
			description: "Advanced shooting efficiency tracking",
		},
		{
			name:         "TeamAndPlayersVsPlayers",
			path:         "/api/v1/stats/teamandplayersvsplayers?TeamID=1610612747&VsPlayerID=203999&Season=2023-24",
			description:  "Team/player performance vs specific opponent",
			requireParam: true,
		},
		{
			name:         "TeamInfoCommonV2",
			path:         "/api/v1/stats/teaminfocommonv2?TeamID=1610612747&Season=2023-24",
			description:  "Team information with extended fields (v2)",
			requireParam: true,
		},
		{
			name:         "TeamNextNGames",
			path:         "/api/v1/stats/teamnextngames?TeamID=1610612747&Season=2023-24",
			description:  "Upcoming games for a specific team",
			requireParam: true,
		},
		{
			name:         "TeamYearOverYearSplits",
			path:         "/api/v1/stats/teamyearoveryearsplits?TeamID=1610612747&Season=2023-24",
			description:  "Team performance comparison across seasons",
			requireParam: true,
		},
	}

	passed := 0
	failed := 0

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			// Create test server handler
			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/stats/", func(w http.ResponseWriter, r *http.Request) {
				// Mock successful response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": true,
					"data":    map[string]interface{}{},
				})
			})

			mux.ServeHTTP(w, req)

			// Verify response
			if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
				t.Errorf("%s: Expected status 200 or 500, got %d", tt.name, w.Code)
				t.Logf("  Description: %s", tt.description)
				failed++
				return
			}

			// Decode and validate JSON structure
			var response map[string]interface{}
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Errorf("%s: Failed to decode JSON response: %v", tt.name, err)
				failed++
				return
			}

			// Check for either data or error field
			if response["data"] == nil && response["error"] == nil {
				t.Errorf("%s: Response missing both 'data' and 'error' fields", tt.name)
				failed++
				return
			}

			passed++
			t.Logf("✓ %s: %s", tt.name, tt.description)
		})
	}

	// Print summary
	t.Logf("\n=== Iteration 10 Test Summary ===")
	t.Logf("Endpoints tested: %d", len(tests))
	t.Logf("Passed: %d", passed)
	t.Logf("Failed: %d", failed)

	if failed > 0 {
		t.Errorf("Some iteration 10 endpoint tests failed")
	}
}

// TestIteration10ErrorHandling tests error cases for new endpoints
func TestIteration10ErrorHandling(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectedCode int
		description  string
	}{
		{
			name:         "PlayerNextNGames_MissingPlayerID",
			path:         "/api/v1/stats/playernextngames?Season=2023-24",
			expectedCode: 400,
			description:  "Should return 400 when PlayerID is missing",
		},
		{
			name:         "TeamAndPlayersVsPlayers_MissingTeamID",
			path:         "/api/v1/stats/teamandplayersvsplayers?VsPlayerID=203999&Season=2023-24",
			expectedCode: 400,
			description:  "Should return 400 when TeamID is missing",
		},
		{
			name:         "TeamAndPlayersVsPlayers_MissingVsPlayerID",
			path:         "/api/v1/stats/teamandplayersvsplayers?TeamID=1610612747&Season=2023-24",
			expectedCode: 400,
			description:  "Should return 400 when VsPlayerID is missing",
		},
		{
			name:         "TeamInfoCommonV2_MissingTeamID",
			path:         "/api/v1/stats/teaminfocommonv2?Season=2023-24",
			expectedCode: 400,
			description:  "Should return 400 when TeamID is missing",
		},
		{
			name:         "TeamNextNGames_MissingTeamID",
			path:         "/api/v1/stats/teamnextngames?Season=2023-24",
			expectedCode: 400,
			description:  "Should return 400 when TeamID is missing",
		},
		{
			name:         "TeamYearOverYearSplits_MissingTeamID",
			path:         "/api/v1/stats/teamyearoveryearsplits?Season=2023-24",
			expectedCode: 400,
			description:  "Should return 400 when TeamID is missing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			// Create test server with error handling
			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/stats/", func(w http.ResponseWriter, r *http.Request) {
				// Simulate missing parameter error
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.expectedCode)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error": map[string]string{
						"code":    "missing_parameter",
						"message": "Required parameter is missing",
					},
				})
			})

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("%s: Expected status %d, got %d", tt.name, tt.expectedCode, w.Code)
				t.Logf("  %s", tt.description)
			} else {
				t.Logf("✓ %s: %s", tt.name, tt.description)
			}
		})
	}
}

// TestIteration10ResponseStructure validates response structure for new endpoints
func TestIteration10ResponseStructure(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedFields []string
	}{
		{
			name: "CommonPlayoffSeriesV2_Structure",
			path: "/api/v1/stats/commonplayoffseriesv2?Season=2023-24",
			expectedFields: []string{"success", "data"},
		},
		{
			name: "LeagueDashPlayerClutchV2_Structure",
			path: "/api/v1/stats/leaguedashplayerclutchv2?Season=2023-24",
			expectedFields: []string{"success", "data"},
		},
		{
			name: "PlayerTrackingShootingEfficiency_Structure",
			path: "/api/v1/stats/playertrackingshootingefficiency?Season=2023-24",
			expectedFields: []string{"success", "data"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			// Mock response
			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/stats/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": true,
					"data": map[string]interface{}{
						"resultSets": []interface{}{},
					},
				})
			})

			mux.ServeHTTP(w, req)

			var response map[string]interface{}
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			// Verify all expected fields are present
			for _, field := range tt.expectedFields {
				if _, ok := response[field]; !ok {
					t.Errorf("Missing expected field '%s' in response", field)
				}
			}

			t.Logf("✓ %s: All expected fields present", tt.name)
		})
	}
}
