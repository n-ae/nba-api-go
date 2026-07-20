# CLAUDE.md

This file provides guidance to Claude Code when working with the nba-api-go repository.

## Repository Overview

**nba-api-go** is a production-ready Go SDK and HTTP API server providing type-safe access to 141 NBA Stats API endpoints (all standard endpoints plus international broadcast schedule). The project emphasizes maintainability, minimal dependencies, and solo engineer viability.

**Current Status**: `main` is at the latest tagged release, `v2.1.0` - a reliability and correctness follow-up to `v2.0.0` (see `CHANGELOG.md`'s `[2.1.0]` section for the full list). All 135 metadata-covered SDK endpoints (including `videoevents.go`, regenerated this release via `tools/generator/fieldname_overrides.json`) now have field names/types verified to match generator output. All 6 hand-written endpoints without generator metadata now validate headers in some form (`commonplayerinfo`/`playergamelog`/`teamgamelog` call `validateHeaders` like generated code; `leagueleaders` needed a different, header-name-driven approach - see below); `playercareerstats` already did before this. **The header-validation piece introduced in `v2.0.0` is real, deliberate new failure-mode surface, and live verification against real NBA.com responses remains only partially discharged** - see `CHANGELOG.md`'s `[2.0.0]` section for the original risk callout. Live-verifying `leagueleaders.go` against real `stats.nba.com` traffic (part of `v2.1.0`) surfaced and fixed two real bugs (a singular-`resultSet` envelope shape unique to this endpoint, and a `TEAM_ID` column the struct lacked) plus a genuine API quirk (its column set varies by `PerMode`). That's 1 of 142 endpoint files checked against live traffic; the other 3 hand-written endpoints and the entire 121+-file generated-endpoint surface remain unverified - `stats.nba.com` began rate-limiting the environment this was done in after a handful of requests, so this is a partial, not complete, discharge of the header-validation risk.
**Grade**: See `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-20_8549390.md` for the current assessment of record - do not hardcode a grade here, it goes stale the moment a new assessment lands and nobody remembers to update this file (see that file's own docs-consolidation section for the fix: archive the superseded assessment in the same commit as the new one).
**Maintenance Burden**: ~1.6 hours/week for the hand-written core; the generated-endpoint surface currently needs more than that until the verification backlog in the current assessment is cleared - see that document's "Is this too complex for one person?" section.

## Project Architecture

### Core Components

- **SDK Library** (`pkg/stats/`): Type-safe Go SDK with 141 endpoints
- **HTTP API Server** (`cmd/nba-api-server/`): REST API exposing all endpoints
- **Code Generator** (`tools/generator/`): Generates endpoint code from hand-written metadata (a separate Go module - see "Code Generation System" below)
- **Static Data** (`pkg/stats/static/`): 5,135 players, 30 teams (no external DB needed)

### Key Design Principles

1. **Boring tech**: stdlib-only HTTP server, 2 dependencies total
2. **Code generation**: 141 endpoints with 43x productivity gain vs manual
3. **Type safety**: No `interface{}` in generated code, compile-time parameter validation
4. **Solo engineer optimized**: Clear docs, test safety net, minimal operational burden

## Essential Commands

### Development Workflow

```bash
# Quick health check (5 minutes)
go test ./...                    # All unit tests
make test-examples              # Verify all 15 examples compile
make lint                       # golangci-lint
go run ./cmd/nba-api-server     # Start HTTP server
curl http://localhost:8080/health

# Integration tests (requires network)
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# Contract tests (offline, requires fixtures)
go test ./tests/contract/... -v

# Record new contract test fixtures
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v
```

### Building

```bash
# Development build
go build -o bin/nba-api-server ./cmd/nba-api-server

# Production build (optimized)
make build

# Generator tool (separate Go module - must build from within its own directory)
cd tools/generator && go build -o bin/generator .
```

### Common Tasks

```bash
# Add new endpoint (if NBA.com adds one) - see "Code Generation System"
# below for the real generator CLI; there is no automated NBA.com
# response analyzer, metadata is written by hand.
cd tools/generator
go run . -endpoint NewEndpoint -dry-run   # preview
go run . -endpoint NewEndpoint            # write pkg/stats/endpoints/newendpoint.go
cd -
go test ./pkg/stats/endpoints/... -run TestNewEndpoint

# Update dependencies (quarterly)
go get -u golang.org/x/text golang.org/x/time
go mod tidy
go test ./...

# Format code
gofmt -w .
make lint
```

## Testing Strategy

### Test Layers

1. **Unit Tests** (`*_test.go` files throughout)
   - Fast, no network calls
   - Test parameter validation, type safety, error handling
   - Run on every commit

2. **Integration Tests** (`tests/integration/`)
   - Smoke tests for critical endpoints
   - Require live NBA.com API access
   - Skip by default (set `INTEGRATION_TESTS=1` to run)
   - Run before releases or when troubleshooting API issues

3. **Contract Tests** (`tests/contract/`)
   - Record/replay system for API responses
   - Detect NBA.com API drift (schema changes)
   - Offline testing with fixtures
   - Run regularly to catch upstream changes early

4. **Example Tests** (`examples/`)
   - Verify all example programs compile
   - Run via `make test-examples`
   - Documentation doubles as integration tests

### When Tests Fail

**Unit tests fail**: Fix immediately, likely a code bug

**Integration tests fail**: Check if NBA.com API changed or network issues
- Review error messages for API response differences
- Check NBA.com website for announcements
- May need to update SDK structs or parameters

**Contract tests fail**: NBA.com API schema changed
- Run `git diff` on fixture to see what changed
- Update SDK structs in `pkg/stats/endpoints/`
- Update HTTP handlers if needed
- Re-record fixture with `UPDATE_FIXTURES=1`
- Document breaking change in CHANGELOG.md

## Code Generation System

### How It Works

The generator (`tools/generator/` - a separate Go module, its own `go.mod`)
turns hand-written JSON metadata (`tools/generator/metadata/*.json`:
endpoint name, parameters, result-set field names) into Go source via
`text/template`. There is no automated NBA.com response analyzer - field
names come from reading the live API response yourself (or an existing
Python `nba_api` endpoint definition) and writing the metadata by hand.

Field **types** now come from `tools/generator/fieldtypes.json` - an
explicit, hand-reviewed `{"FIELD_NAME": "goType"}` dictionary, the
"explicit per-field type metadata" item from the v2.0.0 plan in the
current assessment. `inferGoType` in `tools/generator/generator.go` is
demoted to a fallback used only when a field has no `fieldtypes.json`
entry (falling back prints a warning to stderr); its naming-convention
heuristic gets some field families wrong in ways that corrupt data
(documented rule by rule in `tools/generator/generator_test.go`'s
`TestInferGoType`, with the fix verified in
`TestFieldTypesOverridesKnownWrongInference`).
`TestAllMetadataFieldsHaveExplicitTypes` fails CI if any field referenced
by committed metadata has no `fieldtypes.json` entry. **This dictionary
seeds from `inferGoType`'s own output plus confirmed corrections (48 so
far: 34 found by reading each affected field's context in the committed
metadata, 14 more by cross-referencing every field name against the type
actually used for it across all 143 committed
`pkg/stats/endpoints/*.go` files) - it is not yet independently verified
against 854 live NBA.com responses.** Treat it as a large improvement over
blind inference, not a completed audit; the regeneration in the next
assessment revision is where remaining drift will surface via the
contract tests.

Some field names legitimately need different types in different
endpoints - `OREB`/`DREB`/`REB`/`AST`/`STL`/`BLK`/`PF`/`PTS` are `float64`
(per-game averages) almost everywhere but `int` (single-game counts) in a
handful of box-score/game-log endpoints - which a flat dictionary can't
represent. `tools/generator/fieldtype_overrides.json` holds these as
explicit `(endpoint, result set, field) -> type` exceptions, checked
before `fieldtypes.json`; see `tools/generator/README.md`'s "Per-endpoint
overrides" section.

Generated code looks up each result set **by name** (`findResultSet` in
`pkg/stats/endpoints/types.go`) and **validates its column headers**
against the field order it's about to index positionally
(`validateHeaders`, comparing against `jsonTags(StructType{})` - the
struct's own `json` tags, not a second hand-maintained list) before
parsing any rows. Previously this was `rawResp.ResultSets[0]`-style
positional indexing with no header check at all: if NBA.com ever
reordered result sets or inserted a column, every field after the change
would silently shift into the wrong struct field. A header mismatch now
returns an error instead. **This has not been verified against live
NBA.com responses** - if any endpoint's metadata field order has quietly
drifted from what NBA.com currently returns, this will surface as a new
error on upgrade rather than the previous silent wrong-field behavior;
see `CHANGELOG.md`'s `[2.0.0]` section for the full risk callout.

**Generated files** (not safe to hand-edit and later regenerate over -
regeneration from current metadata does not reproduce several committed
files byte-for-byte; see the current assessment's "regeneration" findings
before assuming any given file matches what the generator would produce
today):
```
pkg/stats/endpoints/
├── playercareerstats.go       # Generated endpoint
├── playergamelog.go           # Generated endpoint
└── ... (140 more)
```

**Template file** (edit this to change generation):
```
tools/generator/templates/
└── endpoint.tmpl               # Endpoint code template (embedded via go:embed)
```

### Regenerating Endpoints

```bash
cd tools/generator

# Preview without writing
go run . -endpoint PlayerCareerStats -dry-run

# Write to pkg/stats/endpoints/ (default -output resolves to
# <repo-root>/pkg/stats/endpoints regardless of your working directory)
go run . -endpoint PlayerCareerStats

# From an existing metadata file covering one or more endpoints
go run . -metadata metadata/tier1_batch.json

cd -
go test ./pkg/stats/endpoints/...
make test-examples
```

### Adding New Endpoints

When NBA.com adds a new endpoint:

1. Inspect the live response (or the equivalent Python `nba_api` endpoint) to get the endpoint path, parameters, and result-set field names.
2. Write a metadata JSON file under `tools/generator/metadata/` (see `tools/generator/README.md` for the format).
3. Generate code: `cd tools/generator && go run . -metadata metadata/newendpoint.json`.
4. Add each new field name's verified type to `tools/generator/fieldtypes.json` - don't trust `inferGoType`'s fallback guess (see "How It Works" above); `TestAllMetadataFieldsHaveExplicitTypes` fails CI until you do.
5. Add integration test in `tests/integration/`.
6. Add example in `examples/`.
7. Update `CHANGELOG.md`.

## HTTP API Server

### Running Locally

The server has no command-line flags. Configuration is via environment
variables only, read in `cmd/nba-api-server/main.go`:

- `PORT` (default `8080`) - listens on `:$PORT`, all interfaces (there is no separate host/bind-address setting - `-host`/`-port`/`-rate-limit`/`-rate-burst` flags don't exist, despite examples that may circulate suggesting otherwise)
- `LOG_LEVEL` (default `info`) - currently read and logged at startup but does not filter or change log output
- `NBA_API_TIMEOUT` (default `30s`) - positive Go duration limiting upstream NBA API requests
- `CORS_ALLOW_ORIGIN` (default `*`) - value returned in `Access-Control-Allow-Origin`
- Per-IP rate limiting is hardcoded (100 req/s, burst 200; see `cmd/nba-api-server/main.go`'s `NewRateLimiter` call) - not configurable without a code change

```bash
# Development mode
go run ./cmd/nba-api-server

# Production mode
PORT=8080 ./bin/nba-api-server
```

### Key Endpoints

- `GET /health` - Health check (always returns 200 OK if server running)
- `GET /metrics` - JSON metrics (request counts, response times, error rates) - not a Prometheus exposition-format endpoint despite `DEPLOYMENT.md`'s Prometheus scrape example; adapt or add a dedicated exporter if you need that
- `GET /api/v1/stats/playercareerstats` - Example stats endpoint (all Stats API routes live under `/api/v1/stats/`)
- `GET /api/v1/stats/scoreboardv2` - Scoreboard data (there is no separate `/live` route; this server only exposes the Stats API's `ScoreboardV2`/`ScoreboardV3`, not `pkg/live`)

### Deployment Options

See `DEPLOYMENT.md` for:
- systemd service setup
- Container deployment (Containerfile included)
- Kubernetes manifests
- Monitoring setup

## Important Development Notes

### Dependencies

**Keep minimal** (currently 2 dependencies):
- `golang.org/x/text` - Unicode normalization for player search
- `golang.org/x/time` - Rate limiting

**Before adding new dependency**:
1. Check if stdlib can do it
2. Assess maintenance burden
3. Document in ADR
4. Update dependency count in README

### API Stability

**Latest tagged release: v2.1.0** - no breaking changes (reliability/correctness follow-ups to `v2.0.0`; see `CHANGELOG.md`'s `[2.1.0]` section). The last major version bump was `v2.0.0`, for real, extensive breaking changes (public struct field types across ~130 generated endpoint files; see `CHANGELOG.md`'s `[2.0.0]` migration guide). Note that v1.2.0 also shipped a source-breaking change, but in a *minor* release (`stats.Config`/`live.Config` field types, `models.NewAPIError`/`HTTPStatusToError` signatures) - see `CHANGELOG.md`'s retroactive compatibility note under `[1.2.0]`. Treat the "strict semver" promise below as the target, not an unconditional guarantee of the historical record.

**Breaking changes** require:
- Major version bump
- Migration guide in CHANGELOG.md
- Deprecation period if possible
- Update all examples

**Stability Promise:**
- All public APIs in `pkg/` are stable
- Minor versions (1.x.0) add features without breaking changes
- Patch versions (1.0.x) fix bugs without breaking changes
- Features deprecated for at least one minor version before removal

### NBA.com API Challenges

The upstream NBA.com API has quirks:

1. **No official documentation** - reverse engineered
2. **Changes without notice** - monitor with contract tests
3. **Rate limiting** - SDK includes automatic retry/backoff
4. **Seasonal data gaps** - some endpoints 404 in offseason
5. **Inconsistent field names** - handled by generator type inference

### Performance Considerations

**SDK Performance**:
- Target: <100ms per request (network excluded)
- JSON parsing is the bottleneck
- Consider caching for static data (players, teams)

**Server Performance**:
- Target: Handle 1000 req/sec on single core
- Rate limiting prevents NBA.com throttling
- Stateless design allows horizontal scaling

## Documentation Files

### For Users

- `README.md` - Quick start, installation, examples
- `docs/API_USAGE.md` - Detailed SDK usage guide
- `docs/MIGRATION_GUIDE.md` - For nba_api (Python) users
- `DEPLOYMENT.md` - Production deployment guide
- `CHANGELOG.md` - Version history and upgrade guides
- `examples/` - 15 working example programs, plus `examples/http-api-client/` (non-Go, curl/Python/JS)

### For Maintainers

- `docs/MAINTENANCE.md` - **START HERE** - Operational runbook
- `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-20_8549390.md` - Current maintainability assessment (supersedes prior assessments, which live in `docs/archive/` with supersession banners; this revision has no separate companion repository-review document - its findings are folded into the one file)
- `docs/adr/` - Architecture decision records
- `CONTRIBUTING.md` - How to contribute

### Key Documents to Reference

**Before making changes**: Read `docs/MAINTENANCE.md` for procedures

**When adding features**: Check ADRs for architectural guidance

**When troubleshooting**: See `docs/MAINTENANCE.md` Emergency Procedures section

## Maintenance Calendar

### Weekly (10 minutes)
- Run quick health check: `go test ./... && make test-examples`
- Check GitHub issues/PRs
- Monitor any production deployments

### Monthly (30 minutes)
- Review dependencies for security advisories
- Check NBA.com for API announcements
- Update static player data if needed

### Quarterly (2-3 hours)
- Run full integration test suite
- Refresh contract test fixtures
- Dependency updates: `go get -u && go mod tidy`
- Review and update documentation
- Consider performance profiling

### Annually (4-6 hours)
- Full maintainability assessment review
- Major version planning if needed
- Dependency major version updates
- Archive outdated documentation

## Troubleshooting

### "Tests are failing"
1. Check if it's a single test or category (unit/integration/contract)
2. Run with `-v` flag for details: `go test -v ./path/to/test`
3. If integration tests fail, check network and NBA.com API status
4. If contract tests fail, likely NBA.com schema changed - see "When Tests Fail" above

### "Examples won't compile"
1. Verify dependencies: `go mod download`
2. Check for breaking changes in recent commits
3. Run `make test-examples` for detailed output
4. Common issue: Import path typos or missing `go mod tidy`

### "HTTP server returns errors"
1. Check if error is from SDK or NBA.com API
2. Review logs for rate limiting (429 status)
3. Verify endpoint parameters are valid
4. Test with integration tests to isolate issue

### "Generator fails"
1. Confirm you're running from `tools/generator/` (`cd tools/generator && go run . ...`) - it's a separate Go module, so `go run ./tools/generator` from the repo root fails with "main module does not contain package"
2. If using `-metadata`, verify the JSON file path and that it parses (`go run . -metadata path/to/file.json -dry-run` to check without writing)
3. If an inferred field type looks wrong, that's a known heuristic limitation, not necessarily a bug - see `TestInferGoType` in `tools/generator/generator_test.go` for the documented cases, and fix the type by hand in the generated file rather than trusting `inferGoType` blindly
4. There is no `-debug` flag or NBA.com-URL-based `analyze` subcommand - see "Code Generation System" above for the real CLI

## Emergency Procedures

### Production Down
1. Check health endpoint: `curl http://server/health`
2. Review logs for panic/crash
3. Check NBA.com API status
4. Rollback if recent deployment
5. See `docs/MAINTENANCE.md` Emergency section for full runbook

### Multiple API Endpoints Failing
Likely NBA.com made breaking changes:
1. Run contract tests to identify affected endpoints
2. Create GitHub issue documenting failures
3. Review NBA.com website/forums for announcements
4. Plan SDK update sprint (may need to regenerate multiple endpoints)
5. Communicate timeline to users

## Architecture Decision Records (ADRs)

### Current ADRs

- **ADR 001**: Go Replication Strategy (why Go, not Python port)
- **ADR 002**: HTTP API Server Architecture (stdlib-only design)
- **ADR 003**: HTTP Transport Policy (why the default transport clones `http.DefaultTransport` with keep-alives enabled)

### When to Create New ADR

Create ADR when making decisions about:
- Technology choices (new dependency, framework)
- API design (breaking changes, new patterns)
- Architectural patterns (caching layer, database)
- Development processes (testing strategy, release process)

Template: See `docs/adr/000-template.md`

## Security Considerations

### API Keys/Secrets
- NBA.com API currently doesn't require keys (public stats)
- If authentication added: Use environment variables, never commit secrets
- Document in `.env.example`, add `.env` to `.gitignore`

### Input Validation
- All user input validated via type-safe parameters
- HTTP server sanitizes query params
- No SQL injection risk (no database)
- XSS risk minimal (JSON API, no HTML rendering)

### Rate Limiting
- SDK includes automatic rate limiting
- Server enforces per-host limits
- Prevents abuse of NBA.com API

## Contributing Guidelines

See `CONTRIBUTING.md` for full guidelines. Quick checklist:

**Before submitting PR**:
- [ ] All tests pass: `go test ./...`
- [ ] Examples compile: `make test-examples`
- [ ] Code formatted: `gofmt -w .`
- [ ] Linter passes: `make lint`
- [ ] Documentation updated if needed
- [ ] CHANGELOG.md updated for user-facing changes

**PR Description Should Include**:
- What problem does this solve?
- How was it tested?
- Any breaking changes?
- Related issues/ADRs

## Quick Reference

### File Structure
```
nba-api-go/                       # root Go module
├── cmd/
│   └── nba-api-server/          # HTTP API server (main)
├── pkg/
│   ├── client/                  # Core HTTP client
│   │   └── middleware/          # Built-in middleware (retry, rate limit, headers)
│   ├── stats/
│   │   ├── endpoints/           # Generated endpoint code
│   │   ├── parameters/          # Type-safe parameters
│   │   └── static/              # Player/team data
│   ├── live/                    # NBA Live Data API client
│   └── models/                  # Common data structures and error types
├── tools/
│   └── generator/                # Code generator (separate Go module, its own go.mod)
├── tests/
│   ├── integration/              # Live API tests
│   └── contract/                 # Schema validation tests (offline, gitignored fixtures)
├── examples/                    # 15 working Go examples, plus a non-Go http-api-client/
└── docs/                        # Documentation
```

### Import Paths
```go
import (
    "github.com/n-ae/nba-api-go/pkg/client"
    "github.com/n-ae/nba-api-go/pkg/client/middleware"
    "github.com/n-ae/nba-api-go/pkg/stats"
    "github.com/n-ae/nba-api-go/pkg/stats/endpoints"
    "github.com/n-ae/nba-api-go/pkg/stats/parameters"
    "github.com/n-ae/nba-api-go/pkg/stats/static"
)
```

### Common Parameter Patterns
```go
// Endpoint functions take a plain (non-pointer) Request struct; optional
// fields use the parameters package's types directly, not *string/*T
// helper constructors.
req := endpoints.PlayerCareerStatsRequest{
    PlayerID: "203999",
    PerMode:  parameters.PerModePerGame,
    LeagueID: parameters.LeagueIDNBA,
}
```

### Error Handling Pattern
```go
response, err := endpoints.PlayerCareerStats(ctx, client, req)
if err != nil {
    // Check if API error vs network error
    return fmt.Errorf("failed to fetch stats: %w", err)
}
```

## Related Projects

- **nba_api** (Python) - Original inspiration for this project
- **balldontlie** - Alternative NBA API (different data source)
- **nba-go** - Another Go implementation (different scope)

## Support and Community

- **Issues**: GitHub Issues for bugs/features
- **Discussions**: GitHub Discussions for questions
- **Documentation**: All docs in `docs/` directory
- **Examples**: See `examples/` for working code

## Version Information

**Latest tagged release**: v2.1.0 (`main` is at this tag; see `CHANGELOG.md`'s `[Unreleased]` section for what's changed since)
**Go Version**: 1.26.5+ (the `go` directive in `go.mod`; older toolchains cannot build this module)
**Stability**: See the Versioning/API Stability section above - v2.1.0 shipped no breaking changes; v2.0.0 was a deliberate major-version break (public struct field types); v1.2.0 also broke source compatibility but in a minor release, documented with a retroactive compatibility note in `CHANGELOG.md`; v1.3.0 shipped no breaking changes

See `CHANGELOG.md` for full version history.

---

**This file last updated**: 2026-07-20 (v2.1.0 release)
**Maintainability grade**: tracked in the current assessment (see the header of this file), not duplicated here
**Next assessment**: the current assessment of record (`docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-20_8549390.md`) predates the `v2.1.0` fixes - a new assessment reviewing `v2.1.0` (including how much of the header-validation risk it actually discharged) is due; otherwise quarterly
