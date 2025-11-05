package contract

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

const (
	// Test players
	lebronJamesID  = "2544"   // Historic player with lots of data
	nikolaJokicID  = "203999" // Recent MVP
	stephenCurryID = "201939" // Active star

	// Test teams
	lakersTeamID   = "1610612747" // Lakers
	warriorsTeamID = "1610612744" // Warriors
	nuggetsTeamID  = "1610612743" // Nuggets

	// Test season
	testSeason = "2023-24"

	// Test game ID (update to recent game for best results)
	testGameID = "0022300001"
)

// =============================================================================
// PLAYER ENDPOINTS
// =============================================================================

// PlayerCareerStats - Career statistics for a player
func TestPlayerCareerStats_Contract(t *testing.T) {
	fixtureName := "playercareerstats_2544.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.PlayerCareerStatsRequest{
			PlayerID: lebronJamesID,
			PerMode:  parameters.PerModePerGame,
		}

		resp, err := endpoints.PlayerCareerStats(ctx, client, req)
		assertNoError(t, err, "Failed to fetch PlayerCareerStats")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "SeasonTotalsRegularSeason")
	t.Log("✓ PlayerCareerStats validated")
}

// PlayerGameLog - Game-by-game logs for a player
func TestPlayerGameLog_Contract(t *testing.T) {
	fixtureName := "playergamelog_203999_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.PlayerGameLogRequest{
			PlayerID:   nikolaJokicID,
			Season:     parameters.Season(testSeason),
			SeasonType: parameters.SeasonTypeRegular,
		}

		resp, err := endpoints.PlayerGameLog(ctx, client, req)
		assertNoError(t, err, "Failed to fetch PlayerGameLog")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "PlayerGameLog")
	t.Log("✓ PlayerGameLog validated")
}

// CommonPlayerInfo - Basic player information
func TestCommonPlayerInfo_Contract(t *testing.T) {
	fixtureName := "commonplayerinfo_201939.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.CommonPlayerInfoRequest{
			PlayerID: stephenCurryID,
		}

		resp, err := endpoints.CommonPlayerInfo(ctx, client, req)
		assertNoError(t, err, "Failed to fetch CommonPlayerInfo")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "CommonPlayerInfo")
	t.Log("✓ CommonPlayerInfo validated")
}

// PlayerProfileV2 - Comprehensive player profile
func TestPlayerProfileV2_Contract(t *testing.T) {
	fixtureName := "playerprofilev2_2544.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.PlayerProfileV2Request{
			PlayerID: lebronJamesID,
			PerMode:  perModePtr(parameters.PerModePerGame),
			LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
		}

		resp, err := endpoints.GetPlayerProfileV2(ctx, client, req)
		assertNoError(t, err, "Failed to fetch PlayerProfileV2")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "SeasonTotalsRegularSeason")
	t.Log("✓ PlayerProfileV2 validated")
}

// =============================================================================
// TEAM ENDPOINTS
// =============================================================================

// TeamGameLog - Game logs for a team
func TestTeamGameLog_Contract(t *testing.T) {
	fixtureName := "teamgamelog_1610612747_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.TeamGameLogRequest{
			TeamID:     lakersTeamID,
			Season:     parameters.Season(testSeason),
			SeasonType: parameters.SeasonTypeRegular,
		}

		resp, err := endpoints.GetTeamGameLog(ctx, client, req)
		assertNoError(t, err, "Failed to fetch TeamGameLog")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "TeamGameLog")
	t.Log("✓ TeamGameLog validated")
}

// TeamInfoCommon - Basic team information
func TestTeamInfoCommon_Contract(t *testing.T) {
	fixtureName := "teaminfocommon_1610612744.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.TeamInfoCommonRequest{
			TeamID:     warriorsTeamID,
			LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
			SeasonType: seasonTypePtr(parameters.SeasonTypeRegular),
		}

		resp, err := endpoints.GetTeamInfoCommon(ctx, client, req)
		assertNoError(t, err, "Failed to fetch TeamInfoCommon")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "TeamInfoCommon")
	t.Log("✓ TeamInfoCommon validated")
}

