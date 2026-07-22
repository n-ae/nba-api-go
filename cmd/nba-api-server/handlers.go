package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

type StatsHandler struct {
	client *stats.Client
}

// NewStatsHandler wraps an existing stats.Client. Callers share one client
// across the handler and any other consumer (e.g. the health checker) so
// they share its connection pool and per-host rate limiter instead of each
// constructing their own.
func NewStatsHandler(client *stats.Client) *StatsHandler {
	return &StatsHandler{
		client: client,
	}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET requests are supported")
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/v1/stats/")
	endpoint := strings.ToLower(path)

	handler, ok := generatedDispatch[endpoint]
	if !ok {
		writeError(w, http.StatusNotFound, "endpoint_not_found", "Endpoint not supported: "+endpoint)
		return
	}
	handler(h, w, r)
}

func getQueryOrDefault(r *http.Request, key, defaultValue string) string {
	if value := r.URL.Query().Get(key); value != "" {
		return value
	}
	return defaultValue
}

// nowFunc is the source of "now" for currentSeasonDefault. Tests override
// it to pin a fixed date instead of depending on the actual wall clock, so
// assertions near the season-rollover boundary (October) don't become
// order- or date-dependent.
var nowFunc = time.Now

// currentSeasonDefault returns the NBA season string (e.g. "2025-26")
// most likely to have the data a caller expects when they omit ?Season=.
// NBA seasons start in October and run into the following June's
// playoffs; before October, the previous season - the most recently
// completed or still-finishing one - is the more useful default than a
// season that hasn't tipped off yet.
func currentSeasonDefault() string {
	now := nowFunc()
	year := now.Year()
	if now.Month() < time.October {
		year--
	}
	return string(parameters.NewSeason(year))
}

// getSeasonOrDefault returns the request's Season query parameter, or
// currentSeasonDefault() if omitted. Centralizes what used to be over a
// hundred call sites each hardcoding a literal season string directly
// (see docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md
// finding #11) - a season rollover, or a change to the default-selection
// rule itself, is now a one-place change instead of a five-file
// search-and-replace.
func getSeasonOrDefault(r *http.Request) string {
	return getQueryOrDefault(r, "Season", currentSeasonDefault())
}

func stringPtr(s string) *string {
	return &s
}

func leagueIDPtr(id parameters.LeagueID) *parameters.LeagueID {
	return &id
}

func perModePtr(pm parameters.PerMode) *parameters.PerMode {
	return &pm
}

func seasonPtr(s parameters.Season) *parameters.Season {
	return &s
}

func seasonTypePtr(st parameters.SeasonType) *parameters.SeasonType {
	return &st
}

// writeEndpointError writes an error response for a failed SDK/endpoint
// call. If err is (or wraps) a *models.APIError - meaning NBA.com actually
// responded, just with a non-2xx status - the response reuses that same
// status code, so e.g. an upstream 404 reads as 404 here instead of a
// misleading blanket 500. Anything else (network failures, decode errors,
// a canceled context, ...) falls back to 500, since those represent
// something going wrong on our side rather than a clean upstream response.
func writeEndpointError(w http.ResponseWriter, err error) {
	var apiErr *models.APIError
	if errors.As(err, &apiErr) && apiErr.StatusCode > 0 {
		writeError(w, apiErr.StatusCode, "api_error", err.Error())
		return
	}
	writeError(w, http.StatusInternalServerError, "api_error", err.Error())
}

func writeSuccess(w http.ResponseWriter, data interface{}) {
	type successResponse struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	resp := successResponse{
		Success: true,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
