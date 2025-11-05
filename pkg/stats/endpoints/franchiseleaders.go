package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// FranchiseLeadersRequest contains parameters for the FranchiseLeaders endpoint
type FranchiseLeadersRequest struct {
	TeamID   string
	LeagueID *parameters.LeagueID
}

// FranchiseLeadersFranchiseLeaders represents the FranchiseLeaders result set for FranchiseLeaders
type FranchiseLeadersFranchiseLeaders struct {
	TEAM_ID       int     `json:"TEAM_ID"`
	PTS           float64 `json:"PTS"`
	PTS_PERSON_ID string  `json:"PTS_PERSON_ID"`
	PTS_PLAYER    float64 `json:"PTS_PLAYER"`
	AST           float64 `json:"AST"`
	AST_PERSON_ID string  `json:"AST_PERSON_ID"`
	AST_PLAYER    float64 `json:"AST_PLAYER"`
	REB           float64 `json:"REB"`
	REB_PERSON_ID string  `json:"REB_PERSON_ID"`
	REB_PLAYER    float64 `json:"REB_PLAYER"`
	BLK           float64 `json:"BLK"`
	BLK_PERSON_ID string  `json:"BLK_PERSON_ID"`
	BLK_PLAYER    float64 `json:"BLK_PLAYER"`
	STL           float64 `json:"STL"`
	STL_PERSON_ID string  `json:"STL_PERSON_ID"`
	STL_PLAYER    float64 `json:"STL_PLAYER"`
}

// FranchiseLeadersResponse contains the response data from the FranchiseLeaders endpoint
type FranchiseLeadersResponse struct {
	FranchiseLeaders []FranchiseLeadersFranchiseLeaders
}

// GetFranchiseLeaders retrieves data from the franchiseleaders endpoint
func GetFranchiseLeaders(ctx context.Context, client *stats.Client, req FranchiseLeadersRequest) (*models.Response[*FranchiseLeadersResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "franchiseleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &FranchiseLeadersResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.FranchiseLeaders = make([]FranchiseLeadersFranchiseLeaders, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 16 {
				item := FranchiseLeadersFranchiseLeaders{
					TEAM_ID:       toInt(row[0]),
					PTS:           toFloat(row[1]),
					PTS_PERSON_ID: toString(row[2]),
					PTS_PLAYER:    toFloat(row[3]),
					AST:           toFloat(row[4]),
					AST_PERSON_ID: toString(row[5]),
					AST_PLAYER:    toFloat(row[6]),
					REB:           toFloat(row[7]),
					REB_PERSON_ID: toString(row[8]),
					REB_PLAYER:    toFloat(row[9]),
					BLK:           toFloat(row[10]),
					BLK_PERSON_ID: toString(row[11]),
					BLK_PLAYER:    toFloat(row[12]),
					STL:           toFloat(row[13]),
					STL_PERSON_ID: toString(row[14]),
					STL_PLAYER:    toFloat(row[15]),
				}
				response.FranchiseLeaders = append(response.FranchiseLeaders, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