// CommonTeamRoster - Team roster
func TestCommonTeamRoster_Contract(t *testing.T) {
	fixtureName := "commonteamroster_1610612743_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.CommonTeamRosterRequest{
			TeamID:   nuggetsTeamID,
			Season:   seasonPtr(parameters.Season(testSeason)),
			LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
		}

		resp, err := endpoints.GetCommonTeamRoster(ctx, client, req)
		assertNoError(t, err, "Failed to fetch CommonTeamRoster")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "CommonTeamRoster")
	t.Log("✓ CommonTeamRoster validated")
}

// TeamDetails - Detailed team information
func TestTeamDetails_Contract(t *testing.T) {
	fixtureName := "teamdetails_1610612743.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.TeamDetailsRequest{
			TeamID: nuggetsTeamID,
		}

		resp, err := endpoints.GetTeamDetails(ctx, client, req)
		assertNoError(t, err, "Failed to fetch TeamDetails")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "TeamBackground")
	t.Log("✓ TeamDetails validated")
}

// =============================================================================
// LEAGUE ENDPOINTS
// =============================================================================

// LeagueLeaders - Top performers
func TestLeagueLeaders_Contract(t *testing.T) {
	fixtureName := "leagueleaders_2023-24_pts.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.LeagueLeadersRequest{
			Season:       parameters.Season(testSeason),
			SeasonType:   parameters.SeasonTypeRegular,
			PerMode:      parameters.PerModePerGame,
			StatCategory: parameters.StatCategory("PTS"),
			LeagueID:     parameters.LeagueIDNBA,
			ActiveFlag:   "",
		}

		resp, err := endpoints.LeagueLeaders(ctx, client, req)
		assertNoError(t, err, "Failed to fetch LeagueLeaders")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "LeagueLeaders")
	t.Log("✓ LeagueLeaders validated")
}

// LeagueStandings - Current league standings
func TestLeagueStandings_Contract(t *testing.T) {
	fixtureName := "leaguestandings_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.LeagueStandingsRequest{
			Season:     seasonPtr(parameters.Season(testSeason)),
			SeasonType: seasonTypePtr(parameters.SeasonTypeRegular),
			LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
		}

		resp, err := endpoints.GetLeagueStandings(ctx, client, req)
		assertNoError(t, err, "Failed to fetch LeagueStandings")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "Standings")
	t.Log("✓ LeagueStandings validated")
}

// LeagueDashPlayerStats - League-wide player stats
func TestLeagueDashPlayerStats_Contract(t *testing.T) {
	fixtureName := "leaguedashplayerstats_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.LeagueDashPlayerStatsRequest{
			Season:     seasonPtr(parameters.Season(testSeason)),
			SeasonType: seasonTypePtr(parameters.SeasonTypeRegular),
			PerMode:    perModePtr(parameters.PerModePerGame),
			LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
		}

		resp, err := endpoints.GetLeagueDashPlayerStats(ctx, client, req)
		assertNoError(t, err, "Failed to fetch LeagueDashPlayerStats")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "LeagueDashPlayerStats")
	t.Log("✓ LeagueDashPlayerStats validated")
}

// LeagueDashTeamStats - League-wide team stats
func TestLeagueDashTeamStats_Contract(t *testing.T) {
	fixtureName := "leaguedashteamstats_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.LeagueDashTeamStatsRequest{
			Season:     seasonPtr(parameters.Season(testSeason)),
			SeasonType: seasonTypePtr(parameters.SeasonTypeRegular),
			PerMode:    perModePtr(parameters.PerModePerGame),
			LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
		}

		resp, err := endpoints.GetLeagueDashTeamStats(ctx, client, req)
		assertNoError(t, err, "Failed to fetch LeagueDashTeamStats")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "LeagueDashTeamStats")
	t.Log("✓ LeagueDashTeamStats validated")
}

// =============================================================================
// GAME/BOXSCORE ENDPOINTS
// =============================================================================

// BoxScoreSummaryV2 - Game summary
func TestBoxScoreSummaryV2_Contract(t *testing.T) {
	fixtureName := "boxscoresummaryv2_0022300001.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.BoxScoreSummaryV2Request{
			GameID: testGameID,
		}

		resp, err := endpoints.GetBoxScoreSummaryV2(ctx, client, req)
		assertNoError(t, err, "Failed to fetch BoxScoreSummaryV2")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "GameSummary")
	t.Log("✓ BoxScoreSummaryV2 validated")
}

