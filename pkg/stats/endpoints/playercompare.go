package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerCompareRequest contains parameters for the PlayerCompare endpoint
type PlayerCompareRequest struct {
	PlayerIDList string
	Season       *parameters.Season
	SeasonType   *parameters.SeasonType
	PerMode      *parameters.PerMode
	LeagueID     *parameters.LeagueID
}

// PlayerCompareOverallCompare represents the OverallCompare result set for PlayerCompare
type PlayerCompareOverallCompare struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	GP          int     `json:"GP"`
	W           string  `json:"W"`
	L           string  `json:"L"`
	W_PCT       float64 `json:"W_PCT"`
	MIN         float64 `json:"MIN"`
	FGM         int     `json:"FGM"`
	FGA         int     `json:"FGA"`
	FG_PCT      float64 `json:"FG_PCT"`
	FG3M        int     `json:"FG3M"`
	FG3A        int     `json:"FG3A"`
	FG3_PCT     float64 `json:"FG3_PCT"`
	FTM         int     `json:"FTM"`
	FTA         int     `json:"FTA"`
	FT_PCT      float64 `json:"FT_PCT"`
	OREB        float64 `json:"OREB"`
	DREB        float64 `json:"DREB"`
	REB         float64 `json:"REB"`
	AST         float64 `json:"AST"`
	TOV         float64 `json:"TOV"`
	STL         float64 `json:"STL"`
	BLK         float64 `json:"BLK"`
	BLKA        int     `json:"BLKA"`
	PF          float64 `json:"PF"`
	PFD         float64 `json:"PFD"`
	PTS         float64 `json:"PTS"`
	PLUS_MINUS  float64 `json:"PLUS_MINUS"`
}

// PlayerCompareResponse contains the response data from the PlayerCompare endpoint
type PlayerCompareResponse struct {
	OverallCompare []PlayerCompareOverallCompare
}

// GetPlayerCompare retrieves data from the playercompare endpoint
func GetPlayerCompare(ctx context.Context, client *stats.Client, req PlayerCompareRequest) (*models.Response[*PlayerCompareResponse], error) {
	params := url.Values{}
	if req.PlayerIDList == "" {
		return nil, fmt.Errorf("PlayerIDList is required")
	}
	params.Set("PlayerIDList", string(req.PlayerIDList))
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
	if err := client.GetJSON(ctx, "/playercompare", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerCompareResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallCompare = make([]PlayerCompareOverallCompare, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := PlayerCompareOverallCompare{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					GP:          toInt(row[2]),
					W:           toString(row[3]),
					L:           toString(row[4]),
					W_PCT:       toFloat(row[5]),
					MIN:         toFloat(row[6]),
					FGM:         toInt(row[7]),
					FGA:         toInt(row[8]),
					FG_PCT:      toFloat(row[9]),
					FG3M:        toInt(row[10]),
					FG3A:        toInt(row[11]),
					FG3_PCT:     toFloat(row[12]),
					FTM:         toInt(row[13]),
					FTA:         toInt(row[14]),
					FT_PCT:      toFloat(row[15]),
					OREB:        toFloat(row[16]),
					DREB:        toFloat(row[17]),
					REB:         toFloat(row[18]),
					AST:         toFloat(row[19]),
					TOV:         toFloat(row[20]),
					STL:         toFloat(row[21]),
					BLK:         toFloat(row[22]),
					BLKA:        toInt(row[23]),
					PF:          toFloat(row[24]),
					PFD:         toFloat(row[25]),
					PTS:         toFloat(row[26]),
					PLUS_MINUS:  toFloat(row[27]),
				}
				response.OverallCompare = append(response.OverallCompare, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
