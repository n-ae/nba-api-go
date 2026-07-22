package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerTrackingElbowTouch_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerTrackingElbowTouch` instead.
func TestGetPlayerTrackingElbowTouch_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrackingElbowTouch", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "ELBOW_TOUCHES", "ELBOW_TOUCH_FGM", "ELBOW_TOUCH_FGA", "ELBOW_TOUCH_FG_PCT", "ELBOW_TOUCH_FTM", "ELBOW_TOUCH_FTA", "ELBOW_TOUCH_FT_PCT", "ELBOW_TOUCH_PTS", "ELBOW_TOUCH_PASS", "ELBOW_TOUCH_AST", "ELBOW_TOUCH_TOV"], "rowSet": [[1, "test", 1, "test", 1, 1.5, "test", 1, 1, 1.5, 1, 1, 1.5, 1.5, "test", 1.5, 1.5]]}
	]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}

	req := PlayerTrackingElbowTouchRequest{}

	resp, err := GetPlayerTrackingElbowTouch(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerTrackingElbowTouch: %v", err)
	}

	if len(resp.Data.PlayerTrackingElbowTouch) == 0 {
		t.Errorf("expected PlayerTrackingElbowTouch to be populated from the synthesized fixture, got empty")
	}
}
