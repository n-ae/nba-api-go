# Repository Assessment

**Reviewed:** 2026-07-19  
**Revision:** `2363f46027713594e55c53074b8361f5c9307588`  
**Method:** Direct repository review; no delegated or specialized assessment agent was used.  
**Scope:** Root module, nested generator module, HTTP server, generated endpoints, tests, CI, container files, and active documentation. No production code was changed.

## Executive summary

The repository has a good small-project architecture: two direct dependencies, a standard-library HTTP server, a reusable core client, context propagation, bounded response reads, retry/rate-limit middleware, a non-root container, and a focused CI workflow. The operational issues identified in the earlier `v1.1.7` reviews—ignored facade configuration, an inaccessible middleware extension point, per-port rate limiting, per-request upstream health probes, and disabled keep-alives—are fixed at the current revision.

The main risk has moved to the generated SDK surface. The generator infers response types from field-name substrings, generated parsers consume result sets and columns by position, and conversion helpers silently replace incompatible values with zero or rounded strings. There are concrete examples in committed source: player display names are declared as `float64`, textual shot-range buckets are declared as `int`, and decimal statistical fields are declared as `string`. This is data corruption, not merely imperfect typing.

The intended safety net does not currently protect that path. No JSON contract fixture is tracked in Git, so a clean checkout skips all 18 contract tests while reporting a passing package. The local ignored fixtures are parsed SDK output rather than raw NBA responses, discard upstream result-set headers, and include obvious zeroed/shifted values that the assertions still accept.

**Overall assessment: C+.** The hand-written core is maintainable and substantially improved, but the project should not claim verified type-safe coverage until generated parsing and its offline verification are repaired.

## Architecture snapshot

Runtime dependencies flow in one direction:

```text
Go consumer ──> pkg/stats or pkg/live ──> pkg/client ──> NBA APIs
HTTP consumer ──> cmd/nba-api-server ──> pkg/stats ──> pkg/client

tools/generator (separate Go module) ──> pkg/stats/endpoints
tests/contract ──> typed endpoint responses (fixtures currently untracked)
```

The package boundaries are sensible. The risky boundary is the development-time generator-to-runtime endpoint edge: roughly 70% of the Go source is generated, but the generator is untested, excluded from root-module linting, and not exercised by a regeneration check.

## Verification results

| Check | Result |
| --- | --- |
| `go test ./...` | Pass |
| `go test -cover ./...` | Pass; endpoint coverage 1.1%, server coverage 8.7% |
| `go vet ./...` | Pass |
| `go test -race ./pkg/client ./internal/middleware ./pkg/stats ./pkg/live ./cmd/nba-api-server` | Pass |
| `make test-examples` | Pass; all 15 Go examples compile |
| `govulncheck ./...` | Pass; no known vulnerabilities found |
| `golangci-lint run ./...` with writable temporary caches | Pass; 0 root-module issues |
| `go test ./...` in `tools/generator` | Pass, but the module has no test files |
| `golangci-lint run ./...` in `tools/generator` | Fail; 5 issues |
| Documented generator dry run | Fail; template path cannot be resolved |
| Contract tests from a clean `git archive` | Package passes, but all 18 tests skip because fixtures are absent |

No live NBA request or container image build was performed. Local ignored fixtures were inspected as supporting evidence, but they are not part of revision `2363f46`.

## Strengths

- `pkg/client`, `pkg/stats`, `pkg/live`, `pkg/models`, and `cmd/nba-api-server` have clear responsibilities.
- The dependency surface is small: only `golang.org/x/text` and `golang.org/x/time` are direct runtime dependencies.
- Request contexts propagate through the server, SDK, middleware, and transport.
- The core client has a response-size limit, useful error types, retry cancellation handling, `Retry-After` support, default-transport cloning, and injectable middleware.
- Liveness and readiness are separated, and the background health probe reuses the shared client.
- The per-IP server limiter strips ephemeral ports and explicitly avoids trusting arbitrary proxy headers.
- CI runs root-module tests, vet, lint, example builds, vulnerability scanning, and a generator build.
- The endpoint inventory test prevents simple count drift between SDK files, HTTP routes, and `/health`.
- The container runs as a non-root user.

## Findings

### P1 — Generated response parsing silently corrupts data

The generator's `inferGoType` function uses substring matching for statistical abbreviations and broad name rules (`tools/generator/generator.go`). That produces deterministic false positives and false negatives:

