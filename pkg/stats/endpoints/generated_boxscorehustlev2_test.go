package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScoreHustleV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScoreHustleV2` instead.
func TestGetBoxScoreHustleV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PLAYER_ID", "PLAYER_NAME", "START_POSITION", "COMMENT", "MIN", "CONTESTED_SHOTS", "CONTESTED_SHOTS_2PT", "CONTESTED_SHOTS_3PT", "DEFLECTIONS", "CHARGES_DRAWN", "SCREEN_ASSISTS", "SCREEN_AST_PTS", "OFF_LOOSE_BALLS_RECOVERED", "DEF_LOOSE_BALLS_RECOVERED", "LOOSE_BALLS_RECOVERED", "OFF_BOXOUTS", "DEF_BOXOUTS", "BOX_OUTS"], "rowSet": [["test", 1, "test", "test", 1, "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test"]]},
		{"name": "TeamStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CITY", "MIN", "CONTESTED_SHOTS", "CONTESTED_SHOTS_2PT", "CONTESTED_SHOTS_3PT", "DEFLECTIONS", "CHARGES_DRAWN", "SCREEN_ASSISTS", "SCREEN_AST_PTS", "OFF_LOOSE_BALLS_RECOVERED", "DEF_LOOSE_BALLS_RECOVERED", "LOOSE_BALLS_RECOVERED", "OFF_BOXOUTS", "DEF_BOXOUTS", "BOX_OUTS"], "rowSet": [["test", 1, "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test"]]}
	]}`

	const wantPath = "/boxscorehustlev2"
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

	req := BoxScoreHustleV2Request{
		GameID: "1",
	}

	resp, err := GetBoxScoreHustleV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScoreHustleV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetBoxScoreHustleV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "boxscorehustlev2")
	}

	if len(resp.Data.PlayerStats) == 0 {
		t.Errorf("expected PlayerStats to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamStats) == 0 {
		t.Errorf("expected TeamStats to be populated from the synthesized fixture, got empty")
	}
}
