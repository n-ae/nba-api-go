package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingElbowTouchRequest contains parameters for the PlayerTrackingElbowTouch endpoint
type PlayerTrackingElbowTouchRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingElbowTouchPlayerTrackingElbowTouch represents the PlayerTrackingElbowTouch result set for PlayerTrackingElbowTouch
type PlayerTrackingElbowTouchPlayerTrackingElbowTouch struct {
	PLAYER_ID          int     `json:"PLAYER_ID"`
	PLAYER_NAME        string  `json:"PLAYER_NAME"`
	TEAM_ID            int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION  string  `json:"TEAM_ABBREVIATION"`
	GP                 int     `json:"GP"`
	MIN                float64 `json:"MIN"`
	ELBOW_TOUCHES      string  `json:"ELBOW_TOUCHES"`
	ELBOW_TOUCH_FGM    int     `json:"ELBOW_TOUCH_FGM"`
	ELBOW_TOUCH_FGA    int     `json:"ELBOW_TOUCH_FGA"`
	ELBOW_TOUCH_FG_PCT float64 `json:"ELBOW_TOUCH_FG_PCT"`
	ELBOW_TOUCH_FTM    int     `json:"ELBOW_TOUCH_FTM"`
	ELBOW_TOUCH_FTA    int     `json:"ELBOW_TOUCH_FTA"`
	ELBOW_TOUCH_FT_PCT float64 `json:"ELBOW_TOUCH_FT_PCT"`
	ELBOW_TOUCH_PTS    float64 `json:"ELBOW_TOUCH_PTS"`
	ELBOW_TOUCH_PASS   string  `json:"ELBOW_TOUCH_PASS"`
	ELBOW_TOUCH_AST    float64 `json:"ELBOW_TOUCH_AST"`
	ELBOW_TOUCH_TOV    float64 `json:"ELBOW_TOUCH_TOV"`
}

// PlayerTrackingElbowTouchResponse contains the response data from the PlayerTrackingElbowTouch endpoint
type PlayerTrackingElbowTouchResponse struct {
	PlayerTrackingElbowTouch []PlayerTrackingElbowTouchPlayerTrackingElbowTouch
}

// GetPlayerTrackingElbowTouch retrieves data from the playertrackingelbowtouch endpoint
func GetPlayerTrackingElbowTouch(ctx context.Context, client *stats.Client, req PlayerTrackingElbowTouchRequest) (*models.Response[*PlayerTrackingElbowTouchResponse], error) {
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
	if err := client.GetJSON(ctx, "/playertrackingelbowtouch", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingElbowTouchResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingElbowTouch = make([]PlayerTrackingElbowTouchPlayerTrackingElbowTouch, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 17 {
				item := PlayerTrackingElbowTouchPlayerTrackingElbowTouch{
					PLAYER_ID:          toInt(row[0]),
					PLAYER_NAME:        toString(row[1]),
					TEAM_ID:            toInt(row[2]),
					TEAM_ABBREVIATION:  toString(row[3]),
					GP:                 toInt(row[4]),
					MIN:                toFloat(row[5]),
					ELBOW_TOUCHES:      toString(row[6]),
					ELBOW_TOUCH_FGM:    toInt(row[7]),
					ELBOW_TOUCH_FGA:    toInt(row[8]),
					ELBOW_TOUCH_FG_PCT: toFloat(row[9]),
					ELBOW_TOUCH_FTM:    toInt(row[10]),
					ELBOW_TOUCH_FTA:    toInt(row[11]),
					ELBOW_TOUCH_FT_PCT: toFloat(row[12]),
					ELBOW_TOUCH_PTS:    toFloat(row[13]),
					ELBOW_TOUCH_PASS:   toString(row[14]),
					ELBOW_TOUCH_AST:    toFloat(row[15]),
					ELBOW_TOUCH_TOV:    toFloat(row[16]),
				}
				response.PlayerTrackingElbowTouch = append(response.PlayerTrackingElbowTouch, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
