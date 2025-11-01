package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerEstimatedAdvancedStatsRequest contains parameters for the PlayerEstimatedAdvancedStats endpoint
type PlayerEstimatedAdvancedStatsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// PlayerEstimatedAdvancedStatsPlayerEstimatedAdvancedStats represents the PlayerEstimatedAdvancedStats result set for PlayerEstimatedAdvancedStats
type PlayerEstimatedAdvancedStatsPlayerEstimatedAdvancedStats struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	E_OFF_RATING      string  `json:"E_OFF_RATING"`
	E_DEF_RATING      string  `json:"E_DEF_RATING"`
	E_NET_RATING      string  `json:"E_NET_RATING"`
	E_AST_RATIO       float64 `json:"E_AST_RATIO"`
	E_OREB_PCT        float64 `json:"E_OREB_PCT"`
	E_DREB_PCT        float64 `json:"E_DREB_PCT"`
	E_REB_PCT         float64 `json:"E_REB_PCT"`
	E_TOV_PCT         float64 `json:"E_TOV_PCT"`
	E_USG_PCT         float64 `json:"E_USG_PCT"`
	E_PACE            string  `json:"E_PACE"`
}

// PlayerEstimatedAdvancedStatsResponse contains the response data from the PlayerEstimatedAdvancedStats endpoint
type PlayerEstimatedAdvancedStatsResponse struct {
	PlayerEstimatedAdvancedStats []PlayerEstimatedAdvancedStatsPlayerEstimatedAdvancedStats
}

// GetPlayerEstimatedAdvancedStats retrieves data from the playerestimatedadvancedstats endpoint
func GetPlayerEstimatedAdvancedStats(ctx context.Context, client *stats.Client, req PlayerEstimatedAdvancedStatsRequest) (*models.Response[*PlayerEstimatedAdvancedStatsResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playerestimatedadvancedstats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerEstimatedAdvancedStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerEstimatedAdvancedStats = make([]PlayerEstimatedAdvancedStatsPlayerEstimatedAdvancedStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 16 {
				item := PlayerEstimatedAdvancedStatsPlayerEstimatedAdvancedStats{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					E_OFF_RATING:      toString(row[6]),
					E_DEF_RATING:      toString(row[7]),
					E_NET_RATING:      toString(row[8]),
					E_AST_RATIO:       toFloat(row[9]),
					E_OREB_PCT:        toFloat(row[10]),
					E_DREB_PCT:        toFloat(row[11]),
					E_REB_PCT:         toFloat(row[12]),
					E_TOV_PCT:         toFloat(row[13]),
					E_USG_PCT:         toFloat(row[14]),
					E_PACE:            toString(row[15]),
				}
				response.PlayerEstimatedAdvancedStats = append(response.PlayerEstimatedAdvancedStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
