# Maintainable-Architect-v4 Assessment: nba-api-go

**Date:** 2026-07-22
**Revision assessed:** `0e400d1` (`main`, tag `v3.1.4`), go1.26.5 darwin/arm64
**Assessor:** maintainable-architect-v4
**Method:** Direct verification against source at HEAD, not against `CHANGELOG.md`'s prose or an unsolicited external review's prose - file reads of `pkg/client/client.go`, `pkg/client/client_test.go`; two throwaway Go programs reproducing `net/url.Parse`'s exact behavior on adversarial inputs, one calling `url.Parse` directly and one calling the real `client.NewClient` from a scratch scope to confirm the leak reaches the actual returned error, not just the standard library's; `git rev-parse`/`git log`; `go build ./...`, `go vet ./...`, `go test ./...`, `golangci-lint run ./...` (root and `tools/generator` modules, run separately); and `gh pr list`/`gh api repos/.../commits/0e400d1/check-runs`/`gh api repos/.../actions/runs/<id>` against the real `n-ae/nba-api-go` GitHub repository to independently check every checkable citation in an external review supplied for this cycle (see §0). All green except the two findings below. No production code was modified while writing this file.

**Why now:** the prior assessment of record (this same file, then covering revision `b3c605d`/tag `v3.1.3`, grade A-) closed with one open finding: `NewClient`'s five explicit `BaseURL` validation checks echoed the complete, unredacted input in their errors. `v3.1.4` (tag `0e400d1`) closed it. This cycle's external review, supplied for `v3.1.4`, found that the fix's own stated scope - "all five rejection paths" - undercounted by one: `NewClient` has a sixth error-return path (the initial `url.Parse` failure, one function call above the five explicit checks) that still leaks. See §0 and §2.

> **Naming convention, unchanged from the last two cycles:** this file stays at this exact path forever - no date, no revision hash. It is always the current assessment of record; every external pointer to it (`CLAUDE.md`, `README.md`, `docs/README.md`, `tests/contract/README.md`) links here once and never needs updating again. **When the next assessment cycle happens:** move *this file's current content* to `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_<date>_<revision>.md` (using this file's own `Date`/`Revision assessed` header values above), prepend the usual supersession banner to that archived copy, and then overwrite *this path* with the new cycle's content. Do not create a new hash-suffixed file for the new cycle - the hash suffix is exclusively an archive-naming convention now.

---

## 0. Reconciling against the external review supplied for this cycle

The user supplied an unsolicited "Senior Software Engineering Review" of `v3.1.4` (8.9/10), consistent with this lineage's standing practice of verifying rather than trusting such input.

### 0.1 Citations, checked directly

| Review cites | Checked | Verdict |
|---|---|---|
| Tag `v3.1.4` → commit `0e400d1` | `git rev-parse v3.1.4^{commit}` | **Correct.** (Tag object itself is `c77f13b`, distinct from the commit - same distinction this lineage has flagged as easy to get wrong for three cycles running; the review cites the commit correctly again.) |
| PRs #58 (secret-echo fix), #59 (release) | `gh pr list --state merged --limit 4` | **Correct**, both merged with matching titles and merge commits. |
| `verify`/`apidiff`/`install-smoke-test` green at `0e400d1`; CI run IDs `29939602666`, `29939602513`, `29939620870` | `gh api repos/n-ae/nba-api-go/commits/0e400d1/check-runs`; `gh api repos/n-ae/nba-api-go/actions/runs/<id>` for each of the three cited IDs | **Correct.** All three named runs are real, `head_sha` `0e400d1`, `conclusion: success`, matching the review's stated run names (`CI`, `API Compatibility`, `Release Install Smoke Test`) and durations. |
| **The central claim: `url.Parse` failures still leak the raw `BaseURL` through the wrapped parser error** | `pkg/client/client.go:82-85` read directly; reproduced with `https://admin:hunter2@example.com/%zz` and `https://example.com/%zz?token=hunter2` against both raw `url.Parse` and the actual `client.NewClient` | **Correct, exactly as described - see §0.2.** |
| **The secondary claim: `https://:443` passes construction because `Host != ""` while `Hostname()` is empty** | `pkg/client/client.go:104` (`baseURL.Host == ""` check) read directly; reproduced with `url.Parse("https://:443")` and `client.NewClient(Config{BaseURL: "https://:443"})` | **Correct, exactly as described - see §0.2.** |

Every specific, checkable citation held up, for a third cycle running. This lineage verifies every time regardless of track record; this is another cycle where that discipline confirms a reliable input, and this time both of the review's substantive findings reproduce byte-for-byte against the actual `NewClient` return value, not just against the standard library in isolation.

