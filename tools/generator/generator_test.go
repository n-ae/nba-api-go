package main

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// TestGenerateFromMetadata_ProducesValidGo is a generator/template
// regression test, not a fidelity check against already-committed
// endpoint code: regenerating from tools/generator/metadata/*.json today
// produces output that differs from pkg/stats/endpoints/*.go in real,
// substantive ways (field types, not just formatting) - see
// docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md's
// v2.0.0 plan item for explicit per-field type metadata and a full
// regeneration of all 141 endpoints, which is the only safe way to
// reconcile that drift (it requires verifying corrected types against
// live NBA.com responses, not blindly trusting either side). What this
// test does check, for every committed metadata file: the generator
// doesn't crash, and its output is syntactically valid Go - so a broken
// template or a change to inferGoType that emits malformed code fails
// CI immediately instead of silently producing something nobody would
// notice was broken until they tried to build it.
func TestGenerateFromMetadata_ProducesValidGo(t *testing.T) {
	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}
	if len(metadataFiles) == 0 {
		t.Fatal("no metadata files found under metadata/ - this test would trivially pass without exercising anything")
	}

	for _, mf := range metadataFiles {
		t.Run(filepath.Base(mf), func(t *testing.T) {
			outDir := t.TempDir()
			g := NewGenerator(outDir)
			if err := g.GenerateFromMetadata(mf, false); err != nil {
				t.Fatalf("GenerateFromMetadata(%s) failed: %v", mf, err)
			}

			entries, err := os.ReadDir(outDir)
			if err != nil {
				t.Fatalf("failed to read output dir: %v", err)
			}
			if len(entries) == 0 {
				t.Fatalf("%s produced no output files", mf)
			}

			fset := token.NewFileSet()
			for _, entry := range entries {
				path := filepath.Join(outDir, entry.Name())
				if _, err := parser.ParseFile(fset, path, nil, parser.AllErrors); err != nil {
					t.Errorf("%s: generated file %s is not valid Go: %v", mf, entry.Name(), err)
				}
			}
		})
	}
}

