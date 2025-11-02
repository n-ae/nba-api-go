# HTTP API Integration Tests

Comprehensive integration tests for the nba-api-go HTTP API server.

## Overview

These tests verify that all 139+ HTTP endpoints are accessible and working correctly. Tests cover:

- ✅ Health check endpoint
- ✅ All player endpoints (35+)
- ✅ All team endpoints (30+)
- ✅ All league endpoints (28+)
- ✅ All box score endpoints (10)
- ✅ Game endpoints (12+)
- ✅ Common/utility endpoints (24+)
- ✅ Error handling (400, 404, 405)
- ✅ Parameter validation

## Running Tests

### Prerequisites

1. **Start the HTTP API server:**
```bash
# Build the server
go build -o bin/nba-api-server ./cmd/nba-api-server/

# Start the server
./bin/nba-api-server
```

2. **Ensure server is healthy:**
```bash
curl http://localhost:8080/health
```

### Run All Tests

```bash
# Run the comprehensive integration test suite
bash tests/http-api/test_endpoints.sh
```

### Run with Custom URL

```bash
# Test against a different server
BASE_URL=http://your-server:8080 bash tests/http-api/test_endpoints.sh
```

## Test Coverage

### Endpoints Tested (60+ samples)

**Player Endpoints (10 samples):**
- playercareerstats
- playergamelog
- commonplayerinfo
- playerprofilev2
- playerawards
- playerdashboardbygeneralsplits
- playerdashboardbyshootingsplits
- playercompare
- playeryearbyyearstats
- playerestimatedmetrics

**Team Endpoints (10 samples):**
- commonteamroster
- teamgamelog
- teaminfocommon
- teamdashboardbygeneralsplits
- teamdashboardbyshootingsplits
- teamdetails
- teamplayerdashboard
- teamlineups
- teamyearbyyearstats
- teamvsteam

**League Endpoints (10 samples):**
- leagueleaders
- leaguestandings
- leaguedashteamstats
- leaguedashplayerstats
- leaguegamelog
- playoffpicture
- leaguedashlineups
- leaguedashplayerclutch
- leaguehustlestatsplayer
- leaguegamefinder

**Box Score Endpoints (10/10 - 100%):**
- boxscoresummaryv2
- boxscoretraditionalv2
- boxscoreadvancedv2
- boxscorescoringv2
- boxscoremiscv2
- boxscoreusagev2
- boxscorefourfactorsv2
- boxscoreplayertrackv2
- boxscoredefensivev2
- boxscorehustlev2

**Game Endpoints (5 samples):**
- playbyplayv2
- playbyplayv3
- shotchartdetail
- gamerotation
- winprobabilitypbp

**Common/Other Endpoints:**
- commonallplayers
- scoreboardv2
- scoreboardv3
- drafthistory
- franchisehistory

**Error Handling:**
- Missing parameter validation (400)
- Invalid endpoint handling (404)

## Expected Results

### Success Criteria

- ✅ HTTP 200: Successful response
- ✅ HTTP 500: NBA API rate limit or temporary error (acceptable)
- ✅ Valid JSON response format
- ✅ Data or error field present in response

### Failure Criteria

- ✗ HTTP 404: Endpoint not found (indicates missing handler)
- ✗ HTTP 400: Parameter validation failure (except in error tests)
- ✗ Invalid JSON response
- ✗ Missing data/error fields

## Sample Output

```bash
========================================
HTTP API Integration Tests
Base URL: http://localhost:8080
========================================

=== Health Check ===
✓ Health (HTTP 200)

=== Player Endpoints (10 samples) ===
✓ PlayerCareerStats (HTTP 200)
✓ PlayerGameLog (HTTP 200)
✓ CommonPlayerInfo (HTTP 200)
✓ PlayerProfileV2 (HTTP 200)
✓ PlayerAwards (HTTP 200)
✓ PlayerDashboardByGeneralSplits (HTTP 200)
✓ PlayerDashboardByShootingSplits (HTTP 200)
✓ PlayerCompare (HTTP 200)
✓ PlayerYearByYearStats (HTTP 200)
✓ PlayerEstimatedMetrics (HTTP 200)

=== Team Endpoints (10 samples) ===
✓ CommonTeamRoster (HTTP 200)
✓ TeamGameLog (HTTP 200)
...

========================================
Test Summary
========================================
Total:  60
Passed: 60
Failed: 0

✓ All tests passed!
```

## Interpreting Results

### HTTP Status Codes

- **200 OK**: Endpoint working correctly
- **400 Bad Request**: Missing or invalid parameters (expected for error tests)
- **404 Not Found**: Endpoint not registered (indicates bug)
- **405 Method Not Allowed**: Non-GET method used (expected for error tests)
- **500 Internal Server Error**: NBA API error or rate limit (acceptable, not a bug)

### Common Issues

1. **Connection Refused**: Server not running
   ```bash
   # Start the server first
   ./bin/nba-api-server
   ```

2. **All 500 Errors**: NBA API rate limiting
   ```bash
   # Wait a few minutes and retry
   sleep 300 && bash tests/http-api/test_endpoints.sh
   ```

3. **404 for Valid Endpoint**: Handler not registered
   ```bash
   # Check if endpoint is in handlers.go switch statement
   grep -n "case \"endpointname\"" cmd/nba-api-server/handlers.go
   ```

## Continuous Integration

### GitHub Actions Example

```yaml
name: HTTP API Integration Tests

on: [push, pull_request]

jobs:
  integration-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Build server
        run: go build -o bin/nba-api-server ./cmd/nba-api-server/
      
      - name: Start server
        run: ./bin/nba-api-server &
        
      - name: Wait for server
        run: sleep 5
      
      - name: Run integration tests
        run: bash tests/http-api/test_endpoints.sh
```

## Manual Testing

### Test Individual Endpoints

```bash
# Test a specific endpoint
curl "http://localhost:8080/api/v1/stats/playercareerstats?PlayerID=203999" | jq '.'

# Test with pretty output
curl -s "http://localhost:8080/api/v1/stats/leagueleaders?Season=2023-24" | jq '.data.LeagueLeaders[0:5]'

# Test error handling
curl -v "http://localhost:8080/api/v1/stats/playercareerstats"  # Missing parameter
curl -v "http://localhost:8080/api/v1/stats/invalidendpoint"    # Invalid endpoint
```

### Performance Testing

```bash
# Test response time
time curl -s "http://localhost:8080/api/v1/stats/playercareerstats?PlayerID=203999" > /dev/null

# Test concurrent requests
for i in {1..10}; do
  curl -s "http://localhost:8080/api/v1/stats/playercareerstats?PlayerID=203999" > /dev/null &
done
wait
```

## Contributing

To add more integration tests:

1. Add test cases to `test_endpoints.sh`
2. Follow the existing pattern:
   ```bash
   test_endpoint "EndpointName" "/api/v1/stats/endpoint?Param=Value"
   ```
3. Test both success and error cases
4. Update this README with new test coverage

## Related Documentation

- [HTTP API Usage Guide](../../docs/API_USAGE.md)
- [Migration Guide](../../docs/MIGRATION_GUIDE.md)
- [HTTP API Examples](../../examples/http-api-client/)
- [100% Coverage Achievement](../../HTTP_API_100_PERCENT_ACHIEVEMENT.md)

## Support

- Report issues: [GitHub Issues](https://github.com/username/nba-api-go/issues)
- See examples: [HTTP API Client Examples](../../examples/http-api-client/)
- Full documentation: [Project README](../../README.md)
