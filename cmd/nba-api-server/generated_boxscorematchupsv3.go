package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
)

// handleBoxScoreMatchupsV3 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint BoxScoreMatchupsV3` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleBoxScoreMatchupsV3(w http.ResponseWriter, r *http.Request) {
	vGameID := r.URL.Query().Get("GameID")
	if vGameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}
	var vStartPeriod *string
	if raw := r.URL.Query().Get("StartPeriod"); raw != "" {
		vStartPeriod = stringPtr(raw)
	}
	var vEndPeriod *string
	if raw := r.URL.Query().Get("EndPeriod"); raw != "" {
		vEndPeriod = stringPtr(raw)
	}

	req := endpoints.BoxScoreMatchupsV3Request{
		GameID:      vGameID,
		StartPeriod: vStartPeriod,
		EndPeriod:   vEndPeriod,
	}

	resp, err := endpoints.GetBoxScoreMatchupsV3(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
