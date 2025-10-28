package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/username/nba-api-go/pkg/live"
	"github.com/username/nba-api-go/pkg/live/endpoints"
)

func main() {
	client := live.NewDefaultClient()

	resp, err := endpoints.Scoreboard(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to get scoreboard: %v", err)
	}

	scoreboard := resp.Data.Scoreboard

	fmt.Printf("=== NBA Scoreboard for %s ===\n\n", scoreboard.GameDate)

	if len(scoreboard.Games) == 0 {
		fmt.Println("No games scheduled for today")
		return
	}

	for i, game := range scoreboard.Games {
		fmt.Printf("Game %d: %s @ %s\n",
			i+1,
			game.AwayTeam.TeamTricode,
			game.HomeTeam.TeamTricode,
		)
		fmt.Printf("Status: %s\n", game.GameStatusText)

		if game.GameStatus > 1 {
			fmt.Printf("Score: %s %d - %d %s",
				game.AwayTeam.TeamTricode,
				game.AwayTeam.Score,
				game.HomeTeam.Score,
				game.HomeTeam.TeamTricode,
			)

			if game.GameStatus == 2 {
				fmt.Printf(" (Period %d, %s)", game.Period, game.GameClock)
			}
			fmt.Println()
		} else {
			fmt.Printf("Scheduled: %s\n", game.GameTimeUTC)
		}

		if game.GameLeaders.HomeLeaders.Points > 0 {
			fmt.Printf("Leaders:\n")
			fmt.Printf("  %s: %s (%d pts, %d reb, %d ast)\n",
				game.HomeTeam.TeamTricode,
				game.GameLeaders.HomeLeaders.Name,
				game.GameLeaders.HomeLeaders.Points,
				game.GameLeaders.HomeLeaders.Rebounds,
				game.GameLeaders.HomeLeaders.Assists,
			)
			fmt.Printf("  %s: %s (%d pts, %d reb, %d ast)\n",
				game.AwayTeam.TeamTricode,
				game.GameLeaders.AwayLeaders.Name,
				game.GameLeaders.AwayLeaders.Points,
				game.GameLeaders.AwayLeaders.Rebounds,
				game.GameLeaders.AwayLeaders.Assists,
			)
		}

		fmt.Println()
	}

	jsonData, err := json.MarshalIndent(resp.Data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Printf("\n=== Full JSON Response ===\n%s\n", string(jsonData))
}
