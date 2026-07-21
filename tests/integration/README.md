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

## Known live-traffic blocks (GitHub Actions runners)

`.github/workflows/live-drift.yml` runs a subset of `TestSimpleSmokeTests`,
not the full suite. Two independent manual runs on 2026-07-21 (workflow runs
`29865194310` and `29865360637`, ~2 minutes apart) hit the identical pattern:

- `PlayerCareerStats`/`PlayerGameLog` (`stats.nba.com`) silently hung to the
  exact 30s client timeout both times - no response, not even a rejection.
- The live `Scoreboard` test's `cdn.nba.com` call got an instant, byte-identical
  Akamai "Access Denied" page both times.
- `LeagueLeaders` (also `stats.nba.com`) and both `InternationalBroadcasterSchedule`
  calls succeeded both times.

Same failures, same successes, twice in a row - that's a structural block on
GitHub Actions' shared runner IP ranges for those specific endpoints/hosts,
not transient rate limiting. Running the full suite from a developer machine
(a residential/office IP, not a well-known cloud CI range) may not hit the
same block. If you're debugging a `live-drift.yml` failure, check whether
it's one of the three known-blocked cases above before assuming SDK/schema
drift; if NBA.com's blocking behavior changes, revisit the `-run` filter in
that workflow.

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
