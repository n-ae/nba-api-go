package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// CommonAllPlayersV2Request contains parameters for the CommonAllPlayersV2 endpoint
type CommonAllPlayersV2Request struct {
	LeagueID            *parameters.LeagueID
	Season              *parameters.Season
	IsOnlyCurrentSeason *string
}

// CommonAllPlayersV2CommonAllPlayers represents the CommonAllPlayers result set for CommonAllPlayersV2
type CommonAllPlayersV2CommonAllPlayers struct {
	PERSON_ID                 string  `json:"PERSON_ID"`
	DISPLAY_LAST_COMMA_FIRST  float64 `json:"DISPLAY_LAST_COMMA_FIRST"`
	DISPLAY_FIRST_LAST        float64 `json:"DISPLAY_FIRST_LAST"`
	ROSTERSTATUS              string  `json:"ROSTERSTATUS"`
	FROM_YEAR                 string  `json:"FROM_YEAR"`
	TO_YEAR                   string  `json:"TO_YEAR"`
	PLAYERCODE                string  `json:"PLAYERCODE"`
	TEAM_ID                   int     `json:"TEAM_ID"`
	TEAM_CITY                 string  `json:"TEAM_CITY"`
	TEAM_NAME                 string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION         string  `json:"TEAM_ABBREVIATION"`
	TEAM_CODE                 string  `json:"TEAM_CODE"`
	GAMES_PLAYED_FLAG         string  `json:"GAMES_PLAYED_FLAG"`
	OTHERLEAGUE_EXPERIENCE_CH string  `json:"OTHERLEAGUE_EXPERIENCE_CH"`
}

// CommonAllPlayersV2Response contains the response data from the CommonAllPlayersV2 endpoint
type CommonAllPlayersV2Response struct {
	CommonAllPlayers []CommonAllPlayersV2CommonAllPlayers
}

// GetCommonAllPlayersV2 retrieves data from the commonallplayersv2 endpoint
func GetCommonAllPlayersV2(ctx context.Context, client *stats.Client, req CommonAllPlayersV2Request) (*models.Response[*CommonAllPlayersV2Response], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.IsOnlyCurrentSeason != nil {
		params.Set("IsOnlyCurrentSeason", string(*req.IsOnlyCurrentSeason))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonallplayersv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonAllPlayersV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.CommonAllPlayers = make([]CommonAllPlayersV2CommonAllPlayers, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := CommonAllPlayersV2CommonAllPlayers{
					PERSON_ID:                 toString(row[0]),
					DISPLAY_LAST_COMMA_FIRST:  toFloat(row[1]),
					DISPLAY_FIRST_LAST:        toFloat(row[2]),
					ROSTERSTATUS:              toString(row[3]),
					FROM_YEAR:                 toString(row[4]),
					TO_YEAR:                   toString(row[5]),
					PLAYERCODE:                toString(row[6]),
					TEAM_ID:                   toInt(row[7]),
					TEAM_CITY:                 toString(row[8]),
					TEAM_NAME:                 toString(row[9]),
					TEAM_ABBREVIATION:         toString(row[10]),
					TEAM_CODE:                 toString(row[11]),
					GAMES_PLAYED_FLAG:         toString(row[12]),
					OTHERLEAGUE_EXPERIENCE_CH: toString(row[13]),
				}
				response.CommonAllPlayers = append(response.CommonAllPlayers, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
