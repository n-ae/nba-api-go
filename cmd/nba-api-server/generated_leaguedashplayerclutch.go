package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueDashPlayerClutch is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueDashPlayerClutch` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueDashPlayerClutch(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vClutchTime *string
	if raw := r.URL.Query().Get("ClutchTime"); raw != "" {
		vClutchTime = stringPtr(raw)
	}
	var vAheadBehind *string
	if raw := r.URL.Query().Get("AheadBehind"); raw != "" {
		vAheadBehind = stringPtr(raw)
	}
	var vPointDiff *string
	if raw := r.URL.Query().Get("PointDiff"); raw != "" {
		vPointDiff = stringPtr(raw)
	}

	req := endpoints.LeagueDashPlayerClutchRequest{
		Season:      vSeason,
		SeasonType:  vSeasonType,
		PerMode:     vPerMode,
		LeagueID:    vLeagueID,
		ClutchTime:  vClutchTime,
		AheadBehind: vAheadBehind,
		PointDiff:   vPointDiff,
	}

	resp, err := endpoints.GetLeagueDashPlayerClutch(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