// TestInferGoType documents inferGoType's current rule-by-rule heuristic
// behavior using real NBA.com field names pulled from committed generated
// code. Several cases below are marked knownWrong: true - they reproduce
// the exact data-corruption bugs cataloged in
// docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md (display
// names typed float64, textual range buckets typed int, decimal
// percentages typed string and truncated). These assertions exist to make
// any future change to the heuristic - the v2.0.0 explicit-type-metadata
// rework the assessment calls for - visible and intentional in a diff, not
// silent. Do not "fix" a knownWrong case by changing its want value
// without also fixing inferGoType and regenerating the affected endpoints;
// until then, this test's job is to pin current behavior, not endorse it.
//
// Two precedence quirks worth knowing when editing inferGoType or this
// test (both verified empirically, not just read from source):
//   - "STATUS_ID"/"GAME_STATUS_ID" never reach the "status substring +
//     _id suffix -> int" rule near the end of the function: the earlier,
//     unconditional "_id suffix" rule already returns "string" for any
//     _id field without "player"/"team" in it. That later rule is dead
//     code today.
//   - "CONF_RANK"/"DIV_RANK" don't reach the dedicated "rank substring ->
//     int" rule either: "_rank" is also in the stat-abbreviation list
//     checked earlier, so they fall through that loop's default
//     ("float64" for most stats) instead. A bare "RANK" (no underscore
//     before it) does reach the dedicated rule and gets "int". Similarly,
//     "SEASON_SEQUENCE" matches the season rule ("contains season, not
//     id") before it can reach the sequence/period/range rule.
func TestInferGoType(t *testing.T) {
	tests := []struct {
		field      string
		want       string
		knownWrong bool
		note       string
	}{
		// --- _PCT/_PERCENTAGE suffix -> float64 (correct) ---
		{field: "FG_PCT", want: "float64"},
		{field: "FT_PCT", want: "float64"},
		{field: "WIN_PERCENTAGE", want: "float64"},

		// --- Decimal percentage fields that DON'T end in _PCT/_PERCENTAGE
		//     fall through to later rules and land on string, losing
		//     decimal precision when toString later formats with %.0f. ---
		{field: "FG_PCT_RA", want: "string", knownWrong: true,
			note: "real committed type in leaguedashplayershotlocations.go; 0.357 becomes \"0\""},
		{field: "FG_PCT_IN_PAINT", want: "string", knownWrong: true,
			note: "same file, same bug class as FG_PCT_RA"},

		// --- _ID suffix: player/team -> int, else -> string (as designed) ---
		{field: "PLAYER_ID", want: "int"},
		{field: "TEAM_ID", want: "int"},
		{field: "GAME_ID", want: "string"},
		{field: "SEASON_ID", want: "string"},
		{field: "LEAGUE_ID", want: "string"},
		{field: "STATUS_ID", want: "string",
			note: "the later status+_id rule is dead code - see doc comment above"},
		{field: "GAME_STATUS_ID", want: "string"},

		// --- Date fields -> string (correct) ---
		{field: "GAME_DATE", want: "string"},
		{field: "GAME_DATE_EST", want: "string"},

		// --- _NAME/_TEXT/etc. suffix -> string, correct for genuine name
		//     fields, but the rule only fires on an exact suffix/substring
		//     match; PLAYER_NAME_LAST_FIRST doesn't end in one of those
		//     suffixes (it ends in "_LAST_FIRST"), so it falls through
		//     instead of matching here. ---
		{field: "TEAM_NAME", want: "string"},

		// --- Fields that read as names but don't match the name rule's
		//     suffix/substring list fall through to the stat-abbreviation
		//     loop and match "ast" hiding inside "last". ---
		{field: "PLAYER_NAME_LAST_FIRST", want: "float64", knownWrong: true,
			note: "real committed type in commonallplayers.go family; a name string becomes 0 via toFloat"},
		{field: "DISPLAY_FIRST_LAST", want: "float64", knownWrong: true,
			note: "real committed type in commonallplayers.go; \"Nikola Jokić\" becomes 0"},
		{field: "DISPLAY_LAST_COMMA_FIRST", want: "float64", knownWrong: true,
			note: "same bug class as DISPLAY_FIRST_LAST"},
		{field: "DISPLAY_FI_LAST", want: "float64", knownWrong: true,
			note: "same bug class as DISPLAY_FIRST_LAST"},

		// --- WL -> string (correct) ---
		{field: "WL", want: "string"},

		// --- Season fields without "id" -> string (correct); this rule's
		//     reach extends past plain "SEASON" to any field containing it,
		//     including ones that arguably belong to a different rule -
		//     see SEASON_SEQUENCE below. ---
		{field: "SEASON", want: "string"},
		{field: "SEASON_YEAR", want: "string"},
		{field: "SEASON_SEQUENCE", want: "string",
			note: "season rule fires before the sequence/period/range rule ever sees it"},

		// --- Stat abbreviations -> float64 by default, with sub-rules ---
		{field: "PTS", want: "float64"},
		{field: "REB", want: "float64"},
		{field: "MIN", want: "float64", note: "\"min\" substring, not \"game\" -> float64 sub-rule"},
		{field: "FGM", want: "int", note: "suffix \"m\" -> int sub-rule"},
		{field: "FGA", want: "int", note: "suffix \"a\" -> int sub-rule"},
		{field: "GP", want: "int", note: "exact \"gp\" -> int sub-rule"},
		{field: "GS", want: "int", note: "exact \"gs\" -> int sub-rule"},
		{field: "CONF_RANK", want: "float64",
			note: "\"_rank\" is itself in the stat-abbreviation list, so this never reaches the dedicated rank rule below"},
		{field: "DIV_RANK", want: "float64"},

		// --- Textual range buckets contain "range" and are caught by the
		//     sequence/period/range rule, typed int even though the actual
		//     values are text buckets like "24-22" or "Very Tight". ---
		{field: "SHOT_CLOCK_RANGE", want: "int", knownWrong: true,
			note: "real committed type in playerdashptshots.go; \"24-22\" becomes 0 via toInt"},
		{field: "CLOSE_DEF_DIST_RANGE", want: "int", knownWrong: true,
			note: "same bug class as SHOT_CLOCK_RANGE"},
		{field: "DRIBBLE_RANGE", want: "int", knownWrong: true,
			note: "same bug class as SHOT_CLOCK_RANGE"},

		// --- Age -> int (correct) ---
		{field: "PLAYER_AGE", want: "int"},
		{field: "AGE", want: "int"},

		// --- Dedicated rank rule -> int, only reached when "_rank" (with
		//     the underscore) isn't already caught by the stat-abbreviation
		//     loop above. ---
		{field: "RANK", want: "int"},

		// --- Sequence/period (non-range, non-season) -> int (correct) ---
		{field: "GAME_SEQUENCE", want: "int"},
		{field: "PERIOD", want: "int"},

		// --- Default fallback -> string ---
		{field: "SOME_UNRECOGNIZED_FIELD", want: "string"},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			got := inferGoType(tt.field)
			if got != tt.want {
				t.Errorf("inferGoType(%q) = %q, want %q (%s)", tt.field, got, tt.want, tt.note)
			}
		})
	}
}

