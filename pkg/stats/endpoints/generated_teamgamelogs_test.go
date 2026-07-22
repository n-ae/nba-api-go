package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// TestGetTeamGameLogs_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamGameLogs` instead.
func TestGetTeamGameLogs_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "TeamGameLogs", "headers": ["SEASON_YEAR", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "GAME_ID", "GAME_DATE", "MATCHUP", "WL", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS", "NBA_FANTASY_PTS", "DD2", "TD3"], "rowSet": [["test", 1, "test", "test", "test", "test", "test", "test", 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.5, 1.5, 1, 1]]}
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

	req := TeamGameLogsRequest{
		Season:     parameters.Season("2023-24"),
		SeasonType: parameters.SeasonType("Regular Season"),
	}

	resp, err := GetTeamGameLogs(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamGameLogs: %v", err)
	}

	if len(resp.Data.TeamGameLogs) == 0 {
		t.Errorf("expected TeamGameLogs to be populated from the synthesized fixture, got empty")
	}
}
