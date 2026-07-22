package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamHistoricalLeaders_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamHistoricalLeaders` instead.
func TestGetTeamHistoricalLeaders_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "CareerLeadersPTS", "headers": ["TEAM_ID", "PLAYER_ID", "PLAYER", "PTS", "PTS_RANK"], "rowSet": [[1, 1, "test", 1.5, 1.5]]},
		{"name": "CareerLeadersAST", "headers": ["TEAM_ID", "PLAYER_ID", "PLAYER", "AST", "AST_RANK"], "rowSet": [[1, 1, "test", 1.5, 1.5]]},
		{"name": "CareerLeadersREB", "headers": ["TEAM_ID", "PLAYER_ID", "PLAYER", "REB", "REB_RANK"], "rowSet": [[1, 1, "test", 1.5, 1.5]]},
		{"name": "CareerLeadersBLK", "headers": ["TEAM_ID", "PLAYER_ID", "PLAYER", "BLK", "BLK_RANK"], "rowSet": [[1, 1, "test", 1.5, 1.5]]},
		{"name": "CareerLeadersSTL", "headers": ["TEAM_ID", "PLAYER_ID", "PLAYER", "STL", "STL_RANK"], "rowSet": [[1, 1, "test", 1.5, 1.5]]}
	]}`

	const wantPath = "/teamhistoricalleaders"
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

	req := TeamHistoricalLeadersRequest{
		TeamID: "1",
	}

	resp, err := GetTeamHistoricalLeaders(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamHistoricalLeaders: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetTeamHistoricalLeaders requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "teamhistoricalleaders")
	}

	if len(resp.Data.CareerLeadersPTS) == 0 {
		t.Errorf("expected CareerLeadersPTS to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.CareerLeadersAST) == 0 {
		t.Errorf("expected CareerLeadersAST to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.CareerLeadersREB) == 0 {
		t.Errorf("expected CareerLeadersREB to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.CareerLeadersBLK) == 0 {
		t.Errorf("expected CareerLeadersBLK to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.CareerLeadersSTL) == 0 {
		t.Errorf("expected CareerLeadersSTL to be populated from the synthesized fixture, got empty")
	}
}
