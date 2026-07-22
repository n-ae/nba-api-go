package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueSeasonMatchups is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueSeasonMatchups` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueSeasonMatchups(w http.ResponseWriter, r *http.Request) {
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

	req := endpoints.LeagueSeasonMatchupsRequest{
		Season:      vSeason,
		SeasonType:  vSeasonType,
		PerMode:     vPerMode,
		LeagueID:    vLeagueID,
		DefPlayerID: vDefPlayerID,
		OffPlayerID: vOffPlayerID,
	}

	resp, err := endpoints.GetLeagueSeasonMatchups(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
