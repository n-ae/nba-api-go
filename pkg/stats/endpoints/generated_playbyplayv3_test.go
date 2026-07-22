package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayByPlayV3_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayByPlayV3` instead.
func TestGetPlayByPlayV3_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayByPlay", "headers": ["gameId", "actionNumber", "clock", "timeActual", "period", "periodType", "teamId", "teamTricode", "actionType", "subType", "descriptor", "qualifiers", "personId", "playerName", "playerNameI", "jerseyNum", "assistPersonId", "assistPlayerNameI", "assistTotal", "officialId", "description", "shotDistance", "shotResult", "isFieldGoal", "scoreHome", "scoreAway", "pointsTotal", "location", "xLegacy", "yLegacy", "isTargetScoreLastPeriod", "orderNumber", "edited"], "rowSet": [["test", "test", "test", "test", 1, 1, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", 1.5, "test", "test"]]}
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

	req := PlayByPlayV3Request{
		GameID: "1",
	}

	resp, err := GetPlayByPlayV3(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayByPlayV3: %v", err)
	}

	if len(resp.Data.PlayByPlay) == 0 {
		t.Errorf("expected PlayByPlay to be populated from the synthesized fixture, got empty")
	}
}
