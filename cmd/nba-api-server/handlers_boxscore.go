package main

import (
	"net/http"

	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
)

func (h *StatsHandler) handleBoxScoreSummaryV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreSummaryV2Request{
		GameID: gameID,
	}

	resp, err := endpoints.GetBoxScoreSummaryV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreTraditionalV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreTraditionalV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreTraditionalV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreAdvancedV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreAdvancedV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreAdvancedV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreScoringV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreScoringV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreScoringV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreMiscV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreMiscV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreMiscV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreUsageV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreUsageV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreUsageV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreFourFactorsV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreFourFactorsV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreFourFactorsV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScorePlayerTrackV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScorePlayerTrackV2Request{
		GameID: gameID,
	}

	resp, err := endpoints.GetBoxScorePlayerTrackV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Player Tracking endpoint handlers

func (h *StatsHandler) handleBoxScoreDefensiveV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreDefensiveV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreDefensiveV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreHustleV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreHustleV2Request{
		GameID: gameID,
	}

	resp, err := endpoints.GetBoxScoreHustleV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Game endpoint handlers

func (h *StatsHandler) handleBoxScoreMatchupsV3(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreMatchupsV3Request{
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

	resp, err := endpoints.GetBoxScoreMatchupsV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}
