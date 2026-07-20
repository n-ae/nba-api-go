# Contract Test Fixtures

This directory contains recorded NBA.com API responses used for contract testing.

## What Are These?

Fixtures are snapshots of this SDK's **already-parsed output** for a real NBA.com API call, captured at a specific point in time - not the raw NBA.com response itself. Recording marshals the typed Go struct the endpoint function returns (`json.MarshalIndent(resp, ...)`); the raw `resultSets`/`headers`/`rowSet` shape NBA.com actually sent is never captured, and replay never re-invokes the endpoint's `Get`/parse/`validateHeaders` code path - it only unmarshals the frozen fixture and checks it's non-empty (`validateBasicSchema` in `../endpoints_test.go`). See `../README.md`'s "Current State" section for the full limitation and what would need to change to close it.

Given that, what these fixtures actually give you:

1. Test without hitting live API (faster, no rate limits)
2. A regression check on the *Go struct's own shape* (e.g. a parser panic, or a field silently going missing between recordings)
3. Document expected response structures, as this SDK models them

What they do **not** give you (despite it being the original design intent - see `../README.md`): detecting an NBA.com schema change, or verifying parsing logic against a fresh raw response. A reordered or renamed upstream column would not be caught by replaying these fixtures, since the parser never runs against them.

## Recording Fixtures

```bash
# Record all fixtures (requires network access)
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ../... -v

# Record specific fixture
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ../... -v -run TestPlayerCareerStats
```

## Using Fixtures

```bash
# Run contract tests (offline, uses these fixtures)
go test ../... -v
```

## File Naming

Format: `{endpoint}_{key_params}.json`

Examples:
- `playercareerstats_203999.json` - Nikola Jokic's career stats
- `playergamelog_203999_2023-24.json` - Jokic's 2023-24 game log
- `leagueleaders_2023-24_points.json` - 2023-24 scoring leaders

## Version Control

**Strategy**: Commit a few key fixtures, ignore the rest

**Why?**
- Fixtures can be large (50-500KB each)
- Most can be regenerated when needed
- Key fixtures document common use cases

**Which to commit?**
- High-value endpoints (PlayerCareerStats, LeagueLeaders)
- Examples for documentation
- Complex response structures

**How to commit:**
```bash
# Force-add specific fixture despite .gitignore
git add -f fixtures/playercareerstats_203999.json
git commit -m "test: add contract test fixture for PlayerCareerStats"
```

## Maintenance

### Quarterly Refresh

Every 3 months, refresh fixtures:

```bash
# Backup old fixtures
cp -r fixtures fixtures.backup

# Record new fixtures
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ../... -v

# Review changes
git diff fixtures/

# Commit if needed
git add fixtures/
git commit -m "test: refresh contract test fixtures (Q1 2025)"
```

### When Tests Fail

**In default (offline replay) mode**, a failure means the SDK's struct shape and the committed fixture have drifted apart - most likely a struct field was renamed/removed without re-recording. It does *not* mean NBA.com changed anything; replay never calls NBA.com, so it can't observe that (see "What Are These?" above).

**Only while re-recording** (`UPDATE_FIXTURES=1`, which does call live NBA.com) can you actually detect an upstream change: `git diff fixtures/` after recording will show it. If you see a real upstream schema change this way:
   - Review the diff
   - Update SDK structs
   - Update HTTP handlers
   - Document the breaking change
   - Commit the updated fixture

## See Also

- [Contract Tests README](../README.md) - Full documentation
- [Integration Tests](../../integration/README.md) - Live API testing
- [Maintenance Runbook](../../../docs/MAINTENANCE.md) - Operations guide
