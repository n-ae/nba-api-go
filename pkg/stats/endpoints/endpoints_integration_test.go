// +build integration

package endpoints

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

func TestPlayerCareerStatsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := stats.NewDefaultClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := PlayerCareerStatsRequest{
		PlayerID: "203999",
		PerMode:  parameters.PerModePerGame,
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := PlayerCareerStats(ctx, client, req)
	if err != nil {
		t.Fatalf("PlayerCareerStats failed: %v", err)
	}

	if len(resp.Data.SeasonTotalsRegularSeason) == 0 {
		t.Error("Expected at least one season in regular season stats")
	}

	firstSeason := resp.Data.SeasonTotalsRegularSeason[0]
	if firstSeason.PlayerID != 203999 {
		t.Errorf("Expected PlayerID 203999, got %d", firstSeason.PlayerID)
	}

	if firstSeason.GP == 0 {
		t.Error("Expected games played > 0")
	}
}

func TestPlayerGameLogIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := stats.NewDefaultClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := PlayerGameLogRequest{
		PlayerID:   "203999",
		Season:     parameters.NewSeason(2023),
		SeasonType: parameters.SeasonTypeRegular,
	}

	resp, err := PlayerGameLog(ctx, client, req)
	if err != nil {
		t.Fatalf("PlayerGameLog failed: %v", err)
	}

	if len(resp.Data.PlayerGameLog) == 0 {
		t.Error("Expected at least one game in game log")
	}

	firstGame := resp.Data.PlayerGameLog[0]
	if firstGame.PlayerID != 203999 {
		t.Errorf("Expected PlayerID 203999, got %d", firstGame.PlayerID)
	}
}

func TestCommonPlayerInfoIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := stats.NewDefaultClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := CommonPlayerInfoRequest{
		PlayerID: "203999",
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := CommonPlayerInfo(ctx, client, req)
	if err != nil {
		t.Fatalf("CommonPlayerInfo failed: %v", err)
	}

	if len(resp.Data.CommonPlayerInfo) == 0 {
		t.Error("Expected player info")
	}

	info := resp.Data.CommonPlayerInfo[0]
	if info.PersonID != 203999 {
		t.Errorf("Expected PersonID 203999, got %d", info.PersonID)
	}

	if info.FirstName == "" || info.LastName == "" {
		t.Error("Expected player to have first and last name")
	}
}

func TestLeagueLeadersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := stats.NewDefaultClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := LeagueLeadersRequest{
		Season:       parameters.NewSeason(2023),
		SeasonType:   parameters.SeasonTypeRegular,
		StatCategory: parameters.StatCategoryPoints,
		PerMode:      parameters.PerModePerGame,
	}

	resp, err := LeagueLeaders(ctx, client, req)
	if err != nil {
		t.Fatalf("LeagueLeaders failed: %v", err)
	}

	if len(resp.Data.LeagueLeaders) == 0 {
		t.Error("Expected at least one leader")
	}

	firstLeader := resp.Data.LeagueLeaders[0]
	if firstLeader.Rank != 1 {
		t.Errorf("Expected first leader to have rank 1, got %d", firstLeader.Rank)
	}

	if firstLeader.Player == "" {
		t.Error("Expected leader to have a player name")
	}

	if firstLeader.PTS <= 0 {
		t.Error("Expected leader to have points > 0")
	}
}

func TestTeamGameLogIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := stats.NewDefaultClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := TeamGameLogRequest{
		TeamID:     "1610612743",
		Season:     parameters.NewSeason(2023),
		SeasonType: parameters.SeasonTypeRegular,
	}

	resp, err := GetTeamGameLog(ctx, client, req)
	if err != nil {
		t.Fatalf("TeamGameLog failed: %v", err)
	}

	if len(resp.Data.TeamGameLog) == 0 {
		t.Error("Expected at least one game in team game log")
	}

	firstGame := resp.Data.TeamGameLog[0]
	if firstGame.TeamID != 1610612743 {
		t.Errorf("Expected TeamID 1610612743, got %d", firstGame.TeamID)
	}
}

func TestMain(m *testing.M) {
	if os.Getenv("INTEGRATION_TESTS") != "1" {
		return
	}
	os.Exit(m.Run())
}
