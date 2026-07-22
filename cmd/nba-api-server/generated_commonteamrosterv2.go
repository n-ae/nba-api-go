package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleCommonTeamRosterV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint CommonTeamRosterV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleCommonTeamRosterV2(w http.ResponseWriter, r *http.Request) {
	vTeamID := r.URL.Query().Get("TeamID")
	if vTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.CommonTeamRosterV2Request{
		TeamID:   vTeamID,
		Season:   vSeason,
		LeagueID: vLeagueID,
	}

	resp, err := endpoints.GetCommonTeamRosterV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
