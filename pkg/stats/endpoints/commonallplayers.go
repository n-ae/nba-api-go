package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// CommonAllPlayersRequest contains parameters for the CommonAllPlayers endpoint
type CommonAllPlayersRequest struct {
	LeagueID            *parameters.LeagueID
	Season              parameters.Season
	IsOnlyCurrentSeason *string
}

// CommonAllPlayersCommonAllPlayers represents the CommonAllPlayers result set for CommonAllPlayers
type CommonAllPlayersCommonAllPlayers struct {
	PERSON_ID                 string `json:"PERSON_ID"`
	DISPLAY_LAST_COMMA_FIRST  string `json:"DISPLAY_LAST_COMMA_FIRST"`
	DISPLAY_FIRST_LAST        string `json:"DISPLAY_FIRST_LAST"`
	ROSTERSTATUS              string `json:"ROSTERSTATUS"`
	FROM_YEAR                 string `json:"FROM_YEAR"`
	TO_YEAR                   string `json:"TO_YEAR"`
	PLAYERCODE                string `json:"PLAYERCODE"`
	TEAM_ID                   int    `json:"TEAM_ID"`
	TEAM_CITY                 string `json:"TEAM_CITY"`
	TEAM_NAME                 string `json:"TEAM_NAME"`
	TEAM_ABBREVIATION         string `json:"TEAM_ABBREVIATION"`
	TEAM_CODE                 string `json:"TEAM_CODE"`
	GAMES_PLAYED_FLAG         string `json:"GAMES_PLAYED_FLAG"`
	OTHERLEAGUE_EXPERIENCE_CH string `json:"OTHERLEAGUE_EXPERIENCE_CH"`
}

// CommonAllPlayersResponse contains the response data from the CommonAllPlayers endpoint
type CommonAllPlayersResponse struct {
	CommonAllPlayers []CommonAllPlayersCommonAllPlayers
}

// GetCommonAllPlayers retrieves data from the commonallplayers endpoint
func GetCommonAllPlayers(ctx context.Context, client *stats.Client, req CommonAllPlayersRequest) (*models.Response[*CommonAllPlayersResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season == "" {
		return nil, fmt.Errorf("%s is required", "Season")
	}
	params.Set("Season", string(req.Season))
	if req.IsOnlyCurrentSeason != nil {
		params.Set("IsOnlyCurrentSeason", *req.IsOnlyCurrentSeason)
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "commonallplayers", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonAllPlayersResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "CommonAllPlayers"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(CommonAllPlayersCommonAllPlayers{})); err != nil {
			return nil, fmt.Errorf("CommonAllPlayers: CommonAllPlayers result set: %w", err)
		}
		response.CommonAllPlayers = make([]CommonAllPlayersCommonAllPlayers, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 14 {
				item := CommonAllPlayersCommonAllPlayers{
					PERSON_ID:                 toString(row[0]),
					DISPLAY_LAST_COMMA_FIRST:  toString(row[1]),
					DISPLAY_FIRST_LAST:        toString(row[2]),
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
