package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerEstimatedMetricsRequest contains parameters for the PlayerEstimatedMetrics endpoint
type PlayerEstimatedMetricsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// PlayerEstimatedMetricsPlayerEstimatedMetrics represents the PlayerEstimatedMetrics result set for PlayerEstimatedMetrics
type PlayerEstimatedMetricsPlayerEstimatedMetrics struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	GP                int     `json:"GP"`
	W                 string  `json:"W"`
	L                 string  `json:"L"`
	W_PCT             float64 `json:"W_PCT"`
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
	GP_RANK           float64 `json:"GP_RANK"`
	W_RANK            float64 `json:"W_RANK"`
	L_RANK            float64 `json:"L_RANK"`
	W_PCT_RANK        float64 `json:"W_PCT_RANK"`
	MIN_RANK          float64 `json:"MIN_RANK"`
	E_OFF_RATING_RANK float64 `json:"E_OFF_RATING_RANK"`
	E_DEF_RATING_RANK float64 `json:"E_DEF_RATING_RANK"`
	E_NET_RATING_RANK float64 `json:"E_NET_RATING_RANK"`
	E_AST_RATIO_RANK  float64 `json:"E_AST_RATIO_RANK"`
	E_OREB_PCT_RANK   float64 `json:"E_OREB_PCT_RANK"`
	E_DREB_PCT_RANK   float64 `json:"E_DREB_PCT_RANK"`
	E_REB_PCT_RANK    float64 `json:"E_REB_PCT_RANK"`
	E_TOV_PCT_RANK    float64 `json:"E_TOV_PCT_RANK"`
	E_USG_PCT_RANK    float64 `json:"E_USG_PCT_RANK"`
	E_PACE_RANK       float64 `json:"E_PACE_RANK"`
}

// PlayerEstimatedMetricsResponse contains the response data from the PlayerEstimatedMetrics endpoint
type PlayerEstimatedMetricsResponse struct {
	PlayerEstimatedMetrics []PlayerEstimatedMetricsPlayerEstimatedMetrics
}

// GetPlayerEstimatedMetrics retrieves data from the playerestimatedmetrics endpoint
func GetPlayerEstimatedMetrics(ctx context.Context, client *stats.Client, req PlayerEstimatedMetricsRequest) (*models.Response[*PlayerEstimatedMetricsResponse], error) {
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
	if err := client.GetJSON(ctx, "playerestimatedmetrics", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerEstimatedMetricsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerEstimatedMetrics = make([]PlayerEstimatedMetricsPlayerEstimatedMetrics, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 32 {
				item := PlayerEstimatedMetricsPlayerEstimatedMetrics{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					GP:                toInt(row[2]),
					W:                 toString(row[3]),
					L:                 toString(row[4]),
					W_PCT:             toFloat(row[5]),
					MIN:               toFloat(row[6]),
					E_OFF_RATING:      toString(row[7]),
					E_DEF_RATING:      toString(row[8]),
					E_NET_RATING:      toString(row[9]),
					E_AST_RATIO:       toFloat(row[10]),
					E_OREB_PCT:        toFloat(row[11]),
					E_DREB_PCT:        toFloat(row[12]),
					E_REB_PCT:         toFloat(row[13]),
					E_TOV_PCT:         toFloat(row[14]),
					E_USG_PCT:         toFloat(row[15]),
					E_PACE:            toString(row[16]),
					GP_RANK:           toFloat(row[17]),
					W_RANK:            toFloat(row[18]),
					L_RANK:            toFloat(row[19]),
					W_PCT_RANK:        toFloat(row[20]),
					MIN_RANK:          toFloat(row[21]),
					E_OFF_RATING_RANK: toFloat(row[22]),
					E_DEF_RATING_RANK: toFloat(row[23]),
					E_NET_RATING_RANK: toFloat(row[24]),
					E_AST_RATIO_RANK:  toFloat(row[25]),
					E_OREB_PCT_RANK:   toFloat(row[26]),
					E_DREB_PCT_RANK:   toFloat(row[27]),
					E_REB_PCT_RANK:    toFloat(row[28]),
					E_TOV_PCT_RANK:    toFloat(row[29]),
					E_USG_PCT_RANK:    toFloat(row[30]),
					E_PACE_RANK:       toFloat(row[31]),
				}
				response.PlayerEstimatedMetrics = append(response.PlayerEstimatedMetrics, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
