package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueDashPlayerShotLocationV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueDashPlayerShotLocationV2` instead.
func TestGetLeagueDashPlayerShotLocationV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "ShotLocationLeague", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "AGE", "GP", "G", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", 1, "test", 1, 1, "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]}
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

	req := LeagueDashPlayerShotLocationV2Request{}

	resp, err := GetLeagueDashPlayerShotLocationV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueDashPlayerShotLocationV2: %v", err)
	}

	if len(resp.Data.ShotLocationLeague) == 0 {
		t.Errorf("expected ShotLocationLeague to be populated from the synthesized fixture, got empty")
	}
}
