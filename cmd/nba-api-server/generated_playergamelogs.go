package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerGameLogs is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerGameLogs` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerGameLogs(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vDateFrom *string
	if raw := r.URL.Query().Get("DateFrom"); raw != "" {
		vDateFrom = stringPtr(raw)
	}
	var vDateTo *string
	if raw := r.URL.Query().Get("DateTo"); raw != "" {
		vDateTo = stringPtr(raw)
	}

	req := endpoints.PlayerGameLogsRequest{
		Season:     vSeason,
		SeasonType: vSeasonType,
		LeagueID:   vLeagueID,
		DateFrom:   vDateFrom,
		DateTo:     vDateTo,
	}

	resp, err := endpoints.GetPlayerGameLogs(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
