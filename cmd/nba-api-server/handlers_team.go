package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func (h *StatsHandler) handleTeamGameLog(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamGameLogRequest{
		TeamID:     teamID,
		Season:     season,
		SeasonType: seasonType,
	}

	resp, err := endpoints.GetTeamGameLog(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamInfoCommon(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamInfoCommonRequest{
		TeamID:     teamID,
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamInfoCommon(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashboardByGeneralSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByGeneralSplitsRequest{
		TeamID:      teamID,
		Season:      season,
		SeasonType:  seasonType,
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamDashboardByGeneralSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashboardByShootingSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByShootingSplitsRequest{
		TeamID:      teamID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamDashboardByShootingSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashboardByOpponent(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByOpponentRequest{
		TeamID:      teamID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamDashboardByOpponent(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDetails(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	req := endpoints.TeamDetailsRequest{
		TeamID: teamID,
	}

	resp, err := endpoints.GetTeamDetails(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamPlayerDashboard(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamPlayerDashboardRequest{
		TeamID:      teamID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamPlayerDashboard(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamLineups(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamLineupsRequest{
		TeamID:        teamID,
		Season:        seasonPtr(season),
		SeasonType:    seasonTypePtr(seasonType),
		PerMode:       perModePtr(perMode),
		MeasureType:   stringPtr("Base"),
		GroupQuantity: stringPtr("5"),
	}

	resp, err := endpoints.GetTeamLineups(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Box Score endpoint handlers


func (h *StatsHandler) handleTeamGameLogs(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	teamID := r.URL.Query().Get("TeamID")

	req := endpoints.TeamGameLogsRequest{
		Season:     season,
		SeasonType: seasonType,
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}
	if teamID != "" {
		req.TeamID = stringPtr(teamID)
	}

	resp, err := endpoints.GetTeamGameLogs(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamYearByYearStats(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamYearByYearStatsRequest{
		TeamID:     teamID,
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamYearByYearStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamVsTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	vsTeamID := r.URL.Query().Get("VsTeamID")
	if teamID == "" || vsTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID and VsTeamID are required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamVsTeamRequest{
		TeamID:     teamID,
		VsTeamID:   vsTeamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamVsTeam(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamHistoricalLeaders(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	req := endpoints.TeamHistoricalLeadersRequest{
		TeamID:   teamID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamHistoricalLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamEstimatedMetrics(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamEstimatedMetricsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamEstimatedMetrics(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashPtShots(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamDashPtShotsRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashPtShots(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Player endpoints (iteration 4)


func (h *StatsHandler) handleTeamDashboardByClutch(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByClutchRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashboardByLastNGames(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByLastNGamesRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByLastNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashboardByYearOverYear(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByYearOverYearRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByYearOverYear(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamVsPlayer(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	vsPlayerID := r.URL.Query().Get("VsPlayerID")
	if teamID == "" || vsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID and VsPlayerID are required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamVsPlayerRequest{
		TeamID:     teamID,
		VsPlayerID: vsPlayerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamVsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Additional endpoints (iteration 5)


func (h *StatsHandler) handleTeamDashboardByGameSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByGameSplitsRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByGameSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamDashboardByTeamPerformance(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByTeamPerformanceRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByTeamPerformance(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamPlayerOnOffSummary(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamPlayerOnOffSummaryRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamPlayerOnOffSummary(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// League endpoints (iteration 6)


func (h *StatsHandler) handleTeamGameStreakFinder(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamGameStreakFinderRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamGameStreakFinder(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamPlayerOnOffDetails(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamPlayerOnOffDetailsRequest{
		TeamID:  teamID,
		PerMode: perModePtr(perMode),
	}

	resp, err := endpoints.GetTeamPlayerOnOffDetails(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamAndPlayersVsPlayers(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	vsPlayerID := r.URL.Query().Get("VsPlayerID")
	if teamID == "" || vsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID and VsPlayerID are required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamAndPlayersVsPlayersRequest{
		TeamID:     teamID,
		VsPlayerID: vsPlayerID,
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamAndPlayersVsPlayers(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamInfoCommonV2(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamInfoCommonV2Request{
		TeamID:     teamID,
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamInfoCommonV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamNextNGames(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamNextNGamesRequest{
		TeamID:     teamID,
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamNextNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleTeamYearOverYearSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamYearOverYearSplitsRequest{
		TeamID:     teamID,
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamYearOverYearSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}



