package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamEstimatedMetrics_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamEstimatedMetrics` instead.
func TestGetTeamEstimatedMetrics_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "TeamEstimatedMetrics", "headers": ["TEAM_ID", "TEAM_NAME", "GP", "W", "L", "W_PCT", "MIN", "E_OFF_RATING", "E_DEF_RATING", "E_NET_RATING", "E_PACE", "E_AST_RATIO", "E_OREB_PCT", "E_DREB_PCT", "E_REB_PCT", "E_TOV_PCT", "GP_RANK", "W_RANK", "L_RANK", "W_PCT_RANK", "MIN_RANK", "E_OFF_RATING_RANK", "E_DEF_RATING_RANK", "E_NET_RATING_RANK", "E_AST_RATIO_RANK", "E_OREB_PCT_RANK", "E_DREB_PCT_RANK", "E_REB_PCT_RANK", "E_TOV_PCT_RANK", "E_PACE_RANK"], "rowSet": [[1, "test", 1, "test", "test", 1.5, 1.5, "test", "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/teamestimatedmetrics"
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

	req := TeamEstimatedMetricsRequest{}

	resp, err := GetTeamEstimatedMetrics(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamEstimatedMetrics: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetTeamEstimatedMetrics requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "teamestimatedmetrics")
	}

	if len(resp.Data.TeamEstimatedMetrics) == 0 {
		t.Errorf("expected TeamEstimatedMetrics to be populated from the synthesized fixture, got empty")
	}
}
