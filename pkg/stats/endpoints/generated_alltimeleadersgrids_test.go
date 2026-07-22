package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetAllTimeLeadersGrids_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint AllTimeLeadersGrids` instead.
func TestGetAllTimeLeadersGrids_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "AllTimeLeadersPTS", "headers": ["PLAYER_ID", "PLAYER_NAME", "PTS", "PTS_RANK"], "rowSet": [[1, "test", 1.5, 1.5]]},
		{"name": "AllTimeLeadersAST", "headers": ["PLAYER_ID", "PLAYER_NAME", "AST", "AST_RANK"], "rowSet": [[1, "test", 1.5, 1.5]]},
		{"name": "AllTimeLeadersREB", "headers": ["PLAYER_ID", "PLAYER_NAME", "REB", "REB_RANK"], "rowSet": [[1, "test", 1.5, 1.5]]},
		{"name": "AllTimeLeadersBLK", "headers": ["PLAYER_ID", "PLAYER_NAME", "BLK", "BLK_RANK"], "rowSet": [[1, "test", 1.5, 1.5]]},
		{"name": "AllTimeLeadersSTL", "headers": ["PLAYER_ID", "PLAYER_NAME", "STL", "STL_RANK"], "rowSet": [[1, "test", 1.5, 1.5]]}
	]}`

	const wantPath = "/alltimeleadersgrids"
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

	req := AllTimeLeadersGridsRequest{}

	resp, err := GetAllTimeLeadersGrids(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetAllTimeLeadersGrids: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetAllTimeLeadersGrids requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "alltimeleadersgrids")
	}

	if len(resp.Data.AllTimeLeadersPTS) == 0 {
		t.Errorf("expected AllTimeLeadersPTS to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AllTimeLeadersAST) == 0 {
		t.Errorf("expected AllTimeLeadersAST to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AllTimeLeadersREB) == 0 {
		t.Errorf("expected AllTimeLeadersREB to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AllTimeLeadersBLK) == 0 {
		t.Errorf("expected AllTimeLeadersBLK to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AllTimeLeadersSTL) == 0 {
		t.Errorf("expected AllTimeLeadersSTL to be populated from the synthesized fixture, got empty")
	}
}
