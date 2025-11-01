package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashOppPtShotRequest contains parameters for the LeagueDashOppPtShot endpoint
type LeagueDashOppPtShotRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueDashOppPtShotLeagueDashOppPtShot represents the LeagueDashOppPtShot result set for LeagueDashOppPtShot
type LeagueDashOppPtShotLeagueDashOppPtShot struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	G                 string  `json:"G"`
	FGA_FREQUENCY     float64 `json:"FGA_FREQUENCY"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	EFG_PCT           float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY    string  `json:"FG2A_FREQUENCY"`
	FG2M              string  `json:"FG2M"`
	FG2A              string  `json:"FG2A"`
	FG2_PCT           float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY    float64 `json:"FG3A_FREQUENCY"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
}

// LeagueDashOppPtShotResponse contains the response data from the LeagueDashOppPtShot endpoint
type LeagueDashOppPtShotResponse struct {
	LeagueDashOppPtShot []LeagueDashOppPtShotLeagueDashOppPtShot
}

// GetLeagueDashOppPtShot retrieves data from the leaguedashoppptshot endpoint
func GetLeagueDashOppPtShot(ctx context.Context, client *stats.Client, req LeagueDashOppPtShotRequest) (*models.Response[*LeagueDashOppPtShotResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguedashoppptshot", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashOppPtShotResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashOppPtShot = make([]LeagueDashOppPtShotLeagueDashOppPtShot, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 18 {
				item := LeagueDashOppPtShotLeagueDashOppPtShot{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					GP:                toInt(row[3]),
					G:                 toString(row[4]),
					FGA_FREQUENCY:     toFloat(row[5]),
					FGM:               toInt(row[6]),
					FGA:               toInt(row[7]),
					FG_PCT:            toFloat(row[8]),
					EFG_PCT:           toFloat(row[9]),
					FG2A_FREQUENCY:    toString(row[10]),
					FG2M:              toString(row[11]),
					FG2A:              toString(row[12]),
					FG2_PCT:           toFloat(row[13]),
					FG3A_FREQUENCY:    toFloat(row[14]),
					FG3M:              toInt(row[15]),
					FG3A:              toInt(row[16]),
					FG3_PCT:           toFloat(row[17]),
				}
				response.LeagueDashOppPtShot = append(response.LeagueDashOppPtShot, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
