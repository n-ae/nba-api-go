package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayoffPicture is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayoffPicture` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayoffPicture(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeasonID := parameters.Season(r.URL.Query().Get("SeasonID"))
	if vSeasonID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "SeasonID is required")
		return
	}

	req := endpoints.PlayoffPictureRequest{
		LeagueID: vLeagueID,
		SeasonID: vSeasonID,
	}

	resp, err := endpoints.GetPlayoffPicture(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
