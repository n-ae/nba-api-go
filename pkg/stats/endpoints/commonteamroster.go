package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// CommonTeamRosterRequest contains parameters for the CommonTeamRoster endpoint
type CommonTeamRosterRequest struct {
	TeamID string
	Season *parameters.Season
	LeagueID *parameters.LeagueID
}


// CommonTeamRosterCommonTeamRoster represents the CommonTeamRoster result set for CommonTeamRoster
type CommonTeamRosterCommonTeamRoster struct {
	TeamID string `json:"TeamID"`
	SEASON string `json:"SEASON"`
	LeagueID string `json:"LeagueID"`
	PLAYER string `json:"PLAYER"`
	NICKNAME string `json:"NICKNAME"`
	PLAYER_SLUG string `json:"PLAYER_SLUG"`
	NUM string `json:"NUM"`
	POSITION string `json:"POSITION"`
	HEIGHT string `json:"HEIGHT"`
	WEIGHT string `json:"WEIGHT"`
	BIRTH_DATE string `json:"BIRTH_DATE"`
	AGE int `json:"AGE"`
	EXP string `json:"EXP"`
	SCHOOL string `json:"SCHOOL"`
	PLAYER_ID int `json:"PLAYER_ID"`
	HOW_ACQUIRED string `json:"HOW_ACQUIRED"`
}

// CommonTeamRosterCoaches represents the Coaches result set for CommonTeamRoster
type CommonTeamRosterCoaches struct {
	TEAM_ID int `json:"TEAM_ID"`
	SEASON string `json:"SEASON"`
	COACH_ID string `json:"COACH_ID"`
	FIRST_NAME string `json:"FIRST_NAME"`
	LAST_NAME string `json:"LAST_NAME"`
	COACH_NAME string `json:"COACH_NAME"`
	COACH_CODE string `json:"COACH_CODE"`
	IS_ASSISTANT string `json:"IS_ASSISTANT"`
	COACH_TYPE string `json:"COACH_TYPE"`
	SCHOOL string `json:"SCHOOL"`
	SORT_SEQUENCE int `json:"SORT_SEQUENCE"`
}


// CommonTeamRosterResponse contains the response data from the CommonTeamRoster endpoint
type CommonTeamRosterResponse struct {
	CommonTeamRoster []CommonTeamRosterCommonTeamRoster
	Coaches []CommonTeamRosterCoaches
}

// GetCommonTeamRoster retrieves data from the commonteamroster endpoint
func GetCommonTeamRoster(ctx context.Context, client *stats.Client, req CommonTeamRosterRequest) (*models.Response[*CommonTeamRosterResponse], error) {
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
	if err := client.GetJSON(ctx, "/commonteamroster", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonTeamRosterResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.CommonTeamRoster = make([]CommonTeamRosterCommonTeamRoster, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 16 {
				item := CommonTeamRosterCommonTeamRoster{
					TeamID: toString(row[0]),
					SEASON: toString(row[1]),
					LeagueID: toString(row[2]),
					PLAYER: toString(row[3]),
					NICKNAME: toString(row[4]),
					PLAYER_SLUG: toString(row[5]),
					NUM: toString(row[6]),
					POSITION: toString(row[7]),
					HEIGHT: toString(row[8]),
					WEIGHT: toString(row[9]),
					BIRTH_DATE: toString(row[10]),
					AGE: toInt(row[11]),
					EXP: toString(row[12]),
					SCHOOL: toString(row[13]),
					PLAYER_ID: toInt(row[14]),
					HOW_ACQUIRED: toString(row[15]),
				}
				response.CommonTeamRoster = append(response.CommonTeamRoster, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.Coaches = make([]CommonTeamRosterCoaches, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				item := CommonTeamRosterCoaches{
					TEAM_ID: toInt(row[0]),
					SEASON: toString(row[1]),
					COACH_ID: toString(row[2]),
					FIRST_NAME: toString(row[3]),
					LAST_NAME: toString(row[4]),
					COACH_NAME: toString(row[5]),
					COACH_CODE: toString(row[6]),
					IS_ASSISTANT: toString(row[7]),
					COACH_TYPE: toString(row[8]),
					SCHOOL: toString(row[9]),
					SORT_SEQUENCE: toInt(row[10]),
				}
				response.Coaches = append(response.Coaches, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
