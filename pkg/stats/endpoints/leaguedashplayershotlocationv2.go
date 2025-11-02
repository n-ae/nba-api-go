package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPlayerShotLocationV2Request contains parameters for the LeagueDashPlayerShotLocationV2 endpoint
type LeagueDashPlayerShotLocationV2Request struct {
	Season        *parameters.Season
	SeasonType    *parameters.SeasonType
	PerMode       *parameters.PerMode
	DistanceRange *string
	LeagueID      *parameters.LeagueID
}

// LeagueDashPlayerShotLocationV2ShotLocationLeague represents the ShotLocationLeague result set for LeagueDashPlayerShotLocationV2
type LeagueDashPlayerShotLocationV2ShotLocationLeague struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	AGE               int     `json:"AGE"`
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

// LeagueDashPlayerShotLocationV2Response contains the response data from the LeagueDashPlayerShotLocationV2 endpoint
type LeagueDashPlayerShotLocationV2Response struct {
	ShotLocationLeague []LeagueDashPlayerShotLocationV2ShotLocationLeague
}

// GetLeagueDashPlayerShotLocationV2 retrieves data from the leaguedashplayershotlocationv2 endpoint
func GetLeagueDashPlayerShotLocationV2(ctx context.Context, client *stats.Client, req LeagueDashPlayerShotLocationV2Request) (*models.Response[*LeagueDashPlayerShotLocationV2Response], error) {
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
	if req.DistanceRange != nil {
		params.Set("DistanceRange", string(*req.DistanceRange))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguedashplayershotlocationv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPlayerShotLocationV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.ShotLocationLeague = make([]LeagueDashPlayerShotLocationV2ShotLocationLeague, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 20 {
				item := LeagueDashPlayerShotLocationV2ShotLocationLeague{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					AGE:               toInt(row[4]),
					GP:                toInt(row[5]),
					G:                 toString(row[6]),
					FGA_FREQUENCY:     toFloat(row[7]),
					FGM:               toInt(row[8]),
					FGA:               toInt(row[9]),
					FG_PCT:            toFloat(row[10]),
					EFG_PCT:           toFloat(row[11]),
					FG2A_FREQUENCY:    toString(row[12]),
					FG2M:              toString(row[13]),
					FG2A:              toString(row[14]),
					FG2_PCT:           toFloat(row[15]),
					FG3A_FREQUENCY:    toFloat(row[16]),
					FG3M:              toInt(row[17]),
					FG3A:              toInt(row[18]),
					FG3_PCT:           toFloat(row[19]),
				}
				response.ShotLocationLeague = append(response.ShotLocationLeague, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
