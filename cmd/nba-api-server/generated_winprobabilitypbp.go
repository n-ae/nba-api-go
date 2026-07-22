package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
)

// handleWinProbabilityPBP is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint WinProbabilityPBP` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleWinProbabilityPBP(w http.ResponseWriter, r *http.Request) {
	vGameID := r.URL.Query().Get("GameID")
	if vGameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}
	var vRunType *string
	if raw := r.URL.Query().Get("RunType"); raw != "" {
		vRunType = stringPtr(raw)
	}

	req := endpoints.WinProbabilityPBPRequest{
		GameID:  vGameID,
		RunType: vRunType,
	}

	resp, err := endpoints.GetWinProbabilityPBP(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
