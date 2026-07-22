package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueDashPlayerBioStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueDashPlayerBioStats` instead.
func TestGetLeagueDashPlayerBioStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeagueDashPlayerBioStats", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "AGE", "PLAYER_HEIGHT", "PLAYER_WEIGHT", "COLLEGE", "COUNTRY", "DRAFT_YEAR", "DRAFT_ROUND", "DRAFT_NUMBER", "GP", "PTS", "REB", "AST", "NET_RATING", "OREB_PCT", "DREB_PCT", "USG_PCT", "TS_PCT", "AST_PCT"], "rowSet": [[1, "test", 1, "test", 1, "test", "test", "test", "test", "test", "test", "test", 1, 1.5, 1.5, 1.5, "test", 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/leaguedashplayerbiostats"
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

	req := LeagueDashPlayerBioStatsRequest{}

	resp, err := GetLeagueDashPlayerBioStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueDashPlayerBioStats: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetLeagueDashPlayerBioStats requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "leaguedashplayerbiostats")
	}

	if len(resp.Data.LeagueDashPlayerBioStats) == 0 {
		t.Errorf("expected LeagueDashPlayerBioStats to be populated from the synthesized fixture, got empty")
	}
}