// BoxScoreTraditionalV2 - Traditional box score
func TestBoxScoreTraditionalV2_Contract(t *testing.T) {
	fixtureName := "boxscoretraditionalv2_0022300001.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.BoxScoreTraditionalV2Request{
			GameID: testGameID,
		}

		resp, err := endpoints.GetBoxScoreTraditionalV2(ctx, client, req)
		assertNoError(t, err, "Failed to fetch BoxScoreTraditionalV2")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "PlayerStats")
	t.Log("✓ BoxScoreTraditionalV2 validated")
}

// PlayByPlayV2 - Play-by-play data
func TestPlayByPlayV2_Contract(t *testing.T) {
	fixtureName := "playbyplayv2_0022300001.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.PlayByPlayV2Request{
			GameID:      testGameID,
			StartPeriod: stringPtr("0"),
			EndPeriod:   stringPtr("10"),
		}

		resp, err := endpoints.GetPlayByPlayV2(ctx, client, req)
		assertNoError(t, err, "Failed to fetch PlayByPlayV2")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "PlayByPlay")
	t.Log("✓ PlayByPlayV2 validated")
}

// =============================================================================
// COMMON/MISC ENDPOINTS
// =============================================================================

// ScoreboardV2 - Live scoreboard
func TestScoreboardV2_Contract(t *testing.T) {
	fixtureName := "scoreboardv2_2024-01-15.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.ScoreboardV2Request{
			GameDate:  "2024-01-15",
			LeagueID:  leagueIDPtr(parameters.LeagueIDNBA),
			DayOffset: stringPtr("0"),
		}

		resp, err := endpoints.GetScoreboardV2(ctx, client, req)
		assertNoError(t, err, "Failed to fetch ScoreboardV2")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "GameHeader")
	t.Log("✓ ScoreboardV2 validated")
}

// CommonAllPlayers - List of all players
func TestCommonAllPlayers_Contract(t *testing.T) {
	fixtureName := "commonallplayers_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.CommonAllPlayersRequest{
			Season:              parameters.Season(testSeason),
			LeagueID:            leagueIDPtr(parameters.LeagueIDNBA),
			IsOnlyCurrentSeason: stringPtr("1"),
		}

		resp, err := endpoints.GetCommonAllPlayers(ctx, client, req)
		assertNoError(t, err, "Failed to fetch CommonAllPlayers")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "CommonAllPlayers")
	t.Log("✓ CommonAllPlayers validated")
}

// ShotChartDetail - Shot chart data
func TestShotChartDetail_Contract(t *testing.T) {
	fixtureName := "shotchartdetail_201939_2023-24.json"

	if shouldUpdateFixtures() {
		skipIfNotIntegration(t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client := stats.NewDefaultClient()
		req := endpoints.ShotChartDetailRequest{
			PlayerID:       stringPtr(stephenCurryID),
			Season:         parameters.Season(testSeason),
			SeasonType:     parameters.SeasonTypeRegular,
			LeagueID:       leagueIDPtr(parameters.LeagueIDNBA),
			TeamID:         stringPtr("0"),
			ContextMeasure: stringPtr("FGA"),
		}

		resp, err := endpoints.GetShotChartDetail(ctx, client, req)
		assertNoError(t, err, "Failed to fetch ShotChartDetail")

		data, err := json.MarshalIndent(resp, "", "  ")
		assertNoError(t, err, "Failed to marshal response")

		saveFixture(t, fixtureName, data)
	}

	fixture := loadFixture(t, fixtureName)
	validateBasicSchema(t, fixture, "Shot_Chart_Detail")
	t.Log("✓ ShotChartDetail validated")
}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

// validateBasicSchema validates that fixture has basic NBA API response structure
func validateBasicSchema(t *testing.T, fixture []byte, expectedField string) {
	t.Helper()

	var resp struct {
		StatusCode int                    `json:"StatusCode"`
		URL        string                 `json:"URL"`
		Data       map[string]interface{} `json:"Data"`
	}

	err := json.Unmarshal(fixture, &resp)
	assertNoError(t, err, "Failed to unmarshal fixture")

	assertEqual(t, 200, resp.StatusCode, "Expected status code 200")
	assert(t, resp.Data != nil && len(resp.Data) > 0, "Expected non-empty Data field")

	if expectedField != "" {
		_, exists := resp.Data[expectedField]
		assert(t, exists, "Expected field '"+expectedField+"' in Data")
	}
}
