package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleHomepageLeaders is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint HomepageLeaders` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleHomepageLeaders(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vPlayerOrTeam *string
	if raw := r.URL.Query().Get("PlayerOrTeam"); raw != "" {
		vPlayerOrTeam = stringPtr(raw)
	}
	var vGameScope *string
	if raw := r.URL.Query().Get("GameScope"); raw != "" {
		vGameScope = stringPtr(raw)
	}
	var vPlayerScope *string
	if raw := r.URL.Query().Get("PlayerScope"); raw != "" {
		vPlayerScope = stringPtr(raw)
	}
	var vStat *string
	if raw := r.URL.Query().Get("Stat"); raw != "" {
		vStat = stringPtr(raw)
	}

	req := endpoints.HomepageLeadersRequest{
		Season:       vSeason,
		SeasonType:   vSeasonType,
		LeagueID:     vLeagueID,
		PlayerOrTeam: vPlayerOrTeam,
		GameScope:    vGameScope,
		PlayerScope:  vPlayerScope,
		Stat:         vStat,
	}

	resp, err := endpoints.GetHomepageLeaders(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
