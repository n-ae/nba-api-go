package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerAwards_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerAwards` instead.
func TestGetPlayerAwards_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerAwards", "headers": ["PERSON_ID", "FIRST_NAME", "LAST_NAME", "TEAM", "DESCRIPTION", "ALL_NBA_TEAM_NUMBER", "SEASON", "MONTH", "WEEK", "CONFERENCE", "TYPE", "SUBTYPE1", "SUBTYPE2", "SUBTYPE3"], "rowSet": [["test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]}
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

	req := PlayerAwardsRequest{
		PlayerID: "1",
	}

	resp, err := GetPlayerAwards(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerAwards: %v", err)
	}

	if len(resp.Data.PlayerAwards) == 0 {
		t.Errorf("expected PlayerAwards to be populated from the synthesized fixture, got empty")
	}
}
