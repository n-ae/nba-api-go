package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// TeamDetailsRequest contains parameters for the TeamDetails endpoint
type TeamDetailsRequest struct {
	TeamID string
}

// TeamDetailsTeamBackground represents the TeamBackground result set for TeamDetails
type TeamDetailsTeamBackground struct {
	TEAM_ID            int    `json:"TEAM_ID"`
	ABBREVIATION       string `json:"ABBREVIATION"`
	NICKNAME           string `json:"NICKNAME"`
	YEARFOUNDED        string `json:"YEARFOUNDED"`
	CITY               string `json:"CITY"`
	ARENA              string `json:"ARENA"`
	ARENACAPACITY      string `json:"ARENACAPACITY"`
	OWNER              string `json:"OWNER"`
	GENERALMANAGER     int    `json:"GENERALMANAGER"`
	HEADCOACH          string `json:"HEADCOACH"`
	DLEAGUEAFFILIATION string `json:"DLEAGUEAFFILIATION"`
}

// TeamDetailsTeamHistory represents the TeamHistory result set for TeamDetails
type TeamDetailsTeamHistory struct {
	TEAM_ID        int    `json:"TEAM_ID"`
	CITY           string `json:"CITY"`
	NICKNAME       string `json:"NICKNAME"`
	YEARFOUNDED    string `json:"YEARFOUNDED"`
	YEARACTIVETILL string `json:"YEARACTIVETILL"`
}

// TeamDetailsTeamSocialSites represents the TeamSocialSites result set for TeamDetails
type TeamDetailsTeamSocialSites struct {
	ACCOUNTTYPE  string `json:"ACCOUNTTYPE"`
	WEBSITE_LINK string `json:"WEBSITE_LINK"`
}

// TeamDetailsTeamAwardsChampionships represents the TeamAwardsChampionships result set for TeamDetails
type TeamDetailsTeamAwardsChampionships struct {
	YEARAWARDED  string `json:"YEARAWARDED"`
	OPPOSITETEAM string `json:"OPPOSITETEAM"`
}

// TeamDetailsTeamAwardsConf represents the TeamAwardsConf result set for TeamDetails
type TeamDetailsTeamAwardsConf struct {
	YEARAWARDED  string `json:"YEARAWARDED"`
	OPPOSITETEAM string `json:"OPPOSITETEAM"`
}

// TeamDetailsTeamAwardsDiv represents the TeamAwardsDiv result set for TeamDetails
type TeamDetailsTeamAwardsDiv struct {
	YEARAWARDED  string `json:"YEARAWARDED"`
	OPPOSITETEAM string `json:"OPPOSITETEAM"`
}

// TeamDetailsTeamHof represents the TeamHof result set for TeamDetails
type TeamDetailsTeamHof struct {
	PLAYERID        string `json:"PLAYERID"`
	PLAYER          string `json:"PLAYER"`
	POSITION        string `json:"POSITION"`
	JERSEY          string `json:"JERSEY"`
	SEASONSWITHTEAM string `json:"SEASONSWITHTEAM"`
	YEAR            string `json:"YEAR"`
}

// TeamDetailsTeamRetired represents the TeamRetired result set for TeamDetails
type TeamDetailsTeamRetired struct {
	PLAYERID        string `json:"PLAYERID"`
	PLAYER          string `json:"PLAYER"`
	POSITION        string `json:"POSITION"`
	JERSEY          string `json:"JERSEY"`
	SEASONSWITHTEAM string `json:"SEASONSWITHTEAM"`
	YEAR            string `json:"YEAR"`
}

// TeamDetailsResponse contains the response data from the TeamDetails endpoint
type TeamDetailsResponse struct {
	TeamBackground          []TeamDetailsTeamBackground
	TeamHistory             []TeamDetailsTeamHistory
	TeamSocialSites         []TeamDetailsTeamSocialSites
	TeamAwardsChampionships []TeamDetailsTeamAwardsChampionships
	TeamAwardsConf          []TeamDetailsTeamAwardsConf
	TeamAwardsDiv           []TeamDetailsTeamAwardsDiv
	TeamHof                 []TeamDetailsTeamHof
	TeamRetired             []TeamDetailsTeamRetired
}

