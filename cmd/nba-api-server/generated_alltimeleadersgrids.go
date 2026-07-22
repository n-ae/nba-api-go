package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleAllTimeLeadersGrids is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint AllTimeLeadersGrids` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleAllTimeLeadersGrids(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	var vTopX *string
	if raw := r.URL.Query().Get("TopX"); raw != "" {
		vTopX = stringPtr(raw)
	}

	req := endpoints.AllTimeLeadersGridsRequest{
		LeagueID:   vLeagueID,
		PerMode:    vPerMode,
		SeasonType: vSeasonType,
		TopX:       vTopX,
	}

	resp, err := endpoints.GetAllTimeLeadersGrids(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
