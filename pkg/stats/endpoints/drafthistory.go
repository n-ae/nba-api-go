package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// DraftHistoryRequest contains parameters for the DraftHistory endpoint
type DraftHistoryRequest struct {
	LeagueID *parameters.LeagueID
	Season   *parameters.Season
}

// DraftHistoryDraftHistory represents the DraftHistory result set for DraftHistory
type DraftHistoryDraftHistory struct {
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
	ORGANIZATION      string `json:"ORGANIZATION"`
	ORGANIZATION_TYPE string `json:"ORGANIZATION_TYPE"`
}

// DraftHistoryResponse contains the response data from the DraftHistory endpoint
type DraftHistoryResponse struct {
	DraftHistory []DraftHistoryDraftHistory
}

// GetDraftHistory retrieves data from the drafthistory endpoint
func GetDraftHistory(ctx context.Context, client *stats.Client, req DraftHistoryRequest) (*models.Response[*DraftHistoryResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/drafthistory", params, &rawResp); err != nil {
		return nil, err
	}

	response := &DraftHistoryResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.DraftHistory = make([]DraftHistoryDraftHistory, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				item := DraftHistoryDraftHistory{
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
					ORGANIZATION:      toString(row[10]),
					ORGANIZATION_TYPE: toString(row[11]),
				}
				response.DraftHistory = append(response.DraftHistory, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
