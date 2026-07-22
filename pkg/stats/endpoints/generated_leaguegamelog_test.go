package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// TestGetLeagueGameLog_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueGameLog` instead.
func TestGetLeagueGameLog_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeagueGameLog", "headers": ["SEASON_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "GAME_ID", "GAME_DATE", "MATCHUP", "WL", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "STL", "BLK", "TOV", "PF", "PTS", "PLUS_MINUS", "VIDEO_AVAILABLE"], "rowSet": [["test", 1, "test", "test", "test", "test", "test", "test", 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, "test"]]}
	]}`

	const wantPath = "/leaguegamelog"
	var gotPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}

	req := LeagueGameLogRequest{
		Season: parameters.Season("2023-24"),
	}

	resp, err := GetLeagueGameLog(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueGameLog: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetLeagueGameLog requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "leaguegamelog")
	}

	if len(resp.Data.LeagueGameLog) == 0 {
		t.Errorf("expected LeagueGameLog to be populated from the synthesized fixture, got empty")
	}
}
