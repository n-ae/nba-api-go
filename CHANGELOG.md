# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.3.0] - 2026-07-20

### Added
- `.github/dependabot.yml` watching the `github-actions` ecosystem and both `gomod` modules (root and `tools/generator`) - automates the check that this session had to do by hand after `golangci-lint-action@v6`/`actions/checkout@v4`/`actions/setup-go@v5` all turned out to be stale. Requires Dependabot to be turned on for this repo in GitHub Settings (Code security and analysis) - the config alone doesn't enable it; as of this change, Dependabot alerts are confirmed disabled at the repo level.
- Facade-level regression tests (`stats.TestNewClient_DefaultHeadersReachTheWire`, `live.TestNewClient_DefaultHeadersReachTheWire`) asserting the actual `User-Agent`/`Referer`/`Accept` bytes an `httptest.Server` receives from a default-config client, not just that some middleware ran.
- `pkg/client/middleware` - the built-in middleware constructors (`WithRetry`, `RetryConfig`/`DefaultRetryConfig`, `WithPerHostRateLimit`, `WithHeaders`, `WithUserAgent`, `WithReferer`, `WithAccept`, the `WithLogging` family) are now importable outside this module. They previously lived at `internal/middleware`, which by Go's own visibility rules was never importable by an external `go get` consumer in the first place - this is a purely additive change, not a relocation of something anyone outside this module could have depended on. `pkg/stats` and `pkg/live` now import the new path; `internal/middleware` no longer exists.
- `stats.Config.AdditionalMiddlewares` / `live.Config.AdditionalMiddlewares` - layers your own middleware on top of whichever chain `Middlewares` resolves to (the defaults if `Middlewares` is empty, or your override if it isn't), instead of requiring `append(stats.DefaultMiddlewares(), yourMiddleware...)` by hand. Purely additive: existing `Middlewares`-only configs are unaffected.
- `tools/generator` unit tests (`TestInferGoType`, 39 subtests) pinning the current field-name-to-Go-type heuristic rule by rule, including the known-wrong cases that produce the data-corruption bugs cataloged in `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md` (display names as `float64`, textual range buckets as `int`, decimal percentages as `string`). These pin current behavior so a future change to the heuristic is a visible, intentional diff, not silent drift.
- `cmd/nba-api-server`'s `Season` query parameter now defaults to the current NBA season (computed from today's date; seasons run October-June) instead of a literal `"2023-24"` that was hardcoded at over 100 call sites and had been stale for multiple seasons. `getSeasonOrDefault`/`currentSeasonDefault` in `cmd/nba-api-server/handlers.go` centralize this - a season rollover, or a change to the default-selection rule, is now a one-place change.
- `tools/generator`'s `TestGenerateFromMetadata_ProducesValidGo` regenerates every committed `metadata/*.json` file into a scratch directory and asserts the output parses as syntactically valid Go. This is a generator/template regression test, not a fidelity check: regenerating today produces output that differs from the committed `pkg/stats/endpoints/*.go` files in real, substantive ways (field types, not just formatting), so a literal "regenerate everything and diff against committed source" CI gate would be permanently red. Reconciling that drift safely (it needs verification against live NBA.com responses, not a blind diff) is the v2.0.0-sized "explicit per-field type metadata, regenerate all 141 endpoints" item tracked in `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md`. What this test does catch now: a broken template or a generator change that emits malformed Go.
- CI: added a `go test -race` step scoped to `pkg/client`, `pkg/stats`, `pkg/live`, and `cmd/nba-api-server` - the packages with actual concurrency (shared clients, middleware, background health polling, per-IP rate limiting) - alongside the existing plain `go test`.

