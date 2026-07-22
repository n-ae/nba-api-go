package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerCareerStats is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerCareerStats` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerCareerStats(w http.ResponseWriter, r *http.Request) {
	vPlayerID := r.URL.Query().Get("PlayerID")
	if vPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	vPerMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	vLeagueID := parameters.LeagueIDNBA

	req := endpoints.PlayerCareerStatsRequest{
		PlayerID: vPlayerID,
		PerMode:  vPerMode,
		LeagueID: vLeagueID,
	}

	resp, err := endpoints.PlayerCareerStats(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
