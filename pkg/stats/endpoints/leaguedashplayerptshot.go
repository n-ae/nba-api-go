package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPlayerPtShotRequest contains parameters for the LeagueDashPlayerPtShot endpoint
type LeagueDashPlayerPtShotRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueDashPlayerPtShotLeagueDashPlayerPtShot represents the LeagueDashPlayerPtShot result set for LeagueDashPlayerPtShot
type LeagueDashPlayerPtShotLeagueDashPlayerPtShot struct {
	PLAYER_ID                     int     `json:"PLAYER_ID"`
	PLAYER_NAME                   string  `json:"PLAYER_NAME"`
	PLAYER_LAST_TEAM_ID           int     `json:"PLAYER_LAST_TEAM_ID"`
	PLAYER_LAST_TEAM_ABBREVIATION string  `json:"PLAYER_LAST_TEAM_ABBREVIATION"`
	AGE                           int     `json:"AGE"`
	GP                            int     `json:"GP"`
	G                             string  `json:"G"`
	FGA_FREQUENCY                 float64 `json:"FGA_FREQUENCY"`
	FGM                           int     `json:"FGM"`
	FGA                           int     `json:"FGA"`
	FG_PCT                        float64 `json:"FG_PCT"`
	EFG_PCT                       float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY                string  `json:"FG2A_FREQUENCY"`
	FG2M                          string  `json:"FG2M"`
	FG2A                          string  `json:"FG2A"`
	FG2_PCT                       float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY                float64 `json:"FG3A_FREQUENCY"`
	FG3M                          int     `json:"FG3M"`
	FG3A                          int     `json:"FG3A"`
	FG3_PCT                       float64 `json:"FG3_PCT"`
}

// LeagueDashPlayerPtShotResponse contains the response data from the LeagueDashPlayerPtShot endpoint
type LeagueDashPlayerPtShotResponse struct {
	LeagueDashPlayerPtShot []LeagueDashPlayerPtShotLeagueDashPlayerPtShot
}

// GetLeagueDashPlayerPtShot retrieves data from the leaguedashplayerptshot endpoint
func GetLeagueDashPlayerPtShot(ctx context.Context, client *stats.Client, req LeagueDashPlayerPtShotRequest) (*models.Response[*LeagueDashPlayerPtShotResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguedashplayerptshot", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPlayerPtShotResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPlayerPtShot = make([]LeagueDashPlayerPtShotLeagueDashPlayerPtShot, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 20 {
				item := LeagueDashPlayerPtShotLeagueDashPlayerPtShot{
					PLAYER_ID:                     toInt(row[0]),
					PLAYER_NAME:                   toString(row[1]),
					PLAYER_LAST_TEAM_ID:           toInt(row[2]),
					PLAYER_LAST_TEAM_ABBREVIATION: toString(row[3]),
					AGE:                           toInt(row[4]),
					GP:                            toInt(row[5]),
					G:                             toString(row[6]),
					FGA_FREQUENCY:                 toFloat(row[7]),
					FGM:                           toInt(row[8]),
					FGA:                           toInt(row[9]),
					FG_PCT:                        toFloat(row[10]),
					EFG_PCT:                       toFloat(row[11]),
					FG2A_FREQUENCY:                toString(row[12]),
					FG2M:                          toString(row[13]),
					FG2A:                          toString(row[14]),
					FG2_PCT:                       toFloat(row[15]),
					FG3A_FREQUENCY:                toFloat(row[16]),
					FG3M:                          toInt(row[17]),
					FG3A:                          toInt(row[18]),
					FG3_PCT:                       toFloat(row[19]),
				}
				response.LeagueDashPlayerPtShot = append(response.LeagueDashPlayerPtShot, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