### Fixed
- **The default `stats`/`live` `User-Agent` never actually applied.** `client.NewClient` unconditionally set `User-Agent: nba-api-go/1.0` into the stored headers at construction; `middleware.WithUserAgent` (used by both facades' `DefaultMiddlewares()` to install a browser-style User-Agent) only sets the header when absent, so it never won. Every default-config request - `stats` and `live` alike - was sending `nba-api-go/1.0` instead of the intended `Mozilla/5.0 (...)` value since `client.Config.Headers`/`Timeout` started forwarding. Fixed by no longer injecting a default User-Agent in the generic core client at all; `DefaultUserAgent` remains exported as a fallback for callers who construct `client.Client` directly without going through a facade.
- `client.NewClient` aliased `config.Headers` directly (only `SetHeaders` cloned). Constructing a `Client` could mutate the caller's map, and later mutations on either side (including the unsynchronized `SetHeader`/`AddHeader`) could reach across. `NewClient` now clones `config.Headers` before storing it.
- `tools/generator`'s documented `cd tools/generator && go run . -endpoint X` command failed outright: `loadTemplate` resolved `tools/generator/templates/<name>.tmpl` relative to the process's working directory, which under that same documented command became `tools/generator/tools/generator/templates/...` and didn't exist. Templates are now embedded via `go:embed`, so loading no longer depends on the working directory. The `-output` flag's default also depended on the working directory in the same way; it now resolves to `<repo-root>/pkg/stats/endpoints` from this source file's own compiled-in location, regardless of where `go run`/the built binary is invoked from. `tools/regenerate_remaining.sh`'s generator build step (`go build ./tools/generator` from the repo root) failed the same way `tools/generator` is a separate Go module, not part of the root module's package tree - fixed to build within that module's own directory.

### Changed
- CI: bumped `actions/checkout` (v4 → v7) and `actions/setup-go` (v5 → v7) - both were triggering a "Node 20 deprecated" warning on every run, the same class of staleness that caused `golangci-lint-action@v6` to silently fail to support this project's v2 lint config. No breaking changes applied to either for this project's minimal usage (plain checkout, `go-version-file: go.mod`).
- README: added a CI status badge linking to the Actions workflow, now that it actually passes.
- CI: pinned `govulncheck` to a fixed version instead of `@latest`, so builds don't change behavior without a repository change.
- CI: `tools/generator` now gets its own `go vet`/`go test`/`golangci-lint` steps (previously only built, never vetted, tested, or linted). Fixing the 5 pre-existing lint issues this surfaced (an unchecked `f.Close()`, dead code, and the `goconst`-flagged literals living inside that dead code) was part of making this pass.
- Documentation consolidation per `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md` section 7: `CLAUDE.md` had its stale `cmd/generator` paths, fictional generator/server CLI flags, fictional SDK API examples, wrong HTTP route paths, self-contradicting version/grade, and stale example count corrected - it now points to the current assessment for grade rather than hardcoding one that goes stale. Eight superseded/historical docs (`MAINTAINABILITY.md`, `docs/MAINTAINABILITY_ASSESSMENT.md`, `docs/MAINTAINABILITY_ASSESSMENT_2026-07-19.md`, `docs/REPOSITORY_REVIEW_2026-07-19.md`, `docs/IMPROVEMENTS_COMPLETED.md`, `docs/IMPLEMENTATION_SUMMARY.md`, `docs/V1.0.0_RELEASE_SUMMARY.md`, `docs/RELEASE_NOTES_v1.0.0.md`) moved to `docs/archive/` with supersession banners where relevant; `MANUAL_REGENERATION_GUIDE.md` archived now that the generator is runnable again. `docs/LINT_CLEANUP_PLAN.md` deleted (its 413-issue snapshot from 2026-07-09 is stale - `golangci-lint run ./...` reports 0 issues today). `tests/contract/README.md` rewritten to state plainly that fixtures aren't committed, so a clean checkout skips all 19 contract tests, and that recorded fixtures capture already-parsed SDK output rather than the raw NBA.com response shape. `docs/adr/002-api-server-architecture.md` amended to note `NBA_API_TIMEOUT` is documented but not read by `cmd/nba-api-server` (same note added to `docker-compose.yml`). `DEPLOYMENT.md`'s Prometheus-scrape claim corrected (`/metrics` returns JSON, not the Prometheus exposition format) and its example Dockerfile's Go version bumped to match `go.mod`. Assorted stale Go-version/dependency-version references fixed in `docs/BENCHMARKS.md`, `docs/MAINTENANCE.md`, and both ADRs. Several dead links removed or fixed (`docs/README.md`, `docs/MIGRATION_GUIDE.md`, `examples/http-api-client/README.md`, and the newly-archived docs' relative paths, which needed adjusting for their new directory depth).

## [1.2.0] - 2026-07-19

