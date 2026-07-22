package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerDashPtShots_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerDashPtShots` instead.
func TestGetPlayerDashPtShots_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "OverallShooting", "headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SORT_ORDER", "GP", "G", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1, "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "GeneralShooting", "headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SHOT_TYPE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "ShotClockShooting", "headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SHOT_CLOCK_RANGE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "DribbleShooting", "headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "DRIBBLE_RANGE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "ClosestDefenderShooting", "headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "CLOSE_DEF_DIST_RANGE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "TouchTimeShooting", "headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "TOUCH_TIME_RANGE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]}
	]}`

	const wantPath = "/playerdashptshots"
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

	req := PlayerDashPtShotsRequest{
		PlayerID: "1",
	}

	resp, err := GetPlayerDashPtShots(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerDashPtShots: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerDashPtShots requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playerdashptshots")
	}

	if len(resp.Data.OverallShooting) == 0 {
		t.Errorf("expected OverallShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.GeneralShooting) == 0 {
		t.Errorf("expected GeneralShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.ShotClockShooting) == 0 {
		t.Errorf("expected ShotClockShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.DribbleShooting) == 0 {
		t.Errorf("expected DribbleShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.ClosestDefenderShooting) == 0 {
		t.Errorf("expected ClosestDefenderShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TouchTimeShooting) == 0 {
		t.Errorf("expected TouchTimeShooting to be populated from the synthesized fixture, got empty")
	}
}
