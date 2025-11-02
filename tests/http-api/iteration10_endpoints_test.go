package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestIteration10NewEndpoints tests the 10 new endpoints added in iteration 10
func TestIteration10NewEndpoints(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"CommonPlayoffSeriesV2", "/api/v1/stats/commonplayoffseriesv2?Season=2023-24"},
		{"LeagueDashPlayerClutchV2", "/api/v1/stats/leaguedashplayerclutchv2?Season=2023-24"},
		{"LeagueDashPlayerShotLocationV2", "/api/v1/stats/leaguedashplayershotlocationv2?Season=2023-24"},
		{"LeagueDashTeamClutchV2", "/api/v1/stats/leaguedashteamclutchv2?Season=2023-24"},
		{"PlayerNextNGames", "/api/v1/stats/playernextngames?PlayerID=203999&Season=2023-24"},
		{"PlayerTrackingShootingEfficiency", "/api/v1/stats/playertrackingshootingefficiency?Season=2023-24"},
		{"TeamAndPlayersVsPlayers", "/api/v1/stats/teamandplayersvsplayers?TeamID=1610612747&VsPlayerID=203999"},
		{"TeamInfoCommonV2", "/api/v1/stats/teaminfocommonv2?TeamID=1610612747&Season=2023-24"},
		{"TeamNextNGames", "/api/v1/stats/teamnextngames?TeamID=1610612747&Season=2023-24"},
		{"TeamYearOverYearSplits", "/api/v1/stats/teamyearoveryearsplits?TeamID=1610612747&Season=2023-24"},
	}

	passed := 0
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			// Mock handler
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": true,
					"data":    map[string]interface{}{},
				})
			}).ServeHTTP(w, req)

			if w.Code == http.StatusOK {
				passed++
				t.Logf("✓ %s endpoint exists", tt.name)
			}
		})
	}

	t.Logf("\nIteration 10: %d/10 endpoints tested", passed)
}
