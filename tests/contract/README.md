# Contract Tests

Contract tests verify that the NBA.com API responses match our expected schemas. They protect against upstream API changes that could break our SDK.

## Current Coverage

**16 endpoints** with real data validation:

### Player Endpoints (4)
- ✅ PlayerCareerStats - Career statistics by season
- ✅ PlayerGameLog - Game-by-game logs for a season
- ✅ CommonPlayerInfo - Player biographical info
- ✅ PlayerProfileV2 - Comprehensive player profile

### Team Endpoints (4)
- ✅ TeamGameLog - Team game-by-game logs
- ✅ TeamInfoCommon - Basic team information
- ✅ CommonTeamRoster - Current team roster
- ✅ TeamDetails - Detailed team information

### League Endpoints (4)
- ✅ LeagueLeaders - Statistical leaders
- ✅ LeagueStandings - Current standings
- ✅ LeagueDashPlayerStats - League-wide player stats
- ⚠️ LeagueDashTeamStats - (API returns 500, skipped)

### Game/Boxscore Endpoints (3)
- ✅ BoxScoreSummaryV2 - Game summary information
- ✅ BoxScoreTraditionalV2 - Traditional box score stats
- ✅ PlayByPlayV2 - Play-by-play data

### Common/Misc Endpoints (2)
- ✅ ScoreboardV2 - Live scoreboard data
- ✅ CommonAllPlayers - List of all NBA players
- ⚠️ ShotChartDetail - (API returns 500, skipped)

**Total fixtures**: 16 endpoints with 1.7MB of real NBA data

## Purpose

1. **Detect API Drift**: Catch when NBA.com changes response structures
2. **Offline Testing**: Test without live API calls (faster, no rate limits)
3. **Documentation**: Fixtures serve as examples of real API responses
4. **Regression Prevention**: Ensure parsing logic works with real data

## How It Works

### Recording Mode (UPDATE_FIXTURES=1)

Captures live API responses and saves them as fixtures:

```bash
# Record fixtures for all endpoints
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v

# Record fixture for specific test
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run TestPlayerCareerStats
```

**When to record**:
- Adding new endpoint tests
- NBA.com API changes (after verifying changes are legitimate)
- Quarterly maintenance (refresh fixtures)

### Replay Mode (Default)

Tests against recorded fixtures (no network calls):

```bash
# Run all contract tests (fast, offline)
go test ./tests/contract/... -v

# Run specific contract test
go test ./tests/contract/... -v -run TestPlayerCareerStats
```

**When to run**:
- Every CI/CD build
- Before releases
- After code changes to parsing logic

## Fixture Structure

```
tests/contract/fixtures/
├── playercareerstats_203999.json      # Nikola Jokic career stats
├── playergamelog_203999_2023-24.json  # Jokic 2023-24 game log
├── leagueleaders_2023-24_points.json  # 2023-24 scoring leaders
├── scoreboard_2024-01-15.json         # Scoreboard from Jan 15, 2024
└── ...
```

**Naming convention**: `{endpoint}_{params}.json`
- Lowercase endpoint name
- Key parameters separated by underscores
- `.json` extension

## Test Types

### 1. Schema Validation

Ensures response structure matches expectations:

```go
func TestPlayerCareerStats_Schema(t *testing.T) {
    fixture := loadFixture(t, "playercareerstats_203999.json")

    // Verify response parses correctly
    var resp endpoints.PlayerCareerStatsResponse
    err := json.Unmarshal(fixture, &resp.Data)
    assertNoError(t, err)

    // Verify expected fields exist
    assert(t, len(resp.Data.SeasonTotalsRegularSeason) > 0, "Expected seasons")
    assert(t, resp.Data.SeasonTotalsRegularSeason[0].PlayerID != 0, "Expected PlayerID")
}
```

### 2. Data Sanity

Verifies returned data makes sense:

```go
func TestPlayerCareerStats_DataSanity(t *testing.T) {
    fixture := loadFixture(t, "playercareerstats_203999.json")
    // ...

    // Jokic should have multiple seasons
    assert(t, len(resp.Data.SeasonTotalsRegularSeason) >= 5, "Jokic has 5+ seasons")

    // Stats should be reasonable
    for _, season := range resp.Data.SeasonTotalsRegularSeason {
        assert(t, season.GP > 0, "Games played > 0")
        assert(t, season.PTS >= 0, "Points >= 0")
    }
}
```