// GetTeamDetails retrieves data from the teamdetails endpoint
func GetTeamDetails(ctx context.Context, client *stats.Client, req TeamDetailsRequest) (*models.Response[*TeamDetailsResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "teamdetails", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDetailsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamBackground = make([]TeamDetailsTeamBackground, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 11 {
				item := TeamDetailsTeamBackground{
					TEAM_ID:            toInt(row[0]),
					ABBREVIATION:       toString(row[1]),
					NICKNAME:           toString(row[2]),
					YEARFOUNDED:        toString(row[3]),
					CITY:               toString(row[4]),
					ARENA:              toString(row[5]),
					ARENACAPACITY:      toString(row[6]),
					OWNER:              toString(row[7]),
					GENERALMANAGER:     toInt(row[8]),
					HEADCOACH:          toString(row[9]),
					DLEAGUEAFFILIATION: toString(row[10]),
				}
				response.TeamBackground = append(response.TeamBackground, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamHistory = make([]TeamDetailsTeamHistory, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 5 {
				item := TeamDetailsTeamHistory{
					TEAM_ID:        toInt(row[0]),
					CITY:           toString(row[1]),
					NICKNAME:       toString(row[2]),
					YEARFOUNDED:    toString(row[3]),
					YEARACTIVETILL: toString(row[4]),
				}
				response.TeamHistory = append(response.TeamHistory, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.TeamSocialSites = make([]TeamDetailsTeamSocialSites, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamSocialSites{
					ACCOUNTTYPE:  toString(row[0]),
					WEBSITE_LINK: toString(row[1]),
				}
				response.TeamSocialSites = append(response.TeamSocialSites, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.TeamAwardsChampionships = make([]TeamDetailsTeamAwardsChampionships, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamAwardsChampionships{
					YEARAWARDED:  toString(row[0]),
					OPPOSITETEAM: toString(row[1]),
				}
				response.TeamAwardsChampionships = append(response.TeamAwardsChampionships, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.TeamAwardsConf = make([]TeamDetailsTeamAwardsConf, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamAwardsConf{
					YEARAWARDED:  toString(row[0]),
					OPPOSITETEAM: toString(row[1]),
				}
				response.TeamAwardsConf = append(response.TeamAwardsConf, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 5 {
		response.TeamAwardsDiv = make([]TeamDetailsTeamAwardsDiv, 0, len(rawResp.ResultSets[5].RowSet))
		for _, row := range rawResp.ResultSets[5].RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamAwardsDiv{
					YEARAWARDED:  toString(row[0]),
					OPPOSITETEAM: toString(row[1]),
				}
				response.TeamAwardsDiv = append(response.TeamAwardsDiv, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 6 {
		response.TeamHof = make([]TeamDetailsTeamHof, 0, len(rawResp.ResultSets[6].RowSet))
		for _, row := range rawResp.ResultSets[6].RowSet {
			if len(row) >= 6 {
				item := TeamDetailsTeamHof{
					PLAYERID:        toString(row[0]),
					PLAYER:          toString(row[1]),
					POSITION:        toString(row[2]),
					JERSEY:          toString(row[3]),
					SEASONSWITHTEAM: toString(row[4]),
					YEAR:            toString(row[5]),
				}
				response.TeamHof = append(response.TeamHof, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 7 {
		response.TeamRetired = make([]TeamDetailsTeamRetired, 0, len(rawResp.ResultSets[7].RowSet))
		for _, row := range rawResp.ResultSets[7].RowSet {
			if len(row) >= 6 {
				item := TeamDetailsTeamRetired{
					PLAYERID:        toString(row[0]),
					PLAYER:          toString(row[1]),
					POSITION:        toString(row[2]),
					JERSEY:          toString(row[3]),
					SEASONSWITHTEAM: toString(row[4]),
					YEAR:            toString(row[5]),
				}
				response.TeamRetired = append(response.TeamRetired, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
