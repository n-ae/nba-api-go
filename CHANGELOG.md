# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- **`.github/workflows/release-install-smoke.yml`'s manual-dispatch tag resolution silently defaulted to the branch it ran on, not the "latest tag" its own input description promised.** The fallback chain (explicit `tag` input -> `github.ref_name` -> `git describe --tags --abbrev=0`) never reached the `git describe` fallback on a `workflow_dispatch` with no `tag` input, since `github.ref_name` is the triggering branch (e.g. `main`) on that event type and is therefore never empty - so an omitted input silently tested `main` instead of the latest release. Now keyed off `github.event_name`: a `workflow_dispatch` with no `tag` input falls back directly to `git describe`, skipping `ref_name` entirely; a real tag push still uses `ref_name`, which is correctly the tag name on that event. Found by the `2026-07-23` (`04537f4`) maintainability assessment (finding #15) - latent so far, since every real dispatch to date has passed an explicit tag.
- **The `git describe` fallback above was itself broken**, found immediately by actually dispatching it: `actions/checkout@v7`'s default shallow clone (`fetch-depth: 1`) fetches no tag history, so `git describe --tags --abbrev=0` failed with `fatal: No names found, cannot describe anything.` on the first real `workflow_dispatch`-with-no-input run - this exact path had never been exercised before (the bug above made it unreachable), so this was latent too. Fixed by checking out with `fetch-depth: 0`. Verified live, not just locally simulated: a real `workflow_dispatch` with no `tag` input against this fix, run twice - the first run reproduced the `git describe` failure above exactly, the second (after the `fetch-depth: 0` fix) correctly resolved `Verifying tag: v3.1.10`, fetched it, and ran the smoke program successfully.

## [3.1.10] - 2026-07-23

**Patch, CI only - no runtime or test source changes**: fixes the corpus-upload condition `v3.1.9` shipped broken, so the scheduled fuzz workflow's failure-artifact mechanism actually works on a genuine invariant violation. No caller-visible behavior changed in this release. Closes the one finding from the `2026-07-23` (`e3ee47c`) maintainability assessment.

### Fixed
- **`.github/workflows/fuzz.yml`'s corpus-upload step could never actually run.** `v3.1.9` scoped the `if:` condition to `steps.fuzz.outcome == 'failure'`, but GitHub Actions implicitly ANDs a bare `if:` with `success()` unless the expression contains an explicit status-check function - so the effective condition was `success() && steps.fuzz.outcome == 'failure'`, which is never true, since the fuzz step failing already puts the job in a failed state by the time the upload step evaluates. The one real CI run cited when `v3.1.9` shipped only exercised the success path (upload skipped, as expected), so this went unnoticed. Now `if: failure() && steps.fuzz.outcome == 'failure'`. Found by the `2026-07-22` (`e3ee47c`) maintainability assessment. Verified live on both paths, not just YAML-validated: a throwaway branch with the fuzz step swapped for one that writes a sentinel corpus file and exits 1 confirmed the upload step now runs and the resulting artifact contains exactly that sentinel file; a second real run on `main` confirmed the success path still correctly skips the upload.

## [3.1.9] - 2026-07-22

**Patch, CI only - no runtime or test source changes**: hardens the scheduled fuzz workflow's failure handling so a real invariant violation can't be confused with infrastructure failure. No caller-visible behavior changed in this release. Closes both findings from the `2026-07-22` (`0e35c33`) maintainability assessment.

### Changed
- **`.github/workflows/fuzz.yml` hardening**: the artifact-upload step is now scoped to the fuzz step's own outcome rather than the job-level `failure()`, which also fired (and could produce a misleading empty artifact) if `checkout`/`setup-go` failed before fuzzing started. The job comment now distinguishes an actual invariant violation (fuzz step failed, corpus artifact present) from an infrastructure failure (no artifact), rather than stating "a red run here means... a real finding" unconditionally. Verified on a real GitHub Actions run after the change. Found by the `2026-07-22` (`0e35c33`) maintainability assessment.

## [3.1.8] - 2026-07-22

**Patch, test/CI only - no runtime source changes**: fixes a false-negative gap in the fuzz test protecting `NewClient`'s `BaseURL` error handling, and adds scheduled fuzzing so this class of gap is caught automatically going forward. No caller-visible behavior changed in this release. Closes the one finding from the `2026-07-22` (`8e85a9c`) maintainability assessment.

### Fixed
- **`FuzzNewClientErrorDoesNotEchoInput`'s scheme-position template now asserts against the value actually inserted into the URL**, not the original fuzzed marker. The scheme template normalizes the marker (`validSchemeMarker`) before inserting it, since URI scheme syntax excludes most characters a fuzzed marker can contain, but the assertion still checked the unnormalized marker - so it only caught a regression by coincidence, for markers that happened to already be valid scheme syntax, and missed realistic secret shapes (e.g. `sk_live_...`, which strips to `skliveabcdef...` under normalization). Found by the `2026-07-22` (`8e85a9c`) maintainability assessment; confirmed to be a test-tooling gap only, not a runtime vulnerability - `TestNewClientErrorMessagesAreFixed` and a direct regression case both independently caught the same reintroduced regression correctly.

### Added
- **A scheduled `Fuzz` CI workflow** (`.github/workflows/fuzz.yml`) runs `FuzzNewClientErrorDoesNotEchoInput` with coverage-guided mutation for 60s daily, uploading any failing input as a build artifact. Ordinary `go test ./...` only replays the fuzz target's seed corpus, not ongoing mutation; this was recommended across three consecutive maintainability-assessment cycles before being added now rather than carried forward a fourth time.

## [3.1.7] - 2026-07-22

**Patch, security-relevant**: closes the last known instance of `BaseURL` values being echoed into `NewClient`'s errors - no caller passing a valid `BaseURL` sees any behavior change; a caller who was passing an unsupported, token-shaped scheme no longer has it echoed into `NewClient`'s error. Also adds a regression test that inventories every rejection message `NewClient` can return, so this class of gap can't reappear silently in the future. Closes the one finding from the `2026-07-22` (`eb62a41`) maintainability assessment.

### Fixed
- **`client.NewClient`'s unsupported-scheme error no longer echoes the rejected scheme.** Unchanged since `v3.1.2` and outside the `url.Parse`-failure branch entirely, this check echoed `baseURL.Scheme` on the assumption a scheme isn't secret-bearing - an assumption never checked against URI scheme grammar (RFC 3986: a letter, then letters, digits, `+`, `-`, `.`), which is permissive enough to hold a token- or secret-shaped string (e.g. `sklive123://host`). Found by the `2026-07-22` (`eb62a41`) maintainability assessment.

