> **Superseded.** This assessed `v1.1.7` (`657b4a5`). The current assessment of record is
> [`docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md`](./MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md),
> covering revision `2363f46` and later. Retained here for history; several findings below were fixed
> in the interim (see that document's "Reconciling the two input reviews" section for what carried
> forward vs. what changed).

# Maintainability Assessment: nba-api-go

**Date**: 2026-07-19
**Assessor**: maintainable-architect-v4
**Perspective**: Solo Engineer Long-Term Viability
**Revision assessed**: `657b4a5` (= tag `v1.1.7`, annotated), clean tree, go1.26.5 darwin/arm64
**Inputs verified**: `docs/REPOSITORY_REVIEW_2026-07-19.md` (internal, grade B) and an external "Senior Software Engineering Review" of the v1.1.7 tag (6.7/10). Every load-bearing claim from both was checked against source before being incorporated. Prior assessment (`docs/MAINTAINABILITY_ASSESSMENT.md`, 2025-11-05, maintainable-architect-v2) retained as history — this document supersedes it.

---

## Executive Summary

**Overall Grade: B- (72/100) — sound skeleton, untrustworthy edges.**

The bones remain excellent: two dependencies, stdlib-only server, single binary, generated endpoint consistency, green test suite, clean SDK/server split. What has eroded is *truthfulness* — the public config silently ignores documented fields, the documented middleware extension point cannot compile outside the module, `/health` reports "healthy" unconditionally, the per-IP rate limit is keyed on `host:port` so it isn't per-IP, a patch release raised the mandatory Go floor against the project's own semver promise, and the repository states five different endpoint counts in five places. None of these are deep defects — the full fix list is roughly 20–25 hours — but together they mean the project's claims can no longer be taken at face value, and for a solo maintainer, docs and claims you can't trust are a direct maintenance cost.

The previous grade of A (93/100) quoted in CLAUDE.md was a same-day self-upgrade in the v2 assessment's addendum and was never a defensible number: it scored Testing A- (92/100) while generated-endpoint coverage sat at 1.1% and the middleware package had (and still has) zero tests.

---

## 1. Verification Ledger

Both prior reviews were checked claim-by-claim. Verdicts: **Confirmed**, **Confirmed + understated** (reality is worse), **Overstated** (technically present, materially milder), or **Wrong**.

### 1.1 Internal repository review (2026-07-19, grade B)

| # | Claim | Verdict | Evidence |
|---|---|---|---|
| 1 | `stats.Config`/`live.Config` silently ignore `Headers`/`Timeout` | **Confirmed** | `pkg/stats/stats.go:29-31` passes only `BaseURL` + `Middlewares` to `client.Config`; `pkg/live/live.go:23-25` identical. `Timeout` is `int`, not `time.Duration`. |
| 2 | Server rate limiting keyed on full `RemoteAddr` (host:port) | **Confirmed** | `cmd/nba-api-server/ratelimit.go:41` — `ip := r.RemoteAddr`. New connection = new token bucket. Map growth is bounded by `CleanupOldLimiters` (5-min sweep), so the memory half of the concern is mitigated; the bypass is not. |
| 3 | `/health` makes a live upstream call yet always returns 200 "healthy" | **Confirmed + understated** | `cmd/nba-api-server/main.go:158-182` — `Status: "healthy"` hardcoded, `WriteHeader(200)` unconditional. Worse than reported: `checkNBAAPI` (`main.go:204`) constructs a **new `stats.NewDefaultClient()` per probe** — fresh transport, fresh rate limiter — so probe traffic bypasses the shared client's rate limiting and pays full TLS setup every time. And `TestHealthEndpoint` (`handlers_test.go:12`) exercises this path, so the "unit" suite makes a live NBA.com attempt on every run. |
| 4 | `DisableKeepAlives: true` with no recorded rationale | **Confirmed** | `pkg/client/client.go:49-55`. No comment, no ADR. Additional consequence (see §2.4): the hand-rolled transport also has no `Proxy` field, so `HTTP_PROXY`/`HTTPS_PROXY` environments are silently unsupported. |
| 5 | Coverage: client 74.2%, static 49.6%, server 7.1%, endpoints 1.1% | **Confirmed exactly** (re-measured) | Plus numbers the review omitted: `pkg/stats/parameters` 22.7%, **`internal/middleware` 0.0%**, **`pkg/live` 0.0%**, **`pkg/live/endpoints` 0.0%**. The review's "`go test -race ./internal/middleware` passes" is vacuously true — there are no test files. |
| 6 | Stale docs/links/version facts | **Confirmed + understated** | Dead link `docs/README.md:35` → `tests/http-api/README.md` (directory removed 2025-11-05). `CONTRIBUTING.md:31` says Go 1.21+ vs `go.mod` 1.26.5. Additionally (new findings, §2.6): CLAUDE.md references `cmd/generator` (actual: `tools/generator`), `docs/DEPLOYMENT.md` (actual: root `DEPLOYMENT.md`), and `docs/PYTHON_MIGRATION.md` (does not exist); CHANGELOG has no entries for tagged releases v1.1.1, v1.1.4, v1.1.5; CLAUDE.md says 14 examples (actual: 15 Go examples compile). |
| 7 | SDK/server duplication is the long-term change hotspot | **Confirmed** | 143 generated endpoint files; `cmd/nba-api-server/handlers.go:32` — one 142-case `switch` dispatching to 142 hand-written ~20-line handlers. |
| 8 | All 15 examples compile | **Confirmed** | 16 example dirs; `examples/http-api-client/` has no `main.go` and is skipped by the Makefile loop. |

