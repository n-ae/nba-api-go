package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerIndex_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerIndex` instead.
func TestGetPlayerIndex_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerIndex", "headers": ["PERSON_ID", "PLAYER_LAST_NAME", "PLAYER_FIRST_NAME", "PLAYER_SLUG", "TEAM_ID", "TEAM_SLUG", "IS_DEFUNCT", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION", "JERSEY_NUMBER", "POSITION", "HEIGHT", "WEIGHT", "COLLEGE", "COUNTRY", "DRAFT_YEAR", "DRAFT_ROUND", "DRAFT_NUMBER", "ROSTER_STATUS", "FROM_YEAR", "TO_YEAR"], "rowSet": [["test", "test", "test", "test", 1, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]}
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

	req := PlayerIndexRequest{}

	resp, err := GetPlayerIndex(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerIndex: %v", err)
	}

	if len(resp.Data.PlayerIndex) == 0 {
		t.Errorf("expected PlayerIndex to be populated from the synthesized fixture, got empty")
	}
}
