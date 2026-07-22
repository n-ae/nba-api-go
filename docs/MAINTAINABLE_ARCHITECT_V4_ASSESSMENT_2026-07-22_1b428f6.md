# Maintainable-Architect-v4 Assessment: nba-api-go

**Date:** 2026-07-22
**Revision assessed:** `1b428f6` (`main`, one commit past tag `v3.1.2`), go1.26.5 darwin/arm64
**Assessor:** maintainable-architect-v4
**Method:** Direct verification against source at HEAD, not against `CHANGELOG.md`'s prose or an unsolicited external review's prose - file reads of `pkg/client/client.go`, `pkg/client/client_test.go`, `tools/generator/templates/endpoint_test.tmpl`, `tools/generator/generator_test.go` (the `TestEndpointPathMatchesNameConvention` addition), `CHANGELOG.md`, `CLAUDE.md`, `README.md`, `docs/README.md`; `git rev-parse`/`git log`; `go build ./...`, `go vet ./...`, `go test ./...`, `go test -cover` (reproducing both headline coverage numbers exactly), `golangci-lint run ./...` (root and `tools/generator` modules, run separately); and `gh api`/`gh pr list`/`gh release list` against the real `n-ae/nba-api-go` GitHub repository to independently check every checkable citation in an external review supplied for this cycle (see §0). All green. No production code was modified while writing this file.

**Why now:** the prior assessment of record (`docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_9eb3a9a.md`, grade A-, reviewed `9eb3a9a`/one commit past `v3.1.1`) closed with exactly two open findings: no test asserted an SDK endpoint's outbound request path, and `NewClient` accepted any `BaseURL` string `url.Parse` could parse regardless of scheme/host. `v3.1.2` (tag `99a6d68`) closed both in one small patch release, and one further doc-currency commit (`1b428f6`, current HEAD) fixed `README.md`/`docs/README.md`'s assessment links - the same "link goes stale the same day a new assessment lands" pattern this lineage has now hit three cycles running (see §1).

---

## 0. Reconciling against the external review supplied for this cycle

The user supplied an unsolicited "Senior Software Engineering Review" of `v3.1.2` (8.9/10) along with an instruction to verify it, consistent with this lineage's standing practice. Unlike the review supplied for the prior cycle (`9eb3a9a`'s §0, which had four independent categories of fabricated citations - a nonexistent directory, a nonexistent package, two wrong workflow filenames, three 404 CI run IDs), **this review's citations check out.**

### 0.1 Citations, checked directly

