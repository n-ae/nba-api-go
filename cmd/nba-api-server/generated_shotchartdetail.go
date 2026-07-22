package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleShotChartDetail is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint ShotChartDetail` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleShotChartDetail(w http.ResponseWriter, r *http.Request) {
	var vPlayerID *string
	if raw := r.URL.Query().Get("PlayerID"); raw != "" {
		vPlayerID = stringPtr(raw)
	}
	var vTeamID *string
	if raw := r.URL.Query().Get("TeamID"); raw != "" {
		vTeamID = stringPtr(raw)
	}
	var vGameID *string
	if raw := r.URL.Query().Get("GameID"); raw != "" {
		vGameID = stringPtr(raw)
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
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	var vContextMeasure *string
	if raw := r.URL.Query().Get("ContextMeasure"); raw != "" {
		vContextMeasure = stringPtr(raw)
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
	var vPeriod *string
	if raw := r.URL.Query().Get("Period"); raw != "" {
		vPeriod = stringPtr(raw)
	}
	var vRookieYear *string
	if raw := r.URL.Query().Get("RookieYear"); raw != "" {
		vRookieYear = stringPtr(raw)
	}
	var vVsConference *string
	if raw := r.URL.Query().Get("VsConference"); raw != "" {
		vVsConference = stringPtr(raw)
	}
	var vVsDivision *string
	if raw := r.URL.Query().Get("VsDivision"); raw != "" {
		vVsDivision = stringPtr(raw)
	}
	var vPosition *string
	if raw := r.URL.Query().Get("Position"); raw != "" {
		vPosition = stringPtr(raw)
	}
	var vGameSegment *string
	if raw := r.URL.Query().Get("GameSegment"); raw != "" {
		vGameSegment = stringPtr(raw)
	}
	var vLastNGames *string
	if raw := r.URL.Query().Get("LastNGames"); raw != "" {
		vLastNGames = stringPtr(raw)
	}
	var vLocation *string
	if raw := r.URL.Query().Get("Location"); raw != "" {
		vLocation = stringPtr(raw)
	}
	var vMonth *string
	if raw := r.URL.Query().Get("Month"); raw != "" {
		vMonth = stringPtr(raw)
	}
	var vOutcome *string
	if raw := r.URL.Query().Get("Outcome"); raw != "" {
		vOutcome = stringPtr(raw)
	}
	var vSeasonSegment *string
	if raw := r.URL.Query().Get("SeasonSegment"); raw != "" {
		vSeasonSegment = stringPtr(raw)
	}
	var vAheadBehind *string
	if raw := r.URL.Query().Get("AheadBehind"); raw != "" {
		vAheadBehind = stringPtr(raw)
	}
	var vClutchTime *string
	if raw := r.URL.Query().Get("ClutchTime"); raw != "" {
		vClutchTime = stringPtr(raw)
	}
	var vPointDiff *string
	if raw := r.URL.Query().Get("PointDiff"); raw != "" {
		vPointDiff = stringPtr(raw)
	}
	var vRangeType *string
	if raw := r.URL.Query().Get("RangeType"); raw != "" {
		vRangeType = stringPtr(raw)
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

	req := endpoints.ShotChartDetailRequest{
		PlayerID:       vPlayerID,
		TeamID:         vTeamID,
		GameID:         vGameID,
		Season:         vSeason,
		SeasonType:     vSeasonType,
		LeagueID:       vLeagueID,
		ContextMeasure: vContextMeasure,
		DateFrom:       vDateFrom,
		DateTo:         vDateTo,
		OpponentTeamID: vOpponentTeamID,
		Period:         vPeriod,
		RookieYear:     vRookieYear,
		VsConference:   vVsConference,
		VsDivision:     vVsDivision,
		Position:       vPosition,
		GameSegment:    vGameSegment,
		LastNGames:     vLastNGames,
		Location:       vLocation,
		Month:          vMonth,
		Outcome:        vOutcome,
		SeasonSegment:  vSeasonSegment,
		AheadBehind:    vAheadBehind,
		ClutchTime:     vClutchTime,
		PointDiff:      vPointDiff,
		RangeType:      vRangeType,
		StartPeriod:    vStartPeriod,
		EndPeriod:      vEndPeriod,
		StartRange:     vStartRange,
		EndRange:       vEndRange,
	}

	resp, err := endpoints.GetShotChartDetail(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