### Added
- `pkg/client.Middleware`, `RoundTripper`, `RoundTripperFunc`, and `Chain` - the middleware seam is now defined in the importable `pkg/client` package instead of only `internal/middleware`, so external consumers can write their own middleware. `internal/middleware`'s identically-named types are now aliases of these for source compatibility.
- `stats.DefaultMiddlewares()` / `live.DefaultMiddlewares()` return each client's default middleware chain, so a custom `Config.Middlewares` can extend the defaults (`append(stats.DefaultMiddlewares(), yourMiddleware...)`) instead of having to silently replace them (this replace-not-extend behavior itself is unchanged - only extension is now possible).
- `middleware.parseRetryAfter` (internal): the retry middleware now honors a `Retry-After` response header (seconds or HTTP-date form, capped at `RetryConfig.MaxBackoff`) instead of always using the calculated exponential backoff.
- `client.Config.MaxResponseBytes` bounds how much of a response body `Get` reads into memory (default 50 MiB via `client.DefaultMaxResponseBytes`); previously `io.ReadAll` had no upper bound at all. Exceeding the limit returns `models.ErrResponseTooLarge` instead of continuing to buffer. Forwarded through `stats.Config`/`live.Config` as `MaxResponseBytes` alongside `Headers`/`Timeout`.
- `models.APIError.Body` - a truncated (2 KiB max) copy of the failing response body, so a 2am debugging session sees NBA.com's actual error text instead of just a status code.
- `cmd/nba-api-server`'s 142 endpoint handlers now map errors through a single `writeEndpointError` helper: if the SDK error is (or wraps) a `*models.APIError`, the server reuses that same upstream status code (a 404 from NBA.com now reads as 404, a 429 as 429, etc.) instead of the previous blanket 500 for every endpoint failure. Non-`APIError` failures (network errors, decode failures) still fall back to 500.
- `TestEndpointInventory` (`cmd/nba-api-server`): statically counts real SDK endpoint files and server route cases from source and cross-checks both against `/health`'s `EndpointsCount` map, so the five-conflicting-counts problem from the 2026-07-19 assessment can't silently reoccur - an endpoint added to one side without the other, or either without updating `EndpointsCount`, now fails this test instead of drifting.

### Fixed
- Retry middleware no longer retries `context.Canceled`/`context.DeadlineExceeded` transport errors - every further attempt would fail the same way immediately, so retrying just burned through `MaxRetries` for nothing.
- `middleware.WithHeaders` used `Header.Add` for every configured value, so on a retried request (which reuses the same `*http.Request`) each attempt piled another copy of the same headers onto the previous attempt's. The first value per key now uses `Header.Set`.
- `client.SetHeaders` aliased the caller's `http.Header` map directly; it now clones it, so later mutations to either side can't unexpectedly affect the other.
- Removed a `sync.Mutex` wrapper around `internal/middleware`'s per-host `rate.Limiter.Wait`/`Allow` - `rate.Limiter` is already safe for concurrent use, and the mutex serialized every caller through `Wait`'s blocking delay.

### Changed
- README's middleware example now imports `pkg/client` (and, via the `stats`/`live` facades, `DefaultMiddlewares()`) instead of `internal/middleware`, so it's actually compilable from an external `go get` consumer - verified against a real external module during this change, not just read.
- **Breaking:** `stats.Config`/`live.Config`'s `Timeout` is now `time.Duration` (was `int`, meaning seconds) and `Headers` is now `http.Header` (was `map[string]string`). `Headers: map[string]string{"K": "V"}` becomes `Headers: http.Header{"K": {"V"}}`; `Timeout: 10` (meaning 10s) becomes `Timeout: 10 * time.Second`. A bare `Timeout: 10` still compiles under the new type (it means 10 *nanoseconds*) - if you set this field, update the literal, don't just leave the number.
- **Breaking:** `models.NewAPIError` and `models.HTTPStatusToError` both gained a `body []byte` parameter (to populate the new `APIError.Body` field).
- `internal/middleware` went from 0% test coverage to ~60%: added tests for retry (retryable-status retries, backoff cap, context-cancellation exit, permanent-vs-transient transport errors, `Retry-After` parsing and capping), per-host rate limiting (hosts are independent, concurrent `Wait` is race-clean), and header idempotency (`WithHeaders` doesn't duplicate across repeated applications to the same request, `WithUserAgent`/`WithReferer`/`WithAccept` don't override an existing value).
- `client`'s default transport now clones `http.DefaultTransport` (restoring `Proxy: ProxyFromEnvironment`, `ForceAttemptHTTP2`, and stdlib connect timeouts) with keep-alives left enabled, instead of a mostly zero-valued `http.Transport{DisableKeepAlives: true, MaxIdleConns: 1, ...}` with no recorded rationale for either choice. `TLSHandshakeTimeout` (30s) and `ResponseHeaderTimeout` (60s) are still explicitly set, generous versus NBA.com/Akamai's occasionally slow responses. See **ADR 003: HTTP Transport Policy** (`docs/adr/003-http-transport-policy.md`) for the full analysis, including why this could not be empirically benchmarked against live NBA.com and the conditions under which it should be revisited.

