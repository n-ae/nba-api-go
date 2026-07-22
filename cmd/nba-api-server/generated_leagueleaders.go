package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueLeaders is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueLeaders` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueLeaders(w http.ResponseWriter, r *http.Request) {
	vSeason := parameters.Season(getSeasonOrDefault(r))
	vSeasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	vPerMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	vStatCategory := parameters.StatCategory(r.URL.Query().Get("StatCategory"))
	vActiveFlag := r.URL.Query().Get("ActiveFlag")
	vLeagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueLeadersRequest{
		Season:       vSeason,
		SeasonType:   vSeasonType,
		PerMode:      vPerMode,
		StatCategory: vStatCategory,
		ActiveFlag:   vActiveFlag,
		LeagueID:     vLeagueID,
	}

	resp, err := endpoints.LeagueLeaders(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
