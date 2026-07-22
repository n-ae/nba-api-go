package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
)

// handleVideoEvents is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint VideoEvents` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleVideoEvents(w http.ResponseWriter, r *http.Request) {
	vGameID := r.URL.Query().Get("GameID")
	if vGameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}
	vGameEventID := r.URL.Query().Get("GameEventID")
	if vGameEventID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameEventID is required")
		return
	}

	req := endpoints.VideoEventsRequest{
		GameID:      vGameID,
		GameEventID: vGameEventID,
	}

	resp, err := endpoints.GetVideoEvents(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
