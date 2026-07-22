package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamGameStreakFinder_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamGameStreakFinder` instead.
func TestGetTeamGameStreakFinder_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "TeamGameStreakFinder", "headers": ["TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "GAME_ID", "GAME_DATE", "MATCHUP", "WL", "MIN", "PTS", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "PF", "PLUS_MINUS"], "rowSet": [[1, "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/teamgamestreakfinder"
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

	req := TeamGameStreakFinderRequest{}

	resp, err := GetTeamGameStreakFinder(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamGameStreakFinder: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetTeamGameStreakFinder requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "teamgamestreakfinder")
	}

	if len(resp.Data.TeamGameStreakFinder) == 0 {
		t.Errorf("expected TeamGameStreakFinder to be populated from the synthesized fixture, got empty")
	}
}
