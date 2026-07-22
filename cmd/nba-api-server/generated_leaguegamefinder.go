package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/v3/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// handleLeagueGameFinder is generated from tools/generator/metadata; see
// tools/generator/templates/handler.tmpl. Do not hand-edit - regenerate via
// `cd tools/generator && go run . -endpoint LeagueGameFinder` (or -all-handlers to
// regenerate every handler plus the dispatch table) instead.
func (h *StatsHandler) handleLeagueGameFinder(w http.ResponseWriter, r *http.Request) {
	vLeagueID := leagueIDPtr(parameters.LeagueIDNBA)
	vSeason := seasonPtr(parameters.Season(getSeasonOrDefault(r)))
	vSeasonType := seasonTypePtr(parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season")))
	var vPlayerOrTeam *string
	if raw := r.URL.Query().Get("PlayerOrTeam"); raw != "" {
		vPlayerOrTeam = stringPtr(raw)
	}
	var vPlayerID *string
	if raw := r.URL.Query().Get("PlayerID"); raw != "" {
		vPlayerID = stringPtr(raw)
	}
	var vTeamID *string
	if raw := r.URL.Query().Get("TeamID"); raw != "" {
		vTeamID = stringPtr(raw)
	}
	var vVsTeamID *string
	if raw := r.URL.Query().Get("VsTeamID"); raw != "" {
		vVsTeamID = stringPtr(raw)
	}
	var vOutcome *string
	if raw := r.URL.Query().Get("Outcome"); raw != "" {
		vOutcome = stringPtr(raw)
	}
	var vLocation *string
	if raw := r.URL.Query().Get("Location"); raw != "" {
		vLocation = stringPtr(raw)
	}
	var vDateFrom *string
	if raw := r.URL.Query().Get("DateFrom"); raw != "" {
		vDateFrom = stringPtr(raw)
	}
	var vDateTo *string
	if raw := r.URL.Query().Get("DateTo"); raw != "" {
		vDateTo = stringPtr(raw)
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
	var vPORound *string
	if raw := r.URL.Query().Get("PORound"); raw != "" {
		vPORound = stringPtr(raw)
	}

	req := endpoints.LeagueGameFinderRequest{
		LeagueID:     vLeagueID,
		Season:       vSeason,
		SeasonType:   vSeasonType,
		PlayerOrTeam: vPlayerOrTeam,
		PlayerID:     vPlayerID,
		TeamID:       vTeamID,
		VsTeamID:     vVsTeamID,
		Outcome:      vOutcome,
		Location:     vLocation,
		DateFrom:     vDateFrom,
		DateTo:       vDateTo,
		VsConference: vVsConference,
		VsDivision:   vVsDivision,
		GameSegment:  vGameSegment,
		Period:       vPeriod,
		LastNGames:   vLastNGames,
		PORound:      vPORound,
	}

	resp, err := endpoints.GetLeagueGameFinder(r.Context(), h.client, req)
	if err != nil {
		writeEndpointError(w, err)
		return
	}

	writeSuccess(w, resp.Data)
}
