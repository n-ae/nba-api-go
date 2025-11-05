package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingPaintTouchRequest contains parameters for the PlayerTrackingPaintTouch endpoint
type PlayerTrackingPaintTouchRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingPaintTouchPlayerTrackingPaintTouch represents the PlayerTrackingPaintTouch result set for PlayerTrackingPaintTouch
type PlayerTrackingPaintTouchPlayerTrackingPaintTouch struct {
	PLAYER_ID          int     `json:"PLAYER_ID"`
	PLAYER_NAME        string  `json:"PLAYER_NAME"`
	TEAM_ID            int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION  string  `json:"TEAM_ABBREVIATION"`
	GP                 int     `json:"GP"`
	MIN                float64 `json:"MIN"`
	PAINT_TOUCHES      string  `json:"PAINT_TOUCHES"`
	PAINT_TOUCH_FGM    int     `json:"PAINT_TOUCH_FGM"`
	PAINT_TOUCH_FGA    int     `json:"PAINT_TOUCH_FGA"`
	PAINT_TOUCH_FG_PCT float64 `json:"PAINT_TOUCH_FG_PCT"`
	PAINT_TOUCH_FTM    int     `json:"PAINT_TOUCH_FTM"`
	PAINT_TOUCH_FTA    int     `json:"PAINT_TOUCH_FTA"`
	PAINT_TOUCH_FT_PCT float64 `json:"PAINT_TOUCH_FT_PCT"`
	PAINT_TOUCH_PTS    float64 `json:"PAINT_TOUCH_PTS"`
	PAINT_TOUCH_PASS   string  `json:"PAINT_TOUCH_PASS"`
	PAINT_TOUCH_AST    float64 `json:"PAINT_TOUCH_AST"`
	PAINT_TOUCH_TOV    float64 `json:"PAINT_TOUCH_TOV"`
}

// PlayerTrackingPaintTouchResponse contains the response data from the PlayerTrackingPaintTouch endpoint
type PlayerTrackingPaintTouchResponse struct {
	PlayerTrackingPaintTouch []PlayerTrackingPaintTouchPlayerTrackingPaintTouch
}

// GetPlayerTrackingPaintTouch retrieves data from the playertrackingpainttouch endpoint
func GetPlayerTrackingPaintTouch(ctx context.Context, client *stats.Client, req PlayerTrackingPaintTouchRequest) (*models.Response[*PlayerTrackingPaintTouchResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingpainttouch", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingPaintTouchResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingPaintTouch = make([]PlayerTrackingPaintTouchPlayerTrackingPaintTouch, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 17 {
				item := PlayerTrackingPaintTouchPlayerTrackingPaintTouch{
					PLAYER_ID:          toInt(row[0]),
					PLAYER_NAME:        toString(row[1]),
					TEAM_ID:            toInt(row[2]),
					TEAM_ABBREVIATION:  toString(row[3]),
					GP:                 toInt(row[4]),
					MIN:                toFloat(row[5]),
					PAINT_TOUCHES:      toString(row[6]),
					PAINT_TOUCH_FGM:    toInt(row[7]),
					PAINT_TOUCH_FGA:    toInt(row[8]),
					PAINT_TOUCH_FG_PCT: toFloat(row[9]),
					PAINT_TOUCH_FTM:    toInt(row[10]),
					PAINT_TOUCH_FTA:    toInt(row[11]),
					PAINT_TOUCH_FT_PCT: toFloat(row[12]),
					PAINT_TOUCH_PTS:    toFloat(row[13]),
					PAINT_TOUCH_PASS:   toString(row[14]),
					PAINT_TOUCH_AST:    toFloat(row[15]),
					PAINT_TOUCH_TOV:    toFloat(row[16]),
				}
				response.PlayerTrackingPaintTouch = append(response.PlayerTrackingPaintTouch, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
