package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetDraftBoard_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint DraftBoard` instead.
func TestGetDraftBoard_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "DraftBoard", "headers": ["PERSON_ID", "PLAYER_NAME", "SEASON", "ROUND_NUMBER", "ROUND_PICK", "OVERALL_PICK", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION"], "rowSet": [["test", "test", "test", "test", "test", "test", 1, "test", "test", "test"]]}
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

	req := DraftBoardRequest{}

	resp, err := GetDraftBoard(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetDraftBoard: %v", err)
	}

	if len(resp.Data.DraftBoard) == 0 {
		t.Errorf("expected DraftBoard to be populated from the synthesized fixture, got empty")
	}
}