### Added
- `TestNewClientErrorMessagesAreFixed` inventories every distinct rejection `NewClient` can return and asserts each is a fixed string, byte-for-byte, regardless of what made the `BaseURL` invalid - a stronger check than "the marker doesn't appear," since it also catches a value unrelated to whatever marker a test happens to choose. `TestNewClientRejectionErrorsDoNotLeakBaseURL` gained a secret-shaped-scheme case, and `FuzzNewClientErrorDoesNotEchoInput` gained a scheme-position template (via a new `validSchemeMarker` normalization helper, since an arbitrary marker isn't valid scheme syntax unmodified). Confirmed all new cases fail against the pre-fix code and pass against the fix; fuzzed 2M+ executions with the expanded template set with no false negatives.

## [3.1.6] - 2026-07-22

**Patch, security-relevant**: closes the `BaseURL` secret-disclosure defect class structurally, ending three consecutive cycles of partial fixes for the same underlying issue - no caller passing a valid `BaseURL` sees any behavior change; a caller who was passing a malformed, credential- or token-bearing `BaseURL` (regardless of the specific way it was malformed) no longer has any part of it echoed into `NewClient`'s error. Closes the one finding from the `2026-07-22` (`f4801ef`) maintainability assessment.

### Fixed
- **`client.NewClient`'s `url.Parse` failure path now returns a fixed, input-free message, closing the `BaseURL` secret-disclosure defect class for good.** `v3.1.5`'s fix unwrapped `*url.Error` to what its own comment called an "input-free" reason - that claim was false. `net/url` builds several of its own error reasons (an invalid port, a malformed IPv6 host) directly from the input, so a credential- or token-bearing `BaseURL` like `https://example.com:sk_live_123/path` still disclosed the secret via the unwrapped reason. This is the third consecutive cycle this defect class recurred (`v3.1.3`'s explicit checks, `v3.1.4`'s wrapped outer error, `v3.1.5`'s unwrapped inner reason); fixed this time by not rendering any parser-derived text at all - any `url.Parse` failure now returns `invalid base URL: malformed`, regardless of why parsing failed, closing every current and future variant of this defect in `net/url`'s error construction, not just the ones found so far. Found by the `2026-07-22` (`f4801ef`) maintainability assessment.

### Added
- `TestNewClientRejectionErrorsDoNotLeakBaseURL` gained two cases for the exact inputs this cycle found (invalid port, malformed IPv6 host), and `FuzzNewClientErrorDoesNotEchoInput` gained port- and host-position templates. Confirmed both new test cases fail against the pre-fix code and pass against the fix; fuzzed 2M+ executions with the expanded template set with no false negatives.

## [3.1.5] - 2026-07-22

**Patch, security-relevant**: closes the remainder of `v3.1.4`'s `BaseURL` secret-disclosure fix - no caller passing a valid, non-secret-bearing `BaseURL` sees any behavior change; a caller who was passing a credential- or token-bearing `BaseURL` that failed to parse no longer has that secret echoed into `NewClient`'s error. Also closes a low-severity validation gap (`https://:443` no longer constructs successfully). Closes both findings from the `2026-07-22` (`0e400d1`) maintainability assessment.

### Fixed
- **`client.NewClient`'s `url.Parse` failure path no longer leaks the raw `BaseURL`.** `v3.1.4` fixed five of `NewClient`'s six `BaseURL`-rejection error paths; the sixth - the initial `url.Parse` failure, wrapped with `%w` - still embedded the complete input via `*url.Error`'s own `Error()` method, so a malformed, credential- or token-bearing `BaseURL` (e.g. one with an invalid percent-escape) still disclosed the secret. Now unwraps to just the parse failure reason, which never contains the input. Found by the `2026-07-22` (`0e400d1`) maintainability assessment, the uncovered remainder of `v3.1.4`'s fix.
- **`client.NewClient` now rejects a `BaseURL` with a port but no hostname** (e.g. `https://:443`). The prior check (`baseURL.Host == ""`) missed this because `Host` includes an optional port; now checks `baseURL.Hostname() == ""`, which correctly reports empty. Found by the same assessment.

### Added
- `FuzzNewClientErrorDoesNotEchoInput` - an invariant-based fuzz test asserting that whenever `NewClient` rejects a `BaseURL`, no injected secret marker appears in the returned error, regardless of which internal path rejected it. Added specifically because the prior fix's own regression test enumerated "all five" rejection paths and missed a sixth; an invariant test doesn't depend on a human having first enumerated every way `NewClient` can fail. Run 2M+ times locally with no false negatives; markers are restricted to those containing a digit to avoid a coincidental-English-word false-positive class discovered while writing this test (a fuzzed marker like `"esca"` matching the stdlib's "invalid URL escape" text is not a real leak).

## [3.1.4] - 2026-07-22

**Patch, security-relevant**: fixes a real secret-disclosure defect in `client.NewClient`'s `BaseURL` rejection errors - no caller passing a valid, non-secret-bearing `BaseURL` sees any behavior change; a caller who was passing a credential- or token-bearing `BaseURL` and logging `NewClient`'s error no longer has that secret echoed into the log. Closes the one finding from the `2026-07-22` (`b3c605d`) maintainability assessment.

### Fixed
- **`client.NewClient`'s `BaseURL` validation errors no longer echo the raw, unredacted `config.BaseURL`.** Every rejection path (invalid scheme, missing host, userinfo, query string, fragment) previously interpolated the complete input string via `%q`, so a caller passing `https://admin:secret@host` or `https://host?token=secret` - exactly the shapes the userinfo/query checks exist to reject - got back an error containing the literal credential or token. Found by the `2026-07-22` (`b3c605d`) maintainability assessment, itself following a lead from an external review of `v3.1.3`; independently verified to predate `v3.1.3` (the scheme/host checks have had the same defect since `v3.1.2`). `TestNewClientRejectionErrorsDoNotLeakBaseURL` covers all five rejection paths, asserting the returned error never contains an injected secret.

## [3.1.3] - 2026-07-22

**Patch, not minor**: the `BaseURL` validation change below only tightens rejection of values that were already unusable (userinfo, a query string, and a fragment are all syntactically legal but never sensible in a `BaseURL`) - no caller passing a real, plain `http`/`https` URL sees any behavior change. Everything else is test coverage and documentation. Closes all three remaining findings from the `2026-07-22` (`1b428f6`) maintainability assessment.

### Added
- `TestNewClientAcceptsValidBaseURL` - the positive-case counterpart to the existing `BaseURL`-rejection tests, covering a bare host, a host with a path, a port, an IPv4 host, an IPv6 host, and a subdomain with both a port and a path. Every prior `BaseURL` test only asserted rejection, so a validation change that became stricter in some new, unintended way could ship without any test noticing a previously-valid configuration broke. Also asserts `buildURL` preserves each base's host/port/path correctly when joining an endpoint. Found by the `2026-07-22` (`1b428f6`) maintainability assessment. No behavior change.

### Changed
- `README.md` and `docs/README.md` both still linked to `MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_180a3db.md` - the assessment of record moved to `..._9eb3a9a.md` when `v3.1.1` shipped, but (as that same `9eb3a9a` assessment predicted about itself) only `CLAUDE.md`'s pointer was kept current at the time; these two were never updated. Both now point at `..._9eb3a9a.md`.
- **The current maintainability assessment now lives at a stable, hash-free path: `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT.md`, permanently - not a new `..._<date>_<revision>.md` file every cycle.** Closes finding #5 (Next #3) from the `2026-07-22` (`1b428f6`) maintainability assessment: the "assessment link goes stale the same day a new assessment supersedes the one it points at" pattern (the entry directly above is itself the third instance) had recurred three cycles running, each time predicted by the outgoing assessment about its own successor. Every external pointer (`CLAUDE.md`, `README.md`, `docs/README.md`, `tests/contract/README.md`) now links to that one stable path and never needs updating again. The hash-suffixed naming convention (`..._<date>_<revision>.md`) still exists, but only in `docs/archive/` - each new cycle now archives the *outgoing* content under that name and overwrites the stable path with the new cycle's content, rather than creating a new hash-named file and repointing every link at it. `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-22_1b428f6.md` is renamed to `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT.md` under this convention (its content and revision-`1b428f6` findings are unchanged); the archived `..._9eb3a9a.md` (and every earlier archived assessment) is untouched except for updating `9eb3a9a`'s own forward-pointing supersession banner to the new stable path.
- **`client.NewClient` now also rejects a `BaseURL` containing userinfo, a query string, or a fragment** - all three are syntactically legal in an absolute `http`/`https` URL (so none were caught by `v3.1.2`'s scheme/host check), but none are a sensible way to configure a `BaseURL`: userinfo (`https://user:pass@host`) risks credentials leaking into logs/errors/metrics wherever `BaseURL` gets printed; a query string is ambiguous with `buildURL`'s own per-request query construction; a fragment is a client-side-only URL component never sent in an HTTP request at all. Found by the `2026-07-22` (`1b428f6`) maintainability assessment, itself following a lead from an external review of `v3.1.2`. Same patch-level reasoning as `v3.1.2`'s original scheme/host check: every newly-rejected value was already nonsensical as a `BaseURL`, so no caller passing a real, plain `https://host/path` sees any behavior change.

## [3.1.2] - 2026-07-22

**Patch, not minor**: the `BaseURL` validation change below only tightens rejection of values that were already unusable (they'd fail on the first `Get` regardless) - no caller passing a real `http`/`https` URL sees any behavior change. Everything else is CI configuration, `tools/generator`'s own test coverage, and documentation. Closes both new findings from the `2026-07-22` (`9eb3a9a`) maintainability assessment, which independently confirmed the same two claims in an unsolicited external review of `v3.1.1`.

### Added
- `.github/workflows/apidiff.yml` gains a `workflow_dispatch` trigger, matching `release-install-smoke.yml`. This job compares `main` against the *latest tag*, so a PR that merges an accepted, documented exception (e.g. `[3.1.1]`'s `DefaultUserAgent` const-value bump) shows red until a new tag catches up - previously that red status stayed on `main`'s check history with no way to re-verify it had resolved short of waiting for the next real push. Confirmed locally: comparing `main` against `v3.1.1` (which already includes that change) now produces zero diff.
- **Every generated endpoint SDK test now asserts the outbound request path**, and `tools/generator` gained a new metadata-wide test, `TestEndpointPathMatchesNameConvention` - two related but structurally different findings from the `2026-07-22` (`9eb3a9a`) maintainability assessment, which independently verified an unsolicited external review's claim that "no test asserts an SDK endpoint's outbound URL path."
  - `tools/generator/templates/endpoint_test.tmpl` now captures `r.URL.Path` on the stub server and asserts it equals `/<endpoint>` from the same metadata that generated the call site. This closes the literal gap (previously nothing checked the path a generated endpoint function actually requests) and guards against a template bug or a hand-edit drifting the generated `.go` file's endpoint string away from its own metadata. **It cannot catch a typo already present in the metadata itself** - both sides of the assertion derive from the identical `.Endpoint` field, so a corrupted metadata value produces a self-consistently "passing" test (verified directly: temporarily corrupting one endpoint's metadata and regenerating still passed this assertion). Regenerated all 135 metadata-covered `generated_*_test.go` files via `cd tools/generator && go run . -metadata metadata/<file>.json` for every committed metadata file; `gamerotation.go`/`leaguedashplayerstats.go` (permanently excluded from regeneration, see `[2.0.0]`) were correctly reverted after the batch run picked them up too, and their existing `generated_*_test.go` files remain accurate since their metadata's `Endpoint` field didn't change.
  - `tools/generator/generator_test.go`'s new `TestEndpointPathMatchesNameConvention` is the automated form of the manual self-consistency check that originally found the 10 malformed endpoint paths in `[3.1.0]`: it reads `metadata/*.json` directly (independent of any generated output) and asserts every entry's `endpoint` field equals its `name` field lowercased, with one documented exception (`TeamYearOverYearSplits`, already explained in that endpoint's own doc comment). **This is the check that can actually catch a metadata typo** - verified directly by temporarily corrupting `AssistLeaders`'s endpoint string in metadata and confirming this test (and only this test, not the per-endpoint generated one above) failed with a clear message.

### Changed
- **`client.NewClient`'s `BaseURL` validation now checks the URL is absolute with an `http`/`https` scheme and a non-empty host**, not just that it parses without error. Found by the `2026-07-22` (`9eb3a9a`) maintainability assessment (confirming the same claim in an unsolicited external review): `url.Parse` alone accepts `""`, a bare hostname with no scheme (`"example.com"`), a protocol-relative reference (`"//example.com"`), and other unusable values without error, so a mistyped `BaseURL` previously constructed a `Client` successfully and only failed confusingly on the first `Get`. `TestNewClientRejectsUnusableBaseURL` covers the previously-uncaught cases.

## [3.1.1] - 2026-07-22

**Patch, not minor**: unlike `[3.1.0]`'s HTTP-API behavior changes or `[2.2.0]`'s `Config.Timeout` fix, nothing in this release changes observable runtime behavior for any consumer - `DefaultUserAgent`'s value bump is explicitly documented as never applied automatically to any request (see its `### Changed` entry below), and everything else is documentation and `tools/generator`'s own internal test coverage. Closes the remaining five findings from the `2026-07-22` (`180a3db`) maintainability assessment's Immediate/Next buckets - the assessment that reviewed the `[3.1.0]` release itself.

### Added
- **Direct unit tests for `generateHandler`/`GenerateDispatchTable`/`processHandlerMetadata`** in `tools/generator`'s own test suite - the "generator's own test suite has no direct coverage of handler/dispatch generation" finding (#15) from the `2026-07-22` (`180a3db`) maintainability assessment. Previously these were only exercised indirectly (a syntactic-validity check across every committed metadata file, plus the root module's cross-package tests on already-committed generated output), so a semantic bug in `handler.tmpl`/`dispatch.tmpl` would have surfaced only on the next `-all-handlers` regeneration, not before. New tests: `TestProcessHandlerMetadata`/`TestProcessHandlerMetadataExplicitOverrides` (per-parameter `HandlerGoType`/`EffectivePointer` resolution, the `SDKFunction`/`ResponseWrapped`/`Pointer` override paths, and the deep-copy invariant documented on `processHandlerMetadata` itself), `TestGenerateHandler` (a handler renders, parses as valid Go, and contains the expected required-parameter validation), and `TestGenerateDispatchTable`/`TestGenerateDispatchTableRequiresMetadataFiles` (the documented first-file-wins dedup behavior for a `Name` appearing in more than one metadata file, using a fixture metadata directory rather than trusting the real 141-endpoint one to exercise this path by accident).

### Changed
- **`client.DefaultUserAgent` bumped from `"nba-api-go/2"` to `"nba-api-go/3"`**, matching its own documented "major-version-only" convention (the same bump `[2.2.0]` did for `1.0`→`2`) and the `v3.0.0` major bump. **Note for anyone consulting `apidiff` CI history**: this shows as an "incompatible" `apidiff` finding (a const's *value*, not its type, changed) - `DefaultUserAgent`'s own doc comment already establishes that its value is expected to shift on major bumps without being source-breaking (it is not applied automatically to any request; see the comment), so this is a documented, accepted exception, not a real breaking change requiring a further version bump.
- **`docs/MAINTENANCE.md` and `tools/generator/README.md` updated to reflect the server-handler and test generation added in `[3.1.0]`** - the two smaller doc-currency findings from the `2026-07-22` (`180a3db`) maintainability assessment. `docs/MAINTENANCE.md` referenced hand-written `cmd/nba-api-server/handlers_*.go` files (deleted this cycle, replaced by generated equivalents) and an `internal/middleware/ratelimit.go` path that never matches the real, public `pkg/client/middleware/ratelimit.go`. `tools/generator/README.md` - the generator's own primary doc - hadn't been updated for the single biggest capability it gained this cycle: its Options/Architecture/Roadmap sections still showed no `-server-output`/`-all-handlers` flags, no `handler.tmpl`/`dispatch.tmpl`/`endpoint_test.tmpl`, and a stale "v0.1" roadmap listing "test skeleton generation" and "batch generation for all 139 endpoints" as unchecked future work that is, in fact, done. Its "Notes" section also said generated code is "a starting point, not production-ready" and to "always review and customize generated code" - directly contradicting the project's actual, established convention (documented elsewhere) that generated files must not be hand-edited because a later regeneration silently discards hand edits.
- **`README.md`'s opening "🏆 100% Coverage Achievement" section now discloses the live-reachability finding.** Previously claimed "100% endpoint coverage"/"World's first complete NBA API implementation in Go!" with no mention that a full reachability sweep found only 5 of 141 endpoints respond at all from any network tested, the other 136 hanging to a hard timeout - a real self-representation gap the `2026-07-22` (`180a3db`) maintainability assessment flagged as its top finding: `CLAUDE.md` and `tests/integration/README.md` already disclosed this precisely, but the most-read document in the repo didn't. Added a callout clarifying "100% coverage" means implemented/type-safe/test-covered, not currently reachable against live traffic, linking to the full sweep results.
- `README.md` and `docs/README.md` both linked to `MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-20_8549390.md` - an assessment archived three cycles ago (superseded by `384e5de`, then `a58d3fe`, then `1592e7e`, then `180a3db`) - despite `CLAUDE.md`'s own assessment pointer being kept current every cycle. Both now point at the current assessment of record. Found by the same `180a3db` assessment.

## [3.1.0] - 2026-07-22

**Minor, not patch**: no Go `pkg/` API changed (the stability promise this project's Versioning section makes), but this release changes real, observable HTTP-API behavior for `cmd/nba-api-server` consumers - see the two `**Breaking (HTTP API)**` entries below - so it's treated as a real behavior change, not a bug-for-bug patch, matching how `[2.2.0]`'s `Config.Timeout` fix was bumped minor for the same reason. Three pieces: a correctness fix (10 endpoints were silently sending requests to nonexistent URL paths), a structural one (the server's 4,358 lines of hand-written HTTP handlers are now generated from the same metadata the SDK already generates from, closing the "decide the server's fate" item carried in every maintainable-architect-v4 assessment), and a coverage one (`pkg/stats/endpoints` goes from 5.2% to 75.1% via generated response-parsing tests - regression-safety-net coverage, not live-verification coverage; see the `### Added` entries below for what that distinction means). This also closes out the "tag lag" left by `v3.0.0`: these three pieces landed on `main` immediately after `v3.0.0` was tagged and sat unreleased for several commits before this tag.

### Fixed
- **10 of 141 endpoints were sending requests to URL paths that don't exist on `stats.nba.com`, silently 404ing (or equivalent) on every call.** Found via a systematic self-consistency check prompted by a live-endpoint audit: for every endpoint, does the URL path string it sends match its own file/type name? 10 didn't. Confirmed against the authoritative Python `nba_api` reference project's own endpoint definitions where a corresponding file exists there (`LeagueHustleStatsPlayer`, `LeagueHustleStatsTeam`, `PlayerCareerByCollegeRollup`, `PlayerDashboardByLastNGames`, `PlayerNextNGames` - all confirmed all-lowercase, no spaces); the remaining 5 (`LeagueHustleStatsTeamLeaders`, `PlayerTrackingRebounding`, `TeamDashboardByLastNGames`, `TeamNextNGames`, `CommonPlayerInfoV2`) have no `nba_api` precedent but were fixed by strong internal-consistency analogy - every other endpoint in this codebase, without exception, uses an all-lowercase URL path, and each of these 5 is either a direct sibling of an already-confirmed case (e.g. `LeagueHustleStatsTeamLeaders` next to the now-confirmed `LeagueHustleStatsTeam`/`Player`) or one of 10 otherwise-identical `v2`-suffixed endpoints where the other 9 all use lowercase `v2`.
  - Two had a literal embedded space in the URL path (`"leaguehustlestatsp layer"`, `"leaguehustlestats team"`) - Go's URL encoding turns this into `%20`, guaranteeing a 404.
  - One had a straightforward typo (`"playertrackingebounding"`, missing the first `r` of "rebounding").
  - One had a different straightforward typo (`"playercareerbyrollegerollup"` - "college" corrupted to "rollege").
  - Six had a stray capital letter in an otherwise-lowercase path (`commonplayerinfoV2`, `leaguehustlestatsTeamleaders`, `playerdashboardbylastnGames`, `playernextnGames`, `teamdashboardbylastnGames`, `teamnextnGames`).
  - Fixed at the source (`tools/generator/metadata/tier2_batch.json`, `tier6_batch.json`, `tier9_batch.json`, `tier10_batch.json`) and regenerated via `tools/generator`, not hand-patched in the generated `.go` files - `git diff` after regenerating + `gofmt` confirms each of the 10 changed files differs from its previous committed version by exactly the corrected endpoint string and its cross-referenced doc comment, nothing else.
  - One additional filename/slug mismatch found by the same check (`teamyearoveryearsplits.go` sending `"teamdashboardbyyearoveryearsplits"`) is **not** a bug - its own doc comment states this is deliberate (a shorter Go-friendly type name for a real, differently-named NBA.com endpoint), left unchanged.
  - Not caught by NBA.com header validation, `TestAllMetadataFieldsHaveExplicitTypes`, or any existing test, because none of them inspect the URL path string itself - only `go vet`/`gofmt`-level correctness and field-name/type coverage are checked today. These 10 endpoints would have failed on first live use for any consumer, silently, since whenever each was added.

### Added
- `tools/generator` now generates `cmd/nba-api-server`'s HTTP handlers, not just `pkg/stats/endpoints`'s SDK code - resolves the "decide the server's fate" item carried in every maintainable-architect-v4 assessment. Returns to ADR 002's original, never-executed "Phase 2: Generate handlers for all endpoints" plan rather than trimming the server's scope. `tools/generator/metadata/handwritten_handlers.json` adds minimal, handler-only metadata (a new `"handler_only": true` field, checked by `generateEndpoint` to refuse ever overwriting hand-written SDK code) for the 6 endpoints whose SDK code is intentionally hand-written, so all 142 HTTP routes are now generated, not 136. `templates/handler.tmpl` (per-endpoint handler) and `templates/dispatch.tmpl` (the route table) are new; `-endpoint`/`-metadata` now generate a handler alongside SDK code automatically, and a new `-all-handlers` flag regenerates every handler plus the dispatch table in one pass. Every one of the 142 generated handlers, plus the dispatch table's routing, was verified by direct dry-run inspection and a live round-trip against `stats.nba.com` (via `leagueleaders`, one of the few endpoints reachable from this environment - see the `[Unreleased]` reachability-sweep entry from earlier this session) before being committed.
- `TestGeneratedHandlers` (`cmd/nba-api-server/generated_handlers_test.go`) - a data-driven test covering all 142 generated handlers: for each, asserts a missing required parameter 400s and a request with every required parameter present succeeds against a stub upstream (`httptest.Server` returning a minimal `{"resultSets": []}` envelope - tolerated by every endpoint's parsing code, including the two with genuinely different raw shapes, `LeagueLeaders`'s singular `resultSet` and `InternationalBroadcasterSchedule`'s bare, unwrapped return type, both confirmed by reading their parsing code before relying on it). Reads `tools/generator/metadata/*.json` directly (the same source of truth generation itself uses) rather than a hand-maintained parallel list, so a newly added endpoint's metadata is automatically covered. Closes the coverage gap generating the handlers opened: `cmd/nba-api-server` goes from 7.2% (real handler-code coverage was ~0% even before generation; the reported "10.0%" pre-generation was entirely infrastructure code) to **76.8%**. Like the `pkg/stats/endpoints` number below, this is regression-safety-net coverage: the stub upstream and the assertions both derive from the same metadata the handler itself was generated from, so it catches a broken template or a wrong dispatch entry immediately, but it cannot catch NBA.com's actual response shape drifting from that metadata.
- **`tools/generator` now generates a response-parsing test alongside every metadata-covered endpoint** - the "endpoint package coverage stuck around 5%" finding carried in every maintainable-architect-v4 assessment since `2026-07-19`. `templates/endpoint_test.tmpl` (new) synthesizes, per endpoint, a fixture `resultSets` JSON body using the endpoint's own real result-set names and column headers from metadata (not hand-picked example data), with a placeholder value typed correctly per field (`"test"` for `string`, `1`/`1.5` for `int`/`float64`) so the generated test exercises the real `findResultSet`/`validateHeaders`/`toInt`/`toFloat`/`toString` parsing path - not just request construction, which `TestGeneratedHandlers` above already covers cross-package. `generateEndpointTest` (`tools/generator/generator.go`) writes `pkg/stats/endpoints/generated_<name>_test.go`, wired into the same `GenerateFromMetadata`/`GenerateSingleEndpoint` paths that already generate the endpoint's SDK code, so a newly added endpoint gets a parsing test for free - no separate step to remember. Skips (not an error) the 6 hand-written endpoints (`HandlerOnly` metadata, no SDK code for it to overwrite) and 2 non-endpoint helper files (`dates.go`, `types.go`) that were never metadata-covered - **135 of 143 endpoint files now have a generated test**, matching the metadata-covered count exactly (verified: `ls pkg/stats/endpoints/generated_*_test.go | wc -l` = 135, and the 8 `.go` files without a corresponding generated test are exactly those 6 + 2, no omissions). `pkg/stats/endpoints` coverage: **5.2% → 75.1%**. **This is regression-safety-net coverage, not live-verification coverage**: the fixture and the assertion both derive from the same metadata that generated the parsing code under test, so it exercises the real parsing path and catches a broken template or a generation-pipeline regression immediately - it cannot catch a field order that's drifted from what NBA.com currently returns, which is a structurally different risk (see the reachability-sweep entry above for the current state of live-verifying that).

### Changed
- **The 142 hand-written HTTP handlers in `cmd/nba-api-server` (`handlers.go`'s switch statement plus 6 `handlers_*.go` files, ~4,358 LOC) are deleted and replaced by generated equivalents.** `StatsHandler.ServeHTTP` now dispatches through a generated `map[string]func(*StatsHandler, http.ResponseWriter, *http.Request)` (`generated_dispatch.go`) instead of a hand-maintained switch statement; `endpoint_inventory_test.go`'s drift-detection meta-test now counts map entries instead of switch cases. The `playertrackingshotdashboard` legacy route alias (a hand-written duplicate of `handlePlayerTrackingShootingEfficiency`'s body under the endpoint's old name) is now a single map entry pointing at the same generated handler function, instead of a ~20-line duplicated function.
- **Breaking (HTTP API): 132 of 142 endpoints' JSON response shape changes.** The hand-written handlers were inconsistent - 132 double-wrapped the response (`{"data": {"Data": {...}, "StatusCode": 200, "URL": "..."}}`, passing the SDK's whole `models.Response[T]` wrapper to `writeSuccess`), while only 10 correctly unwrapped to `{"data": {...}}`. Generation standardizes all 142 on the correct, already-established shape. No GitHub issues suggest any consumer depends on the previous inconsistent shape.
- **Breaking (HTTP API), narrower: a handful of endpoints now return `400 missing_parameter` for an omitted parameter metadata already marked `required` that was previously silently defaulted instead** (e.g. `CommonAllPlayers`'s `Season` - confirmed via a running server: previously silently defaulted, now correctly 400s). Affects every endpoint whose metadata marks `Season` and/or `SeasonType` as `required: true`: `ShotChartDetail`, `PlayerDashboardByGeneralSplits`, `LeagueGameFinder`, `TeamGameLogs`, `TeamDashboardByGeneralSplits`, `CommonAllPlayers`, `CommonPlayoffSeries`, plus `InternationalBroadcasterSchedule` (already required Season; unchanged).
- **Breaking (HTTP API), narrowest: `InternationalBroadcasterSchedule`'s `LeagueID` query parameter is no longer respected** - it was the one endpoint (of 142) that let a caller override `LeagueID`, defaulting to `"00"` only if omitted; every other endpoint has always hardcoded it server-side. Standardized to match all 141 others.
- `PlayerGameLog`/`TeamGameLog` now accept optional `DateFrom`/`DateTo` query parameters (already supported by the underlying SDK `Request` struct, never previously exposed over HTTP); `LeagueLeaders` now accepts optional `StatCategory`/`ActiveFlag` similarly. Additive, not breaking.

## [3.0.0] - 2026-07-22

**Major, breaking.** `NewClient`'s signature change (see `**Breaking:**` below) requires a major version bump; per Go's semantic import versioning this also requires the `/v3` module-path suffix - see the Migration guide below. Everything else in this section is non-breaking and would otherwise have shipped as a minor/patch release.

### Added
- `.github/workflows/apidiff.yml` - a CI job that fails if the module's public API changed incompatibly since the latest tagged release (`golang.org/x/exp/cmd/apidiff`, module-wide comparison). Addresses the "no apidiff/semver-break gate in CI" finding carried in every maintainable-architect-v4 assessment since `2026-07-19`. A red result isn't automatically wrong - this project has shipped deliberate breaking changes before - but it now requires a conscious major-version decision instead of shipping unnoticed. Verified locally against this exact branch's `NewClient` change below before being added (correctly flags exactly the 3 incompatible changes and nothing else). Found by the `2026-07-22` maintainability assessment.
- `.github/workflows/release-install-smoke.yml` - a tag-triggered CI job that `go get`s the just-tagged module into a scratch module (outside this checkout) and builds/runs a small program against it, verifying the module is actually fetchable and usable by an external consumer, not just that this repo's own checkout builds. Addresses the "no tag-triggered CI / no external install smoke test" finding carried in every maintainable-architect-v4 assessment - `v2.0.0` and `v2.1.0` both shipped genuinely unfetchable via `go get` (a module-path mismatch caught by hand, days later; see `[2.1.1]` below), which this would have caught automatically on the tag push itself. Rehearsed locally against the real module proxy (`go get github.com/n-ae/nba-api-go/v2@v2.2.0` into a scratch module, `go mod tidy`, build, run) before being added; its hardcoded `/v2` references were updated to `/v3` in the same change that bumped the module path (below), so it verifies the right thing once `v3.0.0` is actually tagged.
- `TestClientNegativeTimeoutDisablesContextDeadline`/`TestClientNegativeTimeoutSDKBuiltClient` (`pkg/client/client_test.go`) - cover the negative-`Config.Timeout` behavior documented below (disables enforcement entirely, on both the context-deadline and SDK-built-`http.Client.Timeout` paths) with an actual test, closing the "documented but untested" gap the `2026-07-22` assessment flagged.

### Changed
- **Breaking:** `client.NewClient`, `stats.NewClient`, and `live.NewClient` now return `(*Client, error)` instead of `*Client`. Previously an invalid `Config.BaseURL` (a malformed URL) was silently accepted at construction and only surfaced as an error on the first call to `Get` - the "invalid base URL only surfaces on first request" gap carried in every assessment since `a58d3fe`. Construction now fails loudly instead. `stats.NewDefaultClient`/`live.NewDefaultClient` are **unaffected** - their signatures are unchanged (`*Client`, no error) because they always construct against the package's own compile-time-valid `StatsBaseURL`/`LiveBaseURL` constant, which can't fail; they panic internally (an unreachable path, matching the existing `http.DefaultTransport` type-assertion pattern in `pkg/client`) if that invariant is ever violated. Only callers who use `NewClient` directly with an explicit `Config.BaseURL` are affected - see Migration guide below.
- **Breaking, same commit:** `client.Client`'s unexported `buildURL` method (not part of the public API) simplified from `(string, error)` to `string`, since the only error it could ever return - an invalid base URL - is now caught at `NewClient` time instead.
- **Breaking:** `go.mod`'s `module` line is now `github.com/n-ae/nba-api-go/v3`, and every internal import across the repo (185 `.go` files) is updated to match - required by Go's semantic import versioning for the major version bump above. See Migration guide below for the exact consumer-facing change and the `v2.0.0`/`v2.1.0` incident this is deliberately avoiding repeating.
- `tools/generator/templates/endpoint.tmpl` imported `github.com/n-ae/nba-api-go/pkg/...` with no version suffix at all - not `/v2`, despite `go.mod` having required it since `v2.0.0`. A bug independent of this release, discovered while updating this exact line to `/v3`: reproduced live via `go run . -endpoint LeagueGameFinder -dry-run`, which emitted that exact non-compiling import; `TestGenerateFromMetadata_ProducesValidGo` didn't catch it because it only checks generated output parses as syntactically valid Go, not that it compiles. Fixed to `/v3`, verified via the same dry-run command.
- `.github/workflows/live-drift.yml`'s scope narrowed to `LeagueLeaders`/`InternationalBroadcasterSchedule` (commit `1592e7e`). The workflow as tagged in `v2.2.0` ran all six `TestSimpleSmokeTests` subtests; two independent manual runs on the tagged commit (`29865194310`, `29865360637`, ~2 minutes apart) showed `PlayerCareerStats`/`PlayerGameLog` (`stats.nba.com`) hanging to the 30s timeout and `Scoreboard` (`cdn.nba.com`) hitting an immediate Akamai block - all three reproducibly unreachable from GitHub Actions runner IPs, not flaky. Narrowed to the two endpoints confirmed reachable so the weekly signal stays meaningful instead of permanently red. **Decision recorded here rather than left implicit: this is a CI-configuration-only change with no effect on any published module consumer, so it is not being backported into a `v2.2.1` patch tag.** `v2.2.0` as tagged still ships the wider, permanently-red workflow scope; anyone relying on the scheduled check getting a real "no drift detected" signal needs `main`, not the tag. Documented in `tests/integration/README.md`'s "Known live-traffic blocks" section. Found by the `2026-07-22` maintainability assessment.
- `tests/integration/README.md`'s "Test Categories" section rewritten to describe `TestSimpleSmokeTests`'s six actual subtests instead of four `*_test.go` files (`player_test.go`, `team_test.go`, `league_test.go`, `live_test.go`) that don't exist in the directory - a stale-documentation gap flagged by the `2026-07-22` maintainability assessment and, independently, by an external third-party review of `v2.2.0`.
- `docs/MAINTENANCE.md`'s "Code Generation Approach" section corrected: `go run tools/generator/main.go` fails (`undefined: NewGenerator`) because `tools/generator` is a separate Go module - the working invocation is `cd tools/generator && go run . -metadata metadata/<file>.json`. Also corrected its example to copy `leaguegamefinder.json` rather than a nonexistent `playercareerstats.json` (`PlayerCareerStats` is hand-written with no metadata file), and folded in the `fieldtypes.json` step the old "Manual Approach" skipped entirely. Found by the `2026-07-22` maintainability assessment.
- `client.Config.Timeout`'s doc comment now states that a negative value disables timeout enforcement entirely (both the SDK-built client and the per-request context deadline added in `[2.2.0]`), unlike `0` which normalizes to `DefaultTimeout`. No behavior change. Found by the `2026-07-22` maintainability assessment.
- `tests/integration/README.md`'s "Known live-traffic blocks" section corrected: the live-verification backlog's blocking of `PlayerCareerStats`/`PlayerGameLog`/`CommonPlayerInfo`/`TeamGameLog` is **not** GitHub-Actions-runner-IP-specific as previously documented - retested 2026-07-22 via raw `curl` (bypassing the SDK) from a residential/business ISP IP and got the identical hard-timeout pattern on the first request for all four, while `LeagueLeaders` kept succeeding from the same IP. No code change; corrects a wrong assumption ("a developer machine ... may not hit the same block") that would have sent the next live-verification attempt down a dead end.
- `models.ErrTimeout`'s doc comment now states the dual timeout error taxonomy explicitly: it's returned only for a server-reported `408`/`504` HTTP status, while a client-side timeout (`Config.Timeout` elapsing, or a caller's own `ctx` deadline) surfaces from `client.Client.Get` as a wrapped `context.DeadlineExceeded` instead - a caller that wants both needs to check `errors.Is` against each separately. Cross-referenced from the error-wrapping site in `client.go`. No behavior change. Found by the `2026-07-22` maintainability assessment.
- `client.NewClient`'s SDK-built-client path now has a doc comment acknowledging that `Config.Timeout` is deliberately enforced twice (via `http.Client.Timeout` and, uniformly, via `Get`'s per-request context deadline) rather than leaving the redundancy unexplained. No behavior change. Found by the `2026-07-22` maintainability assessment.
- This file's own version-comparison links footer was stale by two releases - no `[2.1.2]` or `[2.2.0]` link existed, and `[Unreleased]` still pointed at `compare/v2.1.1...HEAD`. Fixed alongside adding the `[3.0.0]` link this release needs.

### Migration guide
- **Import path change, required for every consumer.** `go get github.com/n-ae/nba-api-go/v2` becomes `go get github.com/n-ae/nba-api-go/v3`, and every import updates to match, e.g.:
  ```go
  // Before
  import "github.com/n-ae/nba-api-go/v2/pkg/stats"

  // After
  import "github.com/n-ae/nba-api-go/v3/pkg/stats"
  ```
  This is a hard Go modules requirement for any major version bump past 1, not a style choice - see the `[2.1.1]` entry below for what happens when it's missed (`v2.0.0`/`v2.1.0` were unfetchable via `go get` until fixed).
- Any call to `client.NewClient(cfg)`, `stats.NewClient(cfg)`, or `live.NewClient(cfg)` needs a second return value:
  ```go
  // Before
  c := stats.NewClient(stats.Config{BaseURL: myURL})

  // After
  c, err := stats.NewClient(stats.Config{BaseURL: myURL})
  if err != nil {
      log.Fatal(err) // or handle however your application handles startup config errors
  }
  ```
- Calls to `stats.NewDefaultClient()`/`live.NewDefaultClient()` need **no import-statement content change beyond the path above** - their call sites are otherwise unaffected.

## [2.2.0] - 2026-07-21

**Minor, not patch:** the `Config.Timeout` fix below is a real behavior change for custom-`HTTPClient` callers (see its entry), not a pure bug-for-bug patch, so this bumps the minor version per the Versioning policy in `CLAUDE.md`.

### Fixed
- **`client.Config.Timeout` was silently ignored whenever a caller supplied their own `HTTPClient`.** `Client.timeout` was stored at construction but only ever used to set `http.Client.Timeout` on the *SDK-built* client (`HTTPClient == nil`); a caller who passed both a custom `HTTPClient` and a `Timeout` reasonably assumed the timeout applied, but it was dead config. Found by the `2026-07-21` maintainability assessment. `Get` now imposes `Timeout` as a per-request `context` deadline, so it applies uniformly regardless of which `HTTPClient` is used. A caller-provided `ctx` with an earlier deadline still wins (`context.WithTimeout` only tightens, never extends). Covered by `TestClientTimeoutAppliesWithCustomHTTPClient`. Non-breaking for the default-client path (behavior unchanged); custom-`HTTPClient` callers who previously had no SDK-imposed timeout now get the configured one (default `30s`) as a ceiling - raise `Config.Timeout` if a call legitimately needs longer.

### Changed
- `client.DefaultUserAgent` bumped from `"nba-api-go/1.0"` to `"nba-api-go/2"` - a v2 module reporting `1.0` was misleading. Major-version-only so it needn't change every patch. Note this constant is *not* applied automatically (the `pkg/stats`/`pkg/live` facades install a browser-style User-Agent via `middleware.WithUserAgent`); it's only a fallback for callers constructing `client.Client` directly.
- Added `.github/workflows/live-drift.yml` - a weekly (plus manual `workflow_dispatch`) scheduled workflow that runs the `INTEGRATION_TESTS=1` smoke suite against live `stats.nba.com`, to catch upstream schema/header drift out-of-band from push/PR CI (which deliberately doesn't hit the network). Addresses the "no scheduled live-drift workflow" gap carried in every recent assessment; the main CI workflow's header comment already anticipated it.

## [2.1.2] - 2026-07-21

**Docs-only release.** No source, generator, or generated-endpoint changes; `pkg/`, `cmd/`, and `tools/generator/` are byte-identical to `v2.1.1` aside from the version constant below.

### Added
- `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-21_a58d3fe.md` - the current assessment of record (grade B+, up from B). Independently reverifies two claims from `[2.1.1]`'s changelog against source rather than trusting the prose: `playercareerstats.go` does call `validateHeaders` for all 8 result sets now, and its test coverage did go from 0% to 76.7%-85.7%. Also reproduced live against the real Go module proxy that `v2.0.0`/`v2.1.0` were genuinely unfetchable (`go get` failed exactly as `[2.1.1]` describes) and that `v2.1.1` fixes it. Finds one new issue, not previously documented: `Config.Timeout` (`pkg/client`) is silently ignored whenever a caller supplies a custom `HTTPClient` - `Client.timeout` is set at construction but never read anywhere else in the package, so a caller who sets both reasonably but incorrectly assumes `Timeout` still applies.

### Changed
- `docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-20_384e5de.md` (the prior assessment of record, grade B) archived to `docs/archive/` with a supersession banner, per this project's own documented docs-consolidation convention.
- `CLAUDE.md`'s assessment-of-record pointer, grade reference, and "Next assessment" footer updated to point at the new file.
- `cmd/nba-api-server`'s `version` constant bumped from `"2.1.0"` to `"2.1.2"` - missed during the `v2.1.1` release despite `[2.1.0]`'s note that this constant is meant to "track the actual release version going forward."

## [2.1.1] - 2026-07-20

### Fixed
- **`v2.0.0`/`v2.1.0` were unfetchable via `go get` for any new or upgrading consumer.** Go modules requires a module's path to end in `/vN` once its major version reaches 2 or higher (`go.mod`'s own `module` line, and every import of it) - `v2.0.0` shipped this repo's first major-version bump without that suffix. The result: `go get github.com/n-ae/nba-api-go@v2.1.0` (or any v2.x tag) fails immediately with `invalid version: module contains a go.mod file, so module path must match major version ("github.com/n-ae/nba-api-go/v2")`, before any code is even fetched. Confirmed by reproducing it directly against a scratch module. Existing users who pinned a `v1.x` `go.mod` `require` line before this was noticed are unaffected; this blocked anyone trying to adopt or upgrade to `v2.0.0`/`v2.1.0` from the moment `v2.0.0` was tagged. Fixed: `go.mod`'s `module` line is now `github.com/n-ae/nba-api-go/v2`, and every internal import across the repo (185 files) is updated to match. **Import path change for all consumers**: `import "github.com/n-ae/nba-api-go/pkg/..."` becomes `import "github.com/n-ae/nba-api-go/v2/pkg/..."`. Anyone who somehow got a working `v2.0.0`/`v2.1.0` build via a `replace` directive or vendoring workaround will need to update their import paths to `/v2`; anyone on `v1.x`, or anyone who hit the `go get` failure above and is upgrading for the first time now, is unaffected beyond using the corrected path. `docs/RELEASE_CHECKLIST.md`'s Major Release section now calls this requirement out by name so it isn't missed again at `v3.0.0`.

- **`playercareerstats.go` now validates result-set headers**, matching the pattern `commonplayerinfo`/`playergamelog`/`teamgamelog` gained in `[2.1.0]`. Found by the `v2.1.0` maintainability assessment: `CLAUDE.md` claimed this endpoint already validated headers, but it only did name-based result-set dispatch and parsed rows by fixed position with a length guard - the exact silent-corruption gap `validateHeaders` exists to close. All 8 of its result sets (`SeasonTotals*`/`CareerTotals*` x Regular/Post/AllStar/College) now call `findResultSet`/`validateHeaders` against `jsonTags(SeasonStat{})`/`jsonTags(CareerTotalStat{})`.
- **`parseSeasonStats`' row-length guard was off by one** (`len(row) < 28`, found while adding the header-validation test above): `SeasonStat` has 27 fields (indices 0-26), so the guard required one more column than the parser ever reads - a live 27-column NBA.com response would have every row silently dropped. Corrected to `len(row) < 27`. (`CareerTotalStat`'s equivalent guard was already correct.)
- Added `handwritten_headers_test.go`'s first coverage for `playercareerstats.go` (previously 0%).

## [2.1.0] - 2026-07-20

### Added
- `tools/generator/fieldname_overrides.json` - a per-`(endpoint, result set, field)` exception layer for `goFieldName`'s Go-identifier capitalization, mirroring `fieldtype_overrides.json`'s pattern for field *types*. `VideoEvents`' `vl`/`vt`/`gc`/`surl`/`durl`/`vurl`/`purl` fields are hand-committed as fully-uppercase (`VL`, `VT`, `GC`, `SURL`, `DURL`, `VURL`, `PURL`) despite not being recognized initialisms by any standard convention - `v2.0.0` deliberately left these unregenerated rather than guess at a general rule for arbitrary short abbreviations (see `[2.0.0]`'s notes). Rather than add them to the global initialisms list (which would silently affect any future, unrelated field literally named `vl`/`gc`/etc.), they're now a narrow, scoped override consulted before the general capitalization rule. `goFieldName` gained `endpointName`/`resultSetName` parameters to support this (threaded through from `inferFieldTypes`, same as `resolveFieldGoType` already does for types). `TestGoFieldNameOverridesApplyOnlyWithinTheirEndpoint` proves the override doesn't leak beyond its exact scope; `TestGoFieldNameOverridesReferenceRealMetadata` fails CI on a typo'd or stale entry.
- Committed contract test fixtures (`tests/contract/fixtures/`) recorded against live `stats.nba.com` traffic, so the contract test suite is no longer a no-op in a clean checkout.

### Fixed
- **`GetLeagueLeaders` (`leagueleaders.go`, the hand-written v1 endpoint - not `LeagueLeadersV2`) was silently returning zero leaders on every call.** Found and confirmed live against `stats.nba.com` while closing out the header-validation follow-up noted in `[2.0.0]`'s migration guide. Two compounding bugs, both now fixed:
  1. This endpoint's live response wraps its data in a singular `"resultSet"` object, not the `"resultSets"` array every other classic-Stats-API endpoint uses (confirmed live; `LeagueLeadersV2` does use the standard array shape, so this is specific to the older, unversioned endpoint). The old code decoded into `rawStatsResponse` (`json:"resultSets"`), which always came back empty against this shape - `LeagueLeaders` silently stayed `nil`/empty with no error. Now decodes the singular `resultSet` object directly.
  2. The live response also has a `TEAM_ID` column (position 4 of 28) that `LeagueLeader` didn't have a field for - the old fixed-position parsing (`row[3]` for `Team`, etc.) was one column short even before considering the envelope bug above. `LeagueLeader` gains a `TeamID int` field in the correct position.
  - A third, genuine API quirk surfaced while fixing this and shaped the final approach: **the column set itself varies by `PerMode`** - confirmed live, `PerMode=PerGame` (the most common usage) omits `PF`, `AST_TOV`, and `STL_TOV` that `PerMode=Totals` includes (25 columns vs. 28). A strict full-header `validateHeaders` check (the pattern used everywhere else in this codebase) would hard-fail every `PerGame` call. `parseLeagueLeaders` now looks columns up by header name instead of a fixed position/count, so it's correct regardless of which optional columns a given response includes; a column absent from a response decodes as that field's zero value. Header validation here only requires `PLAYER_ID`/`RANK`/`PLAYER` to be present, not an exact match.
  - Non-breaking: `LeagueLeader` only gained a field (`TeamID`); every existing field's name and type is unchanged. Verified with `UPDATE_FIXTURES=1 INTEGRATION_TESTS=1` against live `stats.nba.com` for both `PerMode=Totals` and `PerMode=PerGame`; `tests/contract/fixtures/leagueleaders_2023-24_pts.json` now holds 240 real leaders instead of `null`.
- **Generator `-endpoint` flag**: Fixed a bug where `go run . -endpoint X` (without `-metadata`) would silently generate an empty stub struct, print `✅ Code generation complete`, and exit 0 instead of loading the endpoint's metadata. Every actual regeneration in `[2.0.0]`'s history used `-metadata` explicitly, so generated output was never affected; this only impacted the interactive single-endpoint preview/regenerate workflow documented in `CLAUDE.md`.
- **Client data race** (`SetHeader`/`AddHeader`/`SetHeaders` vs. in-flight `Get`): synchronized all header mutations with a `sync.RWMutex` to eliminate the reported `WARNING: DATA RACE`.
- **`RetryConfig.MaxRetries < 0`** now makes the retry middleware execute the request exactly once (attempt 0) instead of silently returning `(nil, nil)`. Relevant now that `pkg/client/middleware` is a public package and `AdditionalMiddlewares` makes `WithRetry` reachable by any consumer.
- **Oversized response error mapping**: Status-code-to-error mapping is now applied before the body-size check is enforced, preserving e.g. `models.ErrTooManyRequests` for 429 responses that exceed `MaxResponseBytes`.
- **Metrics cardinality** (`cmd/nba-api-server`): the in-memory latency histogram and per-endpoint counters are now bounded so an arbitrary sequence of unique paths/User-Agent strings cannot grow the metrics map without limit.
- **`internationalbroadcasterschedule.go`**: removed a redundant `interface{}` intermediate and a marshal/unmarshal round-trip; now decodes the `"NextGameList"` field directly in a single pass.
- **`cmd/nba-api-server` version constant**: was hardcoded as `"1.2.0"` since the v1.3.0 release; bumped to track the actual release version going forward.

### Changed
- `videoevents.go` regenerated for the first time since the generator became capable of producing correct output for it. Verified precisely: **zero changes** to the public struct or its field-to-row-index mapping - the fieldname override produces exactly the field names already committed. The only diff is the same pre-existing `"X is required"` error-message format change already applied to every other regenerated endpoint in `[2.0.0]`. **135 of 135 metadata-covered SDK endpoints now have field names and types verified to match what the generator produces** (up from 134 of 135 in `[2.0.0]`).
- Closes the `[2.0.0]` follow-up: `commonplayerinfo.go`, `playergamelog.go`, and `teamgamelog.go` now call `findResultSet`/`validateHeaders` against each result set's own `jsonTags`, matching the pattern generated code uses. All 6 hand-written endpoints now validate headers in some form.
- **Live verification of the `[2.0.0]` header-validation risk is partially discharged, not fully closed.** `leagueleaders.go` was exercised end-to-end against live `stats.nba.com` traffic for both `PerMode` variants. `commonplayerinfo.go`/`playergamelog.go`/`teamgamelog.go` and the entire generated-endpoint surface (121+ files) remain unverified against live traffic. Given that the first live-checked endpoint (`LeagueLeaders`) had two previously-invisible bugs, treat the remaining unverified backlog as a priority, not a formality.
- `perf(client)`: base URL is now parsed once during `NewClient` construction rather than on every `Get` call.
- `cmd/nba-api-server`: upstream NBA API timeout and CORS allowed origin are now configurable at runtime via `NBA_API_TIMEOUT` (duration, e.g. `30s`) and `CORS_ALLOW_ORIGIN` environment variables.
- `models.APIError`: response body bytes are now included in the error value so callers and server logs see NBA.com's actual error text, not just a status code.

## [2.0.0] - 2026-07-20

**The correctness release.** Makes this project's central promise - type-safe, verified access to the NBA Stats API - actually true for the large majority of endpoints, per the plan in `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md`. Four pieces, in the order they landed:

1. **Explicit field types** (`tools/generator/fieldtypes.json`, `fieldtype_overrides.json`) replace the generator's unreviewed naming-heuristic (`inferGoType`) as the source of truth for what Go type a field gets - 48 global corrections plus per-endpoint overrides for the handful of fields that legitimately mean different things in different endpoints.
2. **121 of 141 SDK endpoints regenerated** with those corrected types, fixing real data-corruption bugs already living in committed code (decimal values silently truncated to `"0"` via `string` formatting, name strings silently zeroed via `float64` parsing, and more) - not just the type mismatches themselves. A camelCase-metadata generator bug that would have produced **inaccessible, unexported struct fields** is also fixed. 134 of 135 metadata-covered endpoints are verified correct; the remaining one (`VideoEvents`) is documented, not silently left behind.
3. **Result-set-name keying with header validation** replaces positional array indexing (`rawResp.ResultSets[0]`) across every endpoint using the classic row/columnar response shape - the fix for silent corruption if NBA.com ever reorders result sets or columns. **This is real, deliberate new failure-mode surface, not verified against live NBA.com responses** (no network access from the environment that built this release) - see the risk callout below before relying on it in production.
4. Assorted smaller correctness fixes found along the way (an `interface{}`-decoding shape, a corrected task-tracking miscount).

**Read the `Migration guide` section below, and the `**Breaking:**`/`**Real operational risk**` callouts throughout `Changed`, before upgrading** - this is not a drop-in update for every consumer.

### Added
- `tools/generator/fieldtypes.json` - an explicit, hand-reviewed `{"FIELD_NAME": "goType"}` dictionary covering all 854 field names referenced by committed `tools/generator/metadata/*.json`, the "explicit per-field type metadata" item from the v2.0.0 plan in `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md`. The generator's `fieldGoType` now consults this dictionary before ever calling `inferGoType`; `inferGoType` is demoted to a fallback used only for fields with no dictionary entry, and falling back prints a warning to stderr. `TestAllMetadataFieldsHaveExplicitTypes` fails CI if any metadata field lacks an entry, so a new endpoint can't ship on an unreviewed guess.
- The dictionary corrects 34 fields beyond the 9 `knownWrong` cases `TestInferGoType` already documented, found by reading each field's surrounding context in its metadata file (not live-verified against NBA.com - see the caveat below): 6 more textual range-bucket fields typed `int` instead of `string` (`SHOT_DIST_RANGE`, `SHOT_ZONE_RANGE`, `TOUCH_TIME_RANGE`, plus the 3 already documented), `FG_PCT_MID_RANGE` typed `int` instead of `float64` (a percentage, not a range bucket, despite containing `_RANGE`), 4 more zone shooting percentages typed `string` instead of `float64` (`FG_PCT_ABOVE_BREAK_3`, `FG_PCT_BACKCOURT`, `FG_PCT_LEFT_CORNER_3`, `FG_PCT_RIGHT_CORNER_3`, same bug class as the already-documented `FG_PCT_RA`/`FG_PCT_IN_PAINT`), `PCT`/`WinPCT` typed `string` instead of `float64`, 13 `PCT_<STAT>` "share of team total" fields (`PCT_FGM`, `PCT_FGA`, `PCT_FG3M`, `PCT_FG3A`, `PCT_FTM`, `PCT_FTA`, `PCT_AST_2PM`, `PCT_AST_3PM`, `PCT_AST_FGM`, `PCT_UAST_2PM`, `PCT_UAST_3PM`, `PCT_UAST_FGM`, `PCT_BLKA`) typed `int` instead of `float64` because the FGM/FGA suffix sub-rule fires on any field ending in `m`/`a`, `PERCENTILE` typed `string` instead of `float64`, and `GENERALMANAGER` (a person's name) typed `int` instead of `string` because `"generalmanager"` contains the substring `"age"`.
- `tools/generator/generator_test.go`'s `TestFieldTypesOverridesKnownWrongInference` proves the dictionary actually corrects each of the above cases (fails if `inferGoType` alone would already give the right answer, so a stale/redundant entry gets caught too).
- 14 more `fieldtypes.json` corrections found by a different, stronger method: cross-referencing every field name against the Go type actually used for it across all 143 committed `pkg/stats/endpoints/*.go` files and adopting the type where committed code is unanimous (single type across every occurrence) and disagrees with the dictionary's inherited `inferGoType` default. Corrects `TO` (turnovers; was `string`, silently never matched any `inferGoType` rule), `EVENTMSGTYPE`/`EVENTMSGACTIONTYPE`/`EVENTNUM`/`PERSON1TYPE`/`PERSON2TYPE`/`PERSON3TYPE` (play-by-play type codes; were `string`, should be `int`), `CONF_COUNT`/`DIV_COUNT`/`PO_WINS`/`PO_LOSSES`/`VIDEO_AVAILABLE_FLAG` (were `string`/`float64`, should be `int`), `PLAYER_LAST` (was `float64` - same "ast"-in-"last" substring bug as `PLAYER_NAME_LAST_FIRST` - should be `string`), and `PT_DIFF` (was `string`, should be `float64`). `TestFieldTypesMatchCommittedConsensus` proves each correction.

- `tools/generator/fieldtype_overrides.json` - a per-`(endpoint, result set, field)` exception layer on top of `fieldtypes.json`, for the case the flat dictionary genuinely cannot represent: a field name whose correct type legitimately differs by endpoint, not one side being wrong. `OREB`/`DREB`/`REB`/`AST`/`STL`/`BLK`/`PF`/`PTS` are committed as `float64` in ~90 endpoints (per-game averages) but `int` in the 6 box-score/game-log endpoints that have generator metadata (`BoxScoreTraditionalV2`'s 3 result sets, `LeagueGameFinder`, `TeamGameLogs`, `TeamYearByYearStats` - 2 more, `playergamelog`/`teamgamelog`, have no metadata at all and are unaffected). `TeamGameLogs` and `TeamYearByYearStats` also override several other stat fields (`TOV`, `PFD`, `DD2`, `TD3`, `PTS_RANK`, `WINS`, `LOSSES`, `CONF_RANK`, `DIV_RANK`) that are `int` totals in those two result sets specifically (confirmed by every other stat field in the same struct, e.g. `FGM`/`FGA`, also being `int` there) rather than the `float64` per-game averages the global dictionary assumes. `resolveFieldGoType` in `tools/generator/generator.go` checks this file before `fieldtypes.json`, so an override applies only to the exact endpoint/result-set/field triple it names. `TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint` proves it doesn't leak to sibling result sets or other endpoints; `TestFieldTypeOverridesReferenceRealMetadata` fails CI if an override entry doesn't match real committed metadata (typo/stale-entry protection).

### Changed
- **Breaking:** regenerated 121 of the 135 metadata-covered SDK endpoints (64 files actually differ; the rest were already byte-identical to what the corrected generator produces) from `fieldtypes.json`/`fieldtype_overrides.json`, applying the 48 global field-type corrections and the per-endpoint overrides described above to committed code for the first time. This changes public struct field types - see the field-by-field list and migration guidance below. Verified with a precise struct-level diff (not just text diff) against every touched file: every changed field type traces to a specific, named correction or override; zero unexplained/unverified type changes shipped. One real gap this verification caught before it shipped: `LeagueGameFinder`'s `TOV` override was initially missed (only the 8-field `OREB` family had been checked for that endpoint), which would have silently flipped a correct `int` to `float64` - added to the override file above instead.
- The following field-type corrections are now live in generated code (endpoint(s) in parens): `CLOSE_DEF_DIST_RANGE`/`DRIBBLE_RANGE`/`SHOT_CLOCK_RANGE` `int`→`string` (`PlayerDashPtShots`, `TeamDashPtShots`), `TOUCH_TIME_RANGE` `int`→`string` (`PlayerDashPtShots`), `SHOT_DIST_RANGE` `int`→`string` (`PlayerVsPlayer`), `SHOT_ZONE_RANGE` `int`→`string` (`ShotChartLineupDetail`), `FG_PCT_MID_RANGE`/`FG_PCT_ABOVE_BREAK_3`/`FG_PCT_BACKCOURT`/`FG_PCT_LEFT_CORNER_3`/`FG_PCT_RA`/`FG_PCT_RIGHT_CORNER_3` → `float64` (`LeagueDashPlayerShotLocations`, `LeagueDashTeamShotLocations`), `DISPLAY_FIRST_LAST`/`DISPLAY_LAST_COMMA_FIRST`/`DISPLAY_FI_LAST` `float64`→`string` (`CommonPlayerInfoV2`), `PLAYER_NAME_LAST_FIRST` `float64`→`string` (`PlayerDashPtShots`), `PCT`/`WinPCT`/`PERCENTILE` → `float64` and 13 `PCT_<STAT>` fields `int`→`float64` (`BoxScoreScoringV2`, `BoxScoreUsageV2`, `TeamInfoCommonV2`, `TeamYearByYearStats`), `GENERALMANAGER` `int`→`string` (`TeamDetails`), `TO` `string`→`int` (`BoxScoreTraditionalV2` and others).
- **Breaking:** regenerated 3 more endpoints - `PlayerDashboardByGeneralSplits`, `TeamDashboardByGeneralSplits`, `BoxScoreSummaryV2` - that were previously excluded because every result-set field was typed `interface{}` (300, 270, and 93 fields respectively), not because their data was wrong. Checked precisely: the old code assigned the raw decoded row value directly (e.g. `GAME_DATE_EST: row[0]`) with zero calls to `toInt`/`toFloat`/`toString` - `encoding/json` had already decoded each value to its correct dynamic Go type (a JSON number becomes Go's `float64`, a JSON string becomes `string`, etc.), so **no data was corrupted or lost**, unlike the string/int mistyping fixed above. The defect was purely a violation of this project's "no `interface{}` in generated code" design principle (see CLAUDE.md's Key Design Principles): callers had to do their own runtime type assertions (`resp.GameSummary[0].GAME_ID.(string)`) to use any field, with no compile-time safety and a panic risk if a field's dynamic type ever changed. All 663 fields across the three endpoints now have concrete types wired to real `toInt`/`toFloat`/`toString` parsing; field count and result-set shape are unchanged. Verified the same way as the 121-file regeneration above (every field's new type checked against `fieldtypes.json`/`fieldtype_overrides.json`, zero unexplained assignments) plus a direct count check (interface{} count 0, parsed-field count == old field count, for all three files).
- **Fixed a real generator bug, not just a naming-style mismatch**: the template used a metadata field's raw name directly as the Go struct field identifier. NBA Live-Data-style endpoints (`PlayByPlayV3`, `ScoreboardV3`, `VideoEvents`, `LeagueStandings`, `LeagueStandingsV3`) have camelCase metadata field names (e.g. `"gameId"`); generating a Go field named `gameId` produces an **unexported** field, invisible to `encoding/json` and inaccessible to any external caller - a correctness bug, not a style choice. `tools/generator/generator.go`'s new `goFieldName` capitalizes the leading rune of each camelCase word and fully uppercases recognized initialisms (`ID`, `URL`, `UUID`, etc. - the same list `golang.org/x/lint` uses), matching the convention already used by this project's hand-fixed committed code for these endpoints (`"gameId"` → `GameID`, not `GameId`). Already-exported field names (the common `SCREAMING_SNAKE_CASE` case, e.g. `GAME_ID`) pass through completely unchanged. `TestGoFieldName` checks the conversion against real committed field names.
- **Breaking:** regenerated 4 of the 5 endpoints this unblocks - `PlayByPlayV3`, `ScoreboardV3`, `LeagueStandings`, `LeagueStandingsV3` - now byte-identical (modulo the pre-existing error-message-format difference) to what the fixed generator produces. `LeagueStandings`/`LeagueStandingsV3` also picked up the already-known `WinPCT` `string`→`float64` correction, previously blocked by this same naming bug. `VideoEvents` remains excluded: 7 of its fields (`vl`, `vt`, `gc`, `surl`, `durl`, `vurl`, `purl`) are hand-committed as fully-uppercase (`VL`, `VT`, `GC`, `SURL`, `DURL`, `VURL`, `PURL`) despite not being recognized initialisms by any standard convention - deliberately not hardcoded as a generic rule, since guessing which arbitrary short field names "deserve" full capitalization would be fragile and wouldn't generalize to future endpoints.
- **Fixed a real, shipped data-corruption bug found while investigating `shotchartdetail.go`'s remaining drift** (this bug was introduced by this same `[Unreleased]` cycle's own earlier regeneration commit above, not by any tagged release - `ShotChartLineupDetail` had never been correctly typed before this project's `git` history existed, and only the `SHOT_ZONE_RANGE` fix from that commit touched it): `GAME_EVENT_ID`, `MINUTES_REMAINING`, `SHOT_DISTANCE`, `LOC_X`, `LOC_Y`, `SHOT_ATTEMPTED_FLAG`, `SHOT_MADE_FLAG` are corrected globally in `fieldtypes.json` (`int`/`float64`, matching `ShotChartDetail`'s own already-correct committed types - shot chart coordinates, distances, and flags are unambiguously numeric). `SHOT_DISTANCE`/`LOC_X`/`LOC_Y` were typed `string` and formatted via `toString`'s `"%.0f"` - the same corruption class as the already-documented `FG_PCT_RA` bug: a court X coordinate of `23.5` silently became `"24"`. `SECONDS_REMAINING` gets a narrow `fieldtype_overrides.json` entry for `ShotChartDetail` only (its committed value, `int`, disagrees with `ShotChartLineupDetail`/`WinProbabilityPBP`'s `string` without the same unambiguous evidence the other fields have - left as a global `string` default pending further investigation, see `TestShotChartFieldsMatchShotChartDetailPrecedent`'s doc comment). `shotchartdetail.go` is newly regenerated in this change; `shotchartlineupdetail.go` (already among the 121 files from the earlier regeneration commit) is re-regenerated to pick up this correction.
- **Breaking:** regenerated the last 3 files whose only remaining drift was field types: `commonallplayers.go`, `commonallplayersv2.go`, `teaminfocommon.go`. Their apparent "22/12 unexplained lines" in earlier drift analysis turned out to be a false signal from `gofmt` re-flowing column alignment across an entire struct whenever a sibling field's type changes length (e.g. removing `float64` shrinks the type column, shifting every field's alignment even though most field types are unchanged) - a precise struct-level parse (field name/type pairs, ignoring whitespace) showed **zero unexpected type changes** once checked properly. `commonallplayers[v2].go`'s only real change is the already-known `DISPLAY_FIRST_LAST`/`DISPLAY_LAST_COMMA_FIRST` correction. `teaminfocommon.go`'s `W`/`L` (`int`→`string`) and `MIN_YEAR`/`PTS_RANK`/`REB_RANK`/`AST_RANK`/`OPP_PTS_RANK` (`int`→`float64`) were manually verified against codebase-wide committed consensus before regenerating, not just trusted from the dictionary: `W`/`L` are `string` in 85 of 87 committed occurrences (`teaminfocommon.go` and `teamgamelog.go` were the only `int` outliers), the four rank fields are `float64` in 13-14 of 14-15 occurrences each, and `MIN_YEAR` is `float64` in both of `teaminfocommon.go`'s siblings (`commonteamyears.go`, `teaminfocommonv2.go`). None of these changes lose data in either direction (a whole-number value round-trips cleanly through `int`/`string`/`float64`). `TestTeamInfoCommonFieldsMatchCodebaseMajority` documents this.
- **Permanently excluded from regeneration** (not "not yet" - regenerating these would delete real, hand-written value the generator can't reproduce): `gamerotation.go`'s struct comments document a previously-fixed, live-verified bug (a missing `TEAM_CITY` column that shifted every subsequent field, a `PLAYER_LAST` mistyped `float64`, a `PT_DIFF` mistyped `string`) that a fresh generation would silently overwrite with a generic one-line comment; `leaguedashplayerstats.go` sets ~30 parameters to empty-string/default values before the request-specific ones, which the current metadata format has no way to express and a fresh generation would drop entirely, likely breaking the live API call.
- **Not yet regenerated** (1 file, still on the old, partly-wrong types): `videoevents.go` (see above - blocked on 7 non-standard field-name capitalizations, not a data-correctness issue). Verified precisely (parsing every committed struct's field names and types, not just counting regeneration commits): **134 of 135 metadata-covered endpoints now have field names and types matching what the corrected generator would produce** - `videoevents.go` is the sole exception. `gamerotation.go` and `leaguedashplayerstats.go` are counted in the 134: they were already correct before any of this `[Unreleased]` cycle's work and were never regenerated, just left alone (see above for why regenerating them would be a regression despite their types being fine).

- **Result sets are now looked up by name, and their headers validated, instead of being read by array position** - the "result-set-name keying with upstream-header validation that errors on mismatch instead of shifting columns" item from the v2.0.0 plan. Previously every generated parser read `rawResp.ResultSets[0]`, `rawResp.ResultSets[1]`, etc. - purely positional, silently reading the wrong result set into the wrong struct if NBA.com ever returns them in a different order, and blindly trusting `row[i]` positions against a fixed field list with no check against the `headers` array NBA.com actually sent for that response. `pkg/stats/endpoints/types.go` adds three functions: `findResultSet` (looks a result set up by its `name` field), `validateHeaders` (compares a result set's actual headers against the expected field order and returns an error on any mismatch - length or content, in either direction), and `jsonTags` (reflects a generated struct's own `json:"..."` tags to build that expected list, so the expected headers are never a second, separately-maintained copy of the same field names - avoiding exactly the kind of drift-between-two-copies bug this whole file exists to catch). `tools/generator/templates/endpoint.tmpl` now generates `if rs, ok := findResultSet(rawResp.ResultSets, "X"); ok { if err := validateHeaders(rs.Headers, jsonTags(EndpointX{})); err != nil { return nil, fmt.Errorf(...) } ... }` instead of `if len(rawResp.ResultSets) > 0 { ... rawResp.ResultSets[0] ... }`. `types_test.go` adds direct unit coverage (`TestFindResultSet`, including `TestFindResultSet_KeysByNameNotPosition` reproducing the exact bug class this fixes; `TestValidateHeaders`; `TestJSONTags`).
- Applied to all 132 metadata-covered files whose result-set structure could safely regenerate (verified with the same precise struct-level + row-index-mapping diff used elsewhere in this release: **zero** field name, type, or row-index changes anywhere - this is purely a parsing-mechanism change, not a data change) - plus `gamerotation.go`, `leaguedashplayerstats.go`, and `videoevents.go`, hand-patched with the identical pattern since blind regeneration would have destroyed their hand-written content (see above) or their still-open field-naming issue (`videoevents.go`). That's every endpoint using the classic Stats API's row/columnar response shape. Not covered: `internationalbroadcasterschedule.go`, which uses a structurally different (map-keyed, not row-based) response shape unrelated to this mechanism, and 5 of the 6 hand-written endpoints (`commonplayerinfo`, `playercareerstats`, `playergamelog`, `teamgamelog`, `leagueleaders`) which already did name-based lookup by hand (no positional-indexing bug to fix) but don't yet call `validateHeaders` - a smaller, separate follow-up.
- **Real operational risk, stated plainly**: `validateHeaders` returning an error on any header mismatch is new failure-mode surface that did not exist before, and it has **not been verified against live NBA.com responses** (no network access from this environment - see the recurring caveat throughout this release). If any endpoint's metadata field order has silently drifted from what NBA.com currently returns, a call that previously "worked" (by degrading silently - reading data into the wrong fields) will now return an error instead. This is the intended, deliberate trade-off the v2.0.0 plan calls for (fail loud beats corrupt silently), not an oversight - but it means this change should be watched closely once it can be exercised against live traffic, via the contract tests once fixtures exist or via integration testing before the next tagged release.

### Fixed
- `internationalbroadcasterschedule.go`'s `GetInternationalBroadcasterSchedule` decoded its response through an unexported `map[string]interface{}` intermediate, then re-marshaled and re-unmarshaled the `"NextGameList"` value a second time to get typed `[]ScheduledGame` out of it. Since `"NextGameList"` is a fixed, known key, this now decodes directly into a nested struct with a `json:"NextGameList"` tag in a single pass - no `interface{}` anywhere, and no redundant marshal round-trip. Split into a new `parseInternationalBroadcasterScheduleResponse(body []byte)` so this parsing logic is unit-testable without a live HTTP call; `TestParseInternationalBroadcasterScheduleResponse` and `TestParseInternationalBroadcasterScheduleResponse_FieldValues` cover it (populated games, empty/absent `NextGameList`, no result sets, malformed JSON, and field-value correctness).
- **Corrects an earlier miscount in this file**: task tracking for this release claimed "9 `interface{}` fields scattered across the 6 hand-written endpoints." That count came from an unfiltered `grep -c interface{}`, which also matched `[][]interface{}` row-parsing **function parameters** (`parseSeasonStats(rows [][]interface{})` and similar) - the same, unavoidable, correct shape used at the JSON-decode boundary by every endpoint in this codebase, generated or hand-written, not a violation of anything. Rechecked precisely: `internationalbroadcasterschedule.go`'s `map[string]interface{}` above was the only actual `interface{}` struct-field/decoding-shape issue across all 6 files, and it's fixed by this entry. No further work remains from the original "6 hand-written endpoints" item.
- `cmd/nba-api-server`'s `/health` endpoint (and its startup log line) reported a hardcoded `const version = "1.2.0"` - stale since the v1.3.0 release, which should have bumped it and didn't. Bumped to `"2.0.0"` as part of this release; this constant should be updated in every future release's checklist, not just remembered.

### Migration guide
- If your code reads any of the fields listed above from one of the named endpoints, update the Go type you assign it to (e.g. a `float64` variable receiving `DISPLAY_FIRST_LAST` from `CommonPlayerInfoV2` needs to become `string`).
- If you were working around the old wrong type (e.g. parsing a `string`-typed `FG_PCT_RA` back into a number yourself), that workaround is now unnecessary and will fail to compile against the new `float64` field - remove it.
- If your code was previously getting silently-wrong data from an endpoint affected by result-set reordering or a column-header mismatch, calling that endpoint will now return an error instead (see the operational-risk note above) - this is intentional, but check your error handling around SDK calls if you're upgrading. This applies to essentially every endpoint using the classic row/columnar response shape, including `videoevents.go` (whose field *types* are unaffected by this change, but which now validates headers like everything else).
- Endpoints unaffected by *either* the field-type or result-set-keying changes above: `internationalbroadcasterschedule.go` (a structurally different, non-row-based response shape) and 4 of the 6 hand-written endpoints without generator metadata (`commonplayerinfo`, `playergamelog`, `teamgamelog`, `leagueleaders` - `playercareerstats` already did name-based lookup by hand before this release).

### Note
- The 48 global dictionary corrections come from reading field names in context (e.g. `GENERALMANAGER` sitting next to `OWNER`/`HEADCOACH` in `TeamDetails.TeamBackground`) or cross-referencing committed-code consensus, not from replaying live or recorded NBA.com responses. Treat `fieldtypes.json`/`fieldtype_overrides.json` as a large improvement over `inferGoType` alone, not a completed field-by-field audit - the contract tests (`tests/contract/`) are what catches remaining drift once fixtures exist.

## [1.3.0] - 2026-07-20

### Added
- `.github/dependabot.yml` watching the `github-actions` ecosystem and both `gomod` modules (root and `tools/generator`) - automates the check that this session had to do by hand after `golangci-lint-action@v6`/`actions/checkout@v4`/`actions/setup-go@v5` all turned out to be stale. Requires Dependabot to be turned on for this repo in GitHub Settings (Code security and analysis) - the config alone doesn't enable it; as of this change, Dependabot alerts are confirmed disabled at the repo level.
- Facade-level regression tests (`stats.TestNewClient_DefaultHeadersReachTheWire`, `live.TestNewClient_DefaultHeadersReachTheWire`) asserting the actual `User-Agent`/`Referer`/`Accept` bytes an `httptest.Server` receives from a default-config client, not just that some middleware ran.
- `pkg/client/middleware` - the built-in middleware constructors (`WithRetry`, `RetryConfig`/`DefaultRetryConfig`, `WithPerHostRateLimit`, `WithHeaders`, `WithUserAgent`, `WithReferer`, `WithAccept`, the `WithLogging` family) are now importable outside this module. They previously lived at `internal/middleware`, which by Go's own visibility rules was never importable by an external `go get` consumer in the first place - this is a purely additive change, not a relocation of something anyone outside this module could have depended on. `pkg/stats` and `pkg/live` now import the new path; `internal/middleware` no longer exists.
- `stats.Config.AdditionalMiddlewares` / `live.Config.AdditionalMiddlewares` - layers your own middleware on top of whichever chain `Middlewares` resolves to (the defaults if `Middlewares` is empty, or your override if it isn't), instead of requiring `append(stats.DefaultMiddlewares(), yourMiddleware...)` by hand. Purely additive: existing `Middlewares`-only configs are unaffected.
- `tools/generator` unit tests (`TestInferGoType`, 39 subtests) pinning the current field-name-to-Go-type heuristic rule by rule, including the known-wrong cases that produce the data-corruption bugs cataloged in `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md` (display names as `float64`, textual range buckets as `int`, decimal percentages as `string`). These pin current behavior so a future change to the heuristic is a visible, intentional diff, not silent drift.
- `cmd/nba-api-server`'s `Season` query parameter now defaults to the current NBA season (computed from today's date; seasons run October-June) instead of a literal `"2023-24"` that was hardcoded at over 100 call sites and had been stale for multiple seasons. `getSeasonOrDefault`/`currentSeasonDefault` in `cmd/nba-api-server/handlers.go` centralize this - a season rollover, or a change to the default-selection rule, is now a one-place change.
- `tools/generator`'s `TestGenerateFromMetadata_ProducesValidGo` regenerates every committed `metadata/*.json` file into a scratch directory and asserts the output parses as syntactically valid Go. This is a generator/template regression test, not a fidelity check: regenerating today produces output that differs from the committed `pkg/stats/endpoints/*.go` files in real, substantive ways (field types, not just formatting), so a literal "regenerate everything and diff against committed source" CI gate would be permanently red. Reconciling that drift safely (it needs verification against live NBA.com responses, not a blind diff) is the v2.0.0-sized "explicit per-field type metadata, regenerate all 141 endpoints" item tracked in `docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md`. What this test does catch now: a broken template or a generator change that emits malformed Go.
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

[Unreleased]: https://github.com/n-ae/nba-api-go/compare/v3.1.2...HEAD
[3.1.2]: https://github.com/n-ae/nba-api-go/compare/v3.1.1...v3.1.2
[3.1.1]: https://github.com/n-ae/nba-api-go/compare/v3.1.0...v3.1.1
[3.1.0]: https://github.com/n-ae/nba-api-go/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/n-ae/nba-api-go/compare/v2.2.0...v3.0.0
[2.2.0]: https://github.com/n-ae/nba-api-go/compare/v2.1.2...v2.2.0
[2.1.2]: https://github.com/n-ae/nba-api-go/compare/v2.1.1...v2.1.2
[2.1.1]: https://github.com/n-ae/nba-api-go/compare/v2.1.0...v2.1.1
[2.1.0]: https://github.com/n-ae/nba-api-go/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/n-ae/nba-api-go/compare/v1.3.0...v2.0.0
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
