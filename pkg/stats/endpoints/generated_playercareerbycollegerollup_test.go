package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerCareerByCollegeRollup_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerCareerByCollegeRollup` instead.
func TestGetPlayerCareerByCollegeRollup_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "CollegeStats", "headers": ["SCHOOL_NAME", "SEASON_COUNT", "PLAYER_COUNT", "ACTIVE_PLAYER_COUNT", "MIN", "PTS", "REB", "AST", "FG_PCT", "FG3_PCT", "FT_PCT"], "rowSet": [["test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/playercareerbycollegerollup"
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

	req := PlayerCareerByCollegeRollupRequest{}

	resp, err := GetPlayerCareerByCollegeRollup(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerCareerByCollegeRollup: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerCareerByCollegeRollup requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playercareerbycollegerollup")
	}

	if len(resp.Data.CollegeStats) == 0 {
		t.Errorf("expected CollegeStats to be populated from the synthesized fixture, got empty")
	}
}
