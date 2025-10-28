package main

import (
	"context"
	"fmt"
	"log"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/endpoints"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

func main() {
	client := stats.NewDefaultClient()

	fmt.Println("=== NBA Scoring Leaders ===")

	scoringReq := endpoints.LeagueLeadersRequest{
		Season:       parameters.NewSeason(2023),
		SeasonType:   parameters.SeasonTypeRegular,
		StatCategory: parameters.StatCategoryPoints,
		PerMode:      parameters.PerModePerGame,
	}

	scoringResp, err := endpoints.LeagueLeaders(context.Background(), client, scoringReq)
	if err != nil {
		log.Fatalf("Failed to get scoring leaders: %v", err)
	}

	fmt.Println("Top 10 Scorers (PPG):")
	for i, leader := range scoringResp.Data.LeagueLeaders {
		if i >= 10 {
			break
		}
		fmt.Printf("%2d. %-25s (%s) - %.1f PPG\n",
			leader.Rank,
			leader.Player,
			leader.Team,
			leader.PTS,
		)
	}

	fmt.Println("\n=== NBA Rebounding Leaders ===")

	reboundingReq := endpoints.LeagueLeadersRequest{
		Season:       parameters.NewSeason(2023),
		SeasonType:   parameters.SeasonTypeRegular,
		StatCategory: parameters.StatCategoryRebounds,
		PerMode:      parameters.PerModePerGame,
	}

	reboundingResp, err := endpoints.LeagueLeaders(context.Background(), client, reboundingReq)
	if err != nil {
		log.Fatalf("Failed to get rebounding leaders: %v", err)
	}

	fmt.Println("Top 10 Rebounders (RPG):")
	for i, leader := range reboundingResp.Data.LeagueLeaders {
		if i >= 10 {
			break
		}
		fmt.Printf("%2d. %-25s (%s) - %.1f RPG\n",
			leader.Rank,
			leader.Player,
			leader.Team,
			leader.REB,
		)
	}

	fmt.Println("\n=== NBA Assist Leaders ===")

	assistReq := endpoints.LeagueLeadersRequest{
		Season:       parameters.NewSeason(2023),
		SeasonType:   parameters.SeasonTypeRegular,
		StatCategory: parameters.StatCategoryAssists,
		PerMode:      parameters.PerModePerGame,
	}

	assistResp, err := endpoints.LeagueLeaders(context.Background(), client, assistReq)
	if err != nil {
		log.Fatalf("Failed to get assist leaders: %v", err)
	}

	fmt.Println("Top 10 Assist Leaders (APG):")
	for i, leader := range assistResp.Data.LeagueLeaders {
		if i >= 10 {
			break
		}
		fmt.Printf("%2d. %-25s (%s) - %.1f APG\n",
			leader.Rank,
			leader.Player,
			leader.Team,
			leader.AST,
		)
	}
}
