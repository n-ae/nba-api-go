# Integration Tests

This directory contains integration tests that verify the NBA API Go SDK works correctly with the live NBA.com API.

## Running Integration Tests

```bash
# Run all integration tests (requires network access)
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# Run specific test
INTEGRATION_TESTS=1 go test ./tests/integration/... -v -run TestPlayerEndpoints

# Run with timeout
INTEGRATION_TESTS=1 go test ./tests/integration/... -v -timeout 5m
```

## Test Categories

1. **Player Endpoints** (`player_test.go`)
   - PlayerCareerStats
   - PlayerGameLog
   - CommonPlayerInfo
   - PlayerProfileV2

2. **Team Endpoints** (`team_test.go`)
   - TeamGameLog
   - TeamInfoCommon
   - CommonTeamRoster

3. **League Endpoints** (`league_test.go`)
   - LeagueLeaders
   - LeagueStandings
   - LeagueDashPlayerStats

4. **Live Endpoints** (`live_test.go`)
   - Scoreboard

## Test Philosophy

- **Smoke tests**: Verify endpoints respond without errors
- **Schema validation**: Ensure response structure matches expectations
- **Data sanity**: Basic checks that returned data is reasonable
- **No brittle assertions**: Don't assert specific values (NBA data changes)

## Adding New Tests

1. Create test function in appropriate file
2. Use `skipIfNotIntegration(t)` helper
3. Use reasonable timeouts (30s default)
4. Test with known good IDs (LeBron James: 2544, Nikola Jokic: 203999)
5. Handle rate limiting gracefully

## Known Test IDs

```go
const (
    LeBronJamesID   = "2544"
    NikolaJokicID   = "203999"
    LakersTeamID    = 1610612747
    NuggetsTeamID   = 1610612743
    Season2023      = "2023-24"
)
```
