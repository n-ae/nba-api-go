package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleTeamGameLog is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint TeamGameLog` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleTeamGameLog(w http.ResponseWriter, r *http.Request) {
	vTeamID := r.URL.Query().Get("TeamID")
	if vTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}
	vSeason := parameters.Season(getSeasonOrDefault(r))
	vSeasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	vDateFrom := r.URL.Query().Get("DateFrom")
	vDateTo := r.URL.Query().Get("DateTo")

	req := endpoints.TeamGameLogRequest{
		TeamID:     vTeamID,
		Season:     vSeason,
		SeasonType: vSeasonType,
		DateFrom:   vDateFrom,
		DateTo:     vDateTo,
	}

	resp, err := endpoints.GetTeamGameLog(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
