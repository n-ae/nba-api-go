package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerEstimatedMetrics_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerEstimatedMetrics` instead.
func TestGetPlayerEstimatedMetrics_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerEstimatedMetrics", "headers": ["PLAYER_ID", "PLAYER_NAME", "GP", "W", "L", "W_PCT", "MIN", "E_OFF_RATING", "E_DEF_RATING", "E_NET_RATING", "E_AST_RATIO", "E_OREB_PCT", "E_DREB_PCT", "E_REB_PCT", "E_TOV_PCT", "E_USG_PCT", "E_PACE", "GP_RANK", "W_RANK", "L_RANK", "W_PCT_RANK", "MIN_RANK", "E_OFF_RATING_RANK", "E_DEF_RATING_RANK", "E_NET_RATING_RANK", "E_AST_RATIO_RANK", "E_OREB_PCT_RANK", "E_DREB_PCT_RANK", "E_REB_PCT_RANK", "E_TOV_PCT_RANK", "E_USG_PCT_RANK", "E_PACE_RANK"], "rowSet": [[1, "test", 1, "test", "test", 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
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

	req := PlayerEstimatedMetricsRequest{}

	resp, err := GetPlayerEstimatedMetrics(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerEstimatedMetrics: %v", err)
	}

	if len(resp.Data.PlayerEstimatedMetrics) == 0 {
		t.Errorf("expected PlayerEstimatedMetrics to be populated from the synthesized fixture, got empty")
	}
}
