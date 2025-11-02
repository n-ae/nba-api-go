package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashLineupsRequest contains parameters for the LeagueDashLineups endpoint
type LeagueDashLineupsRequest struct {
	Season        *parameters.Season
	SeasonType    *parameters.SeasonType
	MeasureType   *string
	PerMode       *parameters.PerMode
	GroupQuantity *string
	LeagueID      *parameters.LeagueID
}

// LeagueDashLineupsLineups represents the Lineups result set for LeagueDashLineups
type LeagueDashLineupsLineups struct {
	GROUP_ID          string  `json:"GROUP_ID"`
	GROUP_NAME        string  `json:"GROUP_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
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
	OFF_RATING        string  `json:"OFF_RATING"`
	DEF_RATING        string  `json:"DEF_RATING"`
	NET_RATING        string  `json:"NET_RATING"`
}

// LeagueDashLineupsResponse contains the response data from the LeagueDashLineups endpoint
type LeagueDashLineupsResponse struct {
	Lineups []LeagueDashLineupsLineups
}

// GetLeagueDashLineups retrieves data from the leaguedashlineups endpoint
func GetLeagueDashLineups(ctx context.Context, client *stats.Client, req LeagueDashLineupsRequest) (*models.Response[*LeagueDashLineupsResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.MeasureType != nil {
		params.Set("MeasureType", string(*req.MeasureType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.GroupQuantity != nil {
		params.Set("GroupQuantity", string(*req.GroupQuantity))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguedashlineups", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashLineupsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Lineups = make([]LeagueDashLineupsLineups, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 33 {
				item := LeagueDashLineupsLineups{
					GROUP_ID:          toString(row[0]),
					GROUP_NAME:        toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					W:                 toString(row[5]),
					L:                 toString(row[6]),
					W_PCT:             toFloat(row[7]),
					MIN:               toFloat(row[8]),
					FGM:               toInt(row[9]),
					FGA:               toInt(row[10]),
					FG_PCT:            toFloat(row[11]),
					FG3M:              toInt(row[12]),
					FG3A:              toInt(row[13]),
					FG3_PCT:           toFloat(row[14]),
					FTM:               toInt(row[15]),
					FTA:               toInt(row[16]),
					FT_PCT:            toFloat(row[17]),
					OREB:              toFloat(row[18]),
					DREB:              toFloat(row[19]),
					REB:               toFloat(row[20]),
					AST:               toFloat(row[21]),
					TOV:               toFloat(row[22]),
					STL:               toFloat(row[23]),
					BLK:               toFloat(row[24]),
					BLKA:              toInt(row[25]),
					PF:                toFloat(row[26]),
					PFD:               toFloat(row[27]),
					PTS:               toFloat(row[28]),
					PLUS_MINUS:        toFloat(row[29]),
					OFF_RATING:        toString(row[30]),
					DEF_RATING:        toString(row[31]),
					NET_RATING:        toString(row[32]),
				}
				response.Lineups = append(response.Lineups, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
