package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetFranchiseLeaders_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint FranchiseLeaders` instead.
func TestGetFranchiseLeaders_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "FranchiseLeaders", "headers": ["TEAM_ID", "PTS", "PTS_PERSON_ID", "PTS_PLAYER", "AST", "AST_PERSON_ID", "AST_PLAYER", "REB", "REB_PERSON_ID", "REB_PLAYER", "BLK", "BLK_PERSON_ID", "BLK_PLAYER", "STL", "STL_PERSON_ID", "STL_PLAYER"], "rowSet": [[1, 1.5, "test", 1.5, 1.5, "test", 1.5, 1.5, "test", 1.5, 1.5, "test", 1.5, 1.5, "test", 1.5]]}
	]}`

	const wantPath = "/franchiseleaders"
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

	req := FranchiseLeadersRequest{
		TeamID: "1",
	}

	resp, err := GetFranchiseLeaders(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetFranchiseLeaders: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetFranchiseLeaders requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "franchiseleaders")
	}

	if len(resp.Data.FranchiseLeaders) == 0 {
		t.Errorf("expected FranchiseLeaders to be populated from the synthesized fixture, got empty")
	}
}
