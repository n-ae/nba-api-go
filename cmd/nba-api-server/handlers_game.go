package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func (h *StatsHandler) handlePlayByPlayV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.PlayByPlayV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetPlayByPlayV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleShotChartDetail(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	teamID := getQueryOrDefault(r, "TeamID", "0")
	playerID := getQueryOrDefault(r, "PlayerID", "0")

	req := endpoints.ShotChartDetailRequest{
		Season:     season,
		SeasonType: seasonType,
		TeamID:     stringPtr(teamID),
		PlayerID:   stringPtr(playerID),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetShotChartDetail(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleGameRotation(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.GameRotationRequest{
		GameID:   gameID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetGameRotation(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// League endpoint handlers


func (h *StatsHandler) handleWinProbabilityPBP(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.WinProbabilityPBPRequest{
		GameID: gameID,
	}

	resp, err := endpoints.GetWinProbabilityPBP(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleVideoEvents(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	gameEventID := r.URL.Query().Get("GameEventID")
	if gameID == "" || gameEventID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID and GameEventID are required")
		return
	}

	req := endpoints.VideoEventsRequest{
		GameID:      gameID,
		GameEventID: gameEventID,
	}

	resp, err := endpoints.GetVideoEvents(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handlePlayByPlayV3(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.PlayByPlayV3Request{
		GameID: gameID,
	}

	startPeriod := r.URL.Query().Get("StartPeriod")
	if startPeriod != "" {
		req.StartPeriod = &startPeriod
	}

	endPeriod := r.URL.Query().Get("EndPeriod")
	if endPeriod != "" {
		req.EndPeriod = &endPeriod
	}

	resp, err := endpoints.GetPlayByPlayV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}


func (h *StatsHandler) handleShotChartLineupDetail(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	teamID := r.URL.Query().Get("TeamID")
	groupID := r.URL.Query().Get("GroupID")

	req := endpoints.ShotChartLineupDetailRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	if teamID != "" {
		req.TeamID = &teamID
	}
	if groupID != "" {
		req.GroupID = &groupID
	}

	contextMeasure := r.URL.Query().Get("ContextMeasure")
	if contextMeasure != "" {
		req.ContextMeasure = &contextMeasure
	}

	resp, err := endpoints.GetShotChartLineupDetail(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}



