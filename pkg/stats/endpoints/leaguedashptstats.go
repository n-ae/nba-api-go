package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPtStatsRequest contains parameters for the LeagueDashPtStats endpoint
type LeagueDashPtStatsRequest struct {
	Season *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode *parameters.PerMode
	LeagueID *parameters.LeagueID
	PtMeasureType *string
}


// LeagueDashPtStatsLeagueDashPTStats represents the LeagueDashPTStats result set for LeagueDashPtStats
type LeagueDashPtStatsLeagueDashPTStats struct {
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER_NAME string `json:"PLAYER_NAME"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	AGE int `json:"AGE"`
	GP int `json:"GP"`
	W string `json:"W"`
	L string `json:"L"`
	W_PCT float64 `json:"W_PCT"`
	MIN float64 `json:"MIN"`
	DIST_FEET string `json:"DIST_FEET"`
	DIST_MILES string `json:"DIST_MILES"`
	DIST_MILES_OFF string `json:"DIST_MILES_OFF"`
	DIST_MILES_DEF string `json:"DIST_MILES_DEF"`
	AVG_SPEED string `json:"AVG_SPEED"`
	AVG_SPEED_OFF string `json:"AVG_SPEED_OFF"`
	AVG_SPEED_DEF string `json:"AVG_SPEED_DEF"`
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
			if len(row) >= 17 {
				item := LeagueDashPtStatsLeagueDashPTStats{
					PLAYER_ID: toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					TEAM_ID: toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					AGE: toInt(row[4]),
					GP: toInt(row[5]),
					W: toString(row[6]),
					L: toString(row[7]),
					W_PCT: toFloat(row[8]),
					MIN: toFloat(row[9]),
					DIST_FEET: toString(row[10]),
					DIST_MILES: toString(row[11]),
					DIST_MILES_OFF: toString(row[12]),
					DIST_MILES_DEF: toString(row[13]),
					AVG_SPEED: toString(row[14]),
					AVG_SPEED_OFF: toString(row[15]),
					AVG_SPEED_DEF: toString(row[16]),
				}
				response.LeagueDashPTStats = append(response.LeagueDashPTStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
