package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleDraftCombineStats is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint DraftCombineStats` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleDraftCombineStats(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vSeasonYear *string
	if raw := r.URL.Query().Get("SeasonYear"); raw != "" {
		vSeasonYear = stringPtr(raw)
	}

	req := endpoints.DraftCombineStatsRequest{
		LeagueID:   vLeagueID,
		SeasonYear: vSeasonYear,
	}

	resp, err := endpoints.GetDraftCombineStats(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
