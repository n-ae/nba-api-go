package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerCareerByCollege is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerCareerByCollege` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerCareerByCollege(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vCollege *string
	if raw := r.URL.Query().Get("College"); raw != "" {
		vCollege = stringPtr(raw)
	}

	req := endpoints.PlayerCareerByCollegeRequest{
		LeagueID: vLeagueID,
		College:  vCollege,
	}

	resp, err := endpoints.GetPlayerCareerByCollege(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
