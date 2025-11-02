package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamEstimatedMetricsRequest contains parameters for the TeamEstimatedMetrics endpoint
type TeamEstimatedMetricsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// TeamEstimatedMetricsTeamEstimatedMetrics represents the TeamEstimatedMetrics result set for TeamEstimatedMetrics
type TeamEstimatedMetricsTeamEstimatedMetrics struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	GP                int     `json:"GP"`
	W                 string  `json:"W"`
	L                 string  `json:"L"`
	W_PCT             float64 `json:"W_PCT"`
	MIN               float64 `json:"MIN"`
	E_OFF_RATING      string  `json:"E_OFF_RATING"`
	E_DEF_RATING      string  `json:"E_DEF_RATING"`
	E_NET_RATING      string  `json:"E_NET_RATING"`
	E_PACE            string  `json:"E_PACE"`
	E_AST_RATIO       float64 `json:"E_AST_RATIO"`
	E_OREB_PCT        float64 `json:"E_OREB_PCT"`
	E_DREB_PCT        float64 `json:"E_DREB_PCT"`
	E_REB_PCT         float64 `json:"E_REB_PCT"`
	E_TOV_PCT         float64 `json:"E_TOV_PCT"`
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
	E_PACE_RANK       float64 `json:"E_PACE_RANK"`
}

// TeamEstimatedMetricsResponse contains the response data from the TeamEstimatedMetrics endpoint
type TeamEstimatedMetricsResponse struct {
	TeamEstimatedMetrics []TeamEstimatedMetricsTeamEstimatedMetrics
}

// GetTeamEstimatedMetrics retrieves data from the teamestimatedmetrics endpoint
func GetTeamEstimatedMetrics(ctx context.Context, client *stats.Client, req TeamEstimatedMetricsRequest) (*models.Response[*TeamEstimatedMetricsResponse], error) {
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
	if err := client.GetJSON(ctx, "/teamestimatedmetrics", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamEstimatedMetricsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamEstimatedMetrics = make([]TeamEstimatedMetricsTeamEstimatedMetrics, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 30 {
				item := TeamEstimatedMetricsTeamEstimatedMetrics{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					GP:                toInt(row[2]),
					W:                 toString(row[3]),
					L:                 toString(row[4]),
					W_PCT:             toFloat(row[5]),
					MIN:               toFloat(row[6]),
					E_OFF_RATING:      toString(row[7]),
					E_DEF_RATING:      toString(row[8]),
					E_NET_RATING:      toString(row[9]),
					E_PACE:            toString(row[10]),
					E_AST_RATIO:       toFloat(row[11]),
					E_OREB_PCT:        toFloat(row[12]),
					E_DREB_PCT:        toFloat(row[13]),
					E_REB_PCT:         toFloat(row[14]),
					E_TOV_PCT:         toFloat(row[15]),
					GP_RANK:           toFloat(row[16]),
					W_RANK:            toFloat(row[17]),
					L_RANK:            toFloat(row[18]),
					W_PCT_RANK:        toFloat(row[19]),
					MIN_RANK:          toFloat(row[20]),
					E_OFF_RATING_RANK: toFloat(row[21]),
					E_DEF_RATING_RANK: toFloat(row[22]),
					E_NET_RATING_RANK: toFloat(row[23]),
					E_AST_RATIO_RANK:  toFloat(row[24]),
					E_OREB_PCT_RANK:   toFloat(row[25]),
					E_DREB_PCT_RANK:   toFloat(row[26]),
					E_REB_PCT_RANK:    toFloat(row[27]),
					E_TOV_PCT_RANK:    toFloat(row[28]),
					E_PACE_RANK:       toFloat(row[29]),
				}
				response.TeamEstimatedMetrics = append(response.TeamEstimatedMetrics, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
