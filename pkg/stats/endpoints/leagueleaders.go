package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

type LeagueLeader struct {
	PlayerID int     `json:"PLAYER_ID"`
	Rank     int     `json:"RANK"`
	Player   string  `json:"PLAYER"`
	Team     string  `json:"TEAM"`
	GP       int     `json:"GP"`
	MIN      float64 `json:"MIN"`
	FGM      float64 `json:"FGM"`
	FGA      float64 `json:"FGA"`
	FGPct    float64 `json:"FG_PCT"`
	FG3M     float64 `json:"FG3M"`
	FG3A     float64 `json:"FG3A"`
	FG3Pct   float64 `json:"FG3_PCT"`
	FTM      float64 `json:"FTM"`
	FTA      float64 `json:"FTA"`
	FTPct    float64 `json:"FT_PCT"`
	OREB     float64 `json:"OREB"`
	DREB     float64 `json:"DREB"`
	REB      float64 `json:"REB"`
	AST      float64 `json:"AST"`
	STL      float64 `json:"STL"`
	BLK      float64 `json:"BLK"`
	TOV      float64 `json:"TOV"`
	PF       float64 `json:"PF"`
	PTS      float64 `json:"PTS"`
	EFF      float64 `json:"EFF"`
	ASTTOV   float64 `json:"AST_TOV"`
	STLTOV   float64 `json:"STL_TOV"`
}

type LeagueLeadersResponse struct {
	LeagueLeaders []LeagueLeader `json:"LeagueLeaders"`
}

type LeagueLeadersRequest struct {
	LeagueID     parameters.LeagueID
	PerMode      parameters.PerMode
	Season       parameters.Season
	SeasonType   parameters.SeasonType
	StatCategory parameters.StatCategory
	ActiveFlag   string
}

func LeagueLeaders(ctx context.Context, client *stats.Client, req LeagueLeadersRequest) (*models.Response[*LeagueLeadersResponse], error) {
	if req.LeagueID == "" {
		req.LeagueID = parameters.LeagueIDNBA
	}
	if req.PerMode == "" {
		req.PerMode = parameters.PerModeTotals
	}
	if req.SeasonType == "" {
		req.SeasonType = parameters.SeasonTypeRegular
	}
	if req.StatCategory == "" {
		req.StatCategory = parameters.StatCategoryPoints
	}

	params := url.Values{}
	params.Set("LeagueID", req.LeagueID.String())
	params.Set("PerMode", req.PerMode.String())
	if req.Season != "" {
		params.Set("Season", req.Season.String())
	}
	params.Set("SeasonType", req.SeasonType.String())
	params.Set("StatCategory", req.StatCategory.String())
	if req.ActiveFlag != "" {
		params.Set("ActiveFlag", req.ActiveFlag)
	}
	params.Set("Scope", "S")

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leagueleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueLeadersResponse{}
	for _, resultSet := range rawResp.ResultSets {
		if resultSet.Name == "LeagueLeaders" {
			response.LeagueLeaders = parseLeagueLeaders(resultSet.RowSet)
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}

func parseLeagueLeaders(rows [][]interface{}) []LeagueLeader {
	leaders := make([]LeagueLeader, 0, len(rows))
	for _, row := range rows {
		if len(row) < 27 {
			continue
		}
		leader := LeagueLeader{
			PlayerID: toInt(row[0]),
			Rank:     toInt(row[1]),
			Player:   toString(row[2]),
			Team:     toString(row[3]),
			GP:       toInt(row[4]),
			MIN:      toFloat(row[5]),
			FGM:      toFloat(row[6]),
			FGA:      toFloat(row[7]),
			FGPct:    toFloat(row[8]),
			FG3M:     toFloat(row[9]),
			FG3A:     toFloat(row[10]),
			FG3Pct:   toFloat(row[11]),
			FTM:      toFloat(row[12]),
			FTA:      toFloat(row[13]),
			FTPct:    toFloat(row[14]),
			OREB:     toFloat(row[15]),
			DREB:     toFloat(row[16]),
			REB:      toFloat(row[17]),
			AST:      toFloat(row[18]),
			STL:      toFloat(row[19]),
			BLK:      toFloat(row[20]),
			TOV:      toFloat(row[21]),
			PF:       toFloat(row[22]),
			PTS:      toFloat(row[23]),
			EFF:      toFloat(row[24]),
			ASTTOV:   toFloat(row[25]),
			STLTOV:   toFloat(row[26]),
		}
		leaders = append(leaders, leader)
	}
	return leaders
}
