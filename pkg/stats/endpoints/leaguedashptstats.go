package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPtStatsRequest contains parameters for the LeagueDashPtStats endpoint
type LeagueDashPtStatsRequest struct {
	Season        *parameters.Season
	SeasonType    *parameters.SeasonType
	PerMode       *parameters.PerMode
	LeagueID      *parameters.LeagueID
	PlayerOrTeam  *string
	PtMeasureType *string
}

// LeagueDashPtStatsLeagueDashPTStats represents the LeagueDashPTStats result set for LeagueDashPtStats
type LeagueDashPtStatsLeagueDashPTStats struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	DIST_FEET         string  `json:"DIST_FEET"`
	DIST_MILES        string  `json:"DIST_MILES"`
	DIST_MILES_OFF    string  `json:"DIST_MILES_OFF"`
	DIST_MILES_DEF    string  `json:"DIST_MILES_DEF"`
	AVG_SPEED         string  `json:"AVG_SPEED"`
	AVG_SPEED_OFF     string  `json:"AVG_SPEED_OFF"`
	AVG_SPEED_DEF     string  `json:"AVG_SPEED_DEF"`
}

// LeagueDashPtStatsResponse contains the response data from the LeagueDashPtStats endpoint
type LeagueDashPtStatsResponse struct {
	LeagueDashPTStats []LeagueDashPtStatsLeagueDashPTStats
}

// GetLeagueDashPtStats retrieves data from the leaguedashptstats endpoint
func GetLeagueDashPtStats(ctx context.Context, client *stats.Client, req LeagueDashPtStatsRequest) (*models.Response[*LeagueDashPtStatsResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.PlayerOrTeam != nil {
		params.Set("PlayerOrTeam", string(*req.PlayerOrTeam))
	}
	if req.PtMeasureType != nil {
		params.Set("PtMeasureType", string(*req.PtMeasureType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguedashptstats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPtStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPTStats = make([]LeagueDashPtStatsLeagueDashPTStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 13 {
				item := LeagueDashPtStatsLeagueDashPTStats{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					DIST_FEET:         toString(row[6]),
					DIST_MILES:        toString(row[7]),
					DIST_MILES_OFF:    toString(row[8]),
					DIST_MILES_DEF:    toString(row[9]),
					AVG_SPEED:         toString(row[10]),
					AVG_SPEED_OFF:     toString(row[11]),
					AVG_SPEED_DEF:     toString(row[12]),
				}
				response.LeagueDashPTStats = append(response.LeagueDashPTStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
