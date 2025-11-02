package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// DraftBoardRequest contains parameters for the DraftBoard endpoint
type DraftBoardRequest struct {
	LeagueID *parameters.LeagueID
	Season   *parameters.Season
}

// DraftBoardDraftBoard represents the DraftBoard result set for DraftBoard
type DraftBoardDraftBoard struct {
	PERSON_ID         string `json:"PERSON_ID"`
	PLAYER_NAME       string `json:"PLAYER_NAME"`
	SEASON            string `json:"SEASON"`
	ROUND_NUMBER      string `json:"ROUND_NUMBER"`
	ROUND_PICK        string `json:"ROUND_PICK"`
	OVERALL_PICK      string `json:"OVERALL_PICK"`
	TEAM_ID           int    `json:"TEAM_ID"`
	TEAM_CITY         string `json:"TEAM_CITY"`
	TEAM_NAME         string `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
}

// DraftBoardResponse contains the response data from the DraftBoard endpoint
type DraftBoardResponse struct {
	DraftBoard []DraftBoardDraftBoard
}

// GetDraftBoard retrieves data from the draftboard endpoint
func GetDraftBoard(ctx context.Context, client *stats.Client, req DraftBoardRequest) (*models.Response[*DraftBoardResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/draftboard", params, &rawResp); err != nil {
		return nil, err
	}

	response := &DraftBoardResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.DraftBoard = make([]DraftBoardDraftBoard, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 10 {
				item := DraftBoardDraftBoard{
					PERSON_ID:         toString(row[0]),
					PLAYER_NAME:       toString(row[1]),
					SEASON:            toString(row[2]),
					ROUND_NUMBER:      toString(row[3]),
					ROUND_PICK:        toString(row[4]),
					OVERALL_PICK:      toString(row[5]),
					TEAM_ID:           toInt(row[6]),
					TEAM_CITY:         toString(row[7]),
					TEAM_NAME:         toString(row[8]),
					TEAM_ABBREVIATION: toString(row[9]),
				}
				response.DraftBoard = append(response.DraftBoard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
