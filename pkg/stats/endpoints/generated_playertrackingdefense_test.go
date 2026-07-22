package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerTrackingDefense_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerTrackingDefense` instead.
func TestGetPlayerTrackingDefense_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrackingDefense", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "DEF_RIM_FGM", "DEF_RIM_FGA", "DEF_RIM_FG_PCT", "DREB"], "rowSet": [[1, "test", 1, "test", 1, 1.5, 1, 1, 1.5, 1.5]]}
	]}`

	const wantPath = "/playertrackingdefense"
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

	req := PlayerTrackingDefenseRequest{}

	resp, err := GetPlayerTrackingDefense(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerTrackingDefense: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerTrackingDefense requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playertrackingdefense")
	}

	if len(resp.Data.PlayerTrackingDefense) == 0 {
		t.Errorf("expected PlayerTrackingDefense to be populated from the synthesized fixture, got empty")
	}
}
