package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleTeamGameLogs is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint TeamGameLogs` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleTeamGameLogs(w http.ResponseWriter, r *http.Request) {
	vSeason := parameters.Season(r.URL.Query().Get("Season"))
	if vSeason == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Season is required")
		return
	}
	vSeasonType := parameters.SeasonType(r.URL.Query().Get("SeasonType"))
	if vSeasonType == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "SeasonType is required")
		return
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vTeamID *string
	if raw := r.URL.Query().Get("TeamID"); raw != "" {
		vTeamID = stringPtr(raw)
	}
	var vDateFrom *string
	if raw := r.URL.Query().Get("DateFrom"); raw != "" {
		vDateFrom = stringPtr(raw)
	}
	var vDateTo *string
	if raw := r.URL.Query().Get("DateTo"); raw != "" {
		vDateTo = stringPtr(raw)
	}

	req := endpoints.TeamGameLogsRequest{
		Season:     vSeason,
		SeasonType: vSeasonType,
		LeagueID:   vLeagueID,
		TeamID:     vTeamID,
		DateFrom:   vDateFrom,
		DateTo:     vDateTo,
	}

	resp, err := endpoints.GetTeamGameLogs(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
