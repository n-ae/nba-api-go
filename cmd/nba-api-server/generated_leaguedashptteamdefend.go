package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueDashPtTeamDefend is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueDashPtTeamDefend` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueDashPtTeamDefend(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vDefenseCategory *string
	if raw := r.URL.Query().Get("DefenseCategory"); raw != "" {
		vDefenseCategory = stringPtr(raw)
	}

	req := endpoints.LeagueDashPtTeamDefendRequest{
		Season:          vSeason,
		SeasonType:      vSeasonType,
		PerMode:         vPerMode,
		LeagueID:        vLeagueID,
		DefenseCategory: vDefenseCategory,
	}

	resp, err := endpoints.GetLeagueDashPtTeamDefend(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
