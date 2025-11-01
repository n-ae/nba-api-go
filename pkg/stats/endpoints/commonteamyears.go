package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// CommonTeamYearsRequest contains parameters for the CommonTeamYears endpoint
type CommonTeamYearsRequest struct {
	LeagueID *parameters.LeagueID
}

// CommonTeamYearsTeamYears represents the TeamYears result set for CommonTeamYears
type CommonTeamYearsTeamYears struct {
	LEAGUE_ID    string  `json:"LEAGUE_ID"`
	TEAM_ID      int     `json:"TEAM_ID"`
	MIN_YEAR     float64 `json:"MIN_YEAR"`
	MAX_YEAR     string  `json:"MAX_YEAR"`
	ABBREVIATION string  `json:"ABBREVIATION"`
}

// CommonTeamYearsResponse contains the response data from the CommonTeamYears endpoint
type CommonTeamYearsResponse struct {
	TeamYears []CommonTeamYearsTeamYears
}

// GetCommonTeamYears retrieves data from the commonteamyears endpoint
func GetCommonTeamYears(ctx context.Context, client *stats.Client, req CommonTeamYearsRequest) (*models.Response[*CommonTeamYearsResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonteamyears", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonTeamYearsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamYears = make([]CommonTeamYearsTeamYears, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 5 {
				item := CommonTeamYearsTeamYears{
					LEAGUE_ID:    toString(row[0]),
					TEAM_ID:      toInt(row[1]),
					MIN_YEAR:     toFloat(row[2]),
					MAX_YEAR:     toString(row[3]),
					ABBREVIATION: toString(row[4]),
				}
				response.TeamYears = append(response.TeamYears, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
