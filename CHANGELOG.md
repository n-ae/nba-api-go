# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/n-ae/nba-api-go/compare/v2.0.0...HEAD
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
