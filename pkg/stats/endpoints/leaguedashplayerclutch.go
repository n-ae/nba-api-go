package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPlayerClutchRequest contains parameters for the LeagueDashPlayerClutch endpoint
type LeagueDashPlayerClutchRequest struct {
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	PerMode     *parameters.PerMode
	LeagueID    *parameters.LeagueID
	ClutchTime  *string
	AheadBehind *string
	PointDiff   *string
}

// LeagueDashPlayerClutchLeagueDashPlayerClutch represents the LeagueDashPlayerClutch result set for LeagueDashPlayerClutch
type LeagueDashPlayerClutchLeagueDashPlayerClutch struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	AGE               int     `json:"AGE"`
	GP                int     `json:"GP"`
	W                 string  `json:"W"`
	L                 string  `json:"L"`
	W_PCT             float64 `json:"W_PCT"`
	MIN               float64 `json:"MIN"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	FTM               int     `json:"FTM"`
	FTA               int     `json:"FTA"`
	FT_PCT            float64 `json:"FT_PCT"`
	OREB              float64 `json:"OREB"`
	DREB              float64 `json:"DREB"`
	REB               float64 `json:"REB"`
	AST               float64 `json:"AST"`
	TOV               float64 `json:"TOV"`
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	BLKA              int     `json:"BLKA"`
	PF                float64 `json:"PF"`
	PFD               float64 `json:"PFD"`
	PTS               float64 `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// LeagueDashPlayerClutchResponse contains the response data from the LeagueDashPlayerClutch endpoint
type LeagueDashPlayerClutchResponse struct {
	LeagueDashPlayerClutch []LeagueDashPlayerClutchLeagueDashPlayerClutch
}

// GetLeagueDashPlayerClutch retrieves data from the leaguedashplayerclutch endpoint
func GetLeagueDashPlayerClutch(ctx context.Context, client *stats.Client, req LeagueDashPlayerClutchRequest) (*models.Response[*LeagueDashPlayerClutchResponse], error) {
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
	if req.ClutchTime != nil {
		params.Set("ClutchTime", string(*req.ClutchTime))
	}
	if req.AheadBehind != nil {
		params.Set("AheadBehind", string(*req.AheadBehind))
	}
	if req.PointDiff != nil {
		params.Set("PointDiff", string(*req.PointDiff))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguedashplayerclutch", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPlayerClutchResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPlayerClutch = make([]LeagueDashPlayerClutchLeagueDashPlayerClutch, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 31 {
				item := LeagueDashPlayerClutchLeagueDashPlayerClutch{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					AGE:               toInt(row[4]),
					GP:                toInt(row[5]),
					W:                 toString(row[6]),
					L:                 toString(row[7]),
					W_PCT:             toFloat(row[8]),
					MIN:               toFloat(row[9]),
					FGM:               toInt(row[10]),
					FGA:               toInt(row[11]),
					FG_PCT:            toFloat(row[12]),
					FG3M:              toInt(row[13]),
					FG3A:              toInt(row[14]),
					FG3_PCT:           toFloat(row[15]),
					FTM:               toInt(row[16]),
					FTA:               toInt(row[17]),
					FT_PCT:            toFloat(row[18]),
					OREB:              toFloat(row[19]),
					DREB:              toFloat(row[20]),
					REB:               toFloat(row[21]),
					AST:               toFloat(row[22]),
					TOV:               toFloat(row[23]),
					STL:               toFloat(row[24]),
					BLK:               toFloat(row[25]),
					BLKA:              toInt(row[26]),
					PF:                toFloat(row[27]),
					PFD:               toFloat(row[28]),
					PTS:               toFloat(row[29]),
					PLUS_MINUS:        toFloat(row[30]),
				}
				response.LeagueDashPlayerClutch = append(response.LeagueDashPlayerClutch, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
