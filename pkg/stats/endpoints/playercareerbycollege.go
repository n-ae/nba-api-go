package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerCareerByCollegeRequest contains parameters for the PlayerCareerByCollege endpoint
type PlayerCareerByCollegeRequest struct {
	LeagueID *parameters.LeagueID
	College  *string
}

// PlayerCareerByCollegePlayerCareerByCollege represents the PlayerCareerByCollege result set for PlayerCareerByCollege
type PlayerCareerByCollegePlayerCareerByCollege struct {
	PERSON_ID   string  `json:"PERSON_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	SCHOOL_NAME string  `json:"SCHOOL_NAME"`
	FIRST_YEAR  string  `json:"FIRST_YEAR"`
	LAST_YEAR   float64 `json:"LAST_YEAR"`
	SEASONS     string  `json:"SEASONS"`
	GP          int     `json:"GP"`
	MIN         float64 `json:"MIN"`
	PTS         float64 `json:"PTS"`
	REB         float64 `json:"REB"`
	AST         float64 `json:"AST"`
}

// PlayerCareerByCollegeResponse contains the response data from the PlayerCareerByCollege endpoint
type PlayerCareerByCollegeResponse struct {
	PlayerCareerByCollege []PlayerCareerByCollegePlayerCareerByCollege
}

// GetPlayerCareerByCollege retrieves data from the playercareerbycollege endpoint
func GetPlayerCareerByCollege(ctx context.Context, client *stats.Client, req PlayerCareerByCollegeRequest) (*models.Response[*PlayerCareerByCollegeResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.College != nil {
		params.Set("College", string(*req.College))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playercareerbycollege", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerCareerByCollegeResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerCareerByCollege = make([]PlayerCareerByCollegePlayerCareerByCollege, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 11 {
				item := PlayerCareerByCollegePlayerCareerByCollege{
					PERSON_ID:   toString(row[0]),
					PLAYER_NAME: toString(row[1]),
					SCHOOL_NAME: toString(row[2]),
					FIRST_YEAR:  toString(row[3]),
					LAST_YEAR:   toFloat(row[4]),
					SEASONS:     toString(row[5]),
					GP:          toInt(row[6]),
					MIN:         toFloat(row[7]),
					PTS:         toFloat(row[8]),
					REB:         toFloat(row[9]),
					AST:         toFloat(row[10]),
				}
				response.PlayerCareerByCollege = append(response.PlayerCareerByCollege, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