| Review cites | Checked | Verdict |
|---|---|---|
| Tag `v3.1.2` → commit `99a6d68` | `git rev-parse v3.1.2^{commit}` | **Correct.** (`git rev-parse v3.1.2` alone returns the *annotated tag object*'s own hash, `e632003`, not the commit - the review's commit citation specifically is right, worth noting precisely since the two hashes are easy to conflate.) |
| CI run `29927799455` ("CI") | `gh api repos/n-ae/nba-api-go/actions/runs/29927799455` | **Real.** `head_sha` `99a6d68`, `conclusion: success`. |
| CI run `29927793205` ("API Compatibility") | same | **Real.** `head_sha` `99a6d68`, `conclusion: success`. |
| CI run `29927828292` ("Release Install Smoke Test") | same | **Real.** `head_sha` `99a6d68`, `conclusion: success`. |
| `templates/endpoint_test.tmpl` outbound-path assertion | `tools/generator/templates/endpoint_test.tmpl` read in full | **Correct.** `const wantPath = "/{{.Endpoint}}"`, captured via `gotPath = r.URL.Path` in the stub handler, asserted at the end of every generated test. |
| `TestEndpointPathMatchesNameConvention`, sole exception `TeamYearOverYearSplits` | `tools/generator/generator_test.go:878` read in full | **Correct**, including the exception mapping and the doc comment explaining why the two test layers are structurally different (one can't catch a metadata-level typo since both sides derive from the same field; the other independently re-derives the expected value). |
| `client.NewClient`'s new scheme/host check, rejected-example list (`""`, `"not-a-url"`, `"example.com"`, `"//example.com"`, `"ftp://example.com"`) | `pkg/client/client.go`, `pkg/client/client_test.go` read in full | **Correct**, down to the exact five example strings in `TestNewClientRejectsUnusableBaseURL`. |
| `apidiff.yml`'s `workflow_dispatch` | Already verified by the prior assessment (`9eb3a9a`'s finding #7); unchanged this cycle | **Correct**, and correctly attributed to the right prior cycle, not claimed as new in `v3.1.2`. |

Every specific, checkable citation held up. This is worth stating plainly rather than mechanically repeating the prior cycle's skepticism: this lineage's practice is to verify every time, not to extend trust based on a good track record, and every time includes the times the input turns out to be reliable.

### 0.2 Substantive claims not yet spot-checked above

- **`NewClient` doesn't reject `BaseURL` userinfo/query/fragment** - confirmed real. `grep -n "\.User\|RawQuery\|Fragment" pkg/client/client.go` finds no validation logic anywhere in `NewClient`; the only `RawQuery` reference in the file is `buildURL` constructing a request's query from caller-supplied `params`, unrelated to validating the configured `BaseURL` itself. A `BaseURL` like `https://user:pass@stats.nba.com` or `https://stats.nba.com?token=x` is accepted today. Real, previously unflagged by any assessment in this lineage.
- **No positive `BaseURL` test matrix exists** - confirmed. `grep -n "^func Test" pkg/client/client_test.go | grep -i url` finds exactly three: `TestClient_buildURL` (query-encoding logic, not `BaseURL` acceptance), `TestNewClientRejectsInvalidBaseURL`, `TestNewClientRejectsUnusableBaseURL` - all negative-only. Nothing asserts that `https://example.com:8443/base/path`, an IPv6 host, or a bare `http://localhost:8080` construct successfully with the expected `baseURL` fields set. Real; low severity (a false-positive rejection would fail loudly and immediately, not silently), but a legitimate documentation-by-test gap the review correctly identified.
- **Coverage numbers, apidiff/CI/install-smoke all green at the exact release** - reproduced directly this cycle (§ header). Unchanged from `9eb3a9a` since `v3.1.2` didn't touch generation logic, only validation and test-template code.

### 0.3 Bottom line on the external review

Accurate on every checkable claim, including three specific CI run IDs and a commit-vs-tag-object distinction easy to get wrong. Its P1/P2 sections beyond what's re-verified above (HTTP-server-versioning policy, live-reachability framing, ecosystem-maturity commentary, the "confidence tiers" and "repository pattern" architectural proposals) restate positions this lineage has already taken and explained at length in `9eb3a9a`'s §0.2 items 3-4 and its own "not recommending" reasoning - not re-litigated here since nothing in this review's version of those sections presents new evidence changing that reasoning. Two of its P2 items (userinfo/query/fragment, positive-case test matrix) are genuinely new to this lineage's ledger and real; captured in §2 below.

---

## 1. Executive verdict

**Grade: A- (unchanged, third consecutive cycle).** Another small, clean cycle: both findings the prior assessment named were closed with direct, verifiable evidence, nothing regressed, and the fix sizes were proportionate (a ~15-line validation addition, a ~10-line template addition plus one new standalone generator test) rather than over-built.

**What went right:**
- Both `9eb3a9a` findings closed and independently reproduced, not taken on `CHANGELOG.md`'s word: `NewClient` now rejects five categories of unusable `BaseURL` (empty, unparseable, non-`http(s)` scheme, relative/protocol-relative, wrong scheme) with a dedicated test table; every one of the 135 generated endpoint tests now asserts `r.URL.Path` against the endpoint's own metadata, plus a new, structurally independent `TestEndpointPathMatchesNameConvention` that catches a metadata-level typo the per-endpoint test cannot (both sides of that assertion would otherwise derive from the same field).
- The two-layer path-testing design is genuinely good engineering, not just "more tests": the per-endpoint test catches generator/template regressions; the convention test catches the exact class of bug (a metadata typo) that produced the ten broken endpoint paths `v3.1.0` fixed. The two failure modes don't overlap, and the codebase's own test comments explain why, unprompted by any external review.
- `go build`/`go vet`/`go test -race`-scoped/`golangci-lint` (both modules, checked separately)/`make test-examples` all clean, reproduced directly this session.
- The external review supplied for this cycle checked out on every citable claim - a genuinely different outcome from the prior cycle's confabulated one, and treated with the same verification discipline either way (see §0).

**What keeps this at A- rather than moving it up:**
1. **The assessment-link staleness pattern has now recurred a third time, structurally, not by anyone's mistake.** `9eb3a9a` named this exact issue about its own predecessor (`180a3db`) and predicted it would recur for itself - it did (`README.md`/`docs/README.md` briefly pointed at `180a3db` after `9eb3a9a` existed, fixed in `1b428f6`), and by the same logic, this file's own publication makes those two links stale again, on the same day, for the third cycle running. This is worth a structural fix now that it's evidenced three times running, not just named again (see §5).
2. **Two small, real gaps surfaced by this cycle's external review remain open**: `BaseURL` doesn't reject userinfo/query/fragment components, and no positive-case `BaseURL` test exists. Neither is urgent (both fail loudly, not silently, if wrong), but both are cheap and now precisely evidenced (§0.2).

---

## 2. Verification ledger

Status legend: **CONFIRMED** (reproduced/read directly at `1b428f6`), **CLOSED** (carried from a prior assessment, now genuinely done), **NEW** (found independently this cycle).

### Closed this cycle (both findings from `9eb3a9a`)

| # | Item (carried since `9eb3a9a`) | Status | Evidence |
|---|---|---|---|
| 1 | No test asserts an SDK endpoint's outbound URL path/query | **CLOSED** | `tools/generator/templates/endpoint_test.tmpl`: `wantPath`/`gotPath` assertion in every generated test, plus the independent `TestEndpointPathMatchesNameConvention` in `generator_test.go`. Both read in full this cycle. |
| 2 | `NewClient`'s `BaseURL` validation accepts any `url.Parse`-able string | **CLOSED** | `pkg/client/client.go`: scheme-must-be-http(s) and host-must-be-non-empty checks, both reproduced reading the live source. `TestNewClientRejectsUnusableBaseURL` covers five distinct unusable shapes. |

### New this cycle, independently made and via the external review's leads (§0.2)

| # | Finding | Severity | Evidence |
|---|---|---|---|
| 3 | `NewClient` doesn't reject `BaseURL` userinfo, query string, or fragment | Low-medium (credential-leakage risk if userinfo is ever used; low likelihood given no documented use case) | §0.2. `grep` confirms no such check exists. |
| 4 | No positive-case `BaseURL` test matrix (valid host+port, IPv6, subpaths) | Low (a wrongful-rejection bug would be loud and immediate, not silent) | §0.2. Three existing `BaseURL`-related tests are all negative-only. |
| 5 | The assessment-link staleness pattern is now evidenced three cycles running (`180a3db`→`9eb3a9a`, `9eb3a9a`→`1b428f6`, and this file's own publication will make `1b428f6`'s links stale a fourth time) | Low severity, but worth a structural fix given the repeat count | `README.md:31`, `docs/README.md:18` both currently point at `..._9eb3a9a.md`, which this file supersedes as of this commit. |

---

## 3. C4 model

Level 1 unchanged. Level 2 nearly identical to `9eb3a9a`'s - this was another small cycle - with the two closed boxes turning green and the `BaseURL`/doc-link boxes carrying forward at reduced (not zero) severity.

```mermaid
flowchart TD
    subgraph runtime["nba-api-go runtime"]
        server["HTTP API Server\n[cmd/nba-api-server]\n76.8% coverage - unchanged"]
        facades["Facades\n[pkg/stats, pkg/live]\nunchanged, fine"]
        endpoints["Generated + hand-written Endpoints\n[pkg/stats/endpoints]\n75.1% coverage; outbound path now\nasserted per-endpoint AND via an\nindependent metadata convention\ntest (CLOSED, #1)"]
        core["Core Client\n[pkg/client]\nBaseURL scheme/host now\nvalidated (CLOSED, #2);\nuserinfo/query/fragment still\nunchecked (NEW, #3, low severity)"]
        mw["Middleware\n[pkg/client/middleware]\nunchanged, fine"]
        static["Static Data\n[pkg/stats/static]\nunchanged, fine"]
        models["Models/Errors\n[pkg/models]\nunchanged, fine"]
    end

    subgraph devtime["Development-time"]
        gen["Code Generator\n[tools/generator]\nendpoint_test.tmpl now asserts\noutbound path; new standalone\nTestEndpointPathMatchesNameConvention\n(both CLOSED, #1)"]
        contract["Contract Tests\n[tests/contract]\nunchanged, fine"]
        ci["CI\n[ci.yml, apidiff.yml,\nrelease-install-smoke.yml]\nall three green at the exact\nv3.1.2 release commit, verified\ndirectly via gh api"]
        drift["Live-drift workflow\nunchanged this cycle - fine"]
    end

    subgraph docs["Self-representation"]
        readme["README.md, docs/README.md\n[assessment links fixed to 9eb3a9a\nin 1b428f6 - already one cycle\nstale again as of THIS file's own\npublication (NEW, #5, 3rd\nrecurrence - structural fix\nrecommended)]"]
        internal["CLAUDE.md\n[current as of this cycle]"]
    end

    nba2["NBA Stats API\n[stats.nba.com]\n5 of 141 endpoints reachable -\nunchanged, external fact"]

    server -->|"calls SDK"| facades
    facades --> endpoints
    endpoints -->|"GetJSON, path now\nasserted two ways"| core
    core -->|"chained RoundTrip"| mw
    mw -->|"HTTPS, mostly blocked"| nba2
    gen -.->|"generates"| endpoints
    gen -.->|"generates"| server
    contract -.-> endpoints
    ci -.->|"verifies build + API compat +\ninstall, all green at 99a6d68"| runtime
    drift -.->|"weekly, narrow allowlist"| nba2
    endpoints --> models
    core --> models
    facades --> static
    readme -.->|"stale again, 3rd cycle running"| internal

    classDef fixed fill:#2f8f4e,color:#fff
    classDef caution fill:#c9862b,color:#fff
    classDef ext fill:#999999,color:#fff
    class facades,static,models,mw,drift,contract,ci,internal fixed
    class server fixed
    class core,endpoints caution
    class readme caution
    class nba2 ext
```

---

## 4. Where the complexity budget goes (updated)

**Well spent, unchanged:** everything `9eb3a9a` already called well-spent, plus the two-layer path-testing design added this cycle - proportionate to the actual bug class it defends against (the ten broken paths in `v3.0.0`'s predecessor work), not a speculative defense against a hypothetical.

**Newly well spent:** closing both findings from the prior assessment in one small, focused patch, with a test for each closure rather than a bare code change - the `BaseURL` fix shipped with a five-case negative test table; the path-assertion fix shipped with both a per-endpoint test change and a standalone convention test explaining, in its own doc comment, exactly why it's not redundant with the per-endpoint one. That level of self-documentation inside the test code itself is worth naming as a positive pattern, not just a passing check.

**Still leaking, small, newly precise:** `BaseURL` userinfo/query/fragment (finding #3) and the missing positive-case test matrix (finding #4) - both cheap, both low severity, both now named with direct evidence instead of general category.

**Structurally leaking, now evidenced three times:** the assessment-link staleness pattern (finding #5). This is the first cycle worth proposing an actual structural fix rather than repeating the observation a fourth time - see §5.

---

## 5. Recommended order of work

Budget reality unchanged: ~1.6h/week core maintenance.

### Immediate (~20 min)

1. **Update `docs/README.md` and `README.md`'s assessment links** from `9eb3a9a` to `1b428f6` (this file). Same fix as the last two cycles. (~5 min)
2. **Add a `BaseURL` scheme/host check extension**: reject non-empty `URL.User`, `URL.RawQuery`, and `URL.Fragment` in `NewClient`, with a clear error per component (matching the existing per-check error-message pattern already used for scheme/host). Closes finding #3. (~15 min)

### Next (~30-45 min, first cycle this is worth doing rather than deferring again)

3. **Break the assessment-link staleness cycle structurally** instead of fixing it a fourth time: introduce one stable, hash-free filename (e.g. `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT.md`) that `README.md`/`docs/README.md`/`CLAUDE.md` all link to permanently. Keep the revision-suffixed filename (`..._<hash>.md`) as what gets *archived* at each cycle - i.e. flip the naming convention so the stable name is the live one and the hash-suffixed name only ever exists in `docs/archive/`. This is the first cycle recommending this rather than the same "fix the two links" instruction, specifically because the same fix has now been prescribed and executed three times running with the staleness returning every time - the cost of *not* fixing the pattern (repeating the same ~10-minute task every cycle, forever) has exceeded the one-time cost of restructuring it.
4. **Add a positive `BaseURL` test matrix**: `https://example.com`, `https://example.com/stats`, `http://localhost:8080`, `http://127.0.0.1:8080`, `http://[::1]:8080`, `https://sub.example.com:8443/base/path` - assert successful construction and, where relevant, that `buildURL` preserves the base path/port/IPv6 host correctly. Closes finding #4.

### Not urgent, explicitly not a backlog item to keep re-budgeting for

- Everything `9eb3a9a`/`180a3db` already marked not-urgent (live-verifying the 136 unreachable endpoints, HTTP-server independent versioning policy, ecosystem-maturity commentary) remains not-urgent for the same reasons already given in those assessments - not re-litigated here.
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

*Assessment of record for revision `1b428f6` (one commit past tag `v3.1.2`), 2026-07-22. Supersedes `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_9eb3a9a.md` (revision `9eb3a9a`, tag `v3.1.1`, grade A-) as the current maintainability assessment. That file moves to `docs/archive/` in the same changeset as this file.*
