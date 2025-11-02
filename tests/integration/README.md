# Integration Tests for NBA API Go

This directory contains integration tests for all 139 NBA.com API endpoints.

## Overview

Integration tests validate that:
1. Each endpoint can successfully call the real NBA API
2. Response parsing works correctly
3. Type safety is maintained throughout
4. Error handling works as expected

## Running Tests

```bash
# Run all integration tests
go test ./tests/integration/...

# Run specific test suite
go test ./tests/integration/ -run TestPlayerEndpoints

# Skip integration tests (use unit tests only)
go test -short ./...

# Run with verbose output
go test -v ./tests/integration/...
```

## Test Organization

Tests are organized by endpoint category:

- `player_endpoints_test.go` - All player-related endpoints (35+ endpoints)
- `team_endpoints_test.go` - All team-related endpoints (32+ endpoints)
- `league_endpoints_test.go` - League-wide data endpoints (28+ endpoints)
- `game_endpoints_test.go` - Game-level endpoints (12+ endpoints)
- `boxscore_endpoints_test.go` - All box score variants (10 endpoints)
- `tracking_endpoints_test.go` - Player tracking endpoints (10+ endpoints)
- `advanced_endpoints_test.go` - Advanced analytics endpoints (20+ endpoints)

## Framework

The `endpoint_test_framework.go` provides:
- Consistent test setup
- Common validations
- Timeout handling
- Test parameters

## Coverage

Current integration test coverage:
- Player endpoints: 8/35 (23%)
- Team endpoints: 0/32 (0%)
- League endpoints: 0/28 (0%)
- Game endpoints: 0/12 (0%)
- Box score endpoints: 0/10 (0%)
- Tracking endpoints: 0/10 (0%)
- Advanced endpoints: 0/20 (0%)

**Total: 8/139 (5.8%)**

## Contributing

To add tests for new endpoints:

1. Identify the appropriate test file
2. Add a TestEndpoint struct
3. Implement the test function
4. Add to the test suite
5. Run and validate

Example:
```go
{
    Name:        "NewEndpoint",
    Description: "Test NewEndpoint endpoint",
    TestFunc:    testNewEndpoint,
},
```

## Notes

- Tests use real NBA API (require internet)
- Rate limiting applies
- Some endpoints may be seasonal
- Use `testing.Short()` to skip in CI
