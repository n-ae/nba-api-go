package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleDraftHistory is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint DraftHistory` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleDraftHistory(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))

	req := endpoints.DraftHistoryRequest{
		LeagueID: vLeagueID,
		Season:   vSeason,
	}

	resp, err := endpoints.GetDraftHistory(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
