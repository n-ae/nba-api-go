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

All integration tests currently live as subtests of a single test function,
`TestSimpleSmokeTests` in `simple_smoke_test.go` - there are no separate
per-category files yet:

- `PlayerCareerStats`
- `PlayerGameLog`
- `LeagueLeaders`
- `Scoreboard`
- `InternationalBroadcasterSchedule_CurrentSeason`
- `InternationalBroadcasterSchedule_PreviousSeason`

That's 5 endpoint implementations (one, `InternationalBroadcasterSchedule`,
exercised twice with different seasons) out of 141 total. The other 5
hand-written endpoints (`CommonPlayerInfo`, `TeamGameLog`, `TeamInfoCommon`,
`PlayerCareerStats`'s siblings) and the 121+ generated endpoints have no
integration test coverage here - see `CLAUDE.md`'s header for the current
live-verification status.

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

## Known live-traffic blocks

`.github/workflows/live-drift.yml` runs a subset of `TestSimpleSmokeTests`,
not the full suite. Two independent manual runs on 2026-07-21 (workflow runs
`29865194310` and `29865360637`, ~2 minutes apart, from GitHub Actions'
shared runner IP ranges) hit the identical pattern:

- `PlayerCareerStats`/`PlayerGameLog` (`stats.nba.com`) silently hung to the
  exact 30s client timeout both times - no response, not even a rejection.
- The live `Scoreboard` test's `cdn.nba.com` call got an instant, byte-identical
  Akamai "Access Denied" page both times.
- `LeagueLeaders` (also `stats.nba.com`) and both `InternationalBroadcasterSchedule`
  calls succeeded both times.

**This is not GitHub-Actions-specific.** On 2026-07-22, `PlayerCareerStats`,
`PlayerGameLog`, `CommonPlayerInfo`, and `TeamGameLog` were retested directly
against `stats.nba.com` (bypassing the SDK entirely - raw `curl`, several
User-Agent/header combinations, timeouts up to 30s) from a residential/business
ISP IP (Turkish telecom, `AS34296`, not a cloud/CI range) and every one hung
to a hard timeout on the *first* request, no warm-up or backoff involved.
`LeagueLeaders` succeeded reliably from the same IP in the same session. So
the earlier hypothesis in this section - "a developer machine ... may not hit
the same block" - is wrong; it's not (only) a GitHub-Actions-runner-IP-range
block. The pattern looks like per-endpoint Akamai bot-detection (some
`stats.nba.com` paths are defended much more aggressively than others)
rather than a block keyed on the client IP's ASN/class. `LeagueLeaders` and
`InternationalBroadcasterSchedule` remain the only two hand-written endpoints
confirmed reachable from any environment tried so far; `CommonPlayerInfo` and
`TeamGameLog` join `PlayerCareerStats`/`PlayerGameLog` as confirmed-blocked
from at least two independent networks. If you're debugging a `live-drift.yml`
failure, or trying to live-verify one of the still-unverified hand-written
endpoints, check whether it's one of these known-blocked cases before
assuming SDK/schema drift or environment-specific bad luck; if NBA.com's
blocking behavior changes, revisit the `-run` filter in that workflow.

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
