package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// CommonPlayerInfoV2Request contains parameters for the CommonPlayerInfoV2 endpoint
type CommonPlayerInfoV2Request struct {
	PlayerID string
	LeagueID *parameters.LeagueID
}

// CommonPlayerInfoV2CommonPlayerInfo represents the CommonPlayerInfo result set for CommonPlayerInfoV2
type CommonPlayerInfoV2CommonPlayerInfo struct {
	PERSON_ID                        string  `json:"PERSON_ID"`
	FIRST_NAME                       string  `json:"FIRST_NAME"`
	LAST_NAME                        string  `json:"LAST_NAME"`
	DISPLAY_FIRST_LAST               float64 `json:"DISPLAY_FIRST_LAST"`
	DISPLAY_LAST_COMMA_FIRST         float64 `json:"DISPLAY_LAST_COMMA_FIRST"`
	DISPLAY_FI_LAST                  float64 `json:"DISPLAY_FI_LAST"`
	PLAYER_SLUG                      string  `json:"PLAYER_SLUG"`
	BIRTHDATE                        string  `json:"BIRTHDATE"`
	SCHOOL                           string  `json:"SCHOOL"`
	COUNTRY                          string  `json:"COUNTRY"`
	LAST_AFFILIATION                 float64 `json:"LAST_AFFILIATION"`
	HEIGHT                           string  `json:"HEIGHT"`
	WEIGHT                           string  `json:"WEIGHT"`
	SEASON_EXP                       string  `json:"SEASON_EXP"`
	JERSEY                           string  `json:"JERSEY"`
	POSITION                         string  `json:"POSITION"`
	ROSTERSTATUS                     string  `json:"ROSTERSTATUS"`
	GAMES_PLAYED_CURRENT_SEASON_FLAG string  `json:"GAMES_PLAYED_CURRENT_SEASON_FLAG"`
	TEAM_ID                          int     `json:"TEAM_ID"`
	TEAM_NAME                        string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION                string  `json:"TEAM_ABBREVIATION"`
	TEAM_CODE                        string  `json:"TEAM_CODE"`
	TEAM_CITY                        string  `json:"TEAM_CITY"`
	PLAYERCODE                       string  `json:"PLAYERCODE"`
	FROM_YEAR                        string  `json:"FROM_YEAR"`
	TO_YEAR                          string  `json:"TO_YEAR"`
	DLEAGUE_FLAG                     string  `json:"DLEAGUE_FLAG"`
	NBA_FLAG                         string  `json:"NBA_FLAG"`
	GAMES_PLAYED_FLAG                string  `json:"GAMES_PLAYED_FLAG"`
	DRAFT_YEAR                       string  `json:"DRAFT_YEAR"`
	DRAFT_ROUND                      string  `json:"DRAFT_ROUND"`
	DRAFT_NUMBER                     string  `json:"DRAFT_NUMBER"`
	GREATEST_75_FLAG                 string  `json:"GREATEST_75_FLAG"`
}

// CommonPlayerInfoV2PlayerHeadlineStats represents the PlayerHeadlineStats result set for CommonPlayerInfoV2
type CommonPlayerInfoV2PlayerHeadlineStats struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	TimeFrame   string  `json:"TimeFrame"`
	PTS         float64 `json:"PTS"`
	AST         float64 `json:"AST"`
	REB         float64 `json:"REB"`
	PIE         string  `json:"PIE"`
}

// CommonPlayerInfoV2Response contains the response data from the CommonPlayerInfoV2 endpoint
type CommonPlayerInfoV2Response struct {
	CommonPlayerInfo    []CommonPlayerInfoV2CommonPlayerInfo
	PlayerHeadlineStats []CommonPlayerInfoV2PlayerHeadlineStats
}

// GetCommonPlayerInfoV2 retrieves data from the commonplayerinfoV2 endpoint
func GetCommonPlayerInfoV2(ctx context.Context, client *stats.Client, req CommonPlayerInfoV2Request) (*models.Response[*CommonPlayerInfoV2Response], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonplayerinfoV2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonPlayerInfoV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.CommonPlayerInfo = make([]CommonPlayerInfoV2CommonPlayerInfo, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 33 {
				item := CommonPlayerInfoV2CommonPlayerInfo{
					PERSON_ID:                        toString(row[0]),
					FIRST_NAME:                       toString(row[1]),
					LAST_NAME:                        toString(row[2]),
					DISPLAY_FIRST_LAST:               toFloat(row[3]),
					DISPLAY_LAST_COMMA_FIRST:         toFloat(row[4]),
					DISPLAY_FI_LAST:                  toFloat(row[5]),
					PLAYER_SLUG:                      toString(row[6]),
					BIRTHDATE:                        toString(row[7]),
					SCHOOL:                           toString(row[8]),
					COUNTRY:                          toString(row[9]),
					LAST_AFFILIATION:                 toFloat(row[10]),
					HEIGHT:                           toString(row[11]),
					WEIGHT:                           toString(row[12]),
					SEASON_EXP:                       toString(row[13]),
					JERSEY:                           toString(row[14]),
					POSITION:                         toString(row[15]),
					ROSTERSTATUS:                     toString(row[16]),
					GAMES_PLAYED_CURRENT_SEASON_FLAG: toString(row[17]),
					TEAM_ID:                          toInt(row[18]),
					TEAM_NAME:                        toString(row[19]),
					TEAM_ABBREVIATION:                toString(row[20]),
					TEAM_CODE:                        toString(row[21]),
					TEAM_CITY:                        toString(row[22]),
					PLAYERCODE:                       toString(row[23]),
					FROM_YEAR:                        toString(row[24]),
					TO_YEAR:                          toString(row[25]),
					DLEAGUE_FLAG:                     toString(row[26]),
					NBA_FLAG:                         toString(row[27]),
					GAMES_PLAYED_FLAG:                toString(row[28]),
					DRAFT_YEAR:                       toString(row[29]),
					DRAFT_ROUND:                      toString(row[30]),
					DRAFT_NUMBER:                     toString(row[31]),
					GREATEST_75_FLAG:                 toString(row[32]),
				}
				response.CommonPlayerInfo = append(response.CommonPlayerInfo, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.PlayerHeadlineStats = make([]CommonPlayerInfoV2PlayerHeadlineStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 7 {
				item := CommonPlayerInfoV2PlayerHeadlineStats{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					TimeFrame:   toString(row[2]),
					PTS:         toFloat(row[3]),
					AST:         toFloat(row[4]),
					REB:         toFloat(row[5]),
					PIE:         toString(row[6]),
				}
				response.PlayerHeadlineStats = append(response.PlayerHeadlineStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
