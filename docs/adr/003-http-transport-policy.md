# ADR 003: HTTP Transport Policy

## Status

Accepted (2026-07-19)

## Context

`pkg/client.NewClient` builds its own `*http.Transport` when the caller
doesn't supply an `HTTPClient`. Before this ADR, that transport was:

```go
transport := &http.Transport{
    DisableKeepAlives:     true,
    MaxIdleConns:          1,
    IdleConnTimeout:       30 * time.Second,
    TLSHandshakeTimeout:   30 * time.Second,
    ResponseHeaderTimeout: 60 * time.Second,
}
```

Two things about this were worth re-examining:

1. **It unconditionally disables keep-alives.** Every request pays for a
   fresh TCP connection and TLS handshake, even for back-to-back calls to
   the same host (season ingestion, a fantasy optimizer polling several
   endpoints, etc.). Neither the CHANGELOG, code comments, nor the ADRs
   record *why* - there's no documented NBA.com-specific reason for it.
   The actual documented NBA.com-compatibility measures are the
   browser-like `User-Agent`, `Referer: https://www.nba.com/`, and
   `Accept: application/json` headers applied by
   `internal/middleware` - none of which relate to connection reuse.
2. **It's built from a zero-valued `http.Transport{}` instead of cloning
   `http.DefaultTransport`.** That silently drops `Proxy:
   ProxyFromEnvironment` (breaking anyone behind a corporate/dev proxy
   relying on `HTTP_PROXY`/`HTTPS_PROXY`), `ForceAttemptHTTP2`, a
   `DialContext` with sane connect timeouts, and `MaxIdleConnsPerHost` -
   all generally-desirable stdlib defaults with no clear reason to
   discard them here.

### What this ADR could not do

The assessment that raised this item called for benchmarking keep-alives
against live NBA.com. That wasn't possible in the environment this
decision was made in: general internet access worked (confirmed via
`curl https://www.google.com`) and DNS resolved `stats.nba.com` correctly
to an Akamai edge node, but every attempted request to
`stats.nba.com/stats/...` - with the SDK's own browser-like headers -
timed out after the TLS handshake completed with zero bytes of HTTP
response, both via `curl` and through this repository's own integration
tests. This is consistent with NBA.com/Akamai's known bot-detection
silently dropping requests from some source IPs (cloud/datacenter ranges
in particular) rather than anything related to keep-alives specifically -
but it means this decision is reasoned from first principles and stdlib
defaults, not empirical A/B measurement against the real upstream.

## Decision

Clone `http.DefaultTransport` and apply narrow, deliberate overrides,
re-enabling keep-alives in the process:

```go
defaultTransport, ok := http.DefaultTransport.(*http.Transport)
if !ok {
    panic("client: http.DefaultTransport is not *http.Transport")
}
transport := defaultTransport.Clone()
transport.TLSHandshakeTimeout = 30 * time.Second   // stdlib default: 10s
transport.ResponseHeaderTimeout = 60 * time.Second // stdlib default: unset (no timeout)
```

Rationale for each override:

- **`TLSHandshakeTimeout: 30s`** (stdlib default 10s) - kept from the
  pre-existing transport. Worth keeping generous: this ADR's own
  connectivity check saw the TLS handshake to `stats.nba.com` complete
  but the subsequent HTTP response never arrive, suggesting Akamai-fronted
  endpoints can be slower to negotiate than typical sites.
- **`ResponseHeaderTimeout: 60s`** (stdlib default: unset) - also kept
  from the pre-existing transport. NBA.com's API is documented elsewhere
  in this repo as occasionally slow; without this, a hung response would
  rely entirely on `Config.Timeout` (which bounds the whole request, not
  just waiting for headers) to eventually fail.
- **Everything else comes from `http.DefaultTransport.Clone()`**:
  `Proxy: ProxyFromEnvironment`, `ForceAttemptHTTP2: true`,
  `MaxIdleConns: 100`, `IdleConnTimeout: 90s`, a `DialContext` with a 30s
  connect timeout and 30s TCP keep-alive, `ExpectContinueTimeout: 1s`.
  None of these have any documented NBA.com-specific reason to differ
  from the stdlib's own well-reasoned choices.
- **`DisableKeepAlives` is no longer set** (so it defaults to `false`,
  i.e. keep-alives enabled). `MaxIdleConns: 1` from the old transport is
  also gone - it was very likely a vestigial companion to
  `DisableKeepAlives: true` (with keep-alives off, connections aren't
  pooled regardless of this value, so it was effectively a no-op) rather
  than a deliberate cap; the stdlib default of 100 applies instead.

This does **not** change per-host request *rate* - `middleware.WithRetry`,
`WithPerHostRateLimit`, and `Config.Timeout`/`MaxResponseBytes` are
unaffected. Enabling keep-alives only lets consecutive requests to the
same host reuse a connection instead of renegotiating TLS every time; it
doesn't send requests any faster than the existing rate limiter allows.

The escape hatch for anyone who needs different transport behavior
(including reverting to `DisableKeepAlives: true`) is unchanged: pass a
custom `Config.HTTPClient`, which bypasses this default transport
entirely. `TestNewClient_CustomHTTPClientBypassesDefaultTransport`
(`pkg/client/transport_test.go`) guards that this still works.

## Consequences

### Positive

- Fewer TCP/TLS handshakes for any sequential or repeated usage against
  the same host - the common case for this SDK (season ingestion, a
  fantasy tool polling several endpoints, the HTTP server's own shared
  `stats.Client`).
- Proxy support (`HTTP_PROXY`/`HTTPS_PROXY`/`NO_PROXY`) and HTTP/2
  negotiation work again for anyone relying on the default transport.
- The transport's shape is now "stdlib defaults plus two named,
  justified overrides" instead of "a mostly zero-valued struct with no
  comments" - easier for a future maintainer to reason about or extend.

### Negative / Risk

- This decision is **not empirically verified against live NBA.com** (see
  "What this ADR could not do" above). If connection reuse turns out to
  interact badly with NBA.com/Akamai's bot detection in a way single-use
  connections didn't - a real but currently unconfirmed possibility - the
  symptom would likely be intermittent connection resets or degraded
  success rates that don't reproduce with `DisableKeepAlives: true`.

### Revisit trigger

If a maintainer or user reports a reproducible pattern of failures against
NBA.com that goes away when constructing the client with a custom
`HTTPClient` that sets `DisableKeepAlives: true`, treat that as evidence
this decision was wrong for this specific upstream and revert the default
(re-adding `DisableKeepAlives: true` to the cloned transport, with a
comment linking to the report). Absent such a report, there is no reason
to prefer the old behavior.

## References

- [`net/http.DefaultTransport`](https://pkg.go.dev/net/http#RoundTripper) source (`net/http/transport.go`)
- `docs/MAINTAINABILITY_ASSESSMENT_2026-07-19.md` §6.2 item 6 (the backlog item this ADR resolves)
- `pkg/client/transport_test.go` (regression tests for this decision)
