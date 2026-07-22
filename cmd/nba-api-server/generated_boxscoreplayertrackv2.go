package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
)

// handleBoxScorePlayerTrackV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint BoxScorePlayerTrackV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleBoxScorePlayerTrackV2(w http.ResponseWriter, r *http.Request) {
	vGameID := r.URL.Query().Get("GameID")
	if vGameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScorePlayerTrackV2Request{
		GameID: vGameID,
	}

	resp, err := endpoints.GetBoxScorePlayerTrackV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
