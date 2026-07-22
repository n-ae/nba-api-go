package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerGameStreakFinder_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerGameStreakFinder` instead.
func TestGetPlayerGameStreakFinder_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerGameStreakFinder", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GAME_ID", "GAME_DATE", "MATCHUP", "WL", "MIN", "PTS", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "PF", "PLUS_MINUS"], "rowSet": [[1, "test", 1, "test", "test", "test", "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/playergamestreakfinder"
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

	req := PlayerGameStreakFinderRequest{}

	resp, err := GetPlayerGameStreakFinder(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerGameStreakFinder: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerGameStreakFinder requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playergamestreakfinder")
	}

	if len(resp.Data.PlayerGameStreakFinder) == 0 {
		t.Errorf("expected PlayerGameStreakFinder to be populated from the synthesized fixture, got empty")
	}
}
