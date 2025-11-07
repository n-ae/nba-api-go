package integration

import (
	"context"
	"testing"

	"github.com/n-ae/nba-api-go/pkg/live"
	"github.com/n-ae/nba-api-go/pkg/live/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats"
	statsep "github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TestSimpleSmokeTests runs basic smoke tests for the most critical endpoints
func TestSimpleSmokeTests(t *testing.T) {
	skipIfNotIntegration(t)

	t.Run("PlayerCareerStats", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client := stats.NewDefaultClient()
		req := statsep.PlayerCareerStatsRequest{
			PlayerID: NikolaJokicID,
			PerMode:  parameters.PerModePerGame,
		}

		resp, err := statsep.PlayerCareerStats(ctx, client, req)
		assertNoError(t, err, "PlayerCareerStats failed")

		if resp == nil || resp.Data == nil {
			t.Fatal("Expected response data, got nil")
		}

		if len(resp.Data.SeasonTotalsRegularSeason) == 0 {
			t.Error("Expected career stats, got empty response")
		}

		t.Logf("✓ PlayerCareerStats OK: %d seasons",
			len(resp.Data.SeasonTotalsRegularSeason))
	})

	t.Run("PlayerGameLog", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client := stats.NewDefaultClient()
		req := statsep.PlayerGameLogRequest{
			PlayerID:   NikolaJokicID,
			Season:     parameters.Season(Season2023),
			SeasonType: parameters.SeasonTypeRegular,
		}

		resp, err := statsep.PlayerGameLog(ctx, client, req)
		assertNoError(t, err, "PlayerGameLog failed")

		if resp == nil || resp.Data == nil {
			t.Fatal("Expected response data, got nil")
		}

		if len(resp.Data.PlayerGameLog) == 0 {
			t.Error("Expected game log entries, got empty response")
		}

		t.Logf("✓ PlayerGameLog OK: %d games", len(resp.Data.PlayerGameLog))
	})

	t.Run("LeagueLeaders", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client := stats.NewDefaultClient()
		req := statsep.LeagueLeadersRequest{
			Season:       parameters.Season(Season2023),
			SeasonType:   parameters.SeasonTypeRegular,
			StatCategory: parameters.StatCategoryPoints,
			PerMode:      parameters.PerModePerGame,
			LeagueID:     parameters.LeagueIDNBA,
		}

		resp, err := statsep.LeagueLeaders(ctx, client, req)
		assertNoError(t, err, "LeagueLeaders failed")

		if resp == nil || resp.Data == nil {
			t.Fatal("Expected response data, got nil")
		}

		if len(resp.Data.LeagueLeaders) == 0 {
			t.Error("Expected league leaders, got empty response")
		}

		t.Logf("✓ LeagueLeaders OK: %d leaders", len(resp.Data.LeagueLeaders))
	})

	t.Run("Scoreboard", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client := live.NewDefaultClient()
		resp, err := endpoints.Scoreboard(ctx, client)
		assertNoError(t, err, "Scoreboard failed")

		if resp == nil || resp.Data == nil {
			t.Fatal("Expected response data, got nil")
		}

		gameCount := len(resp.Data.Scoreboard.Games)
		t.Logf("✓ Scoreboard OK: %d games today", gameCount)
	})

	t.Run("InternationalBroadcasterSchedule_CurrentSeason", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client := stats.NewDefaultClient()
		req := statsep.InternationalBroadcasterScheduleRequest{
			LeagueID: parameters.LeagueIDNBA,
			Season:   "2025",
		}

		resp, err := statsep.GetInternationalBroadcasterSchedule(ctx, client, req)
		assertNoError(t, err, "InternationalBroadcasterSchedule failed")

		if resp == nil {
			t.Fatal("Expected response, got nil")
		}

		t.Logf("✓ InternationalBroadcasterSchedule OK: %d games", len(resp.Games))
	})

	t.Run("InternationalBroadcasterSchedule_PreviousSeason", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		client := stats.NewDefaultClient()
		req := statsep.InternationalBroadcasterScheduleRequest{
			LeagueID: parameters.LeagueIDNBA,
			Season:   "2024",
		}

		resp, err := statsep.GetInternationalBroadcasterSchedule(ctx, client, req)
		assertNoError(t, err, "InternationalBroadcasterSchedule failed for 2024 season")

		if resp == nil {
			t.Fatal("Expected response, got nil")
		}

		t.Logf("✓ InternationalBroadcasterSchedule 2024 OK: %d games", len(resp.Games))
	})
}