### Compatibility note (added 2026-07-19, retroactive)
- The two "Breaking" entries above originally claimed the old `Headers map[string]string`/`Timeout int`/4-arg-`NewAPIError`/2-arg-`HTTPStatusToError` shapes "have never appeared in a tagged release" or "never appeared in a tagged release with the old signature." **That claim was false and has been removed above.** `git show v1.1.7:pkg/stats/stats.go` and `git show v1.1.7:pkg/models/errors.go` both show the old shapes present and public in the `v1.1.7` tag - the actual, checkable nuance is that `stats.Config.Headers`/`Timeout` were *ignored* (not forwarded to the underlying client) until v1.1.7-and-later commits, which is a behavioral fact, not a source-compatibility one. Code compiled against the `v1.1.7` tag using the old field types or error-constructor signatures does not compile against `v1.2.0` unmodified. This is a real source break in a minor release, contrary to this project's semver guarantee (see the Versioning section below); see the migration examples above for the mechanical fix.
- Separately, this v1.2.0 release also shipped with the User-Agent shadowing bug described under `[Unreleased] / Fixed` above: every default-config `stats`/`live` request sent `nba-api-go/1.0` instead of the intended browser-style User-Agent. That is now fixed on `main`, but the `v1.2.0` tag itself is immutable and still has the bug; if you're pinned to `v1.2.0`, either upgrade or set `Config.Headers`'s `User-Agent` explicitly as a workaround.
- The `v1.2.0` tag commit's CI run also failed at the `golangci-lint` step (an incompatible `golangci-lint-action@v6`/`version: latest` combination against this project's go 1.26.5 + v2 lint config - fixed in the two commits immediately following the tag, and pinned for good under `[Unreleased]`). Tags are immutable, so `v1.2.0`'s release-commit CI result stays red permanently; treat it as a known, already-corrected process gap rather than a defect in the tagged code itself.

## [1.1.7] - 2026-07-11

### Changed
- Bumped `go` directive (both `go.mod` and `tools/generator/go.mod`) from 1.25.3 to 1.26.5.
- Updated `golang.org/x/text` v0.30.0 → v0.40.0 and `golang.org/x/time` v0.14.0 → v0.15.0.
- Updated `Containerfile` base image from `golang:1.25-alpine` to `golang:1.26-alpine`.

### Documentation
- README: added a "Get Player Info (with Date of Birth)" example demonstrating the `DateOfBirth()` accessor methods added in 1.1.6.

### Compatibility note (added 2026-07-19, retroactive)
- The `go` directive bump above (1.25.3 → 1.26.5) is a **minimum-Go-version increase in a patch release**. That's a real compatibility break for anyone on a pinned or hermetic Go 1.25.x toolchain, and it conflicts with this project's own "Backward Compatibility: Minor and patch versions guarantee backward compatibility" promise (see the Versioning section below). `GOTOOLCHAIN=auto` (the Go default) papers over it for most consumers by fetching 1.26.5 automatically, but that's not true for everyone. Future minimum-Go-version increases will ship in a minor release and be called out explicitly in the release title, not folded into a patch alongside dependency bumps.

## [1.1.6] - 2026-07-11

### Added
- `DateOfBirth()` accessor methods on `PlayerInfo` (`CommonPlayerInfo`), `CommonPlayerInfoV2CommonPlayerInfo`, `CommonTeamRosterCommonTeamRoster`, and `CommonTeamRosterV2CommonTeamRoster`. These parse the raw `BIRTHDATE`/`BIRTH_DATE` string (format differs by endpoint) into a `time.Time`. The existing raw string fields are unchanged.

