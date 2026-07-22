package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamPlayerOnOffDetails_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamPlayerOnOffDetails` instead.
func TestGetTeamPlayerOnOffDetails_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "OverallOnOffSummary", "headers": ["TEAM_ID", "TEAM_NAME", "VS_PLAYER_ID", "VS_PLAYER_NAME", "COURT_STATUS", "GP", "MIN", "PLUS_MINUS"], "rowSet": [[1, "test", 1, "test", "test", 1, 1.5, 1.5]]},
		{"name": "PlayersOnCourtTeamPlayerOnOffDetails", "headers": ["TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "VS_PLAYER_ID", "VS_PLAYER_NAME", "COURT_STATUS", "PLAYER_ID", "PLAYER_NAME", "GP", "MIN", "PLUS_MINUS"], "rowSet": [[1, "test", "test", 1, "test", "test", 1, "test", 1, 1.5, 1.5]]}
	]}`

	const wantPath = "/teamplayeronoffdetails"
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

	req := TeamPlayerOnOffDetailsRequest{
		TeamID: "1",
	}

	resp, err := GetTeamPlayerOnOffDetails(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamPlayerOnOffDetails: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetTeamPlayerOnOffDetails requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "teamplayeronoffdetails")
	}

	if len(resp.Data.OverallOnOffSummary) == 0 {
		t.Errorf("expected OverallOnOffSummary to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.PlayersOnCourtTeamPlayerOnOffDetails) == 0 {
		t.Errorf("expected PlayersOnCourtTeamPlayerOnOffDetails to be populated from the synthesized fixture, got empty")
	}
}