### 1.2 External review (v1.1.7 tag, 6.7/10)

| # | Claim | Verdict | Evidence |
|---|---|---|---|
| 1 | **P0**: Public `Config` structs expose `[]middleware.Middleware` from `internal/middleware`; documented extensibility uncompilable downstream | **Confirmed** | All three configs (`pkg/client/client.go:40`, `pkg/stats/stats.go:20`, `pkg/live/live.go:19`). `README.md:276` demonstrates `import "github.com/n-ae/nba-api-go/internal/middleware"` — forbidden by Go's internal rule for any external module. Nuance: the SDK itself remains fully usable (leave `Middlewares` zero-valued); it is specifically the advertised extension point that is dead on arrival outside this repo. |
| 2 | **P0**: Ignored `Headers`/`Timeout`, `Timeout` typed `int` | **Confirmed** | Same evidence as internal #1. |
| 3 | **P1**: Patch v1.1.7 raised mandatory `go` directive 1.25.3 → 1.26.5, violating the semver promise | **Confirmed, with nuance** | `git show v1.1.6:go.mod` = `go 1.25.3`; v1.1.7 = `go 1.26.5`. `CHANGELOG.md:95`: "Minor and patch versions guarantee backward compatibility." Nuance: with the default `GOTOOLCHAIN=auto`, most consumers auto-upgrade transparently; the breakage is real but confined to pinned/hermetic/`GOTOOLCHAIN=local` environments. Still a genuine process breach — the changelog entry doesn't even flag it as compatibility-affecting. |
| 4 | **P1**: `DisableKeepAlives` + transport not cloned from `http.DefaultTransport` | **Confirmed** | `client.go:49-55`. Consequences: no proxy support, no HTTP/2 negotiation attempt, no dialer timeouts. |
| 5 | **P1**: Unbounded `io.ReadAll` of response bodies | **Confirmed** | `client.go:121`. |
| 6 | **P1**: `APIError` discards body/headers/`Retry-After` | **Confirmed** | `client.go:126-130` calls `models.HTTPStatusToError(resp.StatusCode, reqURL)` — body was read at line 121 and is thrown away on the error path. `pkg/models/errors.go:45` accepts only `(statusCode, url)`. |
| 7 | **P1**: No visible CI | **Confirmed** | No `.github/` directory at all; no CI config of any provider. Only `.golangci.yml` + Makefile. Sharpened by history: the changelog records that the lint config silently stopped loading under golangci-lint v2 — exactly the failure CI exists to catch. |
| 8 | P2: Header map not concurrency-safe; `SetHeaders` aliases the caller's map | **Confirmed** | `client.go:179-189`; `Get` iterates the same map at `client.go:108-112`. |
| 9 | P2: `Retry-After` ignored by retry middleware | **Confirmed** | `internal/middleware/retry.go` — no reference to the header. |
| 10 | P2: All transport errors retried, including context cancellation | **Overstated** | `retry.go:53-57` does `continue` on any error, **but** the next iteration's backoff `select` (`retry.go:44-48`) sees the already-cancelled `ctx.Done()` immediately and returns `ctx.Err()` without sleeping. Cancelled contexts cost one loop iteration, not a retry cycle. The valid remainder: permanent errors (bad TLS cert, NXDOMAIN, malformed scheme) do get the full 3-retry backoff, wasting up to ~7s per doomed call. |
| 11 | P2: Custom middleware silently replaces default headers/retry/rate-limit | **Confirmed** | `stats.go:33-43`, `live.go:27-34` — `if len(config.Middlewares) > 0` replaces the entire default chain. Currently mostly theoretical *because of finding #1*: external users can't construct middleware anyway. |
| 12 | P2: Header duplication across retries via `Header.Add` | **Overstated** | The hazard exists in `WithHeaders` (`internal/middleware/headers.go:11-14`, uses `Add`; retry wraps it and reuses the same `*http.Request`). But `WithHeaders` is used by **no default chain** and, being internal, is unreachable by external users. The default `WithUserAgent`/`WithReferer`/`WithAccept` are set-if-empty and idempotent. Latent one-line fix, zero live impact today. |
| 13 | P2: Mutex serializing the rate limiter's `Wait` | **Confirmed** | `internal/middleware/ratelimit.go:35-39` — `rate.Limiter` is already concurrency-safe; the wrapping mutex adds head-of-line blocking, and waiters blocked on `mu.Lock()` are not context-aware. The wrapper is pure deletable complexity. |
| 14 | P2: Constructors defer invalid-base-URL errors | **Confirmed** | `client.go:43` returns `*Client` only; `buildURL` parses per request (`client.go:149`). Practical impact is low — base URLs are package constants. |
| 15 | P2: README claims exceed evidence ("zero bugs", "10,000+ req/min") | **Confirmed verbatim** | `README.md:12`: "Production-ready: Zero bugs, type-safe, fully tested". `README.md:37`: "Handles 10,000+ req/min on 1 vCPU". The changelog itself documents parser bugs fixed in 1.1.3 and 1.1.6, and "fully tested" coexists with 1.1% endpoint coverage. |

