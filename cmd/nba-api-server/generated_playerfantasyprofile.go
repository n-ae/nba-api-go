package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
)

// handlePlayerFantasyProfile is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerFantasyProfile` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerFantasyProfile(w http.ResponseWriter, r *http.Request) {
	vPlayerID := r.URL.Query().Get("PlayerID")
	if vPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerFantasyProfileRequest{
		PlayerID: vPlayerID,
	}

	resp, err := endpoints.GetPlayerFantasyProfile(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
