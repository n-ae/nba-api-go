package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
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
	GENERALMANAGER     string `json:"GENERALMANAGER"`
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
		return nil, fmt.Errorf("%s is required", "TeamID")
	}
	params.Set("TeamID", req.TeamID)

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "teamdetails", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDetailsResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamBackground"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamBackground{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamBackground result set: %w", err)
		}
		response.TeamBackground = make([]TeamDetailsTeamBackground, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
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
					GENERALMANAGER:     toString(row[8]),
					HEADCOACH:          toString(row[9]),
					DLEAGUEAFFILIATION: toString(row[10]),
				}
				response.TeamBackground = append(response.TeamBackground, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamHistory"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamHistory{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamHistory result set: %w", err)
		}
		response.TeamHistory = make([]TeamDetailsTeamHistory, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
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
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamSocialSites"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamSocialSites{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamSocialSites result set: %w", err)
		}
		response.TeamSocialSites = make([]TeamDetailsTeamSocialSites, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamSocialSites{
					ACCOUNTTYPE:  toString(row[0]),
					WEBSITE_LINK: toString(row[1]),
				}
				response.TeamSocialSites = append(response.TeamSocialSites, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamAwardsChampionships"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamAwardsChampionships{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamAwardsChampionships result set: %w", err)
		}
		response.TeamAwardsChampionships = make([]TeamDetailsTeamAwardsChampionships, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamAwardsChampionships{
					YEARAWARDED:  toString(row[0]),
					OPPOSITETEAM: toString(row[1]),
				}
				response.TeamAwardsChampionships = append(response.TeamAwardsChampionships, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamAwardsConf"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamAwardsConf{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamAwardsConf result set: %w", err)
		}
		response.TeamAwardsConf = make([]TeamDetailsTeamAwardsConf, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamAwardsConf{
					YEARAWARDED:  toString(row[0]),
					OPPOSITETEAM: toString(row[1]),
				}
				response.TeamAwardsConf = append(response.TeamAwardsConf, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamAwardsDiv"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamAwardsDiv{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamAwardsDiv result set: %w", err)
		}
		response.TeamAwardsDiv = make([]TeamDetailsTeamAwardsDiv, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 2 {
				item := TeamDetailsTeamAwardsDiv{
					YEARAWARDED:  toString(row[0]),
					OPPOSITETEAM: toString(row[1]),
				}
				response.TeamAwardsDiv = append(response.TeamAwardsDiv, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamHof"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamHof{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamHof result set: %w", err)
		}
		response.TeamHof = make([]TeamDetailsTeamHof, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
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
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamRetired"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDetailsTeamRetired{})); err != nil {
			return nil, fmt.Errorf("TeamDetails: TeamRetired result set: %w", err)
		}
		response.TeamRetired = make([]TeamDetailsTeamRetired, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
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
