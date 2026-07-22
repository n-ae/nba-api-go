package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleTeamAndPlayersVsPlayers is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint TeamAndPlayersVsPlayers` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleTeamAndPlayersVsPlayers(w http.ResponseWriter, r *http.Request) {
	vTeamID := r.URL.Query().Get("TeamID")
	if vTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}
	vVsPlayerID := r.URL.Query().Get("VsPlayerID")
	if vVsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "VsPlayerID is required")
		return
	}
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.TeamAndPlayersVsPlayersRequest{
		TeamID:     vTeamID,
		VsPlayerID: vVsPlayerID,
		Season:     vSeason,
		SeasonType: vSeasonType,
		PerMode:    vPerMode,
		LeagueID:   vLeagueID,
	}

	resp, err := endpoints.GetTeamAndPlayersVsPlayers(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
