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

	players, err := static.SearchPlayers("Stephen Curry")
	if err != nil || len(players) == 0 {
		log.Fatalf("Failed to find player: %v", err)
	}
	player := players[0]

	season := parameters.NewSeason(2023)
	seasonType := parameters.SeasonTypeRegular
	playerID := fmt.Sprintf("%d", player.ID)

	resp, err := endpoints.GetShotChartDetail(ctx, client, endpoints.ShotChartDetailRequest{
		PlayerID:   &playerID,
		Season:     season,
		SeasonType: seasonType,
	})
	if err != nil {
		log.Fatalf("Failed to get shot chart: %v", err)
	}

	fmt.Printf("Shot Chart for %s - %s Season\n", player.FullName, season)
	fmt.Printf("===============================================\n\n")

	shotsMade := 0
	shotsAttempted := len(resp.Data.Shot_Chart_Detail)

	zoneStats := make(map[interface{}]struct {
		made      int
		attempted int
	})

	for _, shot := range resp.Data.Shot_Chart_Detail {
		if shot.SHOT_MADE_FLAG == 1 {
			shotsMade++
		}

		zone := shot.SHOT_ZONE_BASIC
		stats := zoneStats[zone]
		stats.attempted++
		if shot.SHOT_MADE_FLAG == 1 {
			stats.made++
		}
		zoneStats[zone] = stats
	}

	fmt.Printf("Overall Shooting\n")
	fmt.Printf("----------------\n")
	fmt.Printf("Shots Made: %d\n", shotsMade)
	fmt.Printf("Shots Attempted: %d\n", shotsAttempted)
	if shotsAttempted > 0 {
		fmt.Printf("FG%%: %.1f%%\n\n", float64(shotsMade)/float64(shotsAttempted)*100)
	}

	fmt.Printf("Shooting by Zone\n")
	fmt.Printf("----------------\n")
	for zone, stats := range zoneStats {
		pct := 0.0
		if stats.attempted > 0 {
			pct = float64(stats.made) / float64(stats.attempted) * 100
		}
		fmt.Printf("%v: %d/%d (%.1f%%)\n",
			zone, stats.made, stats.attempted, pct)
	}

	if len(resp.Data.LeagueAverages) > 0 {
		fmt.Printf("\nLeague Averages\n")
		fmt.Printf("---------------\n")
		for _, avg := range resp.Data.LeagueAverages[:5] {
			fmt.Printf("%v - %v: %.1f%% (%v/%v)\n",
				avg.SHOT_ZONE_BASIC, avg.SHOT_ZONE_RANGE,
				avg.FG_PCT, avg.FGM, avg.FGA)
		}
	}
}
