package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerDashboardByOpponent is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerDashboardByOpponent` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerDashboardByOpponent(w http.ResponseWriter, r *http.Request) {
	vPlayerID := r.URL.Query().Get("PlayerID")
	if vPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	var vMeasureType *string
	if raw := r.URL.Query().Get("MeasureType"); raw != "" {
		vMeasureType = stringPtr(raw)
	}
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	var vPlusMinus *string
	if raw := r.URL.Query().Get("PlusMinus"); raw != "" {
		vPlusMinus = stringPtr(raw)
	}
	var vPaceAdjust *string
	if raw := r.URL.Query().Get("PaceAdjust"); raw != "" {
		vPaceAdjust = stringPtr(raw)
	}
	var vRank *string
	if raw := r.URL.Query().Get("Rank"); raw != "" {
		vRank = stringPtr(raw)
	}
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.PlayerDashboardByOpponentRequest{
		PlayerID:    vPlayerID,
		MeasureType: vMeasureType,
		PerMode:     vPerMode,
		PlusMinus:   vPlusMinus,
		PaceAdjust:  vPaceAdjust,
		Rank:        vRank,
		Season:      vSeason,
		SeasonType:  vSeasonType,
		LeagueID:    vLeagueID,
	}

	resp, err := endpoints.GetPlayerDashboardByOpponent(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
