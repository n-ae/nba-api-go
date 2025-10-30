package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/endpoints"
	"github.com/username/nba-api-go/pkg/stats/parameters"
	"github.com/username/nba-api-go/pkg/stats/static"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := stats.NewDefaultClient()

	team, err := static.FindTeamByAbbreviation("LAL")
	if err != nil {
		log.Fatalf("Failed to find team: %v", err)
	}

	fmt.Printf("Team Information for %s\n", team.FullName)
	fmt.Println("==============================")

	leagueID := parameters.LeagueIDNBA
	seasonType := parameters.SeasonTypeRegular

	resp, err := endpoints.GetTeamInfoCommon(ctx, client, endpoints.TeamInfoCommonRequest{
		TeamID:     fmt.Sprintf("%d", team.ID),
		LeagueID:   &leagueID,
		SeasonType: &seasonType,
	})
	if err != nil {
		log.Fatalf("Failed to get team info: %v", err)
	}

	if len(resp.Data.TeamInfoCommon) > 0 {
		info := resp.Data.TeamInfoCommon[0]
		fmt.Printf("Season: %v\n", info.SEASON_YEAR)
		fmt.Printf("Conference: %v\n", info.TEAM_CONFERENCE)
		fmt.Printf("Division: %v\n", info.TEAM_DIVISION)
		fmt.Printf("Record: %v-%v (%.3f)\n", info.W, info.L, info.PCT)
		fmt.Printf("Years Active: %v-%v\n", info.MIN_YEAR, info.MAX_YEAR)
	}

	fmt.Println("\nSeason Rankings:")
	if len(resp.Data.TeamSeasonRanks) > 0 {
		ranks := resp.Data.TeamSeasonRanks[0]
		fmt.Printf("Points: Rank #%v (%.1f PPG)\n", ranks.PTS_RANK, ranks.PTS_PG)
		fmt.Printf("Rebounds: Rank #%v (%.1f RPG)\n", ranks.REB_RANK, ranks.REB_PG)
		fmt.Printf("Assists: Rank #%v (%.1f APG)\n", ranks.AST_RANK, ranks.AST_PG)
		fmt.Printf("Opp Points: Rank #%v (%.1f PPG)\n", ranks.OPP_PTS_RANK, ranks.OPP_PTS_PG)
	}
}
