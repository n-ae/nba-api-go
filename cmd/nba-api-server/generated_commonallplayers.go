package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleCommonAllPlayers is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint CommonAllPlayers` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleCommonAllPlayers(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeason := parameters.Season(r.URL.Query().Get("Season"))
	if vSeason == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Season is required")
		return
	}
	var vIsOnlyCurrentSeason *string
	if raw := r.URL.Query().Get("IsOnlyCurrentSeason"); raw != "" {
		vIsOnlyCurrentSeason = stringPtr(raw)
	}

	req := endpoints.CommonAllPlayersRequest{
		LeagueID:            vLeagueID,
		Season:              vSeason,
		IsOnlyCurrentSeason: vIsOnlyCurrentSeason,
	}

	resp, err := endpoints.GetCommonAllPlayers(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
