package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamHistoricalLeadersRequest contains parameters for the TeamHistoricalLeaders endpoint
type TeamHistoricalLeadersRequest struct {
	TeamID string
	LeagueID *parameters.LeagueID
}


// TeamHistoricalLeadersCareerLeadersPTS represents the CareerLeadersPTS result set for TeamHistoricalLeaders
type TeamHistoricalLeadersCareerLeadersPTS struct {
	TEAM_ID int `json:"TEAM_ID"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER string `json:"PLAYER"`
	PTS float64 `json:"PTS"`
	PTS_RANK float64 `json:"PTS_RANK"`
}

// TeamHistoricalLeadersCareerLeadersAST represents the CareerLeadersAST result set for TeamHistoricalLeaders
type TeamHistoricalLeadersCareerLeadersAST struct {
	TEAM_ID int `json:"TEAM_ID"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER string `json:"PLAYER"`
	AST float64 `json:"AST"`
	AST_RANK float64 `json:"AST_RANK"`
}

// TeamHistoricalLeadersCareerLeadersREB represents the CareerLeadersREB result set for TeamHistoricalLeaders
type TeamHistoricalLeadersCareerLeadersREB struct {
	TEAM_ID int `json:"TEAM_ID"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER string `json:"PLAYER"`
	REB float64 `json:"REB"`
	REB_RANK float64 `json:"REB_RANK"`
}

// TeamHistoricalLeadersCareerLeadersBLK represents the CareerLeadersBLK result set for TeamHistoricalLeaders
type TeamHistoricalLeadersCareerLeadersBLK struct {
	TEAM_ID int `json:"TEAM_ID"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER string `json:"PLAYER"`
	BLK float64 `json:"BLK"`
	BLK_RANK float64 `json:"BLK_RANK"`
}

// TeamHistoricalLeadersCareerLeadersSTL represents the CareerLeadersSTL result set for TeamHistoricalLeaders
type TeamHistoricalLeadersCareerLeadersSTL struct {
	TEAM_ID int `json:"TEAM_ID"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER string `json:"PLAYER"`
	STL float64 `json:"STL"`
	STL_RANK float64 `json:"STL_RANK"`
}


// TeamHistoricalLeadersResponse contains the response data from the TeamHistoricalLeaders endpoint
type TeamHistoricalLeadersResponse struct {
	CareerLeadersPTS []TeamHistoricalLeadersCareerLeadersPTS
	CareerLeadersAST []TeamHistoricalLeadersCareerLeadersAST
	CareerLeadersREB []TeamHistoricalLeadersCareerLeadersREB
	CareerLeadersBLK []TeamHistoricalLeadersCareerLeadersBLK
	CareerLeadersSTL []TeamHistoricalLeadersCareerLeadersSTL
}

// GetTeamHistoricalLeaders retrieves data from the teamhistoricalleaders endpoint
func GetTeamHistoricalLeaders(ctx context.Context, client *stats.Client, req TeamHistoricalLeadersRequest) (*models.Response[*TeamHistoricalLeadersResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamhistoricalleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamHistoricalLeadersResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.CareerLeadersPTS = make([]TeamHistoricalLeadersCareerLeadersPTS, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 5 {
				item := TeamHistoricalLeadersCareerLeadersPTS{
					TEAM_ID: toInt(row[0]),
					PLAYER_ID: toInt(row[1]),
					PLAYER: toString(row[2]),
					PTS: toFloat(row[3]),
					PTS_RANK: toFloat(row[4]),
				}
				response.CareerLeadersPTS = append(response.CareerLeadersPTS, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.CareerLeadersAST = make([]TeamHistoricalLeadersCareerLeadersAST, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 5 {
				item := TeamHistoricalLeadersCareerLeadersAST{
					TEAM_ID: toInt(row[0]),
					PLAYER_ID: toInt(row[1]),
					PLAYER: toString(row[2]),
					AST: toFloat(row[3]),
					AST_RANK: toFloat(row[4]),
				}
				response.CareerLeadersAST = append(response.CareerLeadersAST, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.CareerLeadersREB = make([]TeamHistoricalLeadersCareerLeadersREB, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 5 {
				item := TeamHistoricalLeadersCareerLeadersREB{
					TEAM_ID: toInt(row[0]),
					PLAYER_ID: toInt(row[1]),
					PLAYER: toString(row[2]),
					REB: toFloat(row[3]),
					REB_RANK: toFloat(row[4]),
				}
				response.CareerLeadersREB = append(response.CareerLeadersREB, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.CareerLeadersBLK = make([]TeamHistoricalLeadersCareerLeadersBLK, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 5 {
				item := TeamHistoricalLeadersCareerLeadersBLK{
					TEAM_ID: toInt(row[0]),
					PLAYER_ID: toInt(row[1]),
					PLAYER: toString(row[2]),
					BLK: toFloat(row[3]),
					BLK_RANK: toFloat(row[4]),
				}
				response.CareerLeadersBLK = append(response.CareerLeadersBLK, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.CareerLeadersSTL = make([]TeamHistoricalLeadersCareerLeadersSTL, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 5 {
				item := TeamHistoricalLeadersCareerLeadersSTL{
					TEAM_ID: toInt(row[0]),
					PLAYER_ID: toInt(row[1]),
					PLAYER: toString(row[2]),
					STL: toFloat(row[3]),
					STL_RANK: toFloat(row[4]),
				}
				response.CareerLeadersSTL = append(response.CareerLeadersSTL, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
