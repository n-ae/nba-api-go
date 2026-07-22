> **Superseded.** This assessed revision `b3c605d` (grade A-). The current assessment of record is
> [`docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT.md`](../MAINTAINABLE_ARCHITECT_V4_ASSESSMENT.md) - as of
> the follow-up cycle that archived this file, that stable, hash-free path is permanent (see that
> document's naming-convention note near the top): it covered revision `0e400d1` and later (tag
> `v3.1.4`, grade A-, unchanged) at the time it was written, and will cover whatever the current cycle
> is by the time you're reading this. Retained here for history; see that document's section 2
> ("Verification ledger") for the item-by-item status of the one finding below - it turned out to be
> only *partially* closed by `v3.1.4`: five of `NewClient`'s six `BaseURL`-rejection error paths were
> fixed, but the sixth (the wrapped `url.Parse` error) still leaked, found and independently reproduced
> in the follow-up cycle. See that document's section 0 for the full account, including a second, lower
> severity finding (`Host` vs `Hostname()`) found the same cycle.

# Maintainable-Architect-v4 Assessment: nba-api-go

**Date:** 2026-07-22
**Revision assessed:** `b3c605d` (`main`, tag `v3.1.3`), go1.26.5 darwin/arm64
**Assessor:** maintainable-architect-v4
**Method:** Direct verification against source at HEAD, not against `CHANGELOG.md`'s prose or an unsolicited external review's prose - file reads of `pkg/client/client.go`, `pkg/client/client_test.go`, `CHANGELOG.md`, `go.mod`; a throwaway `net/url` script reproducing `url.Parse`'s exact behavior on adversarial inputs; `git rev-parse`/`git log`; `go build ./...`, `go vet ./...`, `go test ./...`, `golangci-lint run ./...` (root and `tools/generator` modules, run separately); and `gh pr list`/`gh api repos/.../commits/b3c605d/check-runs`/`gh run list` against the real `n-ae/nba-api-go` GitHub repository to independently check every checkable citation in an external review supplied for this cycle (see §0). All green. No production code was modified while writing this file.

**Why now:** the prior assessment of record (this same file, then covering revision `1b428f6`/tag `v3.1.2`, grade A-) closed with three open findings: `BaseURL` didn't reject userinfo/query/fragment, no positive-case `BaseURL` test existed, and the assessment-link staleness pattern had recurred three cycles running. `v3.1.3` (tag `b3c605d`) closed all three in one small patch release. This cycle's external review, supplied for `v3.1.3`, surfaced a genuinely new, real, and somewhat ironic finding in the fix that closed the first of those three - see §0 and §2.

> **Naming convention, unchanged from the last cycle:** this file stays at this exact path forever - no date, no revision hash. It is always the current assessment of record; every external pointer to it (`CLAUDE.md`, `README.md`, `docs/README.md`, `tests/contract/README.md`) links here once and never needs updating again. **When the next assessment cycle happens:** move *this file's current content* to `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_<date>_<revision>.md` (using this file's own `Date`/`Revision assessed` header values above), prepend the usual supersession banner to that archived copy, and then overwrite *this path* with the new cycle's content. Do not create a new hash-suffixed file for the new cycle - the hash suffix is exclusively an archive-naming convention now.

---

## 0. Reconciling against the external review supplied for this cycle

The user supplied an unsolicited "Senior Software Engineering Review" of `v3.1.3` (8.8/10), consistent with this lineage's standing practice of verifying rather than trusting such input.

### 0.1 Citations, checked directly

