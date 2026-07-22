package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleCommonPlayerInfoV2 is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint CommonPlayerInfoV2` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleCommonPlayerInfoV2(w http.ResponseWriter, r *http.Request) {
	vPlayerID := r.URL.Query().Get("PlayerID")
	if vPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)

	req := endpoints.CommonPlayerInfoV2Request{
		PlayerID: vPlayerID,
		LeagueID: vLeagueID,
	}

	resp, err := endpoints.GetCommonPlayerInfoV2(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
