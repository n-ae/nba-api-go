package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func (h *StatsHandler) handleLeagueStandings(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueStandingsRequest{
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueStandings(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueLeadersRequest{
		Season:     season,
		SeasonType: seasonType,
		PerMode:    perMode,
		LeagueID:   parameters.LeagueIDNBA,
	}

	resp, err := endpoints.LeagueLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueDashTeamStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashTeamStatsRequest{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashTeamStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueDashPlayerStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashPlayerStatsRequest{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

// Player endpoint handlers (expanded)

func (h *StatsHandler) handleLeagueGameLog(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueGameLogRequest{
		Season:       season,
		SeasonType:   seasonTypePtr(seasonType),
		PlayerOrTeam: stringPtr("P"),
		LeagueID:     leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueGameLog(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashLineups(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashLineupsRequest{
		Season:        seasonPtr(season),
		SeasonType:    seasonTypePtr(seasonType),
		PerMode:       perModePtr(perMode),
		MeasureType:   stringPtr("Base"),
		GroupQuantity: stringPtr("5"),
		LeagueID:      leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashLineups(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Additional League endpoint handlers

func (h *StatsHandler) handleLeagueDashPlayerClutch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPlayerClutchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamClutch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashTeamClutchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPlayerBioStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueDashPlayerBioStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerBioStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamBioStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueDashTeamBioStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamBioStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPtStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPtStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPtStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueHustleStatsPlayer(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueHustleStatsPlayerRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueHustleStatsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueHustleStatsTeam(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueHustleStatsTeamRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueHustleStatsTeam(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPtDefend(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPtDefendRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPtDefend(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueGameFinder(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueGameFinderRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueGameFinder(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueStandingsV3(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.LeagueStandingsV3Request{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueStandingsV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Additional Player endpoint handlers

func (h *StatsHandler) handleLeagueDashPlayerShotLocations(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPlayerShotLocationsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerShotLocations(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamShotLocations(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashTeamShotLocationsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamShotLocations(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueSeasonMatchups(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueSeasonMatchupsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueSeasonMatchups(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPtTeamDefend(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPtTeamDefendRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPtTeamDefend(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoints (iteration 5)

func (h *StatsHandler) handleLeagueDashPlayerPtShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPlayerPtShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerPtShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamPtShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashTeamPtShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamPtShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Advanced analytics endpoints (iteration 7)

func (h *StatsHandler) handleLeagueDashOppPtShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashOppPtShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashOppPtShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Final endpoints (iteration 8)

func (h *StatsHandler) handleLeagueLeadersV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueLeadersV2Request{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueLeadersV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeaguePlayerOnDetails(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeaguePlayerOnDetailsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeaguePlayerOnDetails(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueHustleStatsTeamLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueHustleStatsTeamLeadersRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueHustleStatsTeamLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Iteration 9 endpoints

func (h *StatsHandler) handleLeagueDashPlayerClutchV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashPlayerClutchV2Request{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerClutchV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPlayerShotLocationV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashPlayerShotLocationV2Request{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerShotLocationV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamClutchV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashTeamClutchV2Request{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashTeamClutchV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}
