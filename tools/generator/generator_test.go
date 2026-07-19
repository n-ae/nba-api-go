package main

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
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
