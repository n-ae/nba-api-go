package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// CommonPlayoffSeriesV2Request contains parameters for the CommonPlayoffSeriesV2 endpoint
type CommonPlayoffSeriesV2Request struct {
	Season   *parameters.Season
	LeagueID *parameters.LeagueID
}

// CommonPlayoffSeriesV2PlayoffSeries represents the PlayoffSeries result set for CommonPlayoffSeriesV2
type CommonPlayoffSeriesV2PlayoffSeries struct {
	GAME_ID         string `json:"GAME_ID"`
	HOME_TEAM_ID    int    `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID int    `json:"VISITOR_TEAM_ID"`
	SERIES_ID       string `json:"SERIES_ID"`
	GAME_NUM        string `json:"GAME_NUM"`
}

// CommonPlayoffSeriesV2Response contains the response data from the CommonPlayoffSeriesV2 endpoint
type CommonPlayoffSeriesV2Response struct {
	PlayoffSeries []CommonPlayoffSeriesV2PlayoffSeries
}

// GetCommonPlayoffSeriesV2 retrieves data from the commonplayoffseriesv2 endpoint
func GetCommonPlayoffSeriesV2(ctx context.Context, client *stats.Client, req CommonPlayoffSeriesV2Request) (*models.Response[*CommonPlayoffSeriesV2Response], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonplayoffseriesv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonPlayoffSeriesV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayoffSeries = make([]CommonPlayoffSeriesV2PlayoffSeries, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 5 {
				item := CommonPlayoffSeriesV2PlayoffSeries{
					GAME_ID:         toString(row[0]),
					HOME_TEAM_ID:    toInt(row[1]),
					VISITOR_TEAM_ID: toInt(row[2]),
					SERIES_ID:       toString(row[3]),
					GAME_NUM:        toString(row[4]),
				}
				response.PlayoffSeries = append(response.PlayoffSeries, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
