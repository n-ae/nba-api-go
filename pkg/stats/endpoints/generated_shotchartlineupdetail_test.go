package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetShotChartLineupDetail_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint ShotChartLineupDetail` instead.
func TestGetShotChartLineupDetail_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "Shot_Chart_Detail", "headers": ["GRID_TYPE", "GAME_ID", "GAME_EVENT_ID", "PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_NAME", "PERIOD", "MINUTES_REMAINING", "SECONDS_REMAINING", "EVENT_TYPE", "ACTION_TYPE", "SHOT_TYPE", "SHOT_ZONE_BASIC", "SHOT_ZONE_AREA", "SHOT_ZONE_RANGE", "SHOT_DISTANCE", "LOC_X", "LOC_Y", "SHOT_ATTEMPTED_FLAG", "SHOT_MADE_FLAG", "GAME_DATE", "HTM", "VTM"], "rowSet": [["test", "test", 1, 1, "test", 1, "test", 1, 1, "test", "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1.5, 1, 1, "test", "test", "test"]]},
		{"name": "LeagueAverages", "headers": ["GRID_TYPE", "SHOT_ZONE_BASIC", "SHOT_ZONE_AREA", "SHOT_ZONE_RANGE", "FGA", "FGM", "FG_PCT"], "rowSet": [["test", "test", "test", "test", 1, 1, 1.5]]}
	]}`

	const wantPath = "/shotchartlineupdetail"
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

	req := ShotChartLineupDetailRequest{}

	resp, err := GetShotChartLineupDetail(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetShotChartLineupDetail: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetShotChartLineupDetail requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "shotchartlineupdetail")
	}

	if len(resp.Data.Shot_Chart_Detail) == 0 {
		t.Errorf("expected Shot_Chart_Detail to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.LeagueAverages) == 0 {
		t.Errorf("expected LeagueAverages to be populated from the synthesized fixture, got empty")
	}
}
