package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerIndex is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerIndex` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerIndex(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	var vTeamID *string
	if raw := r.URL.Query().Get("TeamID"); raw != "" {
		vTeamID = stringPtr(raw)
	}
	var vHistorical *string
	if raw := r.URL.Query().Get("Historical"); raw != "" {
		vHistorical = stringPtr(raw)
	}

	req := endpoints.PlayerIndexRequest{
		LeagueID:   vLeagueID,
		Season:     vSeason,
		TeamID:     vTeamID,
		Historical: vHistorical,
	}

	resp, err := endpoints.GetPlayerIndex(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
