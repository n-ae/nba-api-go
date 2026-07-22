package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueDashPlayerClutchV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueDashPlayerClutchV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueDashPlayerClutchV2(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	var vMeasureType *string
	if raw := r.URL.Query().Get("MeasureType"); raw != "" {
		vMeasureType = stringPtr(raw)
	}
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
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
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.LeagueDashPlayerClutchV2Request{
		Season:      vSeason,
		SeasonType:  vSeasonType,
		MeasureType: vMeasureType,
		PerMode:     vPerMode,
		ClutchTime:  vClutchTime,
		AheadBehind: vAheadBehind,
		PointDiff:   vPointDiff,
		LeagueID:    vLeagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerClutchV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
