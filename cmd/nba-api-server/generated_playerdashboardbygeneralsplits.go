package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handlePlayerDashboardByGeneralSplits is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint PlayerDashboardByGeneralSplits` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handlePlayerDashboardByGeneralSplits(w http.ResponseWriter, r *http.Request) {
	vPlayerID := r.URL.Query().Get("PlayerID")
	if vPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	vSeason := parameters.Season(r.URL.Query().Get("Season"))
	if vSeason == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Season is required")
		return
	}
	vSeasonType := parameters.SeasonType(r.URL.Query().Get("SeasonType"))
	if vSeasonType == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "SeasonType is required")
		return
	}
	var vMeasureType *string
	if raw := r.URL.Query().Get("MeasureType"); raw != "" {
		vMeasureType = stringPtr(raw)
	}
	vPerMode := perModePtr(parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame")))
	var vPlusMinus *string
	if raw := r.URL.Query().Get("PlusMinus"); raw != "" {
		vPlusMinus = stringPtr(raw)
	}
	var vPaceAdjust *string
	if raw := r.URL.Query().Get("PaceAdjust"); raw != "" {
		vPaceAdjust = stringPtr(raw)
	}
	var vRank *string
	if raw := r.URL.Query().Get("Rank"); raw != "" {
		vRank = stringPtr(raw)
	}
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vOutcome *string
	if raw := r.URL.Query().Get("Outcome"); raw != "" {
		vOutcome = stringPtr(raw)
	}
	var vLocation *string
	if raw := r.URL.Query().Get("Location"); raw != "" {
		vLocation = stringPtr(raw)
	}
	var vMonth *string
	if raw := r.URL.Query().Get("Month"); raw != "" {
		vMonth = stringPtr(raw)
	}
	var vSeasonSegment *string
	if raw := r.URL.Query().Get("SeasonSegment"); raw != "" {
		vSeasonSegment = stringPtr(raw)
	}
	var vDateFrom *string
	if raw := r.URL.Query().Get("DateFrom"); raw != "" {
		vDateFrom = stringPtr(raw)
	}
	var vDateTo *string
	if raw := r.URL.Query().Get("DateTo"); raw != "" {
		vDateTo = stringPtr(raw)
	}
	var vOpponentTeamID *string
	if raw := r.URL.Query().Get("OpponentTeamID"); raw != "" {
		vOpponentTeamID = stringPtr(raw)
	}
	var vVsConference *string
	if raw := r.URL.Query().Get("VsConference"); raw != "" {
		vVsConference = stringPtr(raw)
	}
	var vVsDivision *string
	if raw := r.URL.Query().Get("VsDivision"); raw != "" {
		vVsDivision = stringPtr(raw)
	}
	var vGameSegment *string
	if raw := r.URL.Query().Get("GameSegment"); raw != "" {
		vGameSegment = stringPtr(raw)
	}
	var vPeriod *string
	if raw := r.URL.Query().Get("Period"); raw != "" {
		vPeriod = stringPtr(raw)
	}
	var vLastNGames *string
	if raw := r.URL.Query().Get("LastNGames"); raw != "" {
		vLastNGames = stringPtr(raw)
	}

	req := endpoints.PlayerDashboardByGeneralSplitsRequest{
		PlayerID:       vPlayerID,
		Season:         vSeason,
		SeasonType:     vSeasonType,
		MeasureType:    vMeasureType,
		PerMode:        vPerMode,
		PlusMinus:      vPlusMinus,
		PaceAdjust:     vPaceAdjust,
		Rank:           vRank,
		LeagueID:       vLeagueID,
		Outcome:        vOutcome,
		Location:       vLocation,
		Month:          vMonth,
		SeasonSegment:  vSeasonSegment,
		DateFrom:       vDateFrom,
		DateTo:         vDateTo,
		OpponentTeamID: vOpponentTeamID,
		VsConference:   vVsConference,
		VsDivision:     vVsDivision,
		GameSegment:    vGameSegment,
		Period:         vPeriod,
		LastNGames:     vLastNGames,
	}

	resp, err := endpoints.GetPlayerDashboardByGeneralSplits(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
