package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueDashPtStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueDashPtStats` instead.
func TestGetLeagueDashPtStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeagueDashPTStats", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "DIST_FEET", "DIST_MILES", "DIST_MILES_OFF", "DIST_MILES_DEF", "AVG_SPEED", "AVG_SPEED_OFF", "AVG_SPEED_DEF"], "rowSet": [[1, "test", 1, "test", 1, 1.5, "test", "test", "test", "test", "test", "test", "test"]]}
	]}`

	const wantPath = "/leaguedashptstats"
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

	req := LeagueDashPtStatsRequest{}

	resp, err := GetLeagueDashPtStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueDashPtStats: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetLeagueDashPtStats requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "leaguedashptstats")
	}

	if len(resp.Data.LeagueDashPTStats) == 0 {
		t.Errorf("expected LeagueDashPTStats to be populated from the synthesized fixture, got empty")
	}
}
