package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerTrackingSpeedDistance_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerTrackingSpeedDistance` instead.
func TestGetPlayerTrackingSpeedDistance_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrackingSpeedDistance", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "DIST_FEET", "DIST_MILES", "DIST_MILES_OFF", "DIST_MILES_DEF", "AVG_SPEED", "AVG_SPEED_OFF", "AVG_SPEED_DEF"], "rowSet": [[1, "test", 1, "test", 1, 1.5, "test", "test", "test", "test", "test", "test", "test"]]}
	]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}

	req := PlayerTrackingSpeedDistanceRequest{}

	resp, err := GetPlayerTrackingSpeedDistance(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerTrackingSpeedDistance: %v", err)
	}

	if len(resp.Data.PlayerTrackingSpeedDistance) == 0 {
		t.Errorf("expected PlayerTrackingSpeedDistance to be populated from the synthesized fixture, got empty")
	}
}
