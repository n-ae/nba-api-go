package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueHustleStatsPlayer_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueHustleStatsPlayer` instead.
func TestGetLeagueHustleStatsPlayer_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "HustleStatsPlayer", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "AGE", "GP", "W", "L", "W_PCT", "MIN", "CONTESTED_SHOTS", "CONTESTED_SHOTS_2PT", "CONTESTED_SHOTS_3PT", "DEFLECTIONS", "CHARGES_DRAWN", "SCREEN_ASSISTS", "SCREEN_AST_PTS", "OFF_LOOSE_BALLS_RECOVERED", "DEF_LOOSE_BALLS_RECOVERED", "LOOSE_BALLS_RECOVERED", "OFF_BOXOUTS", "DEF_BOXOUTS", "BOX_OUT_PLAYER_TEAM_REBS", "BOX_OUT_PLAYER_REBS", "BOX_OUTS"], "rowSet": [[1, "test", 1, "test", 1, 1, "test", "test", 1.5, 1.5, "test", "test", "test", "test", "test", "test", 1.5, "test", "test", "test", "test", "test", 1.5, 1.5, "test"]]}
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

	req := LeagueHustleStatsPlayerRequest{}

	resp, err := GetLeagueHustleStatsPlayer(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueHustleStatsPlayer: %v", err)
	}

	if len(resp.Data.HustleStatsPlayer) == 0 {
		t.Errorf("expected HustleStatsPlayer to be populated from the synthesized fixture, got empty")
	}
}
