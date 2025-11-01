package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerCareerByCollegeRollupRequest contains parameters for the PlayerCareerByCollegeRollup endpoint
type PlayerCareerByCollegeRollupRequest struct {
	LeagueID *parameters.LeagueID
	PerMode  *parameters.PerMode
}

// PlayerCareerByCollegeRollupCollegeStats represents the CollegeStats result set for PlayerCareerByCollegeRollup
type PlayerCareerByCollegeRollupCollegeStats struct {
	SCHOOL_NAME         string  `json:"SCHOOL_NAME"`
	SEASON_COUNT        string  `json:"SEASON_COUNT"`
	PLAYER_COUNT        float64 `json:"PLAYER_COUNT"`
	ACTIVE_PLAYER_COUNT float64 `json:"ACTIVE_PLAYER_COUNT"`
	MIN                 float64 `json:"MIN"`
	PTS                 float64 `json:"PTS"`
	REB                 float64 `json:"REB"`
	AST                 float64 `json:"AST"`
	FG_PCT              float64 `json:"FG_PCT"`
	FG3_PCT             float64 `json:"FG3_PCT"`
	FT_PCT              float64 `json:"FT_PCT"`
}

// PlayerCareerByCollegeRollupResponse contains the response data from the PlayerCareerByCollegeRollup endpoint
type PlayerCareerByCollegeRollupResponse struct {
	CollegeStats []PlayerCareerByCollegeRollupCollegeStats
}

// GetPlayerCareerByCollegeRollup retrieves data from the playercareerbyrollegerollup endpoint
func GetPlayerCareerByCollegeRollup(ctx context.Context, client *stats.Client, req PlayerCareerByCollegeRollupRequest) (*models.Response[*PlayerCareerByCollegeRollupResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playercareerbyrollegerollup", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerCareerByCollegeRollupResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.CollegeStats = make([]PlayerCareerByCollegeRollupCollegeStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 11 {
				item := PlayerCareerByCollegeRollupCollegeStats{
					SCHOOL_NAME:         toString(row[0]),
					SEASON_COUNT:        toString(row[1]),
					PLAYER_COUNT:        toFloat(row[2]),
					ACTIVE_PLAYER_COUNT: toFloat(row[3]),
					MIN:                 toFloat(row[4]),
					PTS:                 toFloat(row[5]),
					REB:                 toFloat(row[6]),
					AST:                 toFloat(row[7]),
					FG_PCT:              toFloat(row[8]),
					FG3_PCT:             toFloat(row[9]),
					FT_PCT:              toFloat(row[10]),
				}
				response.CollegeStats = append(response.CollegeStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
