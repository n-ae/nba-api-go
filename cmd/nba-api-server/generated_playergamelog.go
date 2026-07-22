package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerGameLog is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerGameLog` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerGameLog(w http.ResponseWriter, r *http.Request) {
	vPlayerID := r.URL.Query().Get("PlayerID")
	if vPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	vSeason := parameters.Season(getSeasonOrDefault(r))
	vSeasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	vDateFrom := r.URL.Query().Get("DateFrom")
	vDateTo := r.URL.Query().Get("DateTo")
	vLeagueID := parameters.LeagueIDNBA

	req := endpoints.PlayerGameLogRequest{
		PlayerID:   vPlayerID,
		Season:     vSeason,
		SeasonType: vSeasonType,
		DateFrom:   vDateFrom,
		DateTo:     vDateTo,
		LeagueID:   vLeagueID,
	}

	resp, err := endpoints.PlayerGameLog(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
