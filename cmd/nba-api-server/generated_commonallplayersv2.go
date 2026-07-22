package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleCommonAllPlayersV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint CommonAllPlayersV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleCommonAllPlayersV2(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	var vIsOnlyCurrentSeason *string
	if raw := r.URL.Query().Get("IsOnlyCurrentSeason"); raw != "" {
		vIsOnlyCurrentSeason = stringPtr(raw)
	}

	req := endpoints.CommonAllPlayersV2Request{
		LeagueID:            vLeagueID,
		Season:              vSeason,
		IsOnlyCurrentSeason: vIsOnlyCurrentSeason,
	}

	resp, err := endpoints.GetCommonAllPlayersV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
