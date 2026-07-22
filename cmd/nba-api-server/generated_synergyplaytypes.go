package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleSynergyPlayTypes is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint SynergyPlayTypes` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleSynergyPlayTypes(w http.ResponseWriter, r *http.Request) {
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vPlayerOrTeam *string
	if raw := r.URL.Query().Get("PlayerOrTeam"); raw != "" {
		vPlayerOrTeam = stringPtr(raw)
	}
	var vPlayType *string
	if raw := r.URL.Query().Get("PlayType"); raw != "" {
		vPlayType = stringPtr(raw)
	}

	req := endpoints.SynergyPlayTypesRequest{
		Season:       vSeason,
		SeasonType:   vSeasonType,
		PerMode:      vPerMode,
		LeagueID:     vLeagueID,
		PlayerOrTeam: vPlayerOrTeam,
		PlayType:     vPlayType,
	}

	resp, err := endpoints.GetSynergyPlayTypes(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
