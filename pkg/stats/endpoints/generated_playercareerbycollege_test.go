package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerCareerByCollege_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerCareerByCollege` instead.
func TestGetPlayerCareerByCollege_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerCareerByCollege", "headers": ["PERSON_ID", "PLAYER_NAME", "SCHOOL_NAME", "FIRST_YEAR", "LAST_YEAR", "SEASONS", "GP", "MIN", "PTS", "REB", "AST"], "rowSet": [["test", "test", "test", "test", 1.5, "test", 1, 1.5, 1.5, 1.5, 1.5]]}
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

	req := PlayerCareerByCollegeRequest{}

	resp, err := GetPlayerCareerByCollege(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerCareerByCollege: %v", err)
	}

	if len(resp.Data.PlayerCareerByCollege) == 0 {
		t.Errorf("expected PlayerCareerByCollege to be populated from the synthesized fixture, got empty")
	}
}
