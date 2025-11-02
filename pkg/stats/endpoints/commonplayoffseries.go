package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// CommonPlayoffSeriesRequest contains parameters for the CommonPlayoffSeries endpoint
type CommonPlayoffSeriesRequest struct {
	LeagueID *parameters.LeagueID
	Season   parameters.Season
	SeriesID *string
}

// CommonPlayoffSeriesPlayoffSeries represents the PlayoffSeries result set for CommonPlayoffSeries
type CommonPlayoffSeriesPlayoffSeries struct {
	GAME_ID         string `json:"GAME_ID"`
	HOME_TEAM_ID    int    `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID int    `json:"VISITOR_TEAM_ID"`
	SERIES_ID       string `json:"SERIES_ID"`
	GAME_NUM        string `json:"GAME_NUM"`
}

// CommonPlayoffSeriesResponse contains the response data from the CommonPlayoffSeries endpoint
type CommonPlayoffSeriesResponse struct {
	PlayoffSeries []CommonPlayoffSeriesPlayoffSeries
}

// GetCommonPlayoffSeries retrieves data from the commonplayoffseries endpoint
func GetCommonPlayoffSeries(ctx context.Context, client *stats.Client, req CommonPlayoffSeriesRequest) (*models.Response[*CommonPlayoffSeriesResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season == "" {
		return nil, fmt.Errorf("Season is required")
	}
	params.Set("Season", string(req.Season))
	if req.SeriesID != nil {
		params.Set("SeriesID", string(*req.SeriesID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonplayoffseries", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonPlayoffSeriesResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayoffSeries = make([]CommonPlayoffSeriesPlayoffSeries, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 5 {
				item := CommonPlayoffSeriesPlayoffSeries{
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
