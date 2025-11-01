package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// DraftCombineStatsRequest contains parameters for the DraftCombineStats endpoint
type DraftCombineStatsRequest struct {
	LeagueID   *parameters.LeagueID
	SeasonYear *string
}

// DraftCombineStatsDraftCombineStats represents the DraftCombineStats result set for DraftCombineStats
type DraftCombineStatsDraftCombineStats struct {
	SEASON                     string  `json:"SEASON"`
	PLAYER_ID                  int     `json:"PLAYER_ID"`
	FIRST_NAME                 string  `json:"FIRST_NAME"`
	LAST_NAME                  string  `json:"LAST_NAME"`
	PLAYER_NAME                string  `json:"PLAYER_NAME"`
	POSITION                   string  `json:"POSITION"`
	HEIGHT_WO_SHOES            string  `json:"HEIGHT_WO_SHOES"`
	HEIGHT_WO_SHOES_FT_IN      string  `json:"HEIGHT_WO_SHOES_FT_IN"`
	HEIGHT_W_SHOES             string  `json:"HEIGHT_W_SHOES"`
	HEIGHT_W_SHOES_FT_IN       string  `json:"HEIGHT_W_SHOES_FT_IN"`
	WEIGHT                     string  `json:"WEIGHT"`
	WINGSPAN                   float64 `json:"WINGSPAN"`
	WINGSPAN_FT_IN             float64 `json:"WINGSPAN_FT_IN"`
	STANDING_REACH             string  `json:"STANDING_REACH"`
	STANDING_REACH_FT_IN       string  `json:"STANDING_REACH_FT_IN"`
	BODY_FAT_PCT               float64 `json:"BODY_FAT_PCT"`
	HAND_LENGTH                string  `json:"HAND_LENGTH"`
	HAND_WIDTH                 string  `json:"HAND_WIDTH"`
	STANDING_VERTICAL_LEAP     string  `json:"STANDING_VERTICAL_LEAP"`
	MAX_VERTICAL_LEAP          string  `json:"MAX_VERTICAL_LEAP"`
	LANE_AGILITY_TIME          string  `json:"LANE_AGILITY_TIME"`
	MODIFIED_LANE_AGILITY_TIME string  `json:"MODIFIED_LANE_AGILITY_TIME"`
	THREE_QUARTER_SPRINT       string  `json:"THREE_QUARTER_SPRINT"`
	BENCH_PRESS                string  `json:"BENCH_PRESS"`
}

// DraftCombineStatsResponse contains the response data from the DraftCombineStats endpoint
type DraftCombineStatsResponse struct {
	DraftCombineStats []DraftCombineStatsDraftCombineStats
}

// GetDraftCombineStats retrieves data from the draftcombinestats endpoint
func GetDraftCombineStats(ctx context.Context, client *stats.Client, req DraftCombineStatsRequest) (*models.Response[*DraftCombineStatsResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.SeasonYear != nil {
		params.Set("SeasonYear", string(*req.SeasonYear))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/draftcombinestats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &DraftCombineStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.DraftCombineStats = make([]DraftCombineStatsDraftCombineStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 24 {
				item := DraftCombineStatsDraftCombineStats{
					SEASON:                     toString(row[0]),
					PLAYER_ID:                  toInt(row[1]),
					FIRST_NAME:                 toString(row[2]),
					LAST_NAME:                  toString(row[3]),
					PLAYER_NAME:                toString(row[4]),
					POSITION:                   toString(row[5]),
					HEIGHT_WO_SHOES:            toString(row[6]),
					HEIGHT_WO_SHOES_FT_IN:      toString(row[7]),
					HEIGHT_W_SHOES:             toString(row[8]),
					HEIGHT_W_SHOES_FT_IN:       toString(row[9]),
					WEIGHT:                     toString(row[10]),
					WINGSPAN:                   toFloat(row[11]),
					WINGSPAN_FT_IN:             toFloat(row[12]),
					STANDING_REACH:             toString(row[13]),
					STANDING_REACH_FT_IN:       toString(row[14]),
					BODY_FAT_PCT:               toFloat(row[15]),
					HAND_LENGTH:                toString(row[16]),
					HAND_WIDTH:                 toString(row[17]),
					STANDING_VERTICAL_LEAP:     toString(row[18]),
					MAX_VERTICAL_LEAP:          toString(row[19]),
					LANE_AGILITY_TIME:          toString(row[20]),
					MODIFIED_LANE_AGILITY_TIME: toString(row[21]),
					THREE_QUARTER_SPRINT:       toString(row[22]),
					BENCH_PRESS:                toString(row[23]),
				}
				response.DraftCombineStats = append(response.DraftCombineStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
