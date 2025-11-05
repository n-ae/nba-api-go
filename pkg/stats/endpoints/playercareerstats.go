package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

type SeasonStat struct {
	PlayerID         int     `json:"PLAYER_ID"`
	SeasonID         string  `json:"SEASON_ID"`
	LeagueID         string  `json:"LEAGUE_ID"`
	TeamID           int     `json:"TEAM_ID"`
	TeamAbbreviation string  `json:"TEAM_ABBREVIATION"`
	PlayerAge        int     `json:"PLAYER_AGE"`
	GP               int     `json:"GP"`
	GS               int     `json:"GS"`
	MIN              float64 `json:"MIN"`
	FGM              float64 `json:"FGM"`
	FGA              float64 `json:"FGA"`
	FGPct            float64 `json:"FG_PCT"`
	FG3M             float64 `json:"FG3M"`
	FG3A             float64 `json:"FG3A"`
	FG3Pct           float64 `json:"FG3_PCT"`
	FTM              float64 `json:"FTM"`
	FTA              float64 `json:"FTA"`
	FTPct            float64 `json:"FT_PCT"`
	OREB             float64 `json:"OREB"`
	DREB             float64 `json:"DREB"`
	REB              float64 `json:"REB"`
	AST              float64 `json:"AST"`
	STL              float64 `json:"STL"`
	BLK              float64 `json:"BLK"`
	TOV              float64 `json:"TOV"`
	PF               float64 `json:"PF"`
	PTS              float64 `json:"PTS"`
}

type CareerTotalStat struct {
	PlayerID int     `json:"PLAYER_ID"`
	LeagueID string  `json:"LEAGUE_ID"`
	TeamID   int     `json:"TEAM_ID"`
	GP       int     `json:"GP"`
	GS       int     `json:"GS"`
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
}

type PlayerCareerStatsResponse struct {
	SeasonTotalsRegularSeason []SeasonStat      `json:"SeasonTotalsRegularSeason"`
	CareerTotalsRegularSeason []CareerTotalStat `json:"CareerTotalsRegularSeason"`
	SeasonTotalsPostSeason    []SeasonStat      `json:"SeasonTotalsPostSeason"`
	CareerTotalsPostSeason    []CareerTotalStat `json:"CareerTotalsPostSeason"`
	SeasonTotalsAllStarSeason []SeasonStat      `json:"SeasonTotalsAllStarSeason"`
	CareerTotalsAllStarSeason []CareerTotalStat `json:"CareerTotalsAllStarSeason"`
	SeasonTotalsCollegeSeason []SeasonStat      `json:"SeasonTotalsCollegeSeason"`
	CareerTotalsCollegeSeason []CareerTotalStat `json:"CareerTotalsCollegeSeason"`
}

type PlayerCareerStatsRequest struct {
	PlayerID string
	PerMode  parameters.PerMode
	LeagueID parameters.LeagueID
}

func PlayerCareerStats(ctx context.Context, client *stats.Client, req PlayerCareerStatsRequest) (*models.Response[*PlayerCareerStatsResponse], error) {
	if req.PlayerID == "" {
		return nil, fmt.Errorf("%w: PlayerID is required", models.ErrInvalidRequest)
	}

	if req.PerMode != "" {
		if err := req.PerMode.Validate(); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErrInvalidRequest, err)
		}
	} else {
		req.PerMode = parameters.PerModeTotals
	}

	if err := req.LeagueID.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrInvalidRequest, err)
	}

	params := url.Values{}
	params.Set("PlayerID", req.PlayerID)
	params.Set("PerMode", req.PerMode.String())
	if req.LeagueID != "" {
		params.Set("LeagueID", req.LeagueID.String())
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "playercareerstats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerCareerStatsResponse{}
	for _, resultSet := range rawResp.ResultSets {
		switch resultSet.Name {
		case "SeasonTotalsRegularSeason":
			response.SeasonTotalsRegularSeason = parseSeasonStats(resultSet.RowSet)
		case "CareerTotalsRegularSeason":
			response.CareerTotalsRegularSeason = parseCareerTotals(resultSet.RowSet)
		case "SeasonTotalsPostSeason":
			response.SeasonTotalsPostSeason = parseSeasonStats(resultSet.RowSet)
		case "CareerTotalsPostSeason":
			response.CareerTotalsPostSeason = parseCareerTotals(resultSet.RowSet)
		case "SeasonTotalsAllStarSeason":
			response.SeasonTotalsAllStarSeason = parseSeasonStats(resultSet.RowSet)
		case "CareerTotalsAllStarSeason":
			response.CareerTotalsAllStarSeason = parseCareerTotals(resultSet.RowSet)
		case "SeasonTotalsCollegeSeason":
			response.SeasonTotalsCollegeSeason = parseSeasonStats(resultSet.RowSet)
		case "CareerTotalsCollegeSeason":
			response.CareerTotalsCollegeSeason = parseCareerTotals(resultSet.RowSet)
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}

func parseSeasonStats(rows [][]interface{}) []SeasonStat {
	stats := make([]SeasonStat, 0, len(rows))
	for _, row := range rows {
		if len(row) < 28 {
			continue
		}
		stat := SeasonStat{
			PlayerID:         toInt(row[0]),
			SeasonID:         toString(row[1]),
			LeagueID:         toString(row[2]),
			TeamID:           toInt(row[3]),
			TeamAbbreviation: toString(row[4]),
			PlayerAge:        toInt(row[5]),
			GP:               toInt(row[6]),
			GS:               toInt(row[7]),
			MIN:              toFloat(row[8]),
			FGM:              toFloat(row[9]),
			FGA:              toFloat(row[10]),
			FGPct:            toFloat(row[11]),
			FG3M:             toFloat(row[12]),
			FG3A:             toFloat(row[13]),
			FG3Pct:           toFloat(row[14]),
			FTM:              toFloat(row[15]),
			FTA:              toFloat(row[16]),
			FTPct:            toFloat(row[17]),
			OREB:             toFloat(row[18]),
			DREB:             toFloat(row[19]),
			REB:              toFloat(row[20]),
			AST:              toFloat(row[21]),
			STL:              toFloat(row[22]),
			BLK:              toFloat(row[23]),
			TOV:              toFloat(row[24]),
			PF:               toFloat(row[25]),
			PTS:              toFloat(row[26]),
		}
		stats = append(stats, stat)
	}
	return stats
}

func parseCareerTotals(rows [][]interface{}) []CareerTotalStat {
	stats := make([]CareerTotalStat, 0, len(rows))
	for _, row := range rows {
		if len(row) < 24 {
			continue
		}
		stat := CareerTotalStat{
			PlayerID: toInt(row[0]),
			LeagueID: toString(row[1]),
			TeamID:   toInt(row[2]),
			GP:       toInt(row[3]),
			GS:       toInt(row[4]),
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
		}
		stats = append(stats, stat)
	}
	return stats
}
