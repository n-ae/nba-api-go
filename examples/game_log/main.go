package main

import (
	"context"
	"fmt"
	"log"

	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func main() {
	client := stats.NewDefaultClient()

	req := endpoints.PlayerGameLogRequest{
		PlayerID:   "203999",
		Season:     parameters.NewSeason(2023),
		SeasonType: parameters.SeasonTypeRegular,
	}

	resp, err := endpoints.PlayerGameLog(context.Background(), client, req)
	if err != nil {
		log.Fatalf("Failed to get player game log: %v", err)
	}

	fmt.Printf("=== Nikola Jokić 2023-24 Season Game Log ===\n\n")

	if len(resp.Data.PlayerGameLog) == 0 {
		fmt.Println("No games found")
		return
	}

	fmt.Printf("Found %d games\n\n", len(resp.Data.PlayerGameLog))

	for i, game := range resp.Data.PlayerGameLog {
		if i >= 10 {
			fmt.Printf("\n... and %d more games\n", len(resp.Data.PlayerGameLog)-10)
			break
		}

		result := "W"
		if game.WL != "W" {
			result = "L"
		}

		fmt.Printf("%s | %s (%s) | %d pts, %d reb, %d ast | +/- %d\n",
			game.GameDate,
			game.Matchup,
			result,
			game.PTS,
			game.REB,
			game.AST,
			game.PlusMinus,
		)
	}

	totalPTS := 0
	totalREB := 0
	totalAST := 0
	for _, game := range resp.Data.PlayerGameLog {
		totalPTS += game.PTS
		totalREB += game.REB
		totalAST += game.AST
	}

	gamesPlayed := len(resp.Data.PlayerGameLog)
	fmt.Printf("\n=== Season Averages ===\n")
	fmt.Printf("Games Played: %d\n", gamesPlayed)
	fmt.Printf("PPG: %.1f\n", float64(totalPTS)/float64(gamesPlayed))
	fmt.Printf("RPG: %.1f\n", float64(totalREB)/float64(gamesPlayed))
	fmt.Printf("APG: %.1f\n", float64(totalAST)/float64(gamesPlayed))
}