### 0.2 Both substantive claims, reproduced directly against `NewClient`

**Finding A - the wrapped parser error.** `pkg/client/client.go:82-85`:

```go
baseURL, err := url.Parse(config.BaseURL)
if err != nil {
    return nil, fmt.Errorf("invalid base URL: %w", err)
}
```

`url.Parse`'s own error type renders as `parse "<input>": <reason>` - the complete input is part of `Error()`, and `%w` preserves it verbatim. Reproduced directly against the real `client.NewClient`, not just `url.Parse` in isolation:

```
NewClient(Config{BaseURL: "https://admin:hunter2@example.com/%zz"})
  → invalid base URL: parse "https://admin:hunter2@example.com/%zz": invalid URL escape "%zz"

NewClient(Config{BaseURL: "https://example.com/%zz?token=hunter2"})
  → invalid base URL: parse "https://example.com/%zz?token=hunter2": invalid URL escape "%zz"
```

Both `hunter2` and the surrounding credential/token structure are present in the actual error `NewClient` returns today. This is the same failure mode `v3.1.4` was supposed to close - a rejected, potentially secret-bearing `BaseURL` disclosed via the error raised to reject it - just via a sixth call site the `v3.1.4` fix didn't reach. **This is worth stating plainly rather than softening: the `v3.1.4` fix, its regression test (`TestNewClientRejectionErrorsDoNotLeakBaseURL`), its commit message, and last cycle's assessment (this same document, then covering `b3c605d`) all described the fix as covering "all five rejection paths" in `NewClient`. There are six.** The fifth-and-newly-counted path - `url.Parse`'s own error - is structurally different from the other five (it's a wrapped standard-library error, not a hand-written `fmt.Errorf` string this project controls the wording of), which is plausibly why an enumeration that looked at "the five checks added across `v3.1.2`/`v3.1.3`" missed the parse call sitting one line above them, but the practical effect for a caller is identical either way.

**Finding B - `Host` vs `Hostname()`.** `pkg/client/client.go:104`:

```go
if baseURL.Host == "" {
    return nil, fmt.Errorf("invalid base URL: missing host")
}
```

`url.URL.Host` includes an optional port; `url.URL.Hostname()` strips it. Reproduced directly:

```
url.Parse("https://:443") → Host=":443" Hostname="" Port="443"
NewClient(Config{BaseURL: "https://:443"}) → succeeds, err == nil
```

A `BaseURL` with no actual destination hostname constructs a working-looking `Client` today; the caller only discovers the problem on the first real `Get`, exactly the class of confusing-runtime-failure-instead-of-construction-time-failure this whole validation chain (`v3.1.2` onward) exists to prevent. Lower severity than Finding A - no secret is disclosed, it's a missed validation case, not a leak - but real and cheap to fix (`baseURL.Hostname() == ""` in place of `baseURL.Host == ""`).

### 0.3 Bottom line on the external review

Accurate on every checkable claim for a third cycle running, and both of its substantive findings reproduce exactly as stated against the real `NewClient`, not an approximation. Its P2 items beyond what's re-verified above (typed `ConfigError`, HTTP-server independent versioning, ecosystem-maturity commentary, endpoint-confidence tiers, a fuzz-testing recommendation) mostly restate positions this lineage has already taken and explained in `9eb3a9a`'s §0.2, `1b428f6`'s §0.3, and `b3c605d`'s §0.3 - not re-litigated here, **except the fuzz-testing suggestion, which this assessment is promoting to §5 as the right structural response to Finding A** (see below for why).

---

## 1. Executive verdict

