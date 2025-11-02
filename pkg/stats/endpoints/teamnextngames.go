package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamNextNGamesRequest contains parameters for the TeamNextNGames endpoint
type TeamNextNGamesRequest struct {
	TeamID        string
	Season        *parameters.Season
	SeasonType    *parameters.SeasonType
	NumberOfGames *string
	LeagueID      *parameters.LeagueID
}

// TeamNextNGamesNextNGames represents the NextNGames result set for TeamNextNGames
type TeamNextNGamesNextNGames struct {
	GAME_ID           string `json:"GAME_ID"`
	GAME_DATE         string `json:"GAME_DATE"`
	HOME_TEAM_ID      int    `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID   int    `json:"VISITOR_TEAM_ID"`
	HOME_TEAM_NAME    string `json:"HOME_TEAM_NAME"`
	VISITOR_TEAM_NAME string `json:"VISITOR_TEAM_NAME"`
}

// TeamNextNGamesResponse contains the response data from the TeamNextNGames endpoint
type TeamNextNGamesResponse struct {
	NextNGames []TeamNextNGamesNextNGames
}

// GetTeamNextNGames retrieves data from the teamnextnGames endpoint
func GetTeamNextNGames(ctx context.Context, client *stats.Client, req TeamNextNGamesRequest) (*models.Response[*TeamNextNGamesResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.NumberOfGames != nil {
		params.Set("NumberOfGames", string(*req.NumberOfGames))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamnextnGames", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamNextNGamesResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.NextNGames = make([]TeamNextNGamesNextNGames, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 6 {
				item := TeamNextNGamesNextNGames{
					GAME_ID:           toString(row[0]),
					GAME_DATE:         toString(row[1]),
					HOME_TEAM_ID:      toInt(row[2]),
					VISITOR_TEAM_ID:   toInt(row[3]),
					HOME_TEAM_NAME:    toString(row[4]),
					VISITOR_TEAM_NAME: toString(row[5]),
				}
				response.NextNGames = append(response.NextNGames, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
