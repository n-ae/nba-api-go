package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetDraftCombineStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint DraftCombineStats` instead.
func TestGetDraftCombineStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "DraftCombineStats", "headers": ["SEASON", "PLAYER_ID", "FIRST_NAME", "LAST_NAME", "PLAYER_NAME", "POSITION", "HEIGHT_WO_SHOES", "HEIGHT_WO_SHOES_FT_IN", "HEIGHT_W_SHOES", "HEIGHT_W_SHOES_FT_IN", "WEIGHT", "WINGSPAN", "WINGSPAN_FT_IN", "STANDING_REACH", "STANDING_REACH_FT_IN", "BODY_FAT_PCT", "HAND_LENGTH", "HAND_WIDTH", "STANDING_VERTICAL_LEAP", "MAX_VERTICAL_LEAP", "LANE_AGILITY_TIME", "MODIFIED_LANE_AGILITY_TIME", "THREE_QUARTER_SPRINT", "BENCH_PRESS"], "rowSet": [["test", 1, "test", "test", "test", "test", "test", "test", "test", "test", "test", 1.5, 1.5, "test", "test", 1.5, "test", "test", "test", "test", "test", "test", "test", "test"]]}
	]}`

	const wantPath = "/draftcombinestats"
	var gotPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}

	req := DraftCombineStatsRequest{}

	resp, err := GetDraftCombineStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetDraftCombineStats: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetDraftCombineStats requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "draftcombinestats")
	}

	if len(resp.Data.DraftCombineStats) == 0 {
		t.Errorf("expected DraftCombineStats to be populated from the synthesized fixture, got empty")
	}
}
