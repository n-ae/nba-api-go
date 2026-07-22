package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueDashTeamBioStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueDashTeamBioStats` instead.
func TestGetLeagueDashTeamBioStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeagueDashTeamBioStats", "headers": ["TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "GP", "W", "L", "W_PCT", "PLUS_MINUS", "OFF_RATING", "DEF_RATING", "NET_RATING", "AST_PCT", "AST_TO", "AST_RATIO", "OREB_PCT", "DREB_PCT", "REB_PCT", "TM_TOV_PCT", "EFG_PCT", "TS_PCT", "PACE", "PIE"], "rowSet": [[1, "test", "test", 1, "test", "test", 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, "test", "test"]]}
	]}`

	const wantPath = "/leaguedashteambiostats"
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

	req := LeagueDashTeamBioStatsRequest{}

	resp, err := GetLeagueDashTeamBioStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueDashTeamBioStats: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetLeagueDashTeamBioStats requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "leaguedashteambiostats")
	}

	if len(resp.Data.LeagueDashTeamBioStats) == 0 {
		t.Errorf("expected LeagueDashTeamBioStats to be populated from the synthesized fixture, got empty")
	}
}