// TestFieldTypesOverridesKnownWrongInference verifies that fieldtypes.json
// - the explicit, hand-reviewed dictionary resolveFieldGoType consults
// before ever falling back to inferGoType - actually corrects every
// knownWrong case
// documented in TestInferGoType above (plus other confirmed instances of
// the same bug classes found across the committed metadata: PCT_<STAT>
// "share of team total" fields mistyped by the FGM/FGA suffix sub-rule,
// FG_PCT_MID_RANGE mistyped by the _RANGE substring rule despite being a
// percentage rather than a text bucket, and GENERALMANAGER mistyped int
// because "generalmanager" contains the substring "age"). This is the test
// that proves the override mechanism actually fixes the corruption bugs
// inferGoType alone does not.
func TestFieldTypesOverridesKnownWrongInference(t *testing.T) {
	corrected := map[string]string{
		"CLOSE_DEF_DIST_RANGE":     "string",
		"DRIBBLE_RANGE":            "string",
		"SHOT_CLOCK_RANGE":         "string",
		"SHOT_DIST_RANGE":          "string",
		"SHOT_ZONE_RANGE":          "string",
		"TOUCH_TIME_RANGE":         "string",
		"FG_PCT_MID_RANGE":         "float64",
		"DISPLAY_FIRST_LAST":       "string",
		"DISPLAY_FI_LAST":          "string",
		"DISPLAY_LAST_COMMA_FIRST": "string",
		"PLAYER_NAME_LAST_FIRST":   "string",
		"FG_PCT_ABOVE_BREAK_3":     "float64",
		"FG_PCT_BACKCOURT":         "float64",
		"FG_PCT_IN_PAINT":          "float64",
		"FG_PCT_LEFT_CORNER_3":     "float64",
		"FG_PCT_RA":                "float64",
		"FG_PCT_RIGHT_CORNER_3":    "float64",
		"PCT":                      "float64",
		"WinPCT":                   "float64",
		"PCT_AST_2PM":              "float64",
		"PCT_AST_3PM":              "float64",
		"PCT_AST_FGM":              "float64",
		"PCT_BLKA":                 "float64",
		"PCT_FG3A":                 "float64",
		"PCT_FG3M":                 "float64",
		"PCT_FGA":                  "float64",
		"PCT_FGM":                  "float64",
		"PCT_FTA":                  "float64",
		"PCT_FTM":                  "float64",
		"PCT_UAST_2PM":             "float64",
		"PCT_UAST_3PM":             "float64",
		"PCT_UAST_FGM":             "float64",
		"PERCENTILE":               "float64",
		"GENERALMANAGER":           "string",
	}

	for field, want := range corrected {
		t.Run(field, func(t *testing.T) {
			if inferGoType(field) == want {
				t.Fatalf("inferGoType(%q) already returns %q - this case is no longer knownWrong, remove it from the correction map instead of leaving a redundant assertion", field, want)
			}
			got := resolveFieldGoType("", "", field)
			if got != want {
				t.Errorf("resolveFieldGoType(%q) = %q, want %q (fieldtypes.json entry missing or incorrect)", field, got, want)
			}
		})
	}
}