**Bottom line on the reviews**: both are substantially accurate. The external review overstated two P2s (#10, #12) and, being consumer-oriented, over-prescribes API ceremony (functional options, constructor errors, module splits) that a solo maintainer should decline — see §6.3. Neither review caught the per-probe client construction in `/health`, the 0% middleware coverage, the missing proxy support, the phantom paths in CLAUDE.md, or the changelog gaps.

---

## 2. Findings (maintainer-priority order)

Effort: **S** < 1h, **M** 1–4h, **L** > 4h.

### Critical

**2.1 The public configuration lies** (effort M).
`stats.Config`/`live.Config` advertise `Headers` and `Timeout` and discard both. A user who sets `Timeout: 10` gets the hidden 30s default; a user adding an auth or workaround header sends nothing. Silently ignored configuration is the worst kind of API bug because every support conversation starts from false premises. This is the single highest-leverage fix in the repository. (Wiring note: at the facade the field is `int` — forward it as `time.Duration(cfg.Timeout) * time.Second`, *never* a bare cast, or `Timeout: 5` becomes 5 nanoseconds. Retype to `time.Duration` in v1.2.0.)

**2.2 The documented extension point cannot exist outside this repo** (effort M–L).
Three public structs name `middleware.Middleware` from `internal/`. The README's middleware example (README.md:274-292) does not compile for any downstream module. Fix is mechanical, not architectural: the whole abstraction is 27 lines (`internal/middleware/middleware.go`) — move `Middleware`, `RoundTripper`, `RoundTripperFunc`, `Chain` into `pkg/client` (or a tiny `pkg/transport`), keep implementations internal, and export `stats.DefaultMiddlewares()` so callers can *extend* rather than silently replace the default chain (fixes external P2 #11 with one function).

**2.3 The server's two operational promises are both false** (effort M).
Per-IP rate limiting isn't per-IP (`RemoteAddr` includes the port), and `/health` returns 200 "healthy" regardless of upstream state while burning a fresh client + TLS handshake per probe. A load balancer pointed at `/health` today (a) can never see degradation and (b) turns its probe interval into unthrottled NBA.com traffic. `net.SplitHostPort` fixes the first; a background-cached upstream status (refresh ≤ 1/min, `/health` local-only) fixes the second — and stops the unit suite touching the network.

### Major

**2.4 The transport is undocumented folklore** (effort M).
`DisableKeepAlives: true`, `MaxIdleConns: 1`, built from a zero `http.Transport` — no proxy support, no HTTP/2, a new TCP+TLS connection per request. Maybe NBA.com's CDN genuinely behaves better this way; nobody wrote it down, so nobody can safely change it. This directly contradicts the "10,000+ req/min" claim. Decide once, benchmark once, record it as ADR 003, and either clone `http.DefaultTransport` with deliberate overrides or keep the workaround with a dated comment.

**2.5 The safety net has a hole exactly where the cleverness lives** (effort M).
`internal/middleware` — retry loops, backoff math, per-host limiter double-check locking — is the most behaviorally subtle code in the module and has **zero tests**. `pkg/live` also 0%. Meanwhile the suite is green, which is how 0% hides. Endpoint coverage at 1.1% is less alarming than it sounds (generated, homogeneous code — a handful of fixture-backed families is enough), but middleware at 0% is where a 2am bug actually lives.

**2.6 Release and documentation discipline slipped** (effort S–M).
- Patch release raised the mandatory Go floor (semver breach; unflagged in the changelog).
- No CI of any kind — after a documented incident where lint silently stopped running.
- Tagged releases v1.1.1, v1.1.4, v1.1.5 have no CHANGELOG entries.
- Endpoint count is stated as **139** (ADR-002), **140** (README, CLAUDE.md), **142** (actual routes in `handlers.go`), **143** (endpoint files), **149** (`/health` `http_exposed`, `main.go:171`). Five numbers, at most one of them right.
- CLAUDE.md — the primary maintainer-facing doc — references three paths that don't exist (`cmd/generator`, `docs/DEPLOYMENT.md`, `docs/PYTHON_MIGRATION.md`), a stale example count, "Current Version: 1.0.0" in one place and 1.1.0 in another, and the inflated "Grade: A (93/100)".
- `README.md:12` "Zero bugs" is falsified by the project's own changelog.

### Minor

**2.7 Handlers erase upstream error semantics** (effort S). Every SDK error becomes `500 api_error` (`handlers_player.go:52-54` and its 141 siblings via shared pattern) even though `models.APIError.StatusCode` carries the upstream 400/404/429. One shared helper fix in `writeError` usage, not 142 edits — route it through an `errors.As` check in one place.

**2.8 Small sharp edges** (effort S each): unbounded `io.ReadAll`; `Retry-After` ignored; permanent transport errors retried; `SetHeaders` aliases the caller's map; the deletable mutex wrapper around `rate.Limiter`; `WithHeaders`'s latent `Add`-on-retry duplication.

---

## 3. Grade Reconciliation

| Verdict | Source | What it actually measured |
|---|---|---|
| A (93/100) | v2 assessment addendum, 2025-11-05 | Same-day self-grade after implementing its own recommendations. Scored Testing A- (92) on the *existence* of test frameworks while endpoint coverage was ~1% and middleware 0%. The body of the same document says B+ (85). The 93 was never load-bearing. |
| B | Internal review, 2026-07-19 | Code health, honestly measured. I confirm essentially all of it. |
| 6.7/10 | External review, v1.1.7 tag | Adopter risk. Roughly the same territory as B viewed from outside, weighted down by ecosystem maturity — a dimension the maintainer can't buy and shouldn't chase. |
| **B- (72/100)** | **This assessment** | Code health **plus** process discipline and truth-in-documentation, which I weight because for a solo maintainer, docs are architecture. |

Component view (judgment, no weighting formula):

| Dimension | Grade | One-line justification |
|---|---|---|
| Code structure & dependencies | A- | 2 deps, stdlib server, clean layout, generated consistency — unchanged and still best-in-class. |
| Public API contract | D+ | Ignored fields, internal type in public structs, README example that cannot compile downstream. |
| Testing | C- | Green suite ≠ safety net: middleware 0%, server 7.1%, endpoints 1.1%; contract/integration scaffolding is genuinely good. |
| Server operational honesty | C | Health check lies, rate limit bypassable, perf claim unevidenced. |
| Release/process discipline | C- | Semver breach in a patch, zero CI, three tagged releases missing from CHANGELOG. |
| Documentation accuracy | C | Good structure, five conflicting endpoint counts, phantom paths in CLAUDE.md. |
| Solo viability & ops cost | B+ | Single binary, ~zero infra, and every defect above is shallow — total remediation ≈ 20–25h. |

**Why the grade moved**: the code did not rot — most defects existed at v1.0.0 and were simply never caught, because the tests that would catch them (middleware, config wiring, external-consumer compilation) were never written, and there is no CI to insist. What changed since November is that the *claims* kept inflating (A-grade, "zero bugs", semver guarantees) while a patch release quietly broke one of them. B- is not a crisis grade; it is "stop advertising, start reconciling."

---

## 4. Architecture (C4)

### 4.1 Container view

```d2
# Diagram: c4-container-2026-07-19 | Type: C4 Container
# Date: 2026-07-19 | Related: MAINTAINABILITY_ASSESSMENT_2026-07-19.md
direction: down

sdk_user: "SDK Consumer\n[Go application]" {
  shape: person
  style.fill: "#08427b"
  style.font-color: white
}

http_user: "HTTP Consumer\n[any language]" {
  shape: person
  style.fill: "#08427b"
  style.font-color: white
}

nba_api_go: "nba-api-go\n[single Go module, one binary]" {
  style.stroke-dash: 3

  server: "HTTP API Server\n[cmd/nba-api-server, stdlib net/http]" {
    style.fill: "#438dd5"
    style.font-color: white
  }
  sdk: "Stats + Live SDK\n[pkg/stats, pkg/live — 143 generated endpoint files]" {
    style.fill: "#438dd5"
    style.font-color: white
  }
  core: "Core HTTP Client\n[pkg/client + internal/middleware:\nretry, rate limit, headers]" {
    style.fill: "#438dd5"
    style.font-color: white
  }
  static: "Static Data\n[pkg/stats/static — 5,135 players,\n30 teams, embedded]" {
    shape: cylinder
    style.fill: "#438dd5"
    style.font-color: white
  }
}

generator: "Code Generator\n[tools/generator — separate module,\ndev-time only]" {
  style.fill: "#438dd5"
  style.font-color: white
}

nba_stats: "NBA Stats API\n[stats.nba.com — unofficial,\nchanges without notice]" {
  style.fill: "#999999"
  style.font-color: white
}

nba_cdn: "NBA Live CDN\n[cdn.nba.com]" {
  style.fill: "#999999"
  style.font-color: white
}

sdk_user -> nba_api_go.sdk: "Typed endpoint calls\n[Go, in-process]"
sdk_user -> nba_api_go.static: "Player/team search\n[embedded, no network]"
http_user -> nba_api_go.server: "GET /api/v1/stats/*\n[HTTPS/JSON]"
nba_api_go.server -> nba_api_go.sdk: "142 hand-written handlers\n[duplication hotspot]"
nba_api_go.sdk -> nba_api_go.core: "Get / GetJSON\n[in-process]"
nba_api_go.core -> nba_stats: "Stats requests\n[HTTPS, retry + per-host rate limit,\nkeep-alives disabled]"
nba_api_go.core -> nba_cdn: "Live requests\n[HTTPS]"
generator -> nba_stats: "analyze endpoint (dev-time)\n[HTTPS]"
generator -> nba_api_go.sdk: "emits endpoint files\n[templates + metadata]"
```

Two structural observations the diagram makes obvious: every runtime path funnels through one small core client (so §2.1/§2.2/§2.4 fixes pay off everywhere at once), and the server→SDK edge is the only hand-maintained N-wide surface (142 handlers) — the correct long-term watch point, and still not worth generating until it actually causes a sync bug.

### 4.2 The middleware/config seam (the §2.1–2.2 problem)

```d2
# Diagram: c4-middleware-seam-2026-07-19 | Type: C4 Component
# Date: 2026-07-19 | Related: findings 2.1, 2.2
direction: right

app: "Downstream Module\n[github.com/you/app]" {
  shape: person
  style.fill: "#08427b"
  style.font-color: white
}

mod: "nba-api-go module" {
  style.stroke-dash: 3

  stats_cfg: "stats.Config\n[BaseURL ok |\nHeaders IGNORED | Timeout int IGNORED |\nMiddlewares []middleware.Middleware]" {
    style.fill: "#85bbf0"
  }
  client_cfg: "client.Config\n[Headers + Timeout honored |\nMiddlewares []middleware.Middleware]" {
    style.fill: "#85bbf0"
  }
  internal_mw: "internal/middleware\n[Middleware, RoundTripper, Chain,\nWithRetry, WithPerHostRateLimit]" {
    style.fill: "#85bbf0"
    style.stroke: "#c0392b"
  }
}

app -> mod.stats_cfg: "set BaseURL: compiles.\nset Headers/Timeout: silently dropped" {
  style.stroke: "#e67e22"
}
app -> mod.internal_mw: "import BLOCKED by Go internal rule\n(README example does not compile downstream)" {
  style.stroke: "#c0392b"
  style.stroke-dash: 4
}
mod.stats_cfg -> mod.client_cfg: "forwards BaseURL + Middlewares ONLY\n(stats.go:29-43)"
mod.stats_cfg -> mod.internal_mw: "public field names internal type"
mod.client_cfg -> mod.internal_mw: "public field names internal type"
```

The fix direction: move the 27-line abstraction (`Middleware`, `RoundTripper`, `RoundTripperFunc`, `Chain`) into `pkg/client`, alias it from `internal/middleware` so implementations stay put, forward all facade config fields, and export the default chain. No new packages of consequence, no framework, ~4h.

> If the repo adopts a diagrams convention later, extract these blocks to `docs/diagrams/c4/` as `.d2` sources; for now the assessment is their single home.

---

## 5. Solo-Maintainer Lens

**Burden trajectory**: the claimed ~1.6 h/week steady state is still credible — nothing here adds infrastructure, services, or dependencies. But the current trajectory quietly *increases* future burden three ways: every doc/claim drift compounds (five endpoint counts means every future edit guesses), the absent CI means regressions are discovered by users, and the untested middleware means the next NBA.com behavioral change gets debugged without a harness. The plan below spends ~9h once to bend that curve back down.

**Complexity budget**: healthy. 71% of the Go LOC (23,349 of 33,009) is generated and homogeneous — that's domain size, not complexity. The hand-written core (client + middleware + server plumbing) is small and boring. The correct posture is to *defend* this: every external-review suggestion that adds API surface (options DSLs, error taxonomies, module splits) spends budget this project doesn't need to spend.

**Boring-tech alignment**: exemplary where it counts (stdlib, 2 deps, single binary). The one place the project is insufficiently boring is *process*: no CI is not minimalism, it's a missing seatbelt. One small workflow file is the most boring technology on this entire list.

**Debuggable at 2am?** Mostly yes — single binary, structured-enough logs, in-process everything. The exceptions are exactly §2.3 (a health endpoint that can't tell you anything) and §2.7 (every upstream failure logged as a generic 500).

---

## 6. Action Plan

Budget: ~1.6 h/week ≈ 6.5 h/month. Phase 1 fits in about five weeks.

### 6.1 Before v1.1.8 (~9h total) — restore truth

| # | Action | Effort | Notes |
|---|---|---|---|
| 1 | Wire `stats.Config`/`live.Config` `Headers` + `Timeout` through to `client.Config`; add `httptest`-backed tests observing a custom header and a short timeout | **M** (2h) | Convert `time.Duration(cfg.Timeout) * time.Second` — document seconds; retype in v1.2.0. Changelog: "previously documented-but-ignored fields now honored." |
| 2 | Rate-limit key: `net.SplitHostPort(r.RemoteAddr)`; test with two source ports on one IP | **S** (45m) | Document that proxy-header trust is deliberately *not* implemented. |
| 3 | `/health`: stop per-probe client construction; reuse the server's client; cache upstream status in a background ticker (≤1/min); `/health` answers locally; expose degraded state via status field *and* a `503` on a separate `/readyz` | **M** (2h) | Also removes the live network call from the unit suite. |
| 4 | README truth pass: delete "Zero bugs"/"fully tested", qualify or remove "10,000+ req/min" (or link BENCHMARKS.md methodology), state the Go 1.26.5 floor, mark the middleware example repo-internal until v1.2.0; fix `CONTRIBUTING.md` Go version | **M** (1.5h) | Costs nothing, recovers credibility. |
| 5 | Add minimal CI: one GitHub Actions workflow — `go test ./...`, `go vet`, `golangci-lint`, `make test-examples`, `govulncheck` | **M** (1.5h) | One file. Explicitly *not* a matrix, not scheduled live tests — keep it boring. |
| 6 | CHANGELOG: backfill v1.1.1/v1.1.4/v1.1.5 stubs from git log; add retroactive note to 1.1.7 flagging the Go-floor change as compatibility-affecting | **S** (45m) | Semver promise repair starts with admitting the breach. |
| 7 | Pick the one true endpoint count (routes = 142 today), fix `main.go:169-172`, README, CLAUDE.md, ADR-002 footnote | **S** (30m) | Full drift-proofing lands in v1.2.0 (#8 below). |

### 6.2 v1.2.0 (~14h across the quarter) — make the documented API real

| # | Action | Effort |
|---|---|---|
| 1 | Export the middleware seam into `pkg/client` (types + `Chain`), alias internally; export `stats.DefaultMiddlewares()` / `live.DefaultMiddlewares()` so custom chains extend instead of replace; fix the README example to a compilable one | **L** (4h) |
| 2 | Retype facade `Timeout` to `time.Duration`, `Headers` to `http.Header` (they never worked, so no working behavior changes — say exactly that in the changelog) | **S** (1h) |
| 3 | Middleware tests: retry (retryable statuses, backoff cap, ctx cancel, permanent-error exit), per-host rate limit, header idempotency | **M** (3h) |
| 4 | Retry hygiene: don't retry `context.Canceled`/`DeadlineExceeded` explicitly; honor `Retry-After` (seconds + HTTP-date, capped at `MaxBackoff`) | **M** (1.5h) |
| 5 | Bound response reads: `io.LimitReader` (default ~50 MB, configurable) with explicit error | **S** (1h) |
| 6 | Transport decision: benchmark keep-alives against NBA.com; clone `http.DefaultTransport` + deliberate overrides (restores proxy + HTTP/2) or keep the workaround — either way, write **ADR 003: HTTP transport policy** | **M** (2h) |
| 7 | `APIError`: add truncated `Body` (~2 KB); server maps `APIError.StatusCode` through one shared helper instead of blanket 500 | **M** (1.5h) |
| 8 | Endpoint-inventory test: assert SDK endpoint registry == server route count == documented number; kills the five-counts problem permanently | **M** (2h) |
| 9 | Deletions: remove the mutex wrapper around `rate.Limiter`; make `SetHeaders` clone; fix `WithHeaders` to `Set` | **S** (1h) |

### 6.3 Explicitly declined (over-engineering for this project)

| Suggestion (external review) | Why declined |
|---|---|
| Functional-options API (`stats.WithTimeout(...)` etc.) | A config struct whose fields *work* is boring and sufficient; an options DSL is ceremony plus a second API to document and test. |
| Constructors returning errors | Breaking signature change guarding a failure mode (malformed constant base URL) that tests catch at compile-adjacent time. Revisit only if v2.0.0 ever happens. |
| Splitting SDK/server/generator into separate release modules | Multi-module version skew is a tax a solo engineer pays forever; `tools/generator` is already separate, which is enough. |
| v2.0.0 field-name normalization across generated types | Churns 143 files and breaks every consumer for zero maintenance payback. |
| Fully immutable, concurrency-safe client rewrite | Clone-on-`SetHeaders` + a documented "configure before sharing" contract buys 95% of the value for 2% of the work. |
| External-module compile harness for every README example | Fix the one uncompilable example; a standing harness is CI surface without proportional payback. |
| Expanded `APIError` (Method, RequestID, RetryAfter fields) | Truncated body + existing StatusCode/URL covers real 2am debugging; the rest is speculative taxonomy. |
| OpenAPI spec | Still zero user demand. Same verdict as 2025-11-05. |
| Handler generation for the 142 routes | Still deferred: no sync bug has ever occurred; the v1.2.0 inventory test detects drift for 2h instead of 16–24h. Trigger to revisit: the *first* actual SDK/route sync bug, or the next bulk endpoint expansion. |
| Circuit breakers, per-endpoint metrics, schema-drift alerting inside the SDK | Correct advice *for the consumer's application layer* (the review even says so); building it into the SDK bloats every user's dependency. |

### 6.4 Documentation consolidation

| File | Action | Reason |
|---|---|---|
| `docs/IMPLEMENTATION_SUMMARY.md`, `docs/IMPROVEMENTS_COMPLETED.md`, `docs/LINT_CLEANUP_PLAN.md` | Move to `docs/archive/` | Completed historical status reports; no forward value in the active tree. |
| `docs/RELEASE_NOTES_v1.0.0.md`, `docs/V1.0.0_RELEASE_SUMMARY.md` | Merge into CHANGELOG reference, archive | Duplicate release records; CHANGELOG is the single source. |
| `docs/README.md` | Fix dead `tests/http-api` link; re-index | Broken since 2025-11-05. |
| `CLAUDE.md` | Fix `cmd/generator`→`tools/generator`, `docs/DEPLOYMENT.md`→`DEPLOYMENT.md`, remove `docs/PYTHON_MIGRATION.md`, examples 14→15, version facts, replace "Grade: A (93/100)" with a pointer to this assessment | The maintainer-facing doc must not contain phantom paths. |
| `docs/MAINTAINABILITY_ASSESSMENT.md` | **Keep untouched** as history | Superseded by this document; its addendum's self-grading is part of the record. |
| `docs/adr/` | Add ADR 003 (transport policy) in v1.2.0; footnote ADR-002's endpoint count | Extract the keep-alive decision from folklore into an explicit record. |

---

## 7. Closing Verdict

Still a model solo project in its bones — and a cautionary tale in its claims. The gap between "what the code does" and "what the project says it does" is now the main maintenance liability, and it closes for about nine hours of work. Do §6.1 before shipping anything else; do not let v1.1.8 go out with a health check that lies and a config that ignores its own fields.

**Grade: B- (72/100). Trajectory after Phase 1: B+. Ceiling with v1.2.0 as scoped: A-, honestly earned this time.**

**Next review**: 2026-10-19 (quarterly), or immediately after v1.2.0 ships — whichever comes first.
