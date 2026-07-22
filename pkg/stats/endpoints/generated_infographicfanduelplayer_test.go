package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetInfographicFanDuelPlayer_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint InfographicFanDuelPlayer` instead.
func TestGetInfographicFanDuelPlayer_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "FanDuelPlayer", "headers": ["PLAYER_ID", "PLAYER_NAME", "FD_POSITION", "FD_SALARY", "FD_MINUTES", "FD_FG_PCT", "FD_FT_PCT", "FD_FG3_PCT", "FD_PTS", "FD_REB", "FD_AST", "FD_STL", "FD_BLK", "FD_TOV"], "rowSet": [[1, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
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

	req := InfographicFanDuelPlayerRequest{
		PlayerID: "1",
	}

	resp, err := GetInfographicFanDuelPlayer(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetInfographicFanDuelPlayer: %v", err)
	}

	if len(resp.Data.FanDuelPlayer) == 0 {
		t.Errorf("expected FanDuelPlayer to be populated from the synthesized fixture, got empty")
	}
}
