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

	team, err := static.FindTeamByAbbreviation("GSW")
	if err != nil {
		log.Fatalf("Failed to find team: %v", err)
	}

	fmt.Printf("Historical Stats for %s\n", team.FullName)
	fmt.Printf("====================================\n\n")

	leagueID := parameters.LeagueIDNBA
	perMode := parameters.PerModePerGame

	resp, err := endpoints.GetTeamYearByYearStats(ctx, client, endpoints.TeamYearByYearStatsRequest{
		TeamID:   fmt.Sprintf("%d", team.ID),
		LeagueID: &leagueID,
		PerMode:  &perMode,
	})
	if err != nil {
		log.Fatalf("Failed to get team history: %v", err)
	}

	fmt.Printf("Year-by-Year Performance\n")
	fmt.Printf("========================\n\n")

	recentSeasons := len(resp.Data.TeamStats)
	if recentSeasons > 10 {
		recentSeasons = 10
	}

	fmt.Printf("Last %d Seasons:\n\n", recentSeasons)
	fmt.Printf("%-10s %4s %6s %7s %5s %5s %5s %5s %5s\n",
		"Year", "GP", "Record", "Win%", "PPG", "RPG", "APG", "Rank", "PO")
	fmt.Println("-------------------------------------------------------------------")

	for i := 0; i < recentSeasons; i++ {
		season := resp.Data.TeamStats[i]

		record := fmt.Sprintf("%v-%v", season.WINS, season.LOSSES)
		winPct := fmt.Sprintf("%.3f", season.WIN_PCT)
		playoffs := fmt.Sprintf("%v-%v", season.PO_WINS, season.PO_LOSSES)

		fmt.Printf("%-10v %4v %6s %7s %5.1f %5.1f %5.1f %5v %5s\n",
			season.YEAR,
			season.GP,
			record,
			winPct,
			season.PTS,
			season.REB,
			season.AST,
			season.PTS_RANK,
			playoffs,
		)
	}

	championships := 0
	for _, season := range resp.Data.TeamStats {
		if season.NBA_FINALS_APPEARANCE != "" {
			championships++
		}
	}

	fmt.Printf("\n\nFranchise Summary\n")
	fmt.Printf("=================\n")
	fmt.Printf("Total Seasons: %d\n", len(resp.Data.TeamStats))
	fmt.Printf("NBA Finals Appearances: %d\n", championships)
}
