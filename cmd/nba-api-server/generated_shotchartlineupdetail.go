package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleShotChartLineupDetail is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint ShotChartLineupDetail` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleShotChartLineupDetail(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	var vTeamID *string
	if raw := r.URL.Query().Get("TeamID"); raw != "" {
		vTeamID = stringPtr(raw)
	}
	var vGroupID *string
	if raw := r.URL.Query().Get("GroupID"); raw != "" {
		vGroupID = stringPtr(raw)
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vContextMeasure *string
	if raw := r.URL.Query().Get("ContextMeasure"); raw != "" {
		vContextMeasure = stringPtr(raw)
	}

	req := endpoints.ShotChartLineupDetailRequest{
		Season:         vSeason,
		SeasonType:     vSeasonType,
		TeamID:         vTeamID,
		GroupID:        vGroupID,
		LeagueID:       vLeagueID,
		ContextMeasure: vContextMeasure,
	}

	resp, err := endpoints.GetShotChartLineupDetail(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
