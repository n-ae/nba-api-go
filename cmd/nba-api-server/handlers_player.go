package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func (h *StatsHandler) handlePlayerGameLog(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerGameLogRequest{
		PlayerID:   playerID,
		Season:     season,
		SeasonType: seasonType,
		LeagueID:   parameters.LeagueIDNBA,
	}

	resp, err := endpoints.PlayerGameLog(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}


func (h *StatsHandler) handlePlayerCareerStats(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerCareerStatsRequest{
		PlayerID: playerID,
		PerMode:  perMode,
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := endpoints.PlayerCareerStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}


func (h *StatsHandler) handlePlayerProfileV2(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerProfileV2Request{
		PlayerID: playerID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerProfileV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerAwards(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerAwardsRequest{
		PlayerID: playerID,
	}

	resp, err := endpoints.GetPlayerAwards(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByGeneralSplits(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByGeneralSplitsRequest{
		PlayerID:    playerID,
		Season:      season,
		SeasonType:  seasonType,
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByGeneralSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByShootingSplits(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByShootingSplitsRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByShootingSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByOpponent(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByOpponentRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByOpponent(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByClutch(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByClutchRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerGameLogs(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerGameLogsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerGameLogs(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerVsPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	vsPlayerID := r.URL.Query().Get("VsPlayerID")

	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	if vsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "VsPlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerVsPlayerRequest{
		PlayerID:   playerID,
		VsPlayerID: vsPlayerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(parameters.PerModePerGame),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerVsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoint handlers (expanded)


func (h *StatsHandler) handlePlayerTrackingShotDashboard(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingShootingEfficiencyRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingShootingEfficiency(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingPasses(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPassesRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPasses(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingDefense(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingDefenseRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingDefense(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingRebounding(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingReboundingRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingRebounding(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingSpeedDistance(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingSpeedDistanceRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingSpeedDistance(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingCatchShoot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingCatchShootRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingCatchShoot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingDrives(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingDrivesRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingDrives(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerEstimatedMetrics(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerEstimatedMetricsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerEstimatedMetrics(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerFantasyProfile(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerFantasyProfileRequest{
		PlayerID: playerID,
	}

	resp, err := endpoints.GetPlayerFantasyProfile(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashPtShots(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerDashPtShotsRequest{
		PlayerID:   playerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashPtShots(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByLastNGames(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByLastNGamesRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByLastNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByTeamPerformance(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByTeamPerformanceRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByTeamPerformance(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerDashboardByGameSplits(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByGameSplitsRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByGameSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoints (iteration 4)


func (h *StatsHandler) handlePlayerDashboardByYearOverYear(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByYearOverYearRequest{
		PlayerID:   playerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByYearOverYear(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerCompare(w http.ResponseWriter, r *http.Request) {
	playerIDList := r.URL.Query().Get("PlayerIDList")
	if playerIDList == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerIDList is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerCompareRequest{
		PlayerIDList: playerIDList,
		Season:       seasonPtr(season),
		SeasonType:   seasonTypePtr(seasonType),
		LeagueID:     leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerCompare(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerYearByYearStats(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerYearByYearStatsRequest{
		PlayerID: playerID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerYearByYearStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Common endpoints (iteration 4)


func (h *StatsHandler) handlePlayerTrackingPostTouch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPostTouchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPostTouch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingPaintTouch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPaintTouchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPaintTouch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingElbowTouch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingElbowTouchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingElbowTouch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingPullUpShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPullUpShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPullUpShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoints (iteration 6)


func (h *StatsHandler) handlePlayerIndex(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.PlayerIndexRequest{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerIndex(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerCareerByCollege(w http.ResponseWriter, r *http.Request) {
	college := r.URL.Query().Get("College")

	req := endpoints.PlayerCareerByCollegeRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}
	if college != "" {
		req.College = stringPtr(college)
	}

	resp, err := endpoints.GetPlayerCareerByCollege(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerGameStreakFinder(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerGameStreakFinderRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerGameStreakFinder(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerEstimatedAdvancedStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerEstimatedAdvancedStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerEstimatedAdvancedStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerCareerByCollegeRollup(w http.ResponseWriter, r *http.Request) {
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerCareerByCollegeRollupRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
		PerMode:  perModePtr(perMode),
	}

	resp, err := endpoints.GetPlayerCareerByCollegeRollup(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Iteration 10 endpoints - Final SDK endpoints


func (h *StatsHandler) handlePlayerNextNGames(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.PlayerNextNGamesRequest{
		PlayerID:   playerID,
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetPlayerNextNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayerTrackingShootingEfficiency(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.PlayerTrackingShootingEfficiencyRequest{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetPlayerTrackingShootingEfficiency(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}



