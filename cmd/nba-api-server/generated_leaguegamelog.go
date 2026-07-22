package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueGameLog is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueGameLog` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueGameLog(w http.ResponseWriter, r *http.Request) {
	vSeason := parameters.Season(r.URL.Query().Get("Season"))
	if vSeason == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Season is required")
		return
	}
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vPlayerOrTeam *string
	if raw := r.URL.Query().Get("PlayerOrTeam"); raw != "" {
		vPlayerOrTeam = stringPtr(raw)
	}
	var vCounter *string
	if raw := r.URL.Query().Get("Counter"); raw != "" {
		vCounter = stringPtr(raw)
	}
	var vSorter *string
	if raw := r.URL.Query().Get("Sorter"); raw != "" {
		vSorter = stringPtr(raw)
	}
	var vDirection *string
	if raw := r.URL.Query().Get("Direction"); raw != "" {
		vDirection = stringPtr(raw)
	}
	var vDateFrom *string
	if raw := r.URL.Query().Get("DateFrom"); raw != "" {
		vDateFrom = stringPtr(raw)
	}
	var vDateTo *string
	if raw := r.URL.Query().Get("DateTo"); raw != "" {
		vDateTo = stringPtr(raw)
	}

	req := endpoints.LeagueGameLogRequest{
		Season:       vSeason,
		SeasonType:   vSeasonType,
		LeagueID:     vLeagueID,
		PlayerOrTeam: vPlayerOrTeam,
		Counter:      vCounter,
		Sorter:       vSorter,
		Direction:    vDirection,
		DateFrom:     vDateFrom,
		DateTo:       vDateTo,
	}

	resp, err := endpoints.GetLeagueGameLog(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
