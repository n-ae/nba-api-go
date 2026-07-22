package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// AllTimeLeadersGridsRequest contains parameters for the AllTimeLeadersGrids endpoint
type AllTimeLeadersGridsRequest struct {
	LeagueID   *parameters.LeagueID
	PerMode    *parameters.PerMode
	SeasonType *parameters.SeasonType
	TopX       *string
}

// AllTimeLeadersGridsAllTimeLeadersPTS represents the AllTimeLeadersPTS result set for AllTimeLeadersGrids
type AllTimeLeadersGridsAllTimeLeadersPTS struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	PTS         float64 `json:"PTS"`
	PTS_RANK    float64 `json:"PTS_RANK"`
}

// AllTimeLeadersGridsAllTimeLeadersAST represents the AllTimeLeadersAST result set for AllTimeLeadersGrids
type AllTimeLeadersGridsAllTimeLeadersAST struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	AST         float64 `json:"AST"`
	AST_RANK    float64 `json:"AST_RANK"`
}

// AllTimeLeadersGridsAllTimeLeadersREB represents the AllTimeLeadersREB result set for AllTimeLeadersGrids
type AllTimeLeadersGridsAllTimeLeadersREB struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	REB         float64 `json:"REB"`
	REB_RANK    float64 `json:"REB_RANK"`
}

// AllTimeLeadersGridsAllTimeLeadersBLK represents the AllTimeLeadersBLK result set for AllTimeLeadersGrids
type AllTimeLeadersGridsAllTimeLeadersBLK struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	BLK         float64 `json:"BLK"`
	BLK_RANK    float64 `json:"BLK_RANK"`
}

// AllTimeLeadersGridsAllTimeLeadersSTL represents the AllTimeLeadersSTL result set for AllTimeLeadersGrids
type AllTimeLeadersGridsAllTimeLeadersSTL struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	STL         float64 `json:"STL"`
	STL_RANK    float64 `json:"STL_RANK"`
}

// AllTimeLeadersGridsResponse contains the response data from the AllTimeLeadersGrids endpoint
type AllTimeLeadersGridsResponse struct {
	AllTimeLeadersPTS []AllTimeLeadersGridsAllTimeLeadersPTS
	AllTimeLeadersAST []AllTimeLeadersGridsAllTimeLeadersAST
	AllTimeLeadersREB []AllTimeLeadersGridsAllTimeLeadersREB
	AllTimeLeadersBLK []AllTimeLeadersGridsAllTimeLeadersBLK
	AllTimeLeadersSTL []AllTimeLeadersGridsAllTimeLeadersSTL
}

// GetAllTimeLeadersGrids retrieves data from the alltimeleadersgrids endpoint
func GetAllTimeLeadersGrids(ctx context.Context, client *stats.Client, req AllTimeLeadersGridsRequest) (*models.Response[*AllTimeLeadersGridsResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.TopX != nil {
		params.Set("TopX", *req.TopX)
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "alltimeleadersgrids", params, &rawResp); err != nil {
		return nil, err
	}

	response := &AllTimeLeadersGridsResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "AllTimeLeadersPTS"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(AllTimeLeadersGridsAllTimeLeadersPTS{})); err != nil {
			return nil, fmt.Errorf("AllTimeLeadersGrids: AllTimeLeadersPTS result set: %w", err)
		}
		response.AllTimeLeadersPTS = make([]AllTimeLeadersGridsAllTimeLeadersPTS, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 4 {
				item := AllTimeLeadersGridsAllTimeLeadersPTS{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					PTS:         toFloat(row[2]),
					PTS_RANK:    toFloat(row[3]),
				}
				response.AllTimeLeadersPTS = append(response.AllTimeLeadersPTS, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "AllTimeLeadersAST"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(AllTimeLeadersGridsAllTimeLeadersAST{})); err != nil {
			return nil, fmt.Errorf("AllTimeLeadersGrids: AllTimeLeadersAST result set: %w", err)
		}
		response.AllTimeLeadersAST = make([]AllTimeLeadersGridsAllTimeLeadersAST, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 4 {
				item := AllTimeLeadersGridsAllTimeLeadersAST{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					AST:         toFloat(row[2]),
					AST_RANK:    toFloat(row[3]),
				}
				response.AllTimeLeadersAST = append(response.AllTimeLeadersAST, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "AllTimeLeadersREB"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(AllTimeLeadersGridsAllTimeLeadersREB{})); err != nil {
			return nil, fmt.Errorf("AllTimeLeadersGrids: AllTimeLeadersREB result set: %w", err)
		}
		response.AllTimeLeadersREB = make([]AllTimeLeadersGridsAllTimeLeadersREB, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 4 {
				item := AllTimeLeadersGridsAllTimeLeadersREB{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					REB:         toFloat(row[2]),
					REB_RANK:    toFloat(row[3]),
				}
				response.AllTimeLeadersREB = append(response.AllTimeLeadersREB, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "AllTimeLeadersBLK"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(AllTimeLeadersGridsAllTimeLeadersBLK{})); err != nil {
			return nil, fmt.Errorf("AllTimeLeadersGrids: AllTimeLeadersBLK result set: %w", err)
		}
		response.AllTimeLeadersBLK = make([]AllTimeLeadersGridsAllTimeLeadersBLK, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 4 {
				item := AllTimeLeadersGridsAllTimeLeadersBLK{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					BLK:         toFloat(row[2]),
					BLK_RANK:    toFloat(row[3]),
				}
				response.AllTimeLeadersBLK = append(response.AllTimeLeadersBLK, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "AllTimeLeadersSTL"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(AllTimeLeadersGridsAllTimeLeadersSTL{})); err != nil {
			return nil, fmt.Errorf("AllTimeLeadersGrids: AllTimeLeadersSTL result set: %w", err)
		}
		response.AllTimeLeadersSTL = make([]AllTimeLeadersGridsAllTimeLeadersSTL, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 4 {
				item := AllTimeLeadersGridsAllTimeLeadersSTL{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					STL:         toFloat(row[2]),
					STL_RANK:    toFloat(row[3]),
				}
				response.AllTimeLeadersSTL = append(response.AllTimeLeadersSTL, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
