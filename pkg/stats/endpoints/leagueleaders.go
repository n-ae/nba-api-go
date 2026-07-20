package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

type LeagueLeader struct {
	PlayerID int     `json:"PLAYER_ID"`
	Rank     int     `json:"RANK"`
	Player   string  `json:"PLAYER"`
	TeamID   int     `json:"TEAM_ID"`
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

	// LeagueLeaders is the one classic-Stats-API endpoint that wraps its
	// data in a singular "resultSet" object instead of the "resultSets"
	// array every other endpoint uses - confirmed against a live
	// stats.nba.com response. rawStatsResponse's ResultSets always comes
	// back empty against this shape, so this endpoint silently returned
	// zero leaders on every call until this was verified live.
	var rawResp struct {
		ResultSet resultSet `json:"resultSet"`
	}
	if err := client.GetJSON(ctx, "leagueleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueLeadersResponse{}
	if rawResp.ResultSet.Name == "LeagueLeaders" {
		for _, required := range []string{"PLAYER_ID", "RANK", "PLAYER"} {
			found := false
			for _, h := range rawResp.ResultSet.Headers {
				if h == required {
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("LeagueLeaders: LeagueLeaders result set: missing expected column %q: %v", required, rawResp.ResultSet.Headers)
			}
		}
		response.LeagueLeaders = parseLeagueLeaders(rawResp.ResultSet.Headers, rawResp.ResultSet.RowSet)
	}

	return models.NewResponse(response, 200, "", nil), nil
}

// parseLeagueLeaders looks columns up by header name rather than assuming a
// fixed position/count: confirmed live against stats.nba.com, this
// endpoint's column set genuinely varies by PerMode - PerMode=PerGame
// omits PF, AST_TOV, and STL_TOV that PerMode=Totals includes. A strict
// full-header validateHeaders check (as every other endpoint uses) would
// hard-fail every PerGame call, the most common usage. Columns absent from
// a given response decode as the type's zero value.
func parseLeagueLeaders(headers []string, rows [][]interface{}) []LeagueLeader {
	col := make(map[string]int, len(headers))
	for i, h := range headers {
		col[h] = i
	}
	at := func(row []interface{}, name string) interface{} {
		if i, ok := col[name]; ok && i < len(row) {
			return row[i]
		}
		return nil
	}

	leaders := make([]LeagueLeader, 0, len(rows))
	for _, row := range rows {
		leader := LeagueLeader{
			PlayerID: toInt(at(row, "PLAYER_ID")),
			Rank:     toInt(at(row, "RANK")),
			Player:   toString(at(row, "PLAYER")),
			TeamID:   toInt(at(row, "TEAM_ID")),
			Team:     toString(at(row, "TEAM")),
			GP:       toInt(at(row, "GP")),
			MIN:      toFloat(at(row, "MIN")),
			FGM:      toFloat(at(row, "FGM")),
			FGA:      toFloat(at(row, "FGA")),
			FGPct:    toFloat(at(row, "FG_PCT")),
			FG3M:     toFloat(at(row, "FG3M")),
			FG3A:     toFloat(at(row, "FG3A")),
			FG3Pct:   toFloat(at(row, "FG3_PCT")),
			FTM:      toFloat(at(row, "FTM")),
			FTA:      toFloat(at(row, "FTA")),
			FTPct:    toFloat(at(row, "FT_PCT")),
			OREB:     toFloat(at(row, "OREB")),
			DREB:     toFloat(at(row, "DREB")),
			REB:      toFloat(at(row, "REB")),
			AST:      toFloat(at(row, "AST")),
			STL:      toFloat(at(row, "STL")),
			BLK:      toFloat(at(row, "BLK")),
			TOV:      toFloat(at(row, "TOV")),
			PF:       toFloat(at(row, "PF")),
			PTS:      toFloat(at(row, "PTS")),
			EFF:      toFloat(at(row, "EFF")),
			ASTTOV:   toFloat(at(row, "AST_TOV")),
			STLTOV:   toFloat(at(row, "STL_TOV")),
		}
		leaders = append(leaders, leader)
	}
	return leaders
}