// TestFieldTypesMatchCommittedConsensus documents 14 fieldtypes.json
// corrections found by a different method than TestFieldTypesOverridesKnownWrongInference
// above: cross-referencing every field name against the Go type actually
// used for it across all 143 committed pkg/stats/endpoints/*.go files, and
// adopting the unanimous committed type where fieldtypes.json (inheriting
// inferGoType's unreviewed default) disagreed with it. Each of these fields
// has exactly one Go type across every occurrence in committed code -
// distinct from fields like OREB/REB/AST/PTS, which are float64 in ~90
// committed endpoints (per-game averages) but int in exactly 8 box-score
// endpoints (single-game counts): that's a real per-endpoint semantic
// difference a flat name->type dictionary cannot represent, not a case
// where one side is simply wrong, and is intentionally left alone here
// pending a per-endpoint override mechanism.
func TestFieldTypesMatchCommittedConsensus(t *testing.T) {
	corrected := map[string]string{
		"CONF_COUNT":           "int",
		"DIV_COUNT":            "int",
		"EVENTMSGACTIONTYPE":   "int",
		"EVENTMSGTYPE":         "int",
		"EVENTNUM":             "int",
		"PERSON1TYPE":          "int",
		"PERSON2TYPE":          "int",
		"PERSON3TYPE":          "int",
		"PLAYER_LAST":          "string",
		"PO_LOSSES":            "int",
		"PO_WINS":              "int",
		"PT_DIFF":              "float64",
		"TO":                   "int",
		"VIDEO_AVAILABLE_FLAG": "int",
	}

	for field, want := range corrected {
		t.Run(field, func(t *testing.T) {
			if inferGoType(field) == want {
				t.Fatalf("inferGoType(%q) already returns %q - this correction is redundant, remove it", field, want)
			}
			got := resolveFieldGoType("", "", field)
			if got != want {
				t.Errorf("resolveFieldGoType(%q) = %q, want %q (fieldtypes.json entry missing or incorrect)", field, got, want)
			}
		})
	}
}

// TestAllMetadataFieldsHaveExplicitTypes ensures every field name
// referenced by a committed metadata/*.json file has an explicit entry in
// fieldtypes.json, so generation never silently falls back to inferGoType
// for a field nobody has reviewed. A new endpoint's metadata should add its
// new field names to fieldtypes.json (verified against a live response) in
// the same change, not rely on the fallback.
func TestAllMetadataFieldsHaveExplicitTypes(t *testing.T) {
	types, err := loadFieldTypes()
	if err != nil {
		t.Fatalf("loadFieldTypes() failed: %v", err)
	}

	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}
	if len(metadataFiles) == 0 {
		t.Fatal("no metadata files found under metadata/ - this test would trivially pass without exercising anything")
	}

	missing := map[string][]string{} // field -> metadata files referencing it
	for _, mf := range metadataFiles {
		data, err := os.ReadFile(mf)
		if err != nil {
			t.Fatalf("failed to read %s: %v", mf, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", mf, err)
		}
		for _, ep := range endpoints {
			for _, rs := range ep.ResultSets {
				for _, field := range rs.Fields {
					if _, ok := types[field]; !ok {
						missing[field] = append(missing[field], mf)
					}
				}
			}
		}
	}

	if len(missing) > 0 {
		fields := make([]string, 0, len(missing))
		for f := range missing {
			fields = append(fields, f)
		}
		sort.Strings(fields)
		for _, f := range fields {
			t.Errorf("field %q referenced by %v has no entry in fieldtypes.json", f, missing[f])
		}
	}
}

// TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint proves
// fieldtype_overrides.json actually changes resolution where it's
// declared and leaves the global fieldtypes.json default untouched
// everywhere else. OREB is the motivating case: fieldtypes.json says
// float64 (correct for the ~90 per-game-average endpoints that don't
// override it), but BoxScoreTraditionalV2.PlayerStats overrides it to int
// (correct for that single-game box score). If this test passed with the
// override applying globally, every non-box-score endpoint's OREB would
// silently corrupt back to int.
func TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint(t *testing.T) {
	globalDefault := resolveFieldGoType("", "", "OREB")
	if globalDefault != "float64" {
		t.Fatalf("resolveFieldGoType(_, _, %q) = %q, want %q (fieldtypes.json's global default) - has fieldtypes.json changed?", "OREB", globalDefault, "float64")
	}

	overridden := resolveFieldGoType("BoxScoreTraditionalV2", "PlayerStats", "OREB")
	if overridden != "int" {
		t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want %q (fieldtype_overrides.json entry)", "BoxScoreTraditionalV2", "PlayerStats", "OREB", overridden, "int")
	}

	// Same field, same endpoint, but a result set with no override entry:
	// must fall through to the global default, not leak the sibling
	// result set's override.
	unrelatedResultSet := resolveFieldGoType("BoxScoreTraditionalV2", "SomeOtherResultSetNotInOverrides", "OREB")
	if unrelatedResultSet != globalDefault {
		t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want the global default %q - override leaked to an unlisted result set", "BoxScoreTraditionalV2", "SomeOtherResultSetNotInOverrides", "OREB", unrelatedResultSet, globalDefault)
	}

	// Same field, an endpoint with no override entries at all: must fall
	// through to the global default. PlayerDashboardByGeneralSplits is a
	// real metadata-covered endpoint whose OverallPlayerDashboard result
	// set contains OREB as a per-game average, not a count.
	unrelatedEndpoint := resolveFieldGoType("PlayerDashboardByGeneralSplits", "OverallPlayerDashboard", "OREB")
	if unrelatedEndpoint != globalDefault {
		t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want the global default %q - override leaked to an unrelated endpoint", "PlayerDashboardByGeneralSplits", "OverallPlayerDashboard", "OREB", unrelatedEndpoint, globalDefault)
	}

	// TeamYearByYearStats.TeamStats and TeamGameLogs.TeamGameLogs carry a
	// second, larger set of overrides (season/game totals, not just the
	// OREB family) because every stat field in those two result sets is a
	// whole-number total in committed code, not a per-game average -
	// confirmed by TEAM_ID/GP/FGM/FGA also being int there, not just the
	// fields fieldtypes.json alone would get wrong.
	for _, tc := range []struct{ endpoint, resultSet, field string }{
		{"TeamYearByYearStats", "TeamStats", "WINS"},
		{"TeamYearByYearStats", "TeamStats", "LOSSES"},
		{"TeamYearByYearStats", "TeamStats", "TOV"},
		{"TeamGameLogs", "TeamGameLogs", "TOV"},
		{"TeamGameLogs", "TeamGameLogs", "DD2"},
		{"LeagueGameFinder", "LeagueGameFinderResults", "TOV"},
	} {
		if got := resolveFieldGoType(tc.endpoint, tc.resultSet, tc.field); got != "int" {
			t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want %q", tc.endpoint, tc.resultSet, tc.field, got, "int")
		}
	}
}

// TestFieldTypeOverridesReferenceRealMetadata ensures every
// (endpoint, result set, field) entry in fieldtype_overrides.json
// actually exists in committed metadata - the endpoint name matches some
// EndpointMetadata.Name, the result set name matches one of its
// ResultSets, and the field is in that result set's Fields. Without this,
// a typo'd or stale override entry (e.g. after a metadata file is
// renamed or a result set's fields change) would silently do nothing
// instead of failing loudly.
func TestFieldTypeOverridesReferenceRealMetadata(t *testing.T) {
	overrides, err := loadFieldTypeOverrides()
	if err != nil {
		t.Fatalf("loadFieldTypeOverrides() failed: %v", err)
	}
	if len(overrides) == 0 {
		t.Fatal("fieldtype_overrides.json is empty - this test would trivially pass without exercising anything")
	}

	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}

	// endpoint -> result set -> set of fields, aggregated across every
	// metadata file (some endpoints are duplicated verbatim across more
	// than one committed file).
	realFields := map[string]map[string]map[string]bool{}
	for _, mf := range metadataFiles {
		data, err := os.ReadFile(mf)
		if err != nil {
			t.Fatalf("failed to read %s: %v", mf, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", mf, err)
		}
		for _, ep := range endpoints {
			for _, rs := range ep.ResultSets {
				if realFields[ep.Name] == nil {
					realFields[ep.Name] = map[string]map[string]bool{}
				}
				if realFields[ep.Name][rs.Name] == nil {
					realFields[ep.Name][rs.Name] = map[string]bool{}
				}
				for _, field := range rs.Fields {
					realFields[ep.Name][rs.Name][field] = true
				}
			}
		}
	}

	for endpointName, resultSets := range overrides {
		for resultSetName, fields := range resultSets {
			for field := range fields {
				if !realFields[endpointName][resultSetName][field] {
					t.Errorf("fieldtype_overrides.json has %s.%s.%s, but no committed metadata file has a %q endpoint with a %q result set containing field %q",
						endpointName, resultSetName, field, endpointName, resultSetName, field)
				}
			}
		}
	}
}
