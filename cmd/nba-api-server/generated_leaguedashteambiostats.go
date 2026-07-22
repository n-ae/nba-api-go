package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueDashTeamBioStats is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueDashTeamBioStats` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueDashTeamBioStats(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.LeagueDashTeamBioStatsRequest{
		Season:     vSeason,
		SeasonType: vSeasonType,
		PerMode:    vPerMode,
		LeagueID:   vLeagueID,
	}

	resp, err := endpoints.GetLeagueDashTeamBioStats(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
