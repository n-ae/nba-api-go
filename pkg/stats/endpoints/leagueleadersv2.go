package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueLeadersV2Request contains parameters for the LeagueLeadersV2 endpoint
type LeagueLeadersV2Request struct {
	Season       *parameters.Season
	SeasonType   *parameters.SeasonType
	PerMode      *parameters.PerMode
	Scope        *string
	StatCategory *string
	LeagueID     *parameters.LeagueID
}

// LeagueLeadersV2LeagueLeaders represents the LeagueLeaders result set for LeagueLeadersV2
type LeagueLeadersV2LeagueLeaders struct {
	PLAYER_ID int     `json:"PLAYER_ID"`
	RANK      int     `json:"RANK"`
	PLAYER    string  `json:"PLAYER"`
	TEAM_ID   int     `json:"TEAM_ID"`
	TEAM      string  `json:"TEAM"`
	GP        int     `json:"GP"`
	MIN       float64 `json:"MIN"`
	FGM       int     `json:"FGM"`
	FGA       int     `json:"FGA"`
	FG_PCT    float64 `json:"FG_PCT"`
	FG3M      int     `json:"FG3M"`
	FG3A      int     `json:"FG3A"`
	FG3_PCT   float64 `json:"FG3_PCT"`
	FTM       int     `json:"FTM"`
	FTA       int     `json:"FTA"`
	FT_PCT    float64 `json:"FT_PCT"`
	OREB      float64 `json:"OREB"`
	DREB      float64 `json:"DREB"`
	REB       float64 `json:"REB"`
	AST       float64 `json:"AST"`
	STL       float64 `json:"STL"`
	BLK       float64 `json:"BLK"`
	TOV       float64 `json:"TOV"`
	PF        float64 `json:"PF"`
	PTS       float64 `json:"PTS"`
	EFF       string  `json:"EFF"`
	AST_TOV   float64 `json:"AST_TOV"`
	STL_TOV   float64 `json:"STL_TOV"`
}

// LeagueLeadersV2Response contains the response data from the LeagueLeadersV2 endpoint
type LeagueLeadersV2Response struct {
	LeagueLeaders []LeagueLeadersV2LeagueLeaders
}

// GetLeagueLeadersV2 retrieves data from the leagueleadersv2 endpoint
func GetLeagueLeadersV2(ctx context.Context, client *stats.Client, req LeagueLeadersV2Request) (*models.Response[*LeagueLeadersV2Response], error) {
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
	if req.Scope != nil {
		params.Set("Scope", string(*req.Scope))
	}
	if req.StatCategory != nil {
		params.Set("StatCategory", string(*req.StatCategory))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leagueleadersv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueLeadersV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueLeaders = make([]LeagueLeadersV2LeagueLeaders, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := LeagueLeadersV2LeagueLeaders{
					PLAYER_ID: toInt(row[0]),
					RANK:      toInt(row[1]),
					PLAYER:    toString(row[2]),
					TEAM_ID:   toInt(row[3]),
					TEAM:      toString(row[4]),
					GP:        toInt(row[5]),
					MIN:       toFloat(row[6]),
					FGM:       toInt(row[7]),
					FGA:       toInt(row[8]),
					FG_PCT:    toFloat(row[9]),
					FG3M:      toInt(row[10]),
					FG3A:      toInt(row[11]),
					FG3_PCT:   toFloat(row[12]),
					FTM:       toInt(row[13]),
					FTA:       toInt(row[14]),
					FT_PCT:    toFloat(row[15]),
					OREB:      toFloat(row[16]),
					DREB:      toFloat(row[17]),
					REB:       toFloat(row[18]),
					AST:       toFloat(row[19]),
					STL:       toFloat(row[20]),
					BLK:       toFloat(row[21]),
					TOV:       toFloat(row[22]),
					PF:        toFloat(row[23]),
					PTS:       toFloat(row[24]),
					EFF:       toString(row[25]),
					AST_TOV:   toFloat(row[26]),
					STL_TOV:   toFloat(row[27]),
				}
				response.LeagueLeaders = append(response.LeagueLeaders, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
