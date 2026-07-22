package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleCommonPlayoffSeries is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint CommonPlayoffSeries` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleCommonPlayoffSeries(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeason := parameters.Season(r.URL.Query().Get("Season"))
	if vSeason == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Season is required")
		return
	}
	var vSeriesID *string
	if raw := r.URL.Query().Get("SeriesID"); raw != "" {
		vSeriesID = stringPtr(raw)
	}

	req := endpoints.CommonPlayoffSeriesRequest{
		LeagueID: vLeagueID,
		Season:   vSeason,
		SeriesID: vSeriesID,
	}

	resp, err := endpoints.GetCommonPlayoffSeries(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
