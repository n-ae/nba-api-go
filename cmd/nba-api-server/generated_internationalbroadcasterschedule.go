package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleInternationalBroadcasterSchedule is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint InternationalBroadcasterSchedule` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleInternationalBroadcasterSchedule(w http.ResponseWriter, r *http.Request) {
	vLeagueID := parameters.LeagueIDNBA
	vSeason := r.URL.Query().Get("Season")
	if vSeason == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Season is required")
		return
	}
	var vRegionID *string
	if raw := r.URL.Query().Get("RegionID"); raw != "" {
		vRegionID = stringPtr(raw)
	}
	var vDate *string
	if raw := r.URL.Query().Get("Date"); raw != "" {
		vDate = stringPtr(raw)
	}
	var vEST *string
	if raw := r.URL.Query().Get("EST"); raw != "" {
		vEST = stringPtr(raw)
	}

	req := endpoints.InternationalBroadcasterScheduleRequest{
		LeagueID: vLeagueID,
		Season:   vSeason,
		RegionID: vRegionID,
		Date:     vDate,
		EST:      vEST,
	}

	resp, err := endpoints.GetInternationalBroadcasterSchedule(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp)
}
