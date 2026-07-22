package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamDashboardByOpponent_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamDashboardByOpponent` instead.
func TestGetTeamDashboardByOpponent_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "OverallTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "ConferenceTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "CONF", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "DivisionTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "DIV", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "OpponentTeamDashboard", "headers": ["TEAM_ID", "TEAM_NAME", "VS_TEAM_ID", "VS_TEAM_ABBREVIATION", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", 1, "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]}
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

	req := TeamDashboardByOpponentRequest{
		TeamID: "1",
	}

	resp, err := GetTeamDashboardByOpponent(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamDashboardByOpponent: %v", err)
	}

	if len(resp.Data.OverallTeamDashboard) == 0 {
		t.Errorf("expected OverallTeamDashboard to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.ConferenceTeamDashboard) == 0 {
		t.Errorf("expected ConferenceTeamDashboard to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.DivisionTeamDashboard) == 0 {
		t.Errorf("expected DivisionTeamDashboard to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.OpponentTeamDashboard) == 0 {
		t.Errorf("expected OpponentTeamDashboard to be populated from the synthesized fixture, got empty")
	}
}