### Fixed
- `gamerotation_test.go` failed to compile after `PT_DIFF` was changed to `float64` in 1.1.5 — the test's expected struct literals still used string values (`"-3"`, `"5"`) and its mock JSON response body quoted `PT_DIFF` as a string, which `toFloat` silently parses as `0`. Both now use numeric literals, matching the documented real API shape.
- `cmd/nba-api-server`'s `version` constant was still `"1.1.3"` despite tags through `v1.1.5`; bumped to match this release.

## [1.1.5] - 2026-07-09 (backfilled 2026-07-19)

This entry was reconstructed from git history on 2026-07-19; the release itself predates this changelog entry.

### Fixed
- **`GameRotation` endpoint**: `PT_DIFF` on `GameRotationAwayTeam`/`GameRotationHomeTeam` was typed `string` even though the live API returns it as a number (e.g. `-6`); corrected to `float64`.

### Documentation
- Added a doc comment on `GameRotationAwayTeam` recording the verified live-response column layout (`GAME_ID, TEAM_ID, TEAM_CITY, TEAM_NAME, PERSON_ID, PLAYER_FIRST, PLAYER_LAST, IN_TIME_REAL, OUT_TIME_REAL, PLAYER_PTS, PT_DIFF, USG_PCT` — 12 columns) and confirming `IN_TIME_REAL`/`OUT_TIME_REAL` are tenths-of-a-second elapsed since game start, continuous across periods rather than reset per quarter.

## [1.1.4] - 2026-07-09 (backfilled 2026-07-19)

This entry was reconstructed from git history on 2026-07-19; the release itself predates this changelog entry.

### Fixed
- Resolved the `golangci-lint` v2 issues surfaced by the 1.1.3 config migration (see `docs/LINT_CLEANUP_PLAN.md`): added targeted `//nolint:errcheck` on deliberately ignored close/write errors, removed unused helper code in `tests/contract/helpers.go` and `tests/integration/helpers.go`, and cleaned up smaller `govet`/`staticcheck` findings across the generated endpoints and the generator template.

## [1.1.3] - 2026-07-09

### Fixed
- **`GameRotation` endpoint**: the SDK parsed the `gamerotation` result sets assuming 11 columns, but the live NBA.com API returns 12 columns with `TEAM_CITY` inserted at index 2. This shifted every field from `TEAM_NAME` onward by one column and silently dropped the true `USG_PCT` value entirely. `GameRotationAwayTeam`/`GameRotationHomeTeam` now include `TEAM_CITY` and read all columns at their correct offsets.
- `PLAYER_LAST` on `GameRotationAwayTeam`/`GameRotationHomeTeam` was incorrectly typed `float64` (a name field parsed with `toFloat`, always yielding `0`); corrected to `string`.

### Added
- Regression test (`pkg/stats/endpoints/gamerotation_test.go`) covering the 12-column `gamerotation` response shape, including a short-row guard case.
- `stats.Config.BaseURL` — allows pointing the stats client at a custom base URL (e.g. a test server), enabling the new regression test.

### Changed
- Migrated `.golangci.yml` to the golangci-lint v2 config schema (the v1 config was silently failing to load under the installed v2 toolchain, so `make lint` had not been running). Disabled `govet`'s `fieldalignment` check, which reorders generated struct fields for memory-layout efficiency at the cost of the generator's intentional NBA API column ordering — see `docs/LINT_CLEANUP_PLAN.md` for the full rationale and the plan for the remaining pre-existing lint debt this migration exposed.

## [1.1.1] - 2025-11-15 (backfilled 2026-07-19)

This entry was reconstructed from git history on 2026-07-19; the release itself predates this changelog entry.

### Added
- `middleware.WithRetry(middleware.DefaultRetryConfig())` in the stats client's default middleware chain — retries were implemented but not actually applied to any client until this release.

### Fixed
- `cmd/nba-api-server`'s `version` constant was still `"0.1.0"`; bumped to match this release.
- Suppressed `errcheck` findings on `json.NewEncoder(...).Encode(...)` in the health and metrics handlers (the encode error is logged, not silently dropped; this satisfies the linter without changing behavior).

## [1.1.0] - 2025-11-07

### Added
- **New endpoint**: `InternationalBroadcasterSchedule` - Access international broadcast schedules for NBA games
  - SDK endpoint: `endpoints.GetInternationalBroadcasterSchedule()`
  - HTTP API route: `/api/v1/stats/internationalbroadcasterschedule`
  - Supports filtering by Season, LeagueID, RegionID, Date, and EST parameters
  - Returns detailed game information including broadcasters, teams, dates, and times
  - Useful for tracking which international broadcasters are showing games
