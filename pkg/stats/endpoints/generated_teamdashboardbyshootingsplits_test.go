package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamDashboardByShootingSplits_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamDashboardByShootingSplits` instead.
func TestGetTeamDashboardByShootingSplits_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "OverallTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "Shot5FTTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "SHOT_5FT", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "ShotAreaTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "SHOT_AREA", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]}
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

	req := TeamDashboardByShootingSplitsRequest{
		TeamID: "1",
	}

	resp, err := GetTeamDashboardByShootingSplits(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamDashboardByShootingSplits: %v", err)
	}

	if len(resp.Data.OverallTeamDashboard) == 0 {
		t.Errorf("expected OverallTeamDashboard to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.Shot5FTTeamDashboard) == 0 {
		t.Errorf("expected Shot5FTTeamDashboard to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.ShotAreaTeamDashboard) == 0 {
		t.Errorf("expected ShotAreaTeamDashboard to be populated from the synthesized fixture, got empty")
	}
}
