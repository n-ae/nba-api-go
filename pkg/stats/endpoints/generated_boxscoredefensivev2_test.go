package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScoreDefensiveV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScoreDefensiveV2` instead.
func TestGetBoxScoreDefensiveV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PLAYER_ID", "PLAYER_NAME", "NICKNAME", "START_POSITION", "COMMENT", "MIN", "DEF_RIM_FGM", "DEF_RIM_FGA", "DEF_RIM_FG_PCT"], "rowSet": [["test", 1, "test", "test", 1, "test", "test", "test", "test", 1.5, 1, 1, 1.5]]},
		{"name": "TeamStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CITY", "MIN", "DEF_RIM_FGM", "DEF_RIM_FGA", "DEF_RIM_FG_PCT"], "rowSet": [["test", 1, "test", "test", "test", 1.5, 1, 1, 1.5]]}
	]}`

	const wantPath = "/boxscoredefensivev2"
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

	req := BoxScoreDefensiveV2Request{
		GameID: "1",
	}

	resp, err := GetBoxScoreDefensiveV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScoreDefensiveV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetBoxScoreDefensiveV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "boxscoredefensivev2")
	}

	if len(resp.Data.PlayerStats) == 0 {
		t.Errorf("expected PlayerStats to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamStats) == 0 {
		t.Errorf("expected TeamStats to be populated from the synthesized fixture, got empty")
	}
}
