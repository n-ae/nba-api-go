package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueGameFinder_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueGameFinder` instead.
func TestGetLeagueGameFinder_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeagueGameFinderResults", "headers": ["SEASON_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "GAME_ID", "GAME_DATE", "MATCHUP", "WL", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "STL", "BLK", "TOV", "PF", "PTS", "PLUS_MINUS"], "rowSet": [["test", 1, "test", "test", "test", "test", "test", "test", 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.5]]}
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

	req := LeagueGameFinderRequest{}

	resp, err := GetLeagueGameFinder(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueGameFinder: %v", err)
	}

	if len(resp.Data.LeagueGameFinderResults) == 0 {
		t.Errorf("expected LeagueGameFinderResults to be populated from the synthesized fixture, got empty")
	}
}
