package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleScoreboardV3 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint ScoreboardV3` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleScoreboardV3(w http.ResponseWriter, r *http.Request) {
	vGameDate := r.URL.Query().Get("GameDate")
	if vGameDate == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameDate is required")
		return
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.ScoreboardV3Request{
		GameDate: vGameDate,
		LeagueID: vLeagueID,
	}

	resp, err := endpoints.GetScoreboardV3(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
