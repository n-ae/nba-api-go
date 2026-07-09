# Lint Cleanup Plan

**Status**: Not started (tracking doc only)
**Created**: 2026-07-09, alongside the `GameRotation` column-offset fix (v1.1.3)
**Owner**: unassigned

## Background

`make lint` has silently not been running: `.golangci.yml` was written in the
golangci-lint v1 config schema, but the installed toolchain
(`golangci-lint 2.12.2`) requires the v2 schema and fails to load the old
config with `unsupported version of the configuration: ""`. This was fixed as
part of the `GameRotation` bugfix (v1.1.3) by migrating `.golangci.yml` to
`version: "2"` (see git history for that commit).

Running the migrated config for the first time surfaced **413 pre-existing
issues** across the repo, almost entirely in the 137 generated files under
`pkg/stats/endpoints/`. None of these are new — they've been accumulating
silently since lint stopped running. This doc catalogs them by category with
severity, root cause, and a recommended fix strategy, so they can be tackled
as a separate, deliberate piece of work rather than folded into an unrelated
bugfix.

One category (`fieldalignment`, ~310 instances) was already resolved by
disabling that specific govet check in `.golangci.yml`, not by touching code
— see [Already resolved](#already-resolved-fieldalignment) below.

Re-run `golangci-lint run ./...` at the start of this work to get current
counts; the numbers below are a snapshot from 2026-07-09 and will drift as
the codebase changes.

## Already resolved: fieldalignment

`govet`'s `fieldalignment` check (~310 of the original 724 issues) suggests
reordering struct fields to reduce padding. It was disabled outright rather
than fixed, because the generator (`tools/generator/templates/endpoint.tmpl`)
intentionally emits struct fields in NBA API column order — that ordering is
the readability feature, not a bug. Reordering ~100+ generated files for a
handful of padding bytes on structs that are deserialized once per HTTP
response is not a worthwhile trade. If this is revisited, do it by changing
the generator template (so regeneration doesn't undo it), not by hand-editing
generated files.

## Remaining issues by category

### 1. `structtag` (govet) — 88 instances, 5 files — **real bug, highest priority**

Files: `leaguestandings.go` (13), `leaguestandingsv3.go` (13),
`playbyplayv3.go` (33), `scoreboardv3.go` (21), `videoevents.go` (8).

These structs use **unexported** field names (`gameId`, `actionNumber`,
`strLongHomeStreak`, `vsEast`, ...) with `json:"..."` tags attached, e.g.:

```go
type PlayByPlayV3PlayByPlay struct {
	gameId       string `json:"gameId"`
	actionNumber string `json:"actionNumber"`
	...
}
```

Unlike the `GameRotation` bug, this doesn't corrupt the *parsing* — these
particular endpoints populate fields via manual `toString`/`toInt` calls in a
struct literal within the same package, not via `encoding/json.Unmarshal`, so
the values do get set internally. **The actual bug: every field on these five
result-set types is unexported, so no code outside `pkg/stats/endpoints` can
read them at all.** `GetPlayByPlayV3`, `GetScoreboardV3`,
`GetLeagueStandingsV3`, `GetLeagueStandings`, and `GetVideoEvents` currently
return data that is structurally unusable by SDK consumers and by the HTTP
API server's JSON handlers (which also can't see the fields, so `encoding/json`
silently emits `{}` for these fields in HTTP responses).

**Fix**: capitalize each field name to export it, keeping the original
`json:"..."` tag value (e.g. `gameId string` → `GameID string \`json:"gameId"\`
`, matching this repo's existing `GAME_ID`/`TEAM_ID` naming convention used
elsewhere, or `GameID` if following idiomatic Go naming — pick one convention
and apply consistently). Then:
- Update every struct-literal field reference in the same file (the
  `toString(row[n])` assignments).
- `grep -rn` for any external consumers already reaching into these structs
  (`cmd/nba-api-server/handlers_*.go`, `examples/`) and update field
  references.
- Check whether these files came from `tools/generator` or were hand-written
  — if generator-sourced, fix the source metadata/template too so
  regeneration doesn't reintroduce unexported fields.
- Add/extend unit tests per endpoint (following the pattern in
  `pkg/stats/endpoints/gamerotation_test.go`) asserting the exported fields
  are actually reachable and populated.

Estimated effort: 3–5 hours (5 files, plus consumer grep and tests).

### 2. `unconvert` — 298 instances, ~90 files — **style-only, mechanical**

Root cause: `tools/generator/templates/endpoint.tmpl` unconditionally wraps
every parameter in `string(...)` when building `url.Values`:

```go
params.Set("{{.Name}}", string(req.{{.Name}}))       // required
params.Set("{{.Name}}", string(*req.{{.Name}}))      // optional
```

This is necessary when `.Type` is a named string type (`parameters.Season`,
`parameters.LeagueID`, etc.) but redundant when `.Type` is a plain `string` —
which `unconvert` correctly flags.

**Fix**:
1. Update the template to only emit the `string(...)` wrapper when
   `.Type` is not the literal `"string"`, e.g.:
   ```gotemplate
   {{- if eq .Type "string"}}
   params.Set("{{.Name}}", req.{{.Name}})
   {{- else}}
   params.Set("{{.Name}}", string(req.{{.Name}}))
   {{- end}}
   ```
   (mirror for the optional/pointer branch).
2. Regenerate affected endpoints from their metadata batch files in
   `tools/generator/metadata/`, or run `golangci-lint run --fix ./...`
   directly against the generated files (unconvert supports autofix) — if
   using the autofix path, still fix the template afterward so the next
   regeneration doesn't reintroduce the issue.
3. Diff regenerated files carefully against current ones — the template may
   have drifted in other ways since these files were last generated (e.g.
   the `/gamerotation` vs `gamerotation` leading-slash inconsistency noticed
   while fixing `GameRotation` — check `client.GetJSON(ctx, "/{{.Endpoint}}", ...)`
   in the template against the no-leading-slash convention used in every
   currently-checked-in file before regenerating anything).
4. Run full test suite + `make test-examples` after.

Estimated effort: 2–3 hours, mostly verification/diff review, if the
autofix path is used; more if doing a full template-driven regeneration of
all 137 endpoints.

### 3. `errcheck` — 9 instances — **mostly real, quick**

| File | Issue |
|---|---|
| `cmd/nba-api-server/main.go:182,192` | `_ = json.NewEncoder(w).Encode(...)` still flagged because `.golangci.yml` sets `errcheck.check-blank: true`, which treats blank-identifier discards as still unchecked. Prior commit 8f1a9b4 ("fix: ignore JSON encoder return values in HTTP handlers") added the `_ =` but that doesn't satisfy this stricter setting. |
| `cmd/nba-api-server/main.go:229` | Already checked correctly (`if err := ...`) — not actually an issue, only lines 182/192 need it. |
| `examples/{new_endpoints_demo,tier1,tier2,tier3}_demo/main.go` | `json.MarshalIndent` return value unchecked in example programs. Low risk (examples, not library code) but easy to fix with an `if err != nil { log.Fatal(err) }`. |
| `internal/middleware/retry.go:65` | `resp.Body.Close()` unchecked — standard low-risk pattern, but can silence with `//nolint:errcheck` comment or `_ = resp.Body.Close()` if `check-blank` is scoped down for this pattern (see note below). |
| `pkg/client/client.go:118` | Same `resp.Body.Close()` pattern. |
| `pkg/stats/static/players.go:142` | `transform.String(t, s)` return value (including its error) unchecked in `stripAccents` — this silently swallows Unicode normalization failures during player search. Worth an actual look: does a failure here mean partial/garbled output gets returned as if it succeeded? |

**Decision needed before fixing**: either (a) relax
`errcheck.check-blank` to `false` so `_ = fn()` is treated as
intentionally-acknowledged (simpler, matches what commit 8f1a9b4 apparently
assumed), or (b) keep `check-blank: true` and add real error handling
everywhere it fires. (b) is more correct; (a) is less churn. Recommend (b)
for `players.go` specifically since it's a real correctness question, and
either (a) or explicit `//nolint` comments for the `Body.Close()` /
`Encode()` call sites, which are genuinely low-risk.

Estimated effort: 1 hour.

### 4. `staticcheck` — 8 instances — **quick, cosmetic**

- 7× `ST1005: error strings should not be capitalized` — e.g.
  `fmt.Errorf("GameID is required")` in `commonallplayers.go`,
  `commonplayoffseries.go`, `leaguegamelog.go`,
  `playerdashboardbygeneralsplits.go`, `shotchartdetail.go`,
  `teamdashboardbygeneralsplits.go`, `teamgamelogs.go`. Fix: lowercase the
  first letter (`"GameID is required"` stays fine since `GameID` is a Go
  identifier, but staticcheck wants the *sentence* not to start with a
  capital — check each message individually; some may need rewording rather
  than just lowercasing, e.g. `"gameID is required"` if `GameID` itself
  isn't recognized as a proper noun by the heuristic).
  Note: these come from the generator template's
  `fmt.Errorf("{{.Name}} is required")` — fixing the template fixes all
  future generations, but the 7 existing files still need regenerating or
  hand-editing.
- 1× `S1009` in `tests/contract/endpoints_test.go:651` — redundant nil check
  before `len()` on a map; trivial one-line fix.

Estimated effort: 30–45 minutes.

### 5. `unused` — 10 instances — **quick, delete or justify**

All in test helper files: `tests/contract/helpers.go` (5 functions:
`assertGreaterThan`, `compareSchemas`, `findSchemaDifferences`, `prettyJSON`,
`intPtr`) and `tests/integration/helpers.go` (5 functions:
`assertNotEmpty`, `seasonPtr`, `seasonTypePtr`, `leagueIDPtr`, `perModePtr`).

**Fix**: confirm each is genuinely dead (not reachable via a build tag or
reflection) and delete, or use one of them in the next test that's added to
these packages if it was written in anticipation of near-term use. Default
to deleting — CLAUDE.md's "no half-finished implementations" guidance
applies to test helpers too.

Estimated effort: 20 minutes.

## Suggested order of work

1. **`structtag`** (#1) — real bug, highest user impact, do first.
2. **`errcheck`** (#3) and **`staticcheck`** (#4) — quick, low-risk, can be
   done together in under 2 hours.
3. **`unused`** (#5) — quick cleanup, can be bundled with #3.
4. **`unconvert`** (#2) — mechanical but touches the most files; do last and
   in its own PR so a large, low-risk diff doesn't block or get tangled with
   the real bug fixes above.

## Verifying completion

```bash
golangci-lint run ./...   # should report 0 issues
go test ./...
make test-examples
gofmt -l .                # should print nothing
```
