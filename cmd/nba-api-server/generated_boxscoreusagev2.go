package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
)

// handleBoxScoreUsageV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint BoxScoreUsageV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleBoxScoreUsageV2(w http.ResponseWriter, r *http.Request) {
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
	var vStartRange *string
	if raw := r.URL.Query().Get("StartRange"); raw != "" {
		vStartRange = stringPtr(raw)
	}
	var vEndRange *string
	if raw := r.URL.Query().Get("EndRange"); raw != "" {
		vEndRange = stringPtr(raw)
	}
	var vRangeType *string
	if raw := r.URL.Query().Get("RangeType"); raw != "" {
		vRangeType = stringPtr(raw)
	}

	req := endpoints.BoxScoreUsageV2Request{
		GameID:      vGameID,
		StartPeriod: vStartPeriod,
		EndPeriod:   vEndPeriod,
		StartRange:  vStartRange,
		EndRange:    vEndRange,
		RangeType:   vRangeType,
	}

	resp, err := endpoints.GetBoxScoreUsageV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