| Review cites | Checked | Verdict |
|---|---|---|
| Tag `v3.1.3` → commit `b3c605d` | `git rev-parse v3.1.3^{commit}` | **Correct.** (`git rev-parse v3.1.3` alone returns the annotated tag object's hash, `913b1e4`, not the commit - same tag-object-vs-commit distinction this lineage flagged as easy to get wrong last cycle, and the review gets it right again.) |
| PRs #53 (BaseURL hardening), #54 (positive tests), #55 (stable assessment path), #56 (release) | `gh pr list --state merged --limit 6` | **Correct**, all four merged with matching titles and merge commits, in the order described. |
| `verify`/`apidiff`/`install-smoke-test` all green at `b3c605d` | `gh api repos/n-ae/nba-api-go/commits/b3c605d/check-runs`, `gh run list --branch v3.1.3` | **Correct.** All four checks (`verify`, `apidiff`, `install-smoke-test`, Socket Security) report `success` at the exact tag commit. |
| Workflow filenames `ci.yml`, `apidiff.yml`, `release-install-smoke.yml` | `ls .github/workflows/` | **Correct**, exact filenames. |
| `go.mod`'s `go 1.26.5` | `go.mod:3` | **Correct.** |
| `DefaultMaxResponseBytes` = 50 MiB | `pkg/client/client.go:33` | **Correct**, `50 * 1024 * 1024`. |
| `TestNewClientAcceptsValidBaseURL`'s six cases, including `http://[::1]:8080` | `pkg/client/client_test.go:191-216` read in full | **Correct**, all six cases match verbatim, including the IPv6 case. |
| **The central claim: rejected `BaseURL`s can leak credentials/secrets through error messages** | `pkg/client/client.go:92,95,109,112,115` read in full | **Correct, and worse than described - see §0.2.** |

Every specific, checkable citation held up. This is worth stating plainly rather than mechanically repeating the prior cycle's skepticism: this lineage's practice is to verify every time, not to extend trust based on a good track record, and every time includes the times the input turns out to be reliable.

### 0.2 The central claim, verified and found broader than described

The review's headline finding: `NewClient`'s new userinfo/query/fragment rejection (added in `v3.1.3` to close `#3`) interpolates the complete, unredacted `config.BaseURL` into the returned error via `fmt.Errorf("...%q...", config.BaseURL)`. A caller passing `https://admin:secret@example.com` gets back an error containing the literal string `https://admin:secret@example.com` - the exact credential the check exists to keep out of use, now disclosed in whatever logs, error trackers, or CI output capture that error.

Confirmed directly by reading `pkg/client/client.go:108-116`:

```go
if baseURL.User != nil {
    return nil, fmt.Errorf("invalid base URL %q: must not contain userinfo (e.g. \"user:pass@\") - configure authentication separately", config.BaseURL)
}
if baseURL.RawQuery != "" {
    return nil, fmt.Errorf("invalid base URL %q: must not contain a query string", config.BaseURL)
}
```

What makes this worth calling out precisely: the doc comment two lines above these checks (`client.go:98-100`) already names the exact risk being reintroduced - *"userinfo... risks credentials leaking into logs, error messages, and metrics labels wherever BaseURL or a wrapping error gets printed"* - as the stated justification for adding the check. The fix's own comment predicts the bug it ships in the very next line. No test in `TestNewClientRejectsBaseURLWithUserinfoQueryOrFragment` (`client_test.go:227-239`) asserts anything about the error's content, only that construction fails - so this shipped with full test coverage of "rejects" and zero coverage of "doesn't also leak what it rejects."

**Independently found, beyond what the review states:** this is not new in `v3.1.3` - it's a latent property of the `v3.1.2` scheme/host checks too, and `v3.1.3` just added two more call sites with the same defect. Verified directly with a throwaway script reproducing `url.Parse`:

```
url.Parse("ftp://user:pass@example.com") → scheme="ftp" host="example.com" user="user:pass" err=nil
```

`ftp://user:pass@example.com` fails the *scheme* check at `client.go:92` (`v3.1.2`'s original check, not `v3.1.3`'s), which also interpolates `config.BaseURL` unredacted - so a credential-bearing `BaseURL` with a typo'd scheme has leaked via this code path since `v3.1.2` shipped, one full release before the review's stated scope. All five `fmt.Errorf` call sites in `NewClient` (`client.go:92,95,109,112,115`) share the same pattern and need the same fix, not just the three the review names.

### 0.3 Bottom line on the external review

Accurate on every checkable claim, and its central finding is real, well-reasoned, and - on independent investigation - understates its own scope (pre-existing since `v3.1.2`, not introduced fresh in `v3.1.3`). Its P2 items beyond what's re-verified above (typed error values, HTTP-server independent versioning, ecosystem-maturity commentary, endpoint-confidence tiers, raw-fixture replay infrastructure) restate positions this lineage has already taken and explained in `9eb3a9a`'s §0.2 and `1b428f6`'s §0.3 - not re-litigated here since nothing in this review's version of those sections presents new evidence changing that reasoning.

---

## 1. Executive verdict

**Grade: A- (unchanged, fourth consecutive cycle).** All three findings the prior assessment named were closed with direct, verifiable evidence, and this cycle's external review found one genuinely new, real defect - not a process failure, but a legitimate secret-disclosure bug in code that was itself written to fix a security-adjacent gap. That irony doesn't change the grade calculus this lineage has applied every cycle: small, real, cheap-to-fix findings surfaced by rigorous verification are what a healthy solo-maintained project's assessment cycle should produce, not evidence of decline.

**What went right:**
- All three `1b428f6` findings closed and independently reproduced: `NewClient` now rejects userinfo/query/fragment (`client.go:108-116`), `TestNewClientAcceptsValidBaseURL` covers six positive cases including IPv6 (`client_test.go:191-216`), and the assessment now lives at a stable path that `README.md`/`docs/README.md`/`CLAUDE.md` already point at permanently - no link-staleness fix was needed *this* cycle, closing the loop the structural fix was built for.
- Release engineering reproduced exactly as claimed: `verify`/`apidiff`/`install-smoke-test` all green at the exact tag commit `b3c605d`, four PRs (#53-#56) merged in the described order with matching content.
- `go build`/`go vet`/`go test`/`golangci-lint` (both modules, checked separately) all clean, reproduced directly this session.
- The external review checked out on every citable claim for a second cycle running, and its central substantive finding is real - verified independently rather than taken on the review's word, and found to be broader in scope than the review itself claimed (§0.2).

**What keeps this at A- rather than moving it up:**
1. **A real secret-disclosure defect shipped in the exact code meant to reduce that risk.** `NewClient`'s five `BaseURL` validation error sites (two from `v3.1.2`, three from `v3.1.3`) all interpolate the complete, unredacted input string - so a rejected credential- or token-bearing `BaseURL` is disclosed via the very error raised to prevent its use. Low-likelihood in this project's own documented use case (NBA Stats requires no auth, so no legitimate caller has a reason to put a secret in `BaseURL` today), but real, evidenced, and cheap to fix (§2, §5).
2. Nothing else new this cycle - the remaining P2 commentary in the external review restates already-considered, already-declined architectural proposals (see §0.3).

---

## 2. Verification ledger

Status legend: **CONFIRMED** (reproduced/read directly at `b3c605d`), **CLOSED** (carried from a prior assessment, now genuinely done), **NEW** (found independently this cycle).

### Closed this cycle (all three findings from `1b428f6`)

| # | Item (carried since `1b428f6`) | Status | Evidence |
|---|---|---|---|
| 1 | `NewClient` doesn't reject `BaseURL` userinfo, query string, or fragment | **CLOSED** | `pkg/client/client.go:108-116`, three checks, each read directly. `TestNewClientRejectsBaseURLWithUserinfoQueryOrFragment` covers all three shapes. |
| 2 | No positive-case `BaseURL` test matrix | **CLOSED** | `TestNewClientAcceptsValidBaseURL`, six cases (bare host, host+path, port, IPv4, IPv6, subdomain+port+path), each asserting both successful construction and correct `buildURL` output. |
| 3 | Assessment-link staleness pattern (recurred three cycles running) | **CLOSED** | `README.md`/`docs/README.md`/`CLAUDE.md` all already point at this stable path; no link-fix commit was needed this cycle - the structural fix from last cycle is working as designed. |

### New this cycle, via the external review's lead and independent extension (§0.2)

| # | Finding | Severity | Evidence |
|---|---|---|---|
| 4 | `NewClient`'s `BaseURL` validation errors echo the complete, unredacted input string across all five error sites (`client.go:92,95,109,112,115`), disclosing exactly the credentials/tokens the userinfo/query checks exist to keep out of use. Pre-existing since `v3.1.2` (the scheme/host checks), not introduced fresh in `v3.1.3`. | Medium (a real disclosure, not just a validation gap - the failure is loud *and* leaks, versus prior cycles' "loud but silent-on-secrets" findings; likelihood is low today given no documented use case puts secrets in `BaseURL`, but the defect is unconditional whenever one does) | §0.2. Read `client.go` in full; reproduced the pre-`v3.1.3` leak path (`ftp://user:pass@example.com` failing the `v3.1.2` scheme check) with a throwaway `net/url` script. No test in `client_test.go` asserts anything about error *content*, only that construction fails. |

---

## 3. C4 model

Level 1 unchanged. Level 2 nearly identical to `1b428f6`'s - another small cycle - with the three closed boxes turning green and one new, narrowly-scoped caution box on the core client.

```mermaid
flowchart TD
    subgraph runtime["nba-api-go runtime"]
        server["HTTP API Server\n[cmd/nba-api-server]\n76.8% coverage - unchanged"]
        facades["Facades\n[pkg/stats, pkg/live]\nunchanged, fine"]
        endpoints["Generated + hand-written Endpoints\n[pkg/stats/endpoints]\n75.1% coverage - unchanged,\noutbound path tested two ways (fine)"]
        core["Core Client\n[pkg/client]\nBaseURL userinfo/query/fragment\nnow rejected (CLOSED, #1); all 5\nrejection error sites echo the raw,\nunredacted input incl. secrets\n(NEW, #4, medium severity)"]
        mw["Middleware\n[pkg/client/middleware]\nunchanged, fine"]
        static["Static Data\n[pkg/stats/static]\nunchanged, fine"]
        models["Models/Errors\n[pkg/models]\nunchanged, fine"]
    end

    subgraph devtime["Development-time"]
        gen["Code Generator\n[tools/generator]\nunchanged this cycle, fine"]
        contract["Contract Tests\n[tests/contract]\nunchanged, fine"]
        ci["CI\n[ci.yml, apidiff.yml,\nrelease-install-smoke.yml]\nall three green at the exact\nv3.1.3 release commit, verified\ndirectly via gh api"]
        drift["Live-drift workflow\nunchanged this cycle - fine"]
    end

    subgraph docs["Self-representation"]
        readme["README.md, docs/README.md, CLAUDE.md\n[all point at the stable assessment\npath - CLOSED, #3, structural fix\nfrom last cycle holding as designed]"]
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
    ci -.->|"verifies build + API compat +\ninstall, all green at b3c605d"| runtime
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

**Well spent, unchanged:** everything `1b428f6` already called well-spent - the two-layer outbound-path testing design, proportionate `BaseURL` validation, the stable-plus-archive documentation pattern (now proven working a cycle later with zero link-staleness to fix).

**Newly well spent:** closing all three prior findings in one small, focused patch (`v3.1.3`), each with dedicated test coverage for the closure itself.

**Newly leaking, small, precisely evidenced:** the `BaseURL` error-message secret echo (finding #4) - cheap to fix (sanitize the interpolated value or drop it from the message; the informative part of each error is *which* component is disallowed, not a verbatim echo of the disallowed input), and now evidenced with the exact line numbers and a reproduction of the pre-`v3.1.3` instance of the same pattern.

**Structurally resolved, worth naming as a completed pattern:** the assessment-link staleness cycle. Three prior cycles each fixed two stale links and predicted the next staleness; this cycle required zero link fixes because the stable-path convention introduced last cycle did exactly what it was designed to do. Worth calling out explicitly since it's the first structural fix in this lineage's history that's been verified to hold for a full cycle after being introduced, not just proposed.

---

## 5. Recommended order of work

Budget reality unchanged: ~1.6h/week core maintenance.

### Immediate (~15-20 min)

1. **Stop echoing the raw `BaseURL` in all five `NewClient` validation errors** (`client.go:92,95,109,112,115`). The component-level messages ("scheme must be http or https", "missing host", "must not contain userinfo", etc.) are already informative without the verbatim input; drop `%q, config.BaseURL` from all five `fmt.Errorf` calls, or - if a diagnostic string is worth keeping - build a sanitized copy (`sanitized := *baseURL; sanitized.User = nil; sanitized.RawQuery = ""; sanitized.Fragment = ""`) and print that instead. Closes finding #4, and along with it the latent `v3.1.2`-era instance of the same pattern in the scheme/host checks, not just the three `v3.1.3` checks the external review named.
2. **Add a regression test proving the fix**: construct `NewClient` with a credential- and token-bearing `BaseURL` for each of the five rejection paths, and assert the returned error's `.Error()` string does not contain the injected secret. This is the one test class currently missing from `client_test.go`'s otherwise-thorough `BaseURL` coverage.

### Not urgent, explicitly not a backlog item to keep re-budgeting for

- Everything `9eb3a9a`/`180a3db`/`1b428f6` already marked not-urgent (live-verifying the 136 unreachable endpoints, HTTP-server independent versioning policy, ecosystem-maturity commentary) remains not-urgent for the same reasons already given in those assessments - not re-litigated here.
- The external review's broader architectural proposals (endpoint confidence tiers, raw-fixture replay infrastructure, a pre-tag consumer build gate, structured schema-mismatch error types) are real ideas but not backed by a defect found in this codebase this cycle or last; consistent with this lineage's standing skepticism of scope-expanding suggestions without a specific source-grounded case.

---

## 6. Documentation status

| File | Action taken by this assessment |
|---|---|
| `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_9eb3a9a.md` | Archived to `docs/archive/` in the same changeset as this file, with a supersession banner matching the existing convention |
| This file | New assessment of record |
| `CLAUDE.md` | Updated: header "Grade" line, "For Maintainers" section, and "Next assessment" footer line now point at this file |
| `docs/README.md`, `README.md` | **Not updated by this assessment** - both currently link to `9eb3a9a`, correct as of this cycle's starting state but now one cycle stale (see §1 finding #1, §5 Immediate #1, and finding #5's proposed structural fix) - consistent with this lineage's established scope boundary (the assessment names the fix; a following commit executes it) |
| `CHANGELOG.md`, `go.mod`, version constants | **Not touched** - no new user-facing change is being shipped by this assessment itself |

No docs sprawl introduced this cycle - `docs/` still holds exactly one active assessment plus `adr/`/`archive/`.

---

## 7. Is this too complex for one person?

**Verdict unchanged: no, at the core, and the edges remain small and shrinking.** Three consecutive small, clean cycles now: findings closed, verified independently, zero regressions, proportionate fixes. This cycle also included a rare "unsolicited external review turns out to be fully reliable" outcome - useful precisely because it demonstrates the verification discipline isn't theater against a review that happens to be flawed; it holds up equally when the review is good, confirming genuinely useful, evidenced findings (userinfo/query/fragment, the positive-test-matrix gap) rather than either rubber-stamping or reflexively discounting the input.

The one item worth a different kind of attention than "close it and move on": the assessment-link staleness pattern, now evidenced three times. This is exactly the kind of finding a solo maintainer is prone to under-invest in fixing structurally, because each individual instance is cheap (~10 minutes) and the temptation is to keep paying that small cost rather than spend the slightly larger one-time cost of removing the pattern entirely. §5's recommendation names the actual fix rather than the fourth repetition of the symptom.

---

## 8. Bottom line

`9eb3a9a` → `1b428f6`: a third consecutive small, clean cycle. Both findings the prior assessment named - outbound-path testing, `BaseURL` validation strictness - closed with direct, independently-reproduced evidence, via a genuinely well-designed two-layer test strategy for the path-testing half. An unsolicited external review supplied for this cycle checked out on every citable claim, a different outcome from the previous cycle's partially-fabricated one, verified with the same rigor either way; two of its findings (`BaseURL` userinfo/query/fragment, missing positive-case tests) are real and now on the ledger. The recurring assessment-link staleness pattern is named for the third time, this time with a structural fix recommended rather than the same two-line patch repeated a fourth time. Grade holds at A- - proportionate execution on a small cycle, no new risk, no structural change in either direction.

---

*Assessment of record for revision `b3c605d` (tag `v3.1.3`), 2026-07-22. Supersedes this file's own prior content (revision `1b428f6`, one commit past tag `v3.1.2`, grade A-) as the current maintainability assessment. That prior content moves to `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_1b428f6.md` in the same changeset as this file.*
