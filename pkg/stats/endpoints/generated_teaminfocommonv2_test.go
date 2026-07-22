package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamInfoCommonV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamInfoCommonV2` instead.
func TestGetTeamInfoCommonV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "TeamInfoCommon", "headers": ["TEAM_ID", "SEASON_YEAR", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CONFERENCE", "TEAM_DIVISION", "TEAM_CODE", "W", "L", "PCT", "CONF_RANK", "DIV_RANK", "MIN_YEAR", "MAX_YEAR"], "rowSet": [[1, "test", "test", "test", "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1.5, 1.5, "test"]]},
		{"name": "TeamSeasonRanks", "headers": ["LEAGUE_ID", "SEASON_ID", "TEAM_ID", "PTS_RANK", "PTS_PG", "REB_RANK", "REB_PG", "AST_RANK", "AST_PG", "OPP_PTS_RANK", "OPP_PTS_PG"], "rowSet": [["test", "test", 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/teaminfocommonv2"
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

	req := TeamInfoCommonV2Request{
		TeamID: "1",
	}

	resp, err := GetTeamInfoCommonV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamInfoCommonV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetTeamInfoCommonV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "teaminfocommonv2")
	}

	if len(resp.Data.TeamInfoCommon) == 0 {
		t.Errorf("expected TeamInfoCommon to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamSeasonRanks) == 0 {
		t.Errorf("expected TeamSeasonRanks to be populated from the synthesized fixture, got empty")
	}
}
