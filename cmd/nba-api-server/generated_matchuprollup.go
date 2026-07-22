package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleMatchupRollup is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint MatchupRollup` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleMatchupRollup(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vDefPlayerID *string
	if raw := r.URL.Query().Get("DefPlayerID"); raw != "" {
		vDefPlayerID = stringPtr(raw)
	}
	var vOffPlayerID *string
	if raw := r.URL.Query().Get("OffPlayerID"); raw != "" {
		vOffPlayerID = stringPtr(raw)
	}

	req := endpoints.MatchupRollupRequest{
		Season:      vSeason,
		SeasonType:  vSeasonType,
		PerMode:     vPerMode,
		LeagueID:    vLeagueID,
		DefPlayerID: vDefPlayerID,
		OffPlayerID: vOffPlayerID,
	}

	resp, err := endpoints.GetMatchupRollup(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
