# Repository Review

**Reviewed:** 2026-07-19  
**Revision:** `657b4a5` (`chore: release v1.1.7 - update dependencies and Go toolchain`)  
**Scope:** source, tests, build tooling, container configuration, and active documentation. No production code was changed.

## Summary

The project has a strong foundation: a small dependency surface, a clear SDK/server split, generated endpoint types, context propagation through handlers, and a clean local build. The main maintenance risk is not the number of endpoints; it is keeping the public configuration, HTTP wrapper, tests, and documentation accurate as the API evolves.

**Overall assessment: B.** The codebase is usable and structurally understandable today, but the P1 items below should be addressed before relying on the HTTP server as a broadly exposed production proxy.

## What was verified

| Check | Result |
| --- | --- |
| `go test ./...` | Passes |
| `go test -race ./pkg/client ./internal/middleware ./cmd/nba-api-server ./tests/integration` | Passes; this does not run build-tagged live tests |
| `go vet ./...` | Passes |
| `make test-examples` | All 15 examples compile |
| Go toolchain | `go1.26.5 darwin/arm64` |
| Dependencies | Two direct `golang.org/x/*` dependencies |

Coverage is uneven. The core client reports 74.2%, static-data package 49.6%, server 7.1%, and generated stats endpoints 1.1%. The default suite is therefore a useful compile/regression check, but not evidence that most endpoint parsing and HTTP routes are exercised.

## Strengths

- `pkg/` cleanly separates the reusable client, stats SDK, live SDK, models, and static datasets from the optional server in `cmd/`.
- Requests through the HTTP wrapper consistently use `r.Context()`, so caller cancellation is propagated to upstream requests.
- The middleware chain and `HTTPClient` interface give the core transport a practical seam for tests and customization.
- Endpoint generation keeps the large SDK surface consistent. The generator metadata and templates are versioned with the source.
- The service uses standard-library HTTP, has bounded response-time sampling, and builds as a non-root container user.

## Findings

### P1 — Stats and live client configuration silently ignores documented fields

`stats.Config` declares `Headers` and `Timeout`, but `stats.NewClient` passes only `BaseURL` and middleware to `client.Config` ([`pkg/stats/stats.go`](../pkg/stats/stats.go)). The same applies to `live.Config` ([`pkg/live/live.go`](../pkg/live/live.go)). Consequently, callers cannot set custom headers or a timeout through either public SDK configuration type.

This is particularly problematic for an NBA client, where headers and timeouts are operationally significant. It is also easy to miss because defaults work and there are no tests covering those fields.

**Recommendation:** make the facade configurations match `client.Config`: use `http.Header` and `time.Duration` (or explicitly convert and document an alternate unit), then forward them. Add tests that observe a custom header and enforce a short timeout against `httptest.Server`.

### P1 — Server rate limiting keys on `RemoteAddr`, including the client port

The server stores limiters under `r.RemoteAddr` ([`cmd/nba-api-server/ratelimit.go`](../cmd/nba-api-server/ratelimit.go)). In normal HTTP traffic this value is `host:port`, so a client that opens new connections gets a new token bucket each time. The claimed per-IP limit is therefore bypassable and can grow the limiter map under connection churn.

**Recommendation:** extract the host with `net.SplitHostPort`. If deployed behind a proxy, make trusted-proxy handling an explicit, separately configured policy rather than accepting forwarded headers from arbitrary clients. Add a test using two source ports for the same IP.

### P1 — `/health` is an upstream probe but always returns `200 healthy`

Every health request makes a live NBA API request ([`cmd/nba-api-server/main.go`](../cmd/nba-api-server/main.go)). If it fails, the response still has HTTP 200 and `status: "healthy"`; only `nba_api_status` becomes `"degraded"`. The container health check only evaluates HTTP success, so it cannot detect that degradation. Frequent probes can also consume upstream capacity and contend with normal requests.

**Recommendation:** split liveness (local, no network) from readiness (cached/controlled upstream check), or return an appropriate readiness status when the dependency is unavailable. Do not perform a fresh upstream request for every load-balancer liveness probe.

### P2 — Default transport deliberately disables connection reuse

The default transport sets `DisableKeepAlives: true` ([`pkg/client/client.go`](../pkg/client/client.go)). That imposes a new TCP/TLS connection for every request, at odds with the project's performance claims and the server's shared long-lived `stats.Client`. It may have been a deliberate workaround for upstream behavior, but no rationale is recorded.

**Recommendation:** document the upstream constraint if this is intentional. Otherwise benchmark a transport with keep-alives enabled and reasonable idle-connection bounds; retain an opt-out configuration if NBA.com requires it.

### P2 — The test suite passes while important tests are skipped or absent

- `tests/integration` skips its only smoke test unless `INTEGRATION_TESTS=1` is set ([`tests/integration/simple_smoke_test.go`](../tests/integration/simple_smoke_test.go)).
- The richer endpoint integration tests are behind an `integration` build tag and their `TestMain` returns without calling `m.Run()` unless `INTEGRATION_TESTS=1` is set ([`pkg/stats/endpoints/endpoints_integration_test.go`](../pkg/stats/endpoints/endpoints_integration_test.go)).
- The server's default tests focus on plumbing; `TestHealthEndpoint` makes a real upstream call indirectly, while route coverage is limited.
- Most endpoint code and the server are lightly covered despite the 140-endpoint surface.

**Recommendation:** keep live tests opt-in, but add fixture-backed parsing tests for each generated endpoint family and route-contract tests that inject a fake `stats.Client`/transport. Make the CI test command report skipped live tests clearly rather than letting a green result imply broad endpoint verification.

### P2 — Active documentation contains stale paths and release facts

The docs index links to `tests/http-api/README.md`, which does not exist ([`docs/README.md`](README.md)). It also presents a historical endpoint count of 139, while the README/health response now state 140 and 149 respectively ([`docs/adr/002-api-server-architecture.md`](adr/002-api-server-architecture.md), [`README.md`](../README.md), [`cmd/nba-api-server/main.go`](../cmd/nba-api-server/main.go)). `CONTRIBUTING.md` asks for Go 1.21+, but the module requires Go 1.26.5.

**Recommendation:** choose one source of truth for supported Go version and endpoint inventory, correct dead links, and add a documentation-link check to CI.

### P3 — API wrapper duplication is the long-term change hotspot

The SDK has approximately 140 generated endpoint files while the HTTP server manually maps 142 routes to 142 handler functions. The current organization is readable, but adding or changing an endpoint requires synchronized SDK, handler, tests, docs, and inventory updates.

**Recommendation:** keep the current layout for now, but create a machine-checkable endpoint registry (or generate the route mapping) before the next large endpoint expansion. At minimum, a test should compare the SDK registry to the exposed HTTP route inventory.

## Suggested order of work

1. Forward and test `stats.Config`/`live.Config` headers and timeout values.
2. Correct the rate-limit key and redesign liveness/readiness behavior.
3. Add offline route and parsing coverage for high-use endpoints, then repair test commands and documentation.
4. Decide and document the HTTP keep-alive policy.
5. Introduce a shared endpoint inventory to prevent SDK/server/documentation drift.

## Notes on assessment boundaries

This review did not make live NBA.com requests, build the container image, or run `golangci-lint` (the binary's availability was not established). Passing local tests and compiled examples establish local correctness only; upstream API compatibility remains an opt-in verification concern.
