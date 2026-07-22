package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerTrackingCatchShoot_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerTrackingCatchShoot` instead.
func TestGetPlayerTrackingCatchShoot_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrackingCatchShoot", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "CATCH_SHOOT_FGM", "CATCH_SHOOT_FGA", "CATCH_SHOOT_FG_PCT", "CATCH_SHOOT_PTS", "CATCH_SHOOT_FG3M", "CATCH_SHOOT_FG3A", "CATCH_SHOOT_FG3_PCT", "CATCH_SHOOT_EFG_PCT"], "rowSet": [[1, "test", 1, "test", 1, 1.5, 1, 1, 1.5, 1.5, 1, 1, 1.5, 1.5]]}
	]}`

	const wantPath = "/playertrackingcatchshoot"
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

	req := PlayerTrackingCatchShootRequest{}

	resp, err := GetPlayerTrackingCatchShoot(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerTrackingCatchShoot: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerTrackingCatchShoot requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playertrackingcatchshoot")
	}

	if len(resp.Data.PlayerTrackingCatchShoot) == 0 {
		t.Errorf("expected PlayerTrackingCatchShoot to be populated from the synthesized fixture, got empty")
	}
}
