package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func (h *StatsHandler) handleCommonAllPlayers(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	leagueID := parameters.LeagueIDNBA
	isOnlyCurrent := getQueryOrDefault(r, "IsOnlyCurrentSeason", "0")

	req := endpoints.CommonAllPlayersRequest{
		Season:              season,
		LeagueID:            &leagueID,
		IsOnlyCurrentSeason: &isOnlyCurrent,
	}

	resp, err := endpoints.GetCommonAllPlayers(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}


func (h *StatsHandler) handleScoreboardV2(w http.ResponseWriter, r *http.Request) {
	gameDate := r.URL.Query().Get("GameDate")
	if gameDate == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameDate is required (format: YYYY-MM-DD)")
		return
	}

	leagueID := parameters.LeagueIDNBA
	req := endpoints.ScoreboardV2Request{
		GameDate: gameDate,
		LeagueID: &leagueID,
	}

	resp, err := endpoints.GetScoreboardV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}


func (h *StatsHandler) handleCommonTeamRoster(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.CommonTeamRosterRequest{
		TeamID:   teamID,
		Season:   &season,
		LeagueID: &leagueID,
	}

	resp, err := endpoints.GetCommonTeamRoster(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}


func (h *StatsHandler) handleCommonPlayerInfo(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.CommonPlayerInfoRequest{
		PlayerID: playerID,
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := endpoints.CommonPlayerInfo(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}


func (h *StatsHandler) handlePlayoffPicture(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.PlayoffPictureRequest{
		SeasonID: season,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayoffPicture(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCommonPlayerInfoV2(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.CommonPlayerInfoV2Request{
		PlayerID: playerID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonPlayerInfoV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCommonAllPlayersV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.CommonAllPlayersV2Request{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonAllPlayersV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCommonTeamRosterV2(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.CommonTeamRosterV2Request{
		TeamID:   teamID,
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonTeamRosterV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCommonPlayoffSeries(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.CommonPlayoffSeriesRequest{
		Season:   season,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonPlayoffSeries(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCommonTeamYears(w http.ResponseWriter, r *http.Request) {
	req := endpoints.CommonTeamYearsRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonTeamYears(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Draft & Historical endpoints (iteration 4)


func (h *StatsHandler) handleDraftHistory(w http.ResponseWriter, r *http.Request) {
	req := endpoints.DraftHistoryRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDraftHistory(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleDraftBoard(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.DraftBoardRequest{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDraftBoard(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleDraftCombineStats(w http.ResponseWriter, r *http.Request) {
	seasonYear := getQueryOrDefault(r, "SeasonYear", "2023-24")

	req := endpoints.DraftCombineStatsRequest{
		SeasonYear: stringPtr(seasonYear),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDraftCombineStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleFranchiseHistory(w http.ResponseWriter, r *http.Request) {
	req := endpoints.FranchiseHistoryRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetFranchiseHistory(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleFranchiseLeaders(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	req := endpoints.FranchiseLeadersRequest{
		TeamID:   teamID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetFranchiseLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// League endpoints (iteration 5)


func (h *StatsHandler) handleInfographicFanDuelPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.InfographicFanDuelPlayerRequest{
		PlayerID: playerID,
	}

	resp, err := endpoints.GetInfographicFanDuelPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleHomepageV2(w http.ResponseWriter, r *http.Request) {
	req := endpoints.HomepageV2Request{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetHomepageV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleHomepageLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.HomepageLeadersRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetHomepageLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Player Tracking endpoints (iteration 6)


func (h *StatsHandler) handleScoreboardV3(w http.ResponseWriter, r *http.Request) {
	gameDate := getQueryOrDefault(r, "GameDate", "")

	req := endpoints.ScoreboardV3Request{
		GameDate: gameDate,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetScoreboardV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleAllTimeLeadersGrids(w http.ResponseWriter, r *http.Request) {
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.AllTimeLeadersGridsRequest{
		PerMode:    perModePtr(perMode),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetAllTimeLeadersGrids(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleDefenseHub(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.DefenseHubRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDefenseHub(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleAssistTracker(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.AssistTrackerRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetAssistTracker(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleSynergyPlayTypes(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.SynergyPlayTypesRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetSynergyPlayTypes(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCumeStatsPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.CumeStatsPlayerRequest{
		PlayerID:   playerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCumeStatsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCumeStatsTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.CumeStatsTeamRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCumeStatsTeam(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleOpponentShooting(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.OpponentShootingRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetOpponentShooting(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleShootingEfficiency(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.ShootingEfficiencyRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetShootingEfficiency(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleMatchupRollup(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.MatchupRollupRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetMatchupRollup(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleAssistLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.AssistLeadersRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetAssistLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleCommonPlayoffSeriesV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.CommonPlayoffSeriesV2Request{
		Season:   &season,
		LeagueID: &leagueID,
	}

	resp, err := endpoints.GetCommonPlayoffSeriesV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}



