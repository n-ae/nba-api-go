package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerFantasyProfile_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerFantasyProfile` instead.
func TestGetPlayerFantasyProfile_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LastNGames", "headers": ["PLAYER_ID", "PLAYER_NAME", "LAST_N_GAMES", "GP", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "REB", "AST", "TOV", "STL", "BLK", "PTS", "NBA_FANTASY_PTS"], "rowSet": [[1, "test", 1.5, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/playerfantasyprofile"
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

	req := PlayerFantasyProfileRequest{
		PlayerID: "1",
	}

	resp, err := GetPlayerFantasyProfile(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerFantasyProfile: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerFantasyProfile requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playerfantasyprofile")
	}

	if len(resp.Data.LastNGames) == 0 {
		t.Errorf("expected LastNGames to be populated from the synthesized fixture, got empty")
	}
}
