package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueDashPlayerShotLocationV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueDashPlayerShotLocationV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueDashPlayerShotLocationV2(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	var vDistanceRange *string
	if raw := r.URL.Query().Get("DistanceRange"); raw != "" {
		vDistanceRange = stringPtr(raw)
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.LeagueDashPlayerShotLocationV2Request{
		Season:        vSeason,
		SeasonType:    vSeasonType,
		PerMode:       vPerMode,
		DistanceRange: vDistanceRange,
		LeagueID:      vLeagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerShotLocationV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