- Example program: `examples/international_broadcast_schedule/` demonstrating broadcast schedule usage
- Comprehensive test coverage for new endpoint:
  - Unit tests for parameter validation
  - Integration tests for 2024 and 2025 seasons
  - Contract test with fixture for schema stability
  - HTTP handler tests for error cases and valid requests

### Changed
- HTTP API server version updated to 1.1.0
- Now supports 140/139+ NBA Stats API endpoints (added 1 additional international schedule endpoint)

### Notes
- Season parameter format: "2025" corresponds to 2025-26 season
- Returns 409+ scheduled games with international broadcast information for 2025-26 season
- All tests passing with `go test ./...`

## [1.0.0] - 2025-11-05

**STABLE RELEASE** - This release marks the project as production-ready with comprehensive testing, documentation, and stability guarantees.

### Added
- **Contract test framework** in `tests/contract/` for API drift detection
  - Record/replay system for NBA.com API responses
  - Schema validation to catch upstream changes
  - Data sanity checks for response content
  - Comprehensive documentation and usage guide
- Integration test framework in `tests/integration/` with smoke tests for key endpoints
- Maintainability assessment document analyzing solo engineer viability (Grade: A, 93/100)
- Maintenance runbook (`docs/MAINTENANCE.md`) with operational procedures
- CHANGELOG.md for tracking project changes
- CLAUDE.md for AI assistant guidance
- Comprehensive improvements documentation

### Changed
- Updated ADR 002 status from "Proposed" to "Accepted" with implementation summary
- Archived ROADMAP.md with clear deprecation notice (project reached 100% endpoint coverage)

### Fixed
- Compilation errors in examples (redundant newlines in fmt.Println)
- Import typos throughout codebase (yourn-ae → n-ae)

### Stability Guarantees
- **Semantic Versioning**: Strict semver compliance starting with v1.0.0
- **Breaking Changes**: Only in major version updates (2.0.0, 3.0.0, etc.)
- **Backward Compatibility**: Minor and patch versions guarantee backward compatibility
- **API Stability**: All public APIs in `pkg/` are stable and will not break without major version bump
- **Deprecation Policy**: Features will be deprecated for at least one minor version before removal

## [0.9.0] - 2024-11-04

### Added
- **HISTORIC MILESTONE**: All 139/139 NBA Stats API endpoints implemented (100% coverage)
- HTTP API server in `cmd/nba-api-server/` exposing all endpoints via REST
- Production-ready features:
  - Health check endpoint (`/health`)
  - Metrics endpoint (`/metrics`)
  - Rate limiting per host
  - CORS support
  - Graceful shutdown
- Multi-stage Containerfile for minimal production images (<20MB)
- Comprehensive deployment guide (systemd, Docker, Kubernetes)
- Migration guide for Python nba_api users (887 lines)
- 14 example programs demonstrating SDK usage

### Changed
- Code generation approach for all endpoints (43x productivity gain)
- Type inference system eliminates `interface{}` usage in generated code

## [0.3.0] - 2024-10-31

### Added
- First batch of 8 generated endpoints via code generation tooling
- Batch generation system producing consistent, type-safe code

## [0.2.0] - 2024-10-28

### Added
- Live API support (`pkg/live/`)
- Scoreboard endpoint for real-time game data
- Static player and team data (5,135 players, 30 teams)
- Accent-insensitive player search
- Benchmarking suite

### Changed
- Improved middleware architecture
- Rate limiting implementation

## [0.1.0] - 2024-10-24

### Added
- Initial SDK implementation
- Core HTTP client with middleware support
- First 5 Stats API endpoints:
  - PlayerCareerStats
  - PlayerGameLog
  - CommonPlayerInfo
  - LeagueLeaders
  - TeamGameLog
- Type-safe parameter system
- Response models with generics
- Context-based timeout handling
- Error handling framework
- Documentation:
  - README with examples
  - ADR 001: Go Replication Strategy
  - ADR 002: HTTP API Server Architecture
  - Contributing guidelines
  - API usage guide

### Infrastructure
- Go 1.21+ requirement
- Minimal dependencies (2 total):
  - golang.org/x/text v0.30.0
  - golang.org/x/time v0.14.0
