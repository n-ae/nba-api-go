package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleTeamLineups is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint TeamLineups` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleTeamLineups(w http.ResponseWriter, r *http.Request) {
	vTeamID := r.URL.Query().Get("TeamID")
	if vTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	var vMeasureType *string
	if raw := r.URL.Query().Get("MeasureType"); raw != "" {
		vMeasureType = stringPtr(raw)
	}
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	var vGroupQuantity *string
	if raw := r.URL.Query().Get("GroupQuantity"); raw != "" {
		vGroupQuantity = stringPtr(raw)
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.TeamLineupsRequest{
		TeamID:        vTeamID,
		Season:        vSeason,
		SeasonType:    vSeasonType,
		MeasureType:   vMeasureType,
		PerMode:       vPerMode,
		GroupQuantity: vGroupQuantity,
		LeagueID:      vLeagueID,
	}

	resp, err := endpoints.GetTeamLineups(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
