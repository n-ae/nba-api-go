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

	fmt.Println("=== NBA International Broadcast Schedule (2025-26 Season) ===")

	req := endpoints.InternationalBroadcasterScheduleRequest{
		LeagueID: parameters.LeagueIDNBA,
		Season:   "2025",
	}

	resp, err := endpoints.GetInternationalBroadcasterSchedule(context.Background(), client, req)
	if err != nil {
		log.Fatalf("Failed to get international broadcast schedule: %v", err)
	}

	if len(resp.Games) == 0 {
		fmt.Println("No games found in the international broadcast schedule.")
		return
	}

	fmt.Printf("Found %d scheduled games with international broadcasts\n", len(resp.Games))
	fmt.Println()

	gamesByDate := make(map[string][]endpoints.ScheduledGame)
	for _, game := range resp.Games {
		gamesByDate[game.Date] = append(gamesByDate[game.Date], game)
	}

	uniqueDates := make([]string, 0, len(gamesByDate))
	seenDates := make(map[string]bool)
	for _, game := range resp.Games {
		if !seenDates[game.Date] {
			uniqueDates = append(uniqueDates, game.Date)
			seenDates[game.Date] = true
		}
	}

	displayLimit := 5
	if len(uniqueDates) < displayLimit {
		displayLimit = len(uniqueDates)
	}

	for i := 0; i < displayLimit; i++ {
		date := uniqueDates[i]
		games := gamesByDate[date]

		fmt.Printf("Date: %s (%s)\n", date, games[0].Day)
		fmt.Println("─────────────────────────────────────────")

		processedGames := make(map[string]bool)

		for _, game := range games {
			gameKey := fmt.Sprintf("%s-%s-%s", game.GameID, game.Date, game.Time)
			if processedGames[gameKey] {
				continue
			}
			processedGames[gameKey] = true

			fmt.Printf("  %s @ %s\n", game.VisitorAbbr, game.HomeAbbr)
			fmt.Printf("  Time: %s\n", game.Time)

			broadcasters := make(map[string]bool)
			for _, b := range game.Broadcasters {
				broadcasters[b.BroadcasterName] = true
			}

			if len(broadcasters) > 0 {
				fmt.Printf("  Broadcasters: ")
				count := 0
				for name := range broadcasters {
					if count > 0 {
						fmt.Printf(", ")
					}
					fmt.Printf("%s", name)
					count++
					if count >= 5 {
						remaining := len(broadcasters) - count
						if remaining > 0 {
							fmt.Printf(" (+%d more)", remaining)
						}
						break
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}
		fmt.Println()
	}

	if len(uniqueDates) > displayLimit {
		fmt.Printf("... and %d more dates with scheduled games\n", len(uniqueDates)-displayLimit)
	}

	totalBroadcasters := make(map[string]int)
	for _, game := range resp.Games {
		for _, broadcaster := range game.Broadcasters {
			totalBroadcasters[broadcaster.BroadcasterName]++
		}
	}

	fmt.Printf("\nTotal unique broadcasters: %d\n", len(totalBroadcasters))
	fmt.Println("\nTop 5 broadcasters by game count:")
	type broadcasterCount struct {
		name  string
		count int
	}
	counts := make([]broadcasterCount, 0, len(totalBroadcasters))
	for name, count := range totalBroadcasters {
		counts = append(counts, broadcasterCount{name, count})
	}

	for i := 0; i < len(counts)-1; i++ {
		for j := i + 1; j < len(counts); j++ {
			if counts[j].count > counts[i].count {
				counts[i], counts[j] = counts[j], counts[i]
			}
		}
	}

	for i := 0; i < 5 && i < len(counts); i++ {
		fmt.Printf("  %d. %s (%d games)\n", i+1, counts[i].name, counts[i].count)
	}
}
