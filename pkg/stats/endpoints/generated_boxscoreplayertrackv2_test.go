package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScorePlayerTrackV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScorePlayerTrackV2` instead.
func TestGetBoxScorePlayerTrackV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrack", "headers": ["GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PLAYER_ID", "PLAYER_NAME", "START_POSITION", "COMMENT", "MIN", "SPD", "DIST", "ORBC", "DRBC", "RBC", "TCHS", "SAST", "FTAST", "PASS", "AST", "CFGM", "CFGA", "CFG_PCT", "UFGM", "UFGA", "UFG_PCT", "FG_PCT", "DFGM", "DFGA", "DFG_PCT"], "rowSet": [["test", 1, "test", "test", 1, "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test", 1.5, 1.5, "test", 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1, 1, 1.5]]}
	]}`

	const wantPath = "/boxscoreplayertrackv2"
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

	req := BoxScorePlayerTrackV2Request{
		GameID: "1",
	}

	resp, err := GetBoxScorePlayerTrackV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScorePlayerTrackV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetBoxScorePlayerTrackV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "boxscoreplayertrackv2")
	}

	if len(resp.Data.PlayerTrack) == 0 {
		t.Errorf("expected PlayerTrack to be populated from the synthesized fixture, got empty")
	}
}
