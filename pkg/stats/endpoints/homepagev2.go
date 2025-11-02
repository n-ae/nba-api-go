package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// HomepageV2Request contains parameters for the HomepageV2 endpoint
type HomepageV2Request struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// HomepageV2GameHeader represents the GameHeader result set for HomepageV2
type HomepageV2GameHeader struct {
	GAME_ID                   string `json:"GAME_ID"`
	GAME_DATE                 string `json:"GAME_DATE"`
	HOME_TEAM_ID              int    `json:"HOME_TEAM_ID"`
	HOME_TEAM_NAME            string `json:"HOME_TEAM_NAME"`
	HOME_TEAM_ABBREVIATION    string `json:"HOME_TEAM_ABBREVIATION"`
	HOME_TEAM_SCORE           string `json:"HOME_TEAM_SCORE"`
	VISITOR_TEAM_ID           int    `json:"VISITOR_TEAM_ID"`
	VISITOR_TEAM_NAME         string `json:"VISITOR_TEAM_NAME"`
	VISITOR_TEAM_ABBREVIATION string `json:"VISITOR_TEAM_ABBREVIATION"`
	VISITOR_TEAM_SCORE        string `json:"VISITOR_TEAM_SCORE"`
	GAME_STATUS_TEXT          string `json:"GAME_STATUS_TEXT"`
}

// HomepageV2Response contains the response data from the HomepageV2 endpoint
type HomepageV2Response struct {
	GameHeader []HomepageV2GameHeader
}

// GetHomepageV2 retrieves data from the homepagev2 endpoint
func GetHomepageV2(ctx context.Context, client *stats.Client, req HomepageV2Request) (*models.Response[*HomepageV2Response], error) {
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
	if err := client.GetJSON(ctx, "/homepagev2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &HomepageV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.GameHeader = make([]HomepageV2GameHeader, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 11 {
				item := HomepageV2GameHeader{
					GAME_ID:                   toString(row[0]),
					GAME_DATE:                 toString(row[1]),
					HOME_TEAM_ID:              toInt(row[2]),
					HOME_TEAM_NAME:            toString(row[3]),
					HOME_TEAM_ABBREVIATION:    toString(row[4]),
					HOME_TEAM_SCORE:           toString(row[5]),
					VISITOR_TEAM_ID:           toInt(row[6]),
					VISITOR_TEAM_NAME:         toString(row[7]),
					VISITOR_TEAM_ABBREVIATION: toString(row[8]),
					VISITOR_TEAM_SCORE:        toString(row[9]),
					GAME_STATUS_TEXT:          toString(row[10]),
				}
				response.GameHeader = append(response.GameHeader, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
