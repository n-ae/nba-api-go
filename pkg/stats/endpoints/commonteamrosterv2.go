package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// CommonTeamRosterV2Request contains parameters for the CommonTeamRosterV2 endpoint
type CommonTeamRosterV2Request struct {
	TeamID   string
	Season   *parameters.Season
	LeagueID *parameters.LeagueID
}

// CommonTeamRosterV2CommonTeamRoster represents the CommonTeamRoster result set for CommonTeamRosterV2
type CommonTeamRosterV2CommonTeamRoster struct {
	TeamID       string `json:"TeamID"`
	SEASON       string `json:"SEASON"`
	LeagueID     string `json:"LeagueID"`
	PLAYER       string `json:"PLAYER"`
	NICKNAME     string `json:"NICKNAME"`
	PLAYER_SLUG  string `json:"PLAYER_SLUG"`
	NUM          string `json:"NUM"`
	POSITION     string `json:"POSITION"`
	HEIGHT       string `json:"HEIGHT"`
	WEIGHT       string `json:"WEIGHT"`
	BIRTH_DATE   string `json:"BIRTH_DATE"`
	AGE          int    `json:"AGE"`
	EXP          string `json:"EXP"`
	SCHOOL       string `json:"SCHOOL"`
	PLAYER_ID    int    `json:"PLAYER_ID"`
	HOW_ACQUIRED string `json:"HOW_ACQUIRED"`
}

// CommonTeamRosterV2Coaches represents the Coaches result set for CommonTeamRosterV2
type CommonTeamRosterV2Coaches struct {
	TEAM_ID       int    `json:"TEAM_ID"`
	SEASON        string `json:"SEASON"`
	COACH_ID      string `json:"COACH_ID"`
	FIRST_NAME    string `json:"FIRST_NAME"`
	LAST_NAME     string `json:"LAST_NAME"`
	COACH_NAME    string `json:"COACH_NAME"`
	COACH_CODE    string `json:"COACH_CODE"`
	IS_ASSISTANT  string `json:"IS_ASSISTANT"`
	COACH_TYPE    string `json:"COACH_TYPE"`
	SCHOOL        string `json:"SCHOOL"`
	SORT_SEQUENCE int    `json:"SORT_SEQUENCE"`
}

// CommonTeamRosterV2Response contains the response data from the CommonTeamRosterV2 endpoint
type CommonTeamRosterV2Response struct {
	CommonTeamRoster []CommonTeamRosterV2CommonTeamRoster
	Coaches          []CommonTeamRosterV2Coaches
}

// GetCommonTeamRosterV2 retrieves data from the commonteamrosterv2 endpoint
func GetCommonTeamRosterV2(ctx context.Context, client *stats.Client, req CommonTeamRosterV2Request) (*models.Response[*CommonTeamRosterV2Response], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonteamrosterv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonTeamRosterV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.CommonTeamRoster = make([]CommonTeamRosterV2CommonTeamRoster, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 16 {
				item := CommonTeamRosterV2CommonTeamRoster{
					TeamID:       toString(row[0]),
					SEASON:       toString(row[1]),
					LeagueID:     toString(row[2]),
					PLAYER:       toString(row[3]),
					NICKNAME:     toString(row[4]),
					PLAYER_SLUG:  toString(row[5]),
					NUM:          toString(row[6]),
					POSITION:     toString(row[7]),
					HEIGHT:       toString(row[8]),
					WEIGHT:       toString(row[9]),
					BIRTH_DATE:   toString(row[10]),
					AGE:          toInt(row[11]),
					EXP:          toString(row[12]),
					SCHOOL:       toString(row[13]),
					PLAYER_ID:    toInt(row[14]),
					HOW_ACQUIRED: toString(row[15]),
				}
				response.CommonTeamRoster = append(response.CommonTeamRoster, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.Coaches = make([]CommonTeamRosterV2Coaches, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				item := CommonTeamRosterV2Coaches{
					TEAM_ID:       toInt(row[0]),
					SEASON:        toString(row[1]),
					COACH_ID:      toString(row[2]),
					FIRST_NAME:    toString(row[3]),
					LAST_NAME:     toString(row[4]),
					COACH_NAME:    toString(row[5]),
					COACH_CODE:    toString(row[6]),
					IS_ASSISTANT:  toString(row[7]),
					COACH_TYPE:    toString(row[8]),
					SCHOOL:        toString(row[9]),
					SORT_SEQUENCE: toInt(row[10]),
				}
				response.Coaches = append(response.Coaches, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