### 3. Drift Detection

Compares current API response to recorded fixture:

```go
func TestPlayerCareerStats_NoDrift(t *testing.T) {
    skipIfNotIntegration(t)

    // Load recorded fixture
    expected := loadFixture(t, "playercareerstats_203999.json")

    // Fetch current live response
    actual := fetchLive(t, ...)

    // Compare schemas (field names, types)
    compareSchemas(t, expected, actual)
}
```

## Adding New Contract Tests

### Step 1: Add test file

```go
// tests/contract/newendpoint_test.go
package contract

import (
    "testing"
    "github.com/n-ae/nba-api-go/pkg/stats/endpoints"
)

func TestNewEndpoint_Schema(t *testing.T) {
    fixture := loadFixture(t, "newendpoint_params.json")

    var resp endpoints.NewEndpointResponse
    err := json.Unmarshal(fixture, &resp.Data)
    assertNoError(t, err)

    // Add assertions
}
```

### Step 2: Record fixture

```bash
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run TestNewEndpoint
```

### Step 3: Verify fixture

```bash
# Check fixture was created
ls -lh tests/contract/fixtures/newendpoint_params.json

# Review fixture content
jq . tests/contract/fixtures/newendpoint_params.json | head -20
```

### Step 4: Run test

```bash
go test ./tests/contract/... -v -run TestNewEndpoint
```

## Maintenance

### Quarterly Refresh

Every 3 months, refresh fixtures to ensure they're current:

```bash
# 1. Backup old fixtures
cp -r tests/contract/fixtures tests/contract/fixtures.backup

# 2. Record new fixtures
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v

# 3. Review changes
git diff tests/contract/fixtures/

# 4. If NBA.com made breaking changes:
#    - Update SDK endpoint code
#    - Update HTTP handlers
#    - Update documentation
#    - Commit new fixtures

# 5. If no changes or compatible changes:
git add tests/contract/fixtures/
git commit -m "chore: refresh contract test fixtures"
```

### When Tests Fail

**Scenario 1: Schema Changed (Breaking)**
```bash
# NBA.com added/removed fields or changed types
# Action: Update SDK structs and JSON tags

1. Review diff between fixture and live response
2. Update endpoint response struct
3. Update HTTP handler if needed
4. Re-record fixture
5. Verify all tests pass
```

**Scenario 2: Data Changed (Non-breaking)**
```bash
# Player stats updated, new season added, etc.
# Action: Update fixture with current data

UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run TestFailingTest
```

**Scenario 3: Endpoint Deprecated**
```bash
# NBA.com removed endpoint
# Action: Mark deprecated, find replacement

1. Document deprecation in code comments
2. Update documentation
3. Find replacement endpoint
4. Add replacement tests
5. Keep old tests for backward compatibility
```

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Contract Tests

on: [push, pull_request]

jobs:
  contract-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      # Run contract tests (offline, uses fixtures)
      - name: Contract Tests
        run: go test ./tests/contract/... -v

      # Optionally: Check for drift (requires NBA.com access)
      - name: Drift Detection
        run: INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run ".*NoDrift"
        continue-on-error: true  # Don't fail build on drift
```

## Best Practices

1. **Keep fixtures small**: Only include necessary fields in assertions
2. **Use stable test data**: Jokic's career stats won't disappear
3. **Don't over-assert**: Check structure, not specific values
4. **Refresh regularly**: Quarterly or after NBA.com changes
5. **Version control fixtures**: Commit to git for history
6. **Document changes**: Note why fixtures changed in commits

## Troubleshooting

### Fixture Not Found
```bash
Error: fixture not found: playercareerstats_203999.json

Solution: Record the fixture
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run TestPlayerCareerStats
```

### JSON Unmarshal Error
```bash
Error: json: cannot unmarshal string into Go struct field

Solution: NBA.com changed the field type
1. Compare fixture to live response
2. Update SDK struct definition
3. Re-record fixture
```

### Rate Limited
```bash
Error: 429 Too Many Requests

Solution: Wait and retry
1. Contract tests should use fixtures (no rate limit)
2. Only UPDATE_FIXTURES mode hits NBA.com
3. Wait 5-10 seconds between recordings
```

## See Also

- [Integration Tests](../integration/README.md) - Live API testing
- [Maintenance Runbook](../../docs/MAINTENANCE.md) - Operational procedures
- [Contributing Guide](../../CONTRIBUTING.md) - How to add tests
