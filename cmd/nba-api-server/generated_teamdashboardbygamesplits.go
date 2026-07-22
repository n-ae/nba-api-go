package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleTeamDashboardByGameSplits is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint TeamDashboardByGameSplits` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleTeamDashboardByGameSplits(w http.ResponseWriter, r *http.Request) {
	vTeamID := r.URL.Query().Get("TeamID")
	if vTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}
	var vMeasureType *string
	if raw := r.URL.Query().Get("MeasureType"); raw != "" {
		vMeasureType = stringPtr(raw)
	}
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.TeamDashboardByGameSplitsRequest{
		TeamID:      vTeamID,
		MeasureType: vMeasureType,
		PerMode:     vPerMode,
		Season:      vSeason,
		SeasonType:  vSeasonType,
		LeagueID:    vLeagueID,
	}

	resp, err := endpoints.GetTeamDashboardByGameSplits(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
