# Contract Tests

Contract tests replay recorded NBA.com API responses against the SDK's parsing logic, offline. In principle they protect against upstream API changes breaking the SDK; in their current state, read the limitations below before relying on that.

## Current State (accurate as of this revision - verify before trusting an older description)

**19 `TestX_Contract` functions** in `endpoints_test.go`, covering player, team, league, boxscore, and misc endpoints.

**Fixtures are not committed to git.** `.gitignore` ignores `fixtures/*.json` (the "keep these essential fixtures" exceptions in that file are commented out, not active). Run `git ls-files tests/contract/fixtures/` to check - as of this writing it returns nothing but `fixtures/README.md`. Consequence: **a clean checkout skips all 19 tests** (`loadFixture` calls `t.Skipf` when a fixture file is missing and `UPDATE_FIXTURES` isn't set) and the package reports `ok`, not a failure or a skip warning you'd notice in a scrollback. If you want these tests to actually validate anything, you need to record fixtures locally first (see below) - and keep them out of CI unless you commit them, since CI runs from a clean checkout too.

**Fixtures record already-parsed SDK output, not the raw NBA.com response.** Recording calls the typed endpoint function and marshals the resulting Go struct (`json.MarshalIndent(resp, "", "  ")`) - the raw `resultSets`/`headers`/`rowSet` shape NBA.com actually returns is never captured. Replay therefore checks "does the parser still produce output shaped like its own prior output," which can't catch the class of drift that matters most here - a reordered or renamed upstream column silently shifting fields (`pkg/stats/endpoints` parses positionally; see the current maintainability assessment for the full analysis). Treat these as regression tests for the Go struct shape, not as upstream-contract verification.

**Assertions are shallow.** `validateBasicSchema` only checks that the fixture's `Data` field is non-empty and, optionally, that one named key exists. It does not check types, specific values, or field counts. A fixture with `null` or near-empty data for an endpoint will pass.

None of this makes the tests useless - they do catch a parser panic, a totally broken response, or a renamed top-level field if you keep local fixtures around - but "contract test" currently overstates what's verified. See the current maintainability assessment (`docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md`) for the planned fix: committed raw fixtures, replayed through `httptest.Server` into the real endpoint functions, asserting result-set names and column headers so upstream drift and generator regressions both get caught.

## Purpose (as designed; see limitations above for what's actually delivered today)

1. **Detect API Drift**: Catch when NBA.com changes response structures
2. **Offline Testing**: Test without live API calls (faster, no rate limits)
3. **Documentation**: Fixtures serve as examples of real API responses
4. **Regression Prevention**: Ensure parsing logic works with real data

## How It Works

### Recording Mode (UPDATE_FIXTURES=1)

Captures live API responses (as already-parsed SDK output - see limitations above) and saves them as fixtures:

```bash
# Record fixtures for all endpoints (requires network access to stats.nba.com)
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v

# Record fixture for one test
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run TestPlayerCareerStats_Contract
```

Both `UPDATE_FIXTURES=1` and `INTEGRATION_TESTS=1` are required together - `UPDATE_FIXTURES` alone doesn't skip the `skipIfNotIntegration` guard each recording path calls first.

**When to record**:
- Adding a new contract test
- After verifying an NBA.com API change is legitimate (not just a transient error)
- Periodically, to keep local fixtures from going stale (there's no enforced cadence today)

### Replay Mode (Default)

Tests against locally recorded fixtures, no network calls - but see "fixtures are not committed" above: this only tests anything if you've recorded fixtures yourself first.

```bash
# Run all contract tests
go test ./tests/contract/... -v

# Run one
go test ./tests/contract/... -v -run TestPlayerCareerStats_Contract
```

## Fixture Structure

```
tests/contract/fixtures/
├── playercareerstats_2544.json         # LeBron James career stats
├── playergamelog_203999_2023-24.json   # Jokić 2023-24 game log
└── ...
```

**Naming convention**: `{endpoint}_{params}.json` - lowercase endpoint name, key parameters separated by underscores, `.json` extension. See the `fixtureName` literal at the top of each `TestX_Contract` function in `endpoints_test.go` for the exact name a given test expects.

## Adding New Contract Tests

Follow the existing pattern in `endpoints_test.go` - each `TestX_Contract` function:

1. Under `if shouldUpdateFixtures() { ... }`: calls `skipIfNotIntegration(t)`, fetches live data via the SDK, marshals the response, and calls `saveFixture`.
2. Unconditionally: calls `loadFixture` (which skips the test if the fixture is missing and recording wasn't requested) and runs assertions - typically `validateBasicSchema` with an expected top-level field name.

```go
func TestNewEndpoint_Contract(t *testing.T) {
    fixtureName := "newendpoint_params.json"

    if shouldUpdateFixtures() {
        skipIfNotIntegration(t)
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        client := stats.NewDefaultClient()
        req := endpoints.NewEndpointRequest{ /* ... */ }
        resp, err := endpoints.NewEndpoint(ctx, client, req)
        assertNoError(t, err, "Failed to fetch NewEndpoint")

        data, err := json.MarshalIndent(resp, "", "  ")
        assertNoError(t, err, "Failed to marshal response")
        saveFixture(t, fixtureName, data)
    }

    fixture := loadFixture(t, fixtureName)
    validateBasicSchema(t, fixture, "ExpectedTopLevelField")
}
```

Then record it: `UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v -run TestNewEndpoint_Contract`.

## When Tests Fail (or stay skipped)

**All tests skip on a clean checkout**: expected, given fixtures aren't committed (see "Current State" above) - not a signal anything is broken, but also not a signal anything is verified.

**`json: cannot unmarshal ...` or a parsing panic against a locally recorded fixture**: NBA.com likely changed a field's type or the SDK's positional parsing shifted - compare the fixture against a fresh recording, then update the SDK struct in `pkg/stats/endpoints/`.

**429 / rate limited while recording**: only `UPDATE_FIXTURES=1` mode hits NBA.com; wait a few seconds between recording runs, or record one test at a time with `-run`.

## See Also

- [Integration Tests](../integration/README.md) - Live API testing
- [Maintenance Runbook](../../docs/MAINTENANCE.md) - Operational procedures
- [Contributing Guide](../../CONTRIBUTING.md) - How to add tests
- [Current Maintainability Assessment](../../docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md) - Full analysis and the planned fix for this test layer's real-contract gap