**Grade: A- (unchanged, fifth consecutive cycle) - held deliberately, not automatically; see the caveat below.** The prior finding was closed, but closed incompletely, and the incompleteness was caught the same way everything else in this lineage gets caught: a second, independent look at the same code. That's the system working, but it's also the second cycle running where the specific defect class is "a `BaseURL` rejection path leaks the secret it's rejecting" - first found in `b3c605d` (the explicit checks), now found again in `0e400d1` (the parser-error path the first fix didn't enumerate). One recurrence of a defect class is a normal finding. Two, in immediate succession, on the same function, is the kind of pattern this lineage has previously treated as a signal to fix structurally rather than patch again (see the assessment-link staleness precedent: three recurrences before a structural fix was recommended, grade held throughout). This is only the second recurrence, so the grade holds one more cycle under that same precedent - but the remedy this cycle is explicitly the structural one (an invariant-based test), not another manually-enumerated patch, to try to close this out before a third recurrence would be needed to justify it.

**What went right:**
- The prior finding's explicit-check half was genuinely closed: none of the five hand-written `fmt.Errorf` calls in `NewClient` echo `config.BaseURL` anymore, confirmed by reading `client.go:102-125` directly.
- Release engineering reproduced exactly as claimed: `verify`/`apidiff`/`install-smoke-test` all green at the exact tag commit `0e400d1`, and for the first time this lineage independently re-ran a review's cited CI run IDs directly against the GitHub API (not just checked they existed) and confirmed name, `head_sha`, and conclusion all match.
- `go build`/`go vet`/`go test`/`golangci-lint` (both modules, checked separately) all clean, reproduced directly this session.
- The external review checked out on every citable claim for a third cycle running, and both of its substantive findings reproduce byte-for-byte against the real `NewClient`, not an approximation - the strongest verification outcome yet in this lineage.

**What keeps this at A- rather than moving it down:**
1. **A second instance of the same defect class, in the same function, one cycle after the first was "closed."** This is a legitimate process signal, not a coincidence: enumerating "the rejection checks I just wrote" instead of "every error-return path in this function" undercounted by one, and the undercounting was stated confidently (a regression test named after "all... paths", a commit message and an assessment repeating the same framing) rather than hedged. See §2, §5 for the fix and the structural remedy.
2. **Both severities are still low-likelihood in practice** for the same reason named every cycle since `b3c605d`: this project's own documented use case (NBA Stats, no auth) gives no caller a reason to put a secret in `BaseURL` today. Real, evidenced, cheap to fix - not a sign of runtime harm having occurred.

---

## 2. Verification ledger

Status legend: **CONFIRMED** (reproduced/read directly at `0e400d1`), **CLOSED** (carried from a prior assessment, now genuinely done), **PARTIALLY CLOSED** (the fix landed but didn't cover the full scope of the finding), **NEW** (found independently this cycle).

### From `b3c605d`

| # | Item (carried since `b3c605d`) | Status | Evidence |
|---|---|---|---|
| 4 | `NewClient`'s `BaseURL` validation errors echo the complete, unredacted input across all rejection sites | **PARTIALLY CLOSED** | The five explicit `fmt.Errorf` checks (`client.go:102,105,119,122,125`) no longer interpolate `config.BaseURL` - confirmed by reading the source. The sixth path, the wrapped `url.Parse` error at `client.go:84`, still does - see finding #6 below, which is the uncovered remainder of this same finding, not a new independent one. |

### New this cycle, via the external review (§0.2)

| # | Finding | Severity | Evidence |
|---|---|---|---|
| 6 | `NewClient`'s `url.Parse` failure path (`client.go:82-85`) wraps the standard library's parse error with `%w`, and that error's `Error()` string contains the complete original input - so a malformed, credential- or token-bearing `BaseURL` (e.g. one with an invalid percent-escape) still discloses the secret, via the one error-return path in `NewClient` the `v3.1.4` fix didn't enumerate. | Medium - same disclosure mechanism and same low-likelihood-today caveat as `b3c605d`'s finding #4, but notable for being the *second* instance of the identical failure mode in as many cycles, in the same function | §0.2, Finding A. Reproduced directly against `client.NewClient` (not just `url.Parse`) with two malformed credential/token-bearing inputs; both secrets present verbatim in the returned error. |
| 7 | `baseURL.Host == ""` (`client.go:104`) doesn't catch a host with no actual hostname, e.g. `https://:443` (`Host=":443"`, `Hostname()=""`) - construction succeeds with no real destination. | Low (no secret disclosed; a missed validation case that fails on the first `Get` instead of at construction, the same category of confusing-runtime-symptom every check in this chain since `v3.1.2` exists to prevent) | §0.2, Finding B. Reproduced directly against `client.NewClient("https://:443")`: succeeds, `err == nil`. |

---

## 3. C4 model

Level 1 unchanged. Level 2 nearly identical to `b3c605d`'s, with the core-client box still in caution (finding recurred rather than fully closing) and one new low-severity annotation for the hostname gap.

```mermaid
flowchart TD
    subgraph runtime["nba-api-go runtime"]
        server["HTTP API Server\n[cmd/nba-api-server]\n76.8% coverage - unchanged"]
        facades["Facades\n[pkg/stats, pkg/live]\nunchanged, fine"]
        endpoints["Generated + hand-written Endpoints\n[pkg/stats/endpoints]\n75.1% coverage - unchanged, fine"]
        core["Core Client\n[pkg/client]\n5 of 6 BaseURL rejection paths no\nlonger leak (PARTIAL, #4); the 6th -\nwrapped url.Parse error - still does\n(NEW, #6, medium); Host!=\"\" check\nmisses hostname-less URLs like\nhttps://:443 (NEW, #7, low)"]
        mw["Middleware\n[pkg/client/middleware]\nunchanged, fine"]
        static["Static Data\n[pkg/stats/static]\nunchanged, fine"]
        models["Models/Errors\n[pkg/models]\nunchanged, fine"]
    end

    subgraph devtime["Development-time"]
        gen["Code Generator\n[tools/generator]\nunchanged this cycle, fine"]
        contract["Contract Tests\n[tests/contract]\nunchanged, fine"]
        ci["CI\n[ci.yml, apidiff.yml,\nrelease-install-smoke.yml]\nall three green at the exact\nv3.1.4 release commit; 3 cited run\nIDs independently re-verified\nvia gh api this cycle"]
        drift["Live-drift workflow\nunchanged this cycle - fine"]
    end

    subgraph docs["Self-representation"]
        readme["README.md, docs/README.md, CLAUDE.md\n[all point at the stable assessment\npath - holding, no action needed]"]
        internal["This file\n[stable path, current as of this cycle]"]
    end

    nba2["NBA Stats API\n[stats.nba.com]\n5 of 141 endpoints reachable -\nunchanged, external fact"]

    server -->|"calls SDK"| facades
    facades --> endpoints
    endpoints -->|"GetJSON"| core
    core -->|"chained RoundTrip"| mw
    mw -->|"HTTPS, mostly blocked"| nba2
    gen -.->|"generates"| endpoints
    gen -.->|"generates"| server
    contract -.-> endpoints
    ci -.->|"verifies build + API compat +\ninstall, all green at 0e400d1"| runtime
    drift -.->|"weekly, narrow allowlist"| nba2
    endpoints --> models
    core --> models
    facades --> static
    readme -.->|"stable, no longer stale"| internal

    classDef fixed fill:#2f8f4e,color:#fff
    classDef caution fill:#c9862b,color:#fff
    classDef ext fill:#999999,color:#fff
    class facades,static,models,mw,drift,contract,ci,internal,readme fixed
    class server fixed
    class core caution
    class nba2 ext
```

---

## 4. Where the complexity budget goes (updated)

**Well spent, unchanged:** everything prior cycles already called well-spent - the two-layer outbound-path testing design, the stable-plus-archive documentation pattern, `v3.1.4`'s clean release engineering.

**Genuinely closed this cycle:** the five hand-written `BaseURL` checks no longer echo raw input - real progress, confirmed by direct read, not undone by the finding below.

**Recurred, worth a different kind of attention than "fix and move on":** the `BaseURL`-secret-echo defect class, now found twice in two cycles in the same function. The individual fix (§5 #1) is as cheap as last cycle's. What's worth budgeting differently is *how* the next fix is verified - an enumerated, hand-written test list is exactly the technique that undercounted last cycle, so repeating it a third time carries the same risk of missing a seventh path nobody's looked at yet. §5 promotes the external review's fuzz-testing suggestion for this specific reason: it tests the invariant ("no `NewClient` failure ever echoes a canary planted in `BaseURL`") rather than a fixed list of paths, so it doesn't need a human to have first enumerated every way `NewClient` can fail.

**Newly found, low severity, cheap:** the `Host`/`Hostname()` validation gap (finding #7) - unrelated defect class, one-line fix.

---

## 5. Recommended order of work

Budget reality unchanged: ~1.6h/week core maintenance.

### Immediate (~20-25 min)

1. **Stop wrapping the raw `url.Parse` error in `NewClient`** (`client.go:82-85`). `%w` preserves the parser's own `Error()` string, which contains the complete original input. Replace with a message that doesn't retain the input - e.g. `fmt.Errorf("invalid base URL: %s", err)` still leaks (same reason), so either drop the parser detail entirely (`errors.New("invalid base URL: malformed")`) or, if the specific parse failure reason is worth keeping for debugging, extract just `err.(*url.Error).Err` (the underlying reason, e.g. `invalid URL escape "%zz"`) rather than the whole `*url.Error`, which is the piece that doesn't contain the input. Closes finding #6, the uncovered remainder of `b3c605d`'s finding #4.
2. **Fix the `Host`/`Hostname()` gap**: change `if baseURL.Host == ""` to `if baseURL.Hostname() == ""` at `client.go:104`. Closes finding #7.
3. **Add regression tests for both**, extending `TestNewClientRejectionErrorsDoNotLeakBaseURL` with the parse-failure cases (`https://admin:hunter2@example.com/%zz`, `https://example.com/%zz?token=hunter2`) and adding a new case to `TestNewClientRejectsUnusableBaseURL` for `https://:443`.

### Next (~20-30 min, the structural remedy for the recurrence, not just the patch)

4. **Add an invariant-based (fuzz) test for the secret-echo defect class**, per the external review's suggestion: seed `testing.F` with credential-bearing URLs, query tokens, malformed percent-escapes, and IPv6 edge cases, each containing a high-entropy marker string, and assert the single invariant - *if `NewClient` returns an error, the marker never appears in `err.Error()`* - regardless of which internal path rejected it. This is deliberately not another manually-enumerated list: the whole reason this defect class has now recurred once is that "enumerate every rejection path by hand" already missed one path twice (implicitly, by never being asked to look at the parse-error path at all, until an external review did). An invariant test doesn't need the enumeration to be complete to catch the next unenumerated path.

### Not urgent, explicitly not a backlog item to keep re-budgeting for

- Everything `9eb3a9a`/`180a3db`/`1b428f6`/`b3c605d` already marked not-urgent (live-verifying the 136 unreachable endpoints, HTTP-server independent versioning policy, ecosystem-maturity commentary, a typed `ConfigError`) remains not-urgent for the same reasons already given in those assessments. The one item promoted out of that bucket this cycle is the fuzz test (#4 above), specifically because it's the direct structural response to a now-twice-recurred, source-grounded defect - not a speculative addition.

---

## 6. Documentation status

| File | Action taken by this assessment |
|---|---|
| `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_b3c605d.md` | New: outgoing content of this file (as of revision `b3c605d`) archived here in the same changeset, with a supersession banner matching the existing convention |
| This file | Overwritten with the new assessment of record (revision `0e400d1`, tag `v3.1.4`) |
| `CLAUDE.md`, `README.md`, `docs/README.md`, `tests/contract/README.md` | **Not touched by this assessment** - all four already point at this file's stable path; no update needed |
| `CHANGELOG.md`, `go.mod`, version constants | **Not touched** - no new user-facing change is being shipped by this assessment itself; the recommended fixes in §5 are follow-up commits, not part of this document |

No docs sprawl introduced this cycle - `docs/` still holds exactly one active assessment plus `adr/`/`archive/`.

---

## 7. Is this too complex for one person?

**Verdict unchanged: no, at the core.** Five consecutive cycles where verification, not blind trust, is what caught every real finding - including, this cycle, a gap in this lineage's *own* prior closure claim. That's worth stating without spin: last cycle's assessment (this same document, then covering `b3c605d`) said the fix closed "all five" `BaseURL` rejection paths, and a sixth existed the whole time. The value of running this process every cycle, on an external review or otherwise, is precisely that it catches exactly this kind of thing - a confidently-stated but incomplete closure - rather than letting the confident statement stand unchallenged until something in production surfaces it.

The one item worth a different kind of attention than "close it and move on": this is the second cycle running the same defect class (secret-bearing `BaseURL` leaking through a `NewClient` error) has been found. §5's structural remedy (an invariant fuzz test, not another enumerated list) is offered for the same reason the assessment-link staleness eventually got a structural fix rather than a fourth manual patch - the individual fix is cheap, but paying it repeatedly by the same error-prone method (manual enumeration) is the actual cost worth removing.

---

## 8. Bottom line

`b3c605d` → `0e400d1`: the prior finding was closed, but incompletely - `v3.1.4` fixed five of `NewClient`'s six `BaseURL`-rejection error paths, and the sixth (the wrapped `url.Parse` error) still discloses secrets in malformed input, found by this cycle's external review and independently reproduced directly against the real `NewClient` return value. A second, lower-severity gap (`Host` vs `Hostname()` letting `https://:443` construct successfully) was found the same way. Both are cheap fixes, both are low-likelihood in this project's actual documented use case, and both are now precisely evidenced with line numbers and reproductions. Grade holds at A- for a fifth cycle - deliberately, following the same precedent this lineage applied to the assessment-link staleness pattern (hold through the second recurrence, fix structurally rather than patch a third time) - with the recommended remedy this cycle explicitly promoted from "patch the leak" to "add an invariant test that doesn't depend on a human having enumerated every path."

---

*Assessment of record for revision `0e400d1` (tag `v3.1.4`), 2026-07-22. Supersedes this file's own prior content (revision `b3c605d`, tag `v3.1.3`, grade A-) as the current maintainability assessment. That prior content moves to `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_b3c605d.md` in the same changeset as this file.*