- Makefile for build automation
- golangci-lint configuration
- GitHub-ready project structure

---

## Release Notes

### Version 1.0.0 - Stable Release

**PRODUCTION READY** - This is the first stable release with comprehensive testing, documentation, and long-term support commitment.

**Stability Highlights:**
- ✅ All 139 NBA Stats API endpoints (100% coverage)
- ✅ Comprehensive test coverage (unit + integration + contract tests)
- ✅ Production-grade maintainability (Grade A: 93/100)
- ✅ Complete operational documentation
- ✅ Semver stability guarantees
- ✅ Minimal dependencies (2 total, both from golang.org/x)

**Testing Infrastructure:**
- Integration tests for live API validation
- Contract tests for API drift detection
- Fixture recording/replay system
- Schema validation to catch upstream changes

**Documentation:**
- Maintenance runbook for operational procedures
- Maintainability assessment documenting solo engineer viability
- Complete API usage guides and examples
- Migration guide for Python nba_api users

**What This Means:**
- Stable public API - no breaking changes without major version bump
- Production-ready for serious applications
- Long-term maintenance commitment (~2 hours/week)
- Quarterly maintenance cycle established

### Version 0.9.0 - Production Ready

This release marks feature completeness for the NBA Stats API with all 139 endpoints implemented. The HTTP API server makes the SDK accessible from any programming language.

**Key Achievements:**
- ✅ 100% endpoint coverage (139/139)
- ✅ Production-ready HTTP API
- ✅ Comprehensive documentation
- ✅ Zero technical debt (no TODOs/FIXMEs)
- ✅ Minimal dependencies (2 total)

### Version 0.1.0 - Initial Release

First public release of the NBA API Go SDK. Provides type-safe access to NBA statistics with excellent developer experience.

---

## Upgrade Guide

### From 1.2.0 to 1.3.0

**No breaking changes.** This release rebuilds the verification infrastructure identified as missing in the 2026-07-19 maintainability assessment; the public API is unchanged.

**What's New:**
- The built-in middleware constructors (`WithRetry`, `WithPerHostRateLimit`, `WithHeaders`, `WithUserAgent`, `WithReferer`, `WithAccept`, `WithLogging`) are now importable from `pkg/client/middleware` (previously unexported at `internal/middleware`).
- `stats.Config.AdditionalMiddlewares` / `live.Config.AdditionalMiddlewares` layer custom middleware on top of the default chain without having to reconstruct it by hand.
- The `tools/generator` CLI actually runs now (`cd tools/generator && go run . -endpoint X`); templates are embedded via `go:embed` instead of resolved relative to the working directory.
- `cmd/nba-api-server`'s `Season` query parameter defaults to the current NBA season instead of a hardcoded stale literal.
- CI now runs `go test -race` on the concurrency-bearing packages and lints/tests `tools/generator` as its own module.

