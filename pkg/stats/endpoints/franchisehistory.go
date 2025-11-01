package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// FranchiseHistoryRequest contains parameters for the FranchiseHistory endpoint
type FranchiseHistoryRequest struct {
	LeagueID *parameters.LeagueID
}

// FranchiseHistoryFranchiseHistory represents the FranchiseHistory result set for FranchiseHistory
type FranchiseHistoryFranchiseHistory struct {
	LEAGUE_ID      string  `json:"LEAGUE_ID"`
	TEAM_ID        int     `json:"TEAM_ID"`
	TEAM_CITY      string  `json:"TEAM_CITY"`
	TEAM_NAME      string  `json:"TEAM_NAME"`
	START_YEAR     string  `json:"START_YEAR"`
	END_YEAR       string  `json:"END_YEAR"`
	YEARS          string  `json:"YEARS"`
	GAMES          string  `json:"GAMES"`
	WINS           string  `json:"WINS"`
	LOSSES         string  `json:"LOSSES"`
	WIN_PCT        float64 `json:"WIN_PCT"`
	PO_APPEARANCES string  `json:"PO_APPEARANCES"`
	DIV_TITLES     string  `json:"DIV_TITLES"`
	CONF_TITLES    string  `json:"CONF_TITLES"`
	LEAGUE_TITLES  string  `json:"LEAGUE_TITLES"`
}

// FranchiseHistoryDefunctTeams represents the DefunctTeams result set for FranchiseHistory
type FranchiseHistoryDefunctTeams struct {
	LEAGUE_ID      string  `json:"LEAGUE_ID"`
	TEAM_ID        int     `json:"TEAM_ID"`
	TEAM_CITY      string  `json:"TEAM_CITY"`
	TEAM_NAME      string  `json:"TEAM_NAME"`
	START_YEAR     string  `json:"START_YEAR"`
	END_YEAR       string  `json:"END_YEAR"`
	YEARS          string  `json:"YEARS"`
	GAMES          string  `json:"GAMES"`
	WINS           string  `json:"WINS"`
	LOSSES         string  `json:"LOSSES"`
	WIN_PCT        float64 `json:"WIN_PCT"`
	PO_APPEARANCES string  `json:"PO_APPEARANCES"`
	DIV_TITLES     string  `json:"DIV_TITLES"`
	CONF_TITLES    string  `json:"CONF_TITLES"`
	LEAGUE_TITLES  string  `json:"LEAGUE_TITLES"`
}

// FranchiseHistoryResponse contains the response data from the FranchiseHistory endpoint
type FranchiseHistoryResponse struct {
	FranchiseHistory []FranchiseHistoryFranchiseHistory
	DefunctTeams     []FranchiseHistoryDefunctTeams
}

// GetFranchiseHistory retrieves data from the franchisehistory endpoint
func GetFranchiseHistory(ctx context.Context, client *stats.Client, req FranchiseHistoryRequest) (*models.Response[*FranchiseHistoryResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/franchisehistory", params, &rawResp); err != nil {
		return nil, err
	}

	response := &FranchiseHistoryResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.FranchiseHistory = make([]FranchiseHistoryFranchiseHistory, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 15 {
				item := FranchiseHistoryFranchiseHistory{
					LEAGUE_ID:      toString(row[0]),
					TEAM_ID:        toInt(row[1]),
					TEAM_CITY:      toString(row[2]),
					TEAM_NAME:      toString(row[3]),
					START_YEAR:     toString(row[4]),
					END_YEAR:       toString(row[5]),
					YEARS:          toString(row[6]),
					GAMES:          toString(row[7]),
					WINS:           toString(row[8]),
					LOSSES:         toString(row[9]),
					WIN_PCT:        toFloat(row[10]),
					PO_APPEARANCES: toString(row[11]),
					DIV_TITLES:     toString(row[12]),
					CONF_TITLES:    toString(row[13]),
					LEAGUE_TITLES:  toString(row[14]),
				}
				response.FranchiseHistory = append(response.FranchiseHistory, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.DefunctTeams = make([]FranchiseHistoryDefunctTeams, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 15 {
				item := FranchiseHistoryDefunctTeams{
					LEAGUE_ID:      toString(row[0]),
					TEAM_ID:        toInt(row[1]),
					TEAM_CITY:      toString(row[2]),
					TEAM_NAME:      toString(row[3]),
					START_YEAR:     toString(row[4]),
					END_YEAR:       toString(row[5]),
					YEARS:          toString(row[6]),
					GAMES:          toString(row[7]),
					WINS:           toString(row[8]),
					LOSSES:         toString(row[9]),
					WIN_PCT:        toFloat(row[10]),
					PO_APPEARANCES: toString(row[11]),
					DIV_TITLES:     toString(row[12]),
					CONF_TITLES:    toString(row[13]),
					LEAGUE_TITLES:  toString(row[14]),
				}
				response.DefunctTeams = append(response.DefunctTeams, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
