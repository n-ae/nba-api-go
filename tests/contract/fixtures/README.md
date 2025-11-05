# Contract Test Fixtures

This directory contains recorded NBA.com API responses used for contract testing.

## What Are These?

Fixtures are **snapshots of real NBA.com API responses** captured at a specific point in time. They allow us to:

1. Test without hitting live API (faster, no rate limits)
2. Detect when NBA.com changes their API schema
3. Document expected response structures
4. Ensure parsing logic works with real data

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

If a contract test fails, it means either:

1. **NBA.com changed their API** (action required)
   - Review the diff
   - Update SDK structs
   - Update HTTP handlers
   - Re-record fixture
   - Document the breaking change

2. **Fixture is stale** (non-breaking)
   - Re-record fixture
   - Verify tests pass
   - Commit updated fixture

## See Also

- [Contract Tests README](../README.md) - Full documentation
- [Integration Tests](../../integration/README.md) - Live API testing
- [Maintenance Runbook](../../../docs/MAINTENANCE.md) - Operations guide
