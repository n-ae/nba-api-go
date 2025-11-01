package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// OpponentShootingRequest contains parameters for the OpponentShooting endpoint
type OpponentShootingRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// OpponentShootingOpponentShooting represents the OpponentShooting result set for OpponentShooting
type OpponentShootingOpponentShooting struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	OPP_FGM           int     `json:"OPP_FGM"`
	OPP_FGA           int     `json:"OPP_FGA"`
	OPP_FG_PCT        float64 `json:"OPP_FG_PCT"`
	OPP_FG2M          string  `json:"OPP_FG2M"`
	OPP_FG2A          string  `json:"OPP_FG2A"`
	OPP_FG2_PCT       float64 `json:"OPP_FG2_PCT"`
	OPP_FG3M          int     `json:"OPP_FG3M"`
	OPP_FG3A          int     `json:"OPP_FG3A"`
	OPP_FG3_PCT       float64 `json:"OPP_FG3_PCT"`
}

// OpponentShootingResponse contains the response data from the OpponentShooting endpoint
type OpponentShootingResponse struct {
	OpponentShooting []OpponentShootingOpponentShooting
}

// GetOpponentShooting retrieves data from the opponentshooting endpoint
func GetOpponentShooting(ctx context.Context, client *stats.Client, req OpponentShootingRequest) (*models.Response[*OpponentShootingResponse], error) {
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
	if err := client.GetJSON(ctx, "/opponentshooting", params, &rawResp); err != nil {
		return nil, err
	}

	response := &OpponentShootingResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OpponentShooting = make([]OpponentShootingOpponentShooting, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 15 {
				item := OpponentShootingOpponentShooting{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					OPP_FGM:           toInt(row[6]),
					OPP_FGA:           toInt(row[7]),
					OPP_FG_PCT:        toFloat(row[8]),
					OPP_FG2M:          toString(row[9]),
					OPP_FG2A:          toString(row[10]),
					OPP_FG2_PCT:       toFloat(row[11]),
					OPP_FG3M:          toInt(row[12]),
					OPP_FG3A:          toInt(row[13]),
					OPP_FG3_PCT:       toFloat(row[14]),
				}
				response.OpponentShooting = append(response.OpponentShooting, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
