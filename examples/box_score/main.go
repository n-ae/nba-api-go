package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/endpoints"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := stats.NewDefaultClient()

	gameID := "0022300001"

	resp, err := endpoints.GetBoxScoreSummaryV2(ctx, client, endpoints.BoxScoreSummaryV2Request{
		GameID: gameID,
	})
	if err != nil {
		log.Fatalf("Failed to get box score: %v", err)
	}

	if len(resp.Data.GameSummary) > 0 {
		summary := resp.Data.GameSummary[0]
		fmt.Printf("Game Summary\n")
		fmt.Printf("============\n")
		fmt.Printf("Game ID: %v\n", summary.GAME_ID)
		fmt.Printf("Date: %v\n", summary.GAME_DATE_EST)
		fmt.Printf("Status: %v\n", summary.GAME_STATUS_TEXT)
		fmt.Printf("Season: %v\n", summary.SEASON)
		fmt.Println()
	}

	if len(resp.Data.LineScore) > 0 {
		fmt.Printf("Line Score\n")
		fmt.Printf("==========\n")
		for _, team := range resp.Data.LineScore {
			fmt.Printf("%s %s: %v points\n",
				team.TEAM_CITY_NAME, team.TEAM_ABBREVIATION, team.PTS)
			fmt.Printf("  Record: %v\n", team.TEAM_WINS_LOSSES)
			fmt.Printf("  Q1: %v, Q2: %v, Q3: %v, Q4: %v\n",
				team.PTS_QTR1, team.PTS_QTR2, team.PTS_QTR3, team.PTS_QTR4)
			fmt.Printf("  FG%%: %.3f, FT%%: %.3f, 3P%%: %.3f\n",
				team.FG_PCT, team.FT_PCT, team.FG3_PCT)
			fmt.Println()
		}
	}

	if len(resp.Data.OtherStats) > 0 {
		fmt.Printf("Other Stats\n")
		fmt.Printf("===========\n")
		for _, stats := range resp.Data.OtherStats {
			fmt.Printf("%s %s\n", stats.TEAM_CITY, stats.TEAM_ABBREVIATION)
			fmt.Printf("  Points in Paint: %v\n", stats.PTS_PAINT)
			fmt.Printf("  2nd Chance Points: %v\n", stats.PTS_2ND_CHANCE)
			fmt.Printf("  Fast Break Points: %v\n", stats.PTS_FB)
			fmt.Printf("  Largest Lead: %v\n", stats.LARGEST_LEAD)
			fmt.Printf("  Turnovers: %v\n", stats.TOTAL_TURNOVERS)
			fmt.Println()
		}
	}

	if len(resp.Data.Officials) > 0 {
		fmt.Printf("Officials\n")
		fmt.Printf("=========\n")
		for _, official := range resp.Data.Officials {
			fmt.Printf("  #%v %s %s\n",
				official.JERSEY_NUM, official.FIRST_NAME, official.LAST_NAME)
		}
	}
}
