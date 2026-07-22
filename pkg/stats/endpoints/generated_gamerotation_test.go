package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetGameRotation_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint GameRotation` instead.
func TestGetGameRotation_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "AwayTeam", "headers": ["GAME_ID", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "PERSON_ID", "PLAYER_FIRST", "PLAYER_LAST", "IN_TIME_REAL", "OUT_TIME_REAL", "PLAYER_PTS", "PT_DIFF", "USG_PCT"], "rowSet": [["test", 1, "test", "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1.5]]},
		{"name": "HomeTeam", "headers": ["GAME_ID", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "PERSON_ID", "PLAYER_FIRST", "PLAYER_LAST", "IN_TIME_REAL", "OUT_TIME_REAL", "PLAYER_PTS", "PT_DIFF", "USG_PCT"], "rowSet": [["test", 1, "test", "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1.5]]}
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

	req := GameRotationRequest{
		GameID: "1",
	}

	resp, err := GetGameRotation(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetGameRotation: %v", err)
	}

	if len(resp.Data.AwayTeam) == 0 {
		t.Errorf("expected AwayTeam to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.HomeTeam) == 0 {
		t.Errorf("expected HomeTeam to be populated from the synthesized fixture, got empty")
	}
}
