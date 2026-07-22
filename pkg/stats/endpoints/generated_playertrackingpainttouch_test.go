package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerTrackingPaintTouch_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerTrackingPaintTouch` instead.
func TestGetPlayerTrackingPaintTouch_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerTrackingPaintTouch", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "PAINT_TOUCHES", "PAINT_TOUCH_FGM", "PAINT_TOUCH_FGA", "PAINT_TOUCH_FG_PCT", "PAINT_TOUCH_FTM", "PAINT_TOUCH_FTA", "PAINT_TOUCH_FT_PCT", "PAINT_TOUCH_PTS", "PAINT_TOUCH_PASS", "PAINT_TOUCH_AST", "PAINT_TOUCH_TOV"], "rowSet": [[1, "test", 1, "test", 1, 1.5, "test", 1, 1, 1.5, 1, 1, 1.5, 1.5, "test", 1.5, 1.5]]}
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

	req := PlayerTrackingPaintTouchRequest{}

	resp, err := GetPlayerTrackingPaintTouch(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerTrackingPaintTouch: %v", err)
	}

	if len(resp.Data.PlayerTrackingPaintTouch) == 0 {
		t.Errorf("expected PlayerTrackingPaintTouch to be populated from the synthesized fixture, got empty")
	}
}