**Fixed:**
- Default `stats`/`live` clients were sending `User-Agent: nba-api-go/1.0` instead of the intended browser-style value, since v1.1.7 - see the `[1.3.0]` section above for the mechanism.
- `client.NewClient` could alias a caller's `Config.Headers` map; it's now cloned at construction.

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.3.0`
2. No code changes required.
3. If you were relying on the incorrect default `User-Agent` value (`nba-api-go/1.0`) for some reason, set it explicitly via `Config.Headers`.

See the `[1.3.0]` section above for the full list of changes.

### From 1.1.7 to 1.2.0

**Two source-breaking changes, both narrowly scoped and neither ever shipped working in a tagged release before now:**

1. `stats.Config`/`live.Config`'s `Timeout` field is `time.Duration` (was `int` seconds) and `Headers` is `http.Header` (was `map[string]string`). If you set either field:
   - `Headers: map[string]string{"K": "V"}` → `Headers: http.Header{"K": {"V"}}`
   - `Timeout: 10` (meaning 10s) → `Timeout: 10 * time.Second` (a bare `Timeout: 10` still compiles but now means 10 *nanoseconds*)
2. If you call `models.NewAPIError` or `models.HTTPStatusToError` directly (most users don't - they're used internally by `pkg/client`), both now take an additional `body []byte` parameter.

If you never set `stats.Config.Timeout`/`Headers` or called the `models` functions above directly, **there is nothing to change** - upgrade with `go get github.com/n-ae/nba-api-go@v1.2.0`.

**What's New:**
- The middleware seam (`Middleware`, `RoundTripper`, `RoundTripperFunc`, `Chain`) is now importable from `pkg/client`, plus `stats.DefaultMiddlewares()`/`live.DefaultMiddlewares()` - see the README's "Middleware" section for a compilable example of writing your own.
- `client.Config.MaxResponseBytes` bounds response body reads (default 50 MiB) instead of the previous unbounded read.
- `models.APIError.Body` carries a truncated copy of the failing response for diagnostics.
- The HTTP server now returns the real upstream status code for endpoint failures (404, 429, etc.) instead of a blanket 500.
- `Retry-After` response headers are now honored by the retry middleware.

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.2.0`
2. If you set `stats.Config.Timeout`/`Headers` or `live.Config.Timeout`/`Headers`, update those literals per above.
3. If you call `models.NewAPIError`/`HTTPStatusToError` directly, add the `body` argument (pass `nil` if you don't have one).

See the `[1.2.0]` section above for the full list of changes.

### From 1.0.0 to 1.1.0

**No breaking changes!** This is a minor release adding a new endpoint.

**What's New:**
- `InternationalBroadcasterSchedule` endpoint for accessing international broadcast schedules
- Example program demonstrating broadcast schedule usage
- Comprehensive tests for the new endpoint

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.1.0`
2. No code changes required for existing code
3. Optionally use the new endpoint: `endpoints.GetInternationalBroadcasterSchedule()`

**New Features:**
- Track which international broadcasters are showing NBA games
- Filter by Season, LeagueID, RegionID, Date, and EST
- Access game schedules with detailed broadcaster information

### From 0.9.0 to 1.0.0

**No breaking changes!** This is a stability release with no API changes.

**What's New:**
- Comprehensive testing infrastructure (integration + contract tests)
- Maintenance runbook for operational procedures
- Maintainability assessment confirming production-readiness
- Stability guarantees and semver commitment

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.0.0`
2. No code changes required
3. Review `docs/MAINTENANCE.md` if you're maintaining a fork
4. Consider running contract tests to validate your integration

**Recommended Actions:**
- Set `INTEGRATION_TESTS=1` and run tests to validate your environment
- Review `tests/contract/README.md` for API drift detection strategies
- Check `docs/MAINTENANCE.md` for operational best practices

### From 0.1.0 to 0.9.0

No breaking changes! All endpoints maintain backward compatibility. New endpoints are purely additive.

**New Features Available:**
- 134 additional endpoints
- HTTP API server (opt-in)
- Container deployment option

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v0.9.0`
2. No code changes required
3. Optionally explore new endpoints in `pkg/stats/endpoints/`

---

## Versioning Policy

This project follows [Semantic Versioning](https://semver.org/):

- **Major version (1.x.x)**: Breaking API changes
- **Minor version (x.1.x)**: New features, backward compatible
- **Patch version (x.x.1)**: Bug fixes, backward compatible

**Post-1.0 Guarantees:** Starting with v1.0.0, strict semver guarantees apply. Breaking changes will only occur in major version updates.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to suggest changes or report issues.

[Unreleased]: https://github.com/n-ae/nba-api-go/compare/v1.3.0...HEAD
[1.3.0]: https://github.com/n-ae/nba-api-go/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/n-ae/nba-api-go/compare/v1.1.7...v1.2.0
[1.1.7]: https://github.com/n-ae/nba-api-go/compare/v1.1.6...v1.1.7
[1.1.6]: https://github.com/n-ae/nba-api-go/compare/v1.1.5...v1.1.6
[1.1.5]: https://github.com/n-ae/nba-api-go/compare/v1.1.4...v1.1.5
[1.1.4]: https://github.com/n-ae/nba-api-go/compare/v1.1.3...v1.1.4
[1.1.3]: https://github.com/n-ae/nba-api-go/compare/v1.1.1...v1.1.3
[1.1.1]: https://github.com/n-ae/nba-api-go/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/n-ae/nba-api-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/n-ae/nba-api-go/compare/v0.9.0...v1.0.0
[0.9.0]: https://github.com/n-ae/nba-api-go/compare/v0.3.0...v0.9.0
[0.3.0]: https://github.com/n-ae/nba-api-go/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/n-ae/nba-api-go/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/n-ae/nba-api-go/releases/tag/v0.1.0
