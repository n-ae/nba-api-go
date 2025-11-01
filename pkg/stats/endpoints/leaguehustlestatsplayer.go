package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueHustleStatsPlayerRequest contains parameters for the LeagueHustleStatsPlayer endpoint
type LeagueHustleStatsPlayerRequest struct {
	Season *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode *parameters.PerMode
	LeagueID *parameters.LeagueID
}


// LeagueHustleStatsPlayerHustleStatsPlayer represents the HustleStatsPlayer result set for LeagueHustleStatsPlayer
type LeagueHustleStatsPlayerHustleStatsPlayer struct {
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER_NAME string `json:"PLAYER_NAME"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	AGE int `json:"AGE"`
	GP int `json:"GP"`
	W string `json:"W"`
	L string `json:"L"`
	W_PCT float64 `json:"W_PCT"`
	MIN float64 `json:"MIN"`
	CONTESTED_SHOTS string `json:"CONTESTED_SHOTS"`
	CONTESTED_SHOTS_2PT string `json:"CONTESTED_SHOTS_2PT"`
	CONTESTED_SHOTS_3PT string `json:"CONTESTED_SHOTS_3PT"`
	DEFLECTIONS string `json:"DEFLECTIONS"`
	CHARGES_DRAWN string `json:"CHARGES_DRAWN"`
	SCREEN_ASSISTS string `json:"SCREEN_ASSISTS"`
	SCREEN_AST_PTS float64 `json:"SCREEN_AST_PTS"`
	OFF_LOOSE_BALLS_RECOVERED string `json:"OFF_LOOSE_BALLS_RECOVERED"`
	DEF_LOOSE_BALLS_RECOVERED string `json:"DEF_LOOSE_BALLS_RECOVERED"`
	LOOSE_BALLS_RECOVERED string `json:"LOOSE_BALLS_RECOVERED"`
	OFF_BOXOUTS string `json:"OFF_BOXOUTS"`
	DEF_BOXOUTS string `json:"DEF_BOXOUTS"`
	BOX_OUT_PLAYER_TEAM_REBS float64 `json:"BOX_OUT_PLAYER_TEAM_REBS"`
	BOX_OUT_PLAYER_REBS float64 `json:"BOX_OUT_PLAYER_REBS"`
	BOX_OUTS string `json:"BOX_OUTS"`
}


// LeagueHustleStatsPlayerResponse contains the response data from the LeagueHustleStatsPlayer endpoint
type LeagueHustleStatsPlayerResponse struct {
	HustleStatsPlayer []LeagueHustleStatsPlayerHustleStatsPlayer
}

// GetLeagueHustleStatsPlayer retrieves data from the leaguehustlestatsp layer endpoint
func GetLeagueHustleStatsPlayer(ctx context.Context, client *stats.Client, req LeagueHustleStatsPlayerRequest) (*models.Response[*LeagueHustleStatsPlayerResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguehustlestatsp layer", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueHustleStatsPlayerResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.HustleStatsPlayer = make([]LeagueHustleStatsPlayerHustleStatsPlayer, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 25 {
				item := LeagueHustleStatsPlayerHustleStatsPlayer{
					PLAYER_ID: toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					TEAM_ID: toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					AGE: toInt(row[4]),
					GP: toInt(row[5]),
					W: toString(row[6]),
					L: toString(row[7]),
					W_PCT: toFloat(row[8]),
					MIN: toFloat(row[9]),
					CONTESTED_SHOTS: toString(row[10]),
					CONTESTED_SHOTS_2PT: toString(row[11]),
					CONTESTED_SHOTS_3PT: toString(row[12]),
					DEFLECTIONS: toString(row[13]),
					CHARGES_DRAWN: toString(row[14]),
					SCREEN_ASSISTS: toString(row[15]),
					SCREEN_AST_PTS: toFloat(row[16]),
					OFF_LOOSE_BALLS_RECOVERED: toString(row[17]),
					DEF_LOOSE_BALLS_RECOVERED: toString(row[18]),
					LOOSE_BALLS_RECOVERED: toString(row[19]),
					OFF_BOXOUTS: toString(row[20]),
					DEF_BOXOUTS: toString(row[21]),
					BOX_OUT_PLAYER_TEAM_REBS: toFloat(row[22]),
					BOX_OUT_PLAYER_REBS: toFloat(row[23]),
					BOX_OUTS: toString(row[24]),
				}
				response.HustleStatsPlayer = append(response.HustleStatsPlayer, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