- `DISPLAY_FIRST_LAST`, `DISPLAY_LAST_COMMA_FIRST`, `DISPLAY_FI_LAST`, `LAST_AFFILIATION`, and `PLAYER_NAME_LAST_FIRST` contain `ast`, so they are inferred as statistics and emitted as `float64`.
- `SHOT_CLOCK_RANGE`, `DRIBBLE_RANGE`, `CLOSE_DEF_DIST_RANGE`, `TOUCH_TIME_RANGE`, `SHOT_DIST_RANGE`, and `SHOT_ZONE_RANGE` are textual buckets but are emitted as `int`.
- At least 59 declarations with names such as ratings, frequencies, and non-terminal percentage fields are emitted as `string`, including fields where decimal precision is expected.

The conversion helpers make these mistakes silent (`pkg/stats/endpoints/types.go`):

- `toFloat("Stephen Curry")` returns `0`.
- `toInt("24+ Feet")` returns `0`.
- `toString(0.357)` uses `%.0f` and returns `"0"`.

Concrete committed examples include:

- `pkg/stats/endpoints/commonallplayers.go`
- `pkg/stats/endpoints/commonallplayersv2.go`
- `pkg/stats/endpoints/commonplayerinfov2.go`
- `pkg/stats/endpoints/playerdashptshots.go`
- `pkg/stats/endpoints/shotchartlineupdetail.go`
- `pkg/stats/endpoints/leaguedashplayershotlocations.go`

Generated parsing is also positional. About 135 endpoint files access `rawResp.ResultSets[n]`, and the template maps `row[n]` without checking `ResultSets[n].Name` or comparing the upstream `Headers` with metadata. A reordered result set or inserted column can therefore shift every later field without returning an error.

The ignored local fixtures expose the failure mode: `commonallplayers_2023-24.json` contains zero for every display name, while `commonplayerinfo_201939.json` contains a team ID under the team-name field and subsequent values shifted into the wrong fields. Those files are not proof of the current upstream schema, but they are proof that the parser can emit obviously invalid typed output without failing.

**Recommendation:** replace heuristic-only result typing with explicit field types in metadata, retaining inference only as a reviewed bootstrap. Generate parsing keyed by result-set name and validate/map columns using the upstream header list. Add table tests for every inference rule and raw-response parser tests for representative endpoint families. Because correcting exported field types is source-breaking, plan the public-type repair explicitly under the project's compatibility policy rather than regenerating blindly.

### P1 — Contract tests are a no-op in CI and do not validate the raw contract

`tests/contract/.gitignore` ignores every `fixtures/*.json`, and `git ls-files` confirms that zero JSON fixtures are tracked. `loadFixture` calls `t.Skipf` when a file is missing. In a clean archive of the reviewed revision, all 18 contract tests skipped and the package returned `PASS`.

Even when local fixtures exist, the recording flow calls a typed endpoint and marshals its already-parsed `models.Response`. It does not preserve the raw NBA `resultSets`, `headers`, or `rowSet`. This prevents the fixtures from detecting column additions, removals, or reorderings—the exact drift that positional parsing needs protection against.

Most replay assertions call `validateBasicSchema`, which only checks that `Data` is non-empty and a named key exists. Local fixtures with `null` or empty endpoint data pass that check; examples include league leaders, league standings, and team game log.

**Recommendation:** commit a compact set of raw upstream fixtures, make missing required fixtures fail in CI, and feed those raw responses through `httptest.Server` into the real endpoint functions. Assert exact representative values, types, result-set names, and columns. Keep live fixture recording opt-in, but separate it from offline replay so recording cannot normalize away the defect being tested.

### P1 — The generator cannot be run using its documented workflows

The generator is a nested module. Its README instructs maintainers to run:

```text
cd tools/generator
go run . -endpoint PlayerGameLog
```

That command fails because `loadTemplate` looks for `tools/generator/templates/endpoint.tmpl` relative to the already-changed working directory. The default output directory is likewise relative to the current directory, so it would target `tools/generator/pkg/stats/endpoints` rather than the root package.

Root-level alternatives also fail:

- `go run ./tools/generator ...` fails because the root module does not contain the nested module package.
- `go run tools/generator/main.go ...` compiles only `main.go` and cannot find `NewGenerator`.
- `tools/regenerate_remaining.sh` tries `go build ... ./tools/generator` from the root module, which fails in a clean checkout.

CI builds the nested module but does not execute a dry run, compare generated output, run generator tests, or lint it. The module has no tests, and its standalone lint currently reports five issues.

**Recommendation:** embed the template with `go:embed` or resolve it independently of the process working directory; require/derive an unambiguous output root; correct all maintenance commands and scripts; then add generator unit tests and a clean regeneration/diff check. CI should run `go test` and lint inside `tools/generator`, not only build it.

### P1 — The HTTP API defaults 101 handlers to the 2023-24 season

