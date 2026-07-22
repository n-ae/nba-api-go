package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetWinProbabilityPBP_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint WinProbabilityPBP` instead.
func TestGetWinProbabilityPBP_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "WinProbPBP", "headers": ["GAME_ID", "EVENT_NUM", "HOME_PCT", "VISITOR_PCT", "HOME_SCORE", "VISITOR_SCORE", "SCORE_MARGIN", "HOME_PTS_EST", "VISITOR_PTS_EST", "HOME_PTS_RANGE", "VISITOR_PTS_RANGE", "PERIOD", "SECONDS_REMAINING"], "rowSet": [["test", "test", 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1, "test"]]},
		{"name": "GameInfo", "headers": ["GAME_ID", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "HOME_TEAM_ABR", "VISITOR_TEAM_ABR", "HOME_FINAL_SCORE", "VISITOR_FINAL_SCORE"], "rowSet": [["test", 1, 1, "test", "test", "test", "test"]]}
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

	req := WinProbabilityPBPRequest{
		GameID: "1",
	}

	resp, err := GetWinProbabilityPBP(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetWinProbabilityPBP: %v", err)
	}

	if len(resp.Data.WinProbPBP) == 0 {
		t.Errorf("expected WinProbPBP to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.GameInfo) == 0 {
		t.Errorf("expected GameInfo to be populated from the synthesized fixture, got empty")
	}
}
