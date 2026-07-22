package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerTrackingPostTouch_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerTrackingPostTouch` instead.
func TestGetPlayerTrackingPostTouch_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrackingPostTouch", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "POST_TOUCHES", "POST_TOUCH_FGM", "POST_TOUCH_FGA", "POST_TOUCH_FG_PCT", "POST_TOUCH_FTM", "POST_TOUCH_FTA", "POST_TOUCH_FT_PCT", "POST_TOUCH_PTS", "POST_TOUCH_PASS", "POST_TOUCH_AST", "POST_TOUCH_TOV"], "rowSet": [[1, "test", 1, "test", 1, 1.5, "test", 1, 1, 1.5, 1, 1, 1.5, 1.5, "test", 1.5, 1.5]]}
	]}`

	const wantPath = "/playertrackingposttouch"
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

	req := PlayerTrackingPostTouchRequest{}

	resp, err := GetPlayerTrackingPostTouch(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerTrackingPostTouch: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerTrackingPostTouch requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playertrackingposttouch")
	}

	if len(resp.Data.PlayerTrackingPostTouch) == 0 {
		t.Errorf("expected PlayerTrackingPostTouch to be populated from the synthesized fixture, got empty")
	}
}