There are 101 `getQueryOrDefault(..., "Season", "2023-24")` calls across the five handler files. On the review date in 2026, omitting `Season` therefore selects historical data. The API usage guide documents the same stale default, and the health probe also uses 2023-24.

The repetition makes seasonal maintenance a five-file search-and-replace and increases the chance of inconsistent updates.

**Recommendation:** centralize season policy. Either require `Season` at the HTTP boundary or derive a configurable default through one tested function with an injected clock. Keep historical seasons explicit in examples and fixtures rather than using them as production defaults.

### P2 — The server adapter remains a lightly tested manual duplication surface

The HTTP server contains 141 handler methods and 142 route cases. Most handlers manually:

1. parse query parameters,
2. supply defaults,
3. build a generated request,
4. call one endpoint, and
5. wrap its response.

Server statement coverage is 8.7%. The inventory test verifies counts, but it does not prove that each route calls the corresponding SDK function, forwards the complete parameter set, or applies correct required/default rules. It can still pass if two routes point to the same endpoint or one handler omits a parameter.

**Recommendation:** introduce a machine-readable endpoint/route registry and generate or table-test route dispatch and parameter mapping. A fake upstream transport can exercise every route offline without 142 bespoke test servers.

### P2 — Metrics and operational configuration make claims the implementation does not fulfill

- `Metrics.requestsByPath` retains every raw path forever. Requests to arbitrary unsupported suffixes under `/api/v1/stats/` can grow the map without a bound.
- Response-time statistics stop collecting after the first 1,000 requests rather than using a rolling sample, so long-running metrics describe startup traffic.
- `/metrics` returns JSON, while `DEPLOYMENT.md` says Prometheus can scrape it directly. Prometheus expects an exposition format unless an adapter is configured.
- `LOG_LEVEL` is read and logged but does not filter or change logging.
- `docker-compose.yml` and ADR 002 set/document `NBA_API_TIMEOUT`, but the server never reads it.
- CORS is hard-coded to `*`, despite the README and ADR describing it as configurable.

**Recommendation:** bound or normalize path labels, use a rolling latency window or histogram, and either expose Prometheus text or document the JSON endpoint accurately. Remove nonfunctional configuration or implement and test it.

### P2 — Active documentation still contains executable inaccuracies

- The Dockerfile example in `DEPLOYMENT.md` uses Go 1.21, but `go.mod` requires Go 1.26.5.
- `docs/README.md` links to the deleted `tests/http-api/README.md`.
- `CLAUDE.md` still points to `cmd/generator`, `docs/DEPLOYMENT.md`, and `docs/PYTHON_MIGRATION.md`, none of which are current paths.
- `docs/MAINTAINABILITY_ASSESSMENT_2026-07-19.md` and `docs/REPOSITORY_REVIEW_2026-07-19.md` present resolved `v1.1.7` findings without a prominent superseded notice.
- The contract-test documentation says fixtures are version-controlled and CI provides offline replay, contrary to the tracked tree.

**Recommendation:** add a superseded banner to historical active assessments or archive them, run a local-link checker in CI, and test documentation commands from a clean checkout.

### P3 — Typed response metadata is placeholder data

`models.Response` exposes `StatusCode`, `URL`, and `Headers`, but 140 stats endpoint functions construct it with `(200, "", nil)`; the live scoreboard does the same. `GetJSON` discards the successful `RawResponse` metadata before endpoint code constructs the typed response.

**Recommendation:** decode from a returned `RawResponse` and carry its metadata into `models.Response`, or remove/document the placeholder fields if they are not intended to be meaningful.

## Recommended order of work

1. Commit raw fixtures for a small, high-value endpoint set and make missing fixtures fail in CI.
2. Add regression tests that reproduce the display-name, decimal-stat, textual-range, and column-shift failures.
3. Replace heuristic-only field typing with explicit metadata overrides and header-aware parsing; audit generated public fields before release.
4. Make the generator runnable from a clean checkout, test it in its own module, and add deterministic regeneration checks.
5. Centralize HTTP season defaults and add route/parameter contract tests.
6. Bound metrics cardinality and reconcile operational configuration/documentation.
7. Repair active links, commands, and superseded assessment status.

## Bottom line

The repository's core transport and server lifecycle are in better shape than the older reviews suggest. The architectural idea—small hand-written core plus generated endpoint breadth—is still the right one for a solo-maintained SDK. The missing control is trust at the generation boundary. Until raw fixtures, header-aware parsing, explicit field types, and a working generator pipeline are in place, a green build establishes compilation and core-client behavior, not correctness across 141 NBA endpoints.
