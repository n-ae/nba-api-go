package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetCommonTeamRosterV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint CommonTeamRosterV2` instead.
func TestGetCommonTeamRosterV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "CommonTeamRoster", "headers": ["TeamID", "SEASON", "LeagueID", "PLAYER", "NICKNAME", "PLAYER_SLUG", "NUM", "POSITION", "HEIGHT", "WEIGHT", "BIRTH_DATE", "AGE", "EXP", "SCHOOL", "PLAYER_ID", "HOW_ACQUIRED"], "rowSet": [["test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", 1, "test", "test", 1, "test"]]},
		{"name": "Coaches", "headers": ["TEAM_ID", "SEASON", "COACH_ID", "FIRST_NAME", "LAST_NAME", "COACH_NAME", "COACH_CODE", "IS_ASSISTANT", "COACH_TYPE", "SCHOOL", "SORT_SEQUENCE"], "rowSet": [[1, "test", "test", "test", "test", "test", "test", "test", "test", "test", 1]]}
	]}`

	const wantPath = "/commonteamrosterv2"
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

	req := CommonTeamRosterV2Request{
		TeamID: "1",
	}

	resp, err := GetCommonTeamRosterV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetCommonTeamRosterV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetCommonTeamRosterV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "commonteamrosterv2")
	}

	if len(resp.Data.CommonTeamRoster) == 0 {
		t.Errorf("expected CommonTeamRoster to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.Coaches) == 0 {
		t.Errorf("expected Coaches to be populated from the synthesized fixture, got empty")
	}
}
