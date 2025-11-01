package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeaguePlayerOnDetailsRequest contains parameters for the LeaguePlayerOnDetails endpoint
type LeaguePlayerOnDetailsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
	TeamID     *string
	PlayerID   *string
}

// LeaguePlayerOnDetailsLeaguePlayerOnDetails represents the LeaguePlayerOnDetails result set for LeaguePlayerOnDetails
type LeaguePlayerOnDetailsLeaguePlayerOnDetails struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	VS_PLAYER_ID      int     `json:"VS_PLAYER_ID"`
	VS_PLAYER_NAME    string  `json:"VS_PLAYER_NAME"`
	COURT_STATUS      string  `json:"COURT_STATUS"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
	NET_RATING        string  `json:"NET_RATING"`
	OFF_RATING        string  `json:"OFF_RATING"`
	DEF_RATING        string  `json:"DEF_RATING"`
}

// LeaguePlayerOnDetailsResponse contains the response data from the LeaguePlayerOnDetails endpoint
type LeaguePlayerOnDetailsResponse struct {
	LeaguePlayerOnDetails []LeaguePlayerOnDetailsLeaguePlayerOnDetails
}

// GetLeaguePlayerOnDetails retrieves data from the leagueplayerondetails endpoint
func GetLeaguePlayerOnDetails(ctx context.Context, client *stats.Client, req LeaguePlayerOnDetailsRequest) (*models.Response[*LeaguePlayerOnDetailsResponse], error) {
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
	if req.TeamID != nil {
		params.Set("TeamID", string(*req.TeamID))
	}
	if req.PlayerID != nil {
		params.Set("PlayerID", string(*req.PlayerID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leagueplayerondetails", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeaguePlayerOnDetailsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeaguePlayerOnDetails = make([]LeaguePlayerOnDetailsLeaguePlayerOnDetails, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				item := LeaguePlayerOnDetailsLeaguePlayerOnDetails{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					VS_PLAYER_ID:      toInt(row[3]),
					VS_PLAYER_NAME:    toString(row[4]),
					COURT_STATUS:      toString(row[5]),
					GP:                toInt(row[6]),
					MIN:               toFloat(row[7]),
					PLUS_MINUS:        toFloat(row[8]),
					NET_RATING:        toString(row[9]),
					OFF_RATING:        toString(row[10]),
					DEF_RATING:        toString(row[11]),
				}
				response.LeaguePlayerOnDetails = append(response.LeaguePlayerOnDetails, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
