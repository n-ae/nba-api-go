package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueHustleStatsTeamLeaders_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueHustleStatsTeamLeaders` instead.
func TestGetLeagueHustleStatsTeamLeaders_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "HustleStatsTeamLeaders", "headers": ["TEAM_ID", "TEAM_NAME", "GP", "MIN", "SCREEN_ASSISTS", "SCREEN_AST_PTS", "OFF_LOOSE_BALLS_RECOVERED", "DEF_LOOSE_BALLS_RECOVERED", "LOOSE_BALLS_RECOVERED", "OFF_BOXOUTS", "DEF_BOXOUTS", "BOX_OUTS", "CONTESTED_SHOTS", "CONTESTED_SHOTS_2PT", "CONTESTED_SHOTS_3PT", "CHARGES_DRAWN", "DEFLECTIONS"], "rowSet": [[1, "test", 1, 1.5, "test", 1.5, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]}
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

	req := LeagueHustleStatsTeamLeadersRequest{}

	resp, err := GetLeagueHustleStatsTeamLeaders(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueHustleStatsTeamLeaders: %v", err)
	}

	if len(resp.Data.HustleStatsTeamLeaders) == 0 {
		t.Errorf("expected HustleStatsTeamLeaders to be populated from the synthesized fixture, got empty")
	}
}
