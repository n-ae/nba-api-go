package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerEstimatedMetrics is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerEstimatedMetrics` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerEstimatedMetrics(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.PlayerEstimatedMetricsRequest{
		Season:     vSeason,
		SeasonType: vSeasonType,
		LeagueID:   vLeagueID,
	}

	resp, err := endpoints.GetPlayerEstimatedMetrics(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
