# NBA API Go - Code Generator

This tool generates Go endpoint code from metadata about NBA.com API endpoints.

## Overview

The code generator helps automate the creation of endpoint wrappers for all 141 stats API endpoints. From one metadata file, it generates:
- Request structs with typed parameters
- Response structs for result sets
- Endpoint functions with validation
- Parsing functions for data conversion
- The endpoint's `cmd/nba-api-server` HTTP handler (`generated_<name>.go`) and its entry in the route dispatch table (`generated_dispatch.go`)
- A response-parsing test (`generated_<name>_test.go` under `pkg/stats/endpoints/`), synthesized from the endpoint's own result-set field names/types

Handler and test generation happen automatically alongside SDK generation for any `-endpoint`/`-metadata` run - there's no separate step to remember. See "Options" below for `-all-handlers`, which regenerates every handler and the dispatch table in one pass (needed after adding a brand-new endpoint, so its route actually gets wired in).

## Usage

### Generate Single Endpoint

```bash
cd tools/generator
go run . -endpoint PlayerGameLog
```

### Generate from Metadata File

```bash
go run . -metadata endpoints.json
```

### Dry Run (Print Without Writing)

```bash
go run . -endpoint TeamInfoCommon -dry-run
```

### Options

- `-endpoint <name>` - Generate a single endpoint's SDK code, HTTP handler, and parsing test
- `-metadata <file>` - Generate from metadata JSON file (same three outputs, for every entry in the file)
- `-output <dir>` - Output directory for SDK code (default: `<repo-root>/pkg/stats/endpoints`, resolved from this file's own location so it's correct regardless of the working directory you run `go run .` from)
- `-server-output <dir>` - Output directory for generated HTTP handler files (default: `<repo-root>/cmd/nba-api-server`, resolved the same way as `-output`)
- `-all-handlers` - Regenerate every endpoint's HTTP handler plus the dispatch table (`cmd/nba-api-server/generated_dispatch.go`) from every `metadata/*.json` file - run this after adding a new endpoint's metadata so its route is actually wired into the server
- `-dry-run` - Print generated code without writing files

## Metadata Format

The generator expects JSON metadata in this format:

```json
[
  {
    "name": "PlayerGameLog",
    "endpoint": "playergamelog",
    "parameters": [
      {
        "name": "PlayerID",
        "type": "string",
        "required": true
      },
      {
        "name": "Season",
        "type": "Season",
        "required": false,
        "default": ""
      }
    ],
    "result_sets": [
      {
        "name": "PlayerGameLog",
        "fields": ["SEASON_ID", "Player_ID", "Game_ID", "GAME_DATE", ...]
      }
    ]
  }
]
```

## Field Types

Result-set field types come from `fieldtypes.json` - a flat, hand-reviewed
`{"FIELD_NAME": "goType"}` dictionary embedded into the generator binary.
This is the canonical, explicit source of truth for what Go type a given
NBA API field name maps to (`string`, `int`, or `float64`).

`inferGoType` in `generator.go` still exists as a **fallback only**, used
when a field name has no entry in `fieldtypes.json`. Its name-pattern
heuristic gets some field families wrong in ways that silently corrupt
data - see `TestInferGoType` in `generator_test.go` for the documented,
confirmed-wrong cases (e.g. `DISPLAY_FIRST_LAST` inferred as `float64`
turns `"Nikola Jokić"` into `0`; `SHOT_CLOCK_RANGE` inferred as `int` turns
`"24-22"` into `0`). Falling back prints a warning to stderr - if you see
one, it means a field is running on an unreviewed guess.

`TestAllMetadataFieldsHaveExplicitTypes` (in `generator_test.go`) fails CI
if any field referenced by a committed `metadata/*.json` file has no
`fieldtypes.json` entry, so this can't silently regress.

**Adding a new field**: verify its true type against a live (or recorded)
NBA.com response - not by trusting `inferGoType`'s guess - then add it to
`fieldtypes.json`.

### Per-endpoint overrides

Some field names mean different things in different endpoints. `OREB`
(offensive rebounds) is `float64` in `fieldtypes.json` because it's a
per-game average in the vast majority of endpoints that have it - but in a
handful of box-score/game-log endpoints it's a single-game count, and
should be `int`. A flat `{"FIELD_NAME": "goType"}` dictionary cannot
represent both; picking either as the global default silently breaks the
other case.

`fieldtype_overrides.json` holds these exceptions, keyed by endpoint name,
then result-set name, then field name:

```json
{
  "BoxScoreTraditionalV2": {
    "PlayerStats": {
      "OREB": "int"
    }
  }
}
```

`resolveFieldGoType` in `generator.go` checks this file first, then
`fieldtypes.json`, then falls back to `inferGoType`. An override only
applies to the exact `(endpoint, result set, field)` triple it names -
it does not affect the same field name anywhere else, including other
result sets within the same endpoint.

`TestFieldTypeOverridesReferenceRealMetadata` fails CI if an override
entry's endpoint/result-set/field doesn't match anything in committed
metadata (catches typos and stale entries after a metadata file changes).
`TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint` proves an override
doesn't leak beyond the triple it's declared for.

**Adding an override**: only do this when you've confirmed (e.g. by
comparing committed struct definitions across endpoints, or against live
responses) that the same field name genuinely needs different types in
different places - not as a shortcut for a `fieldtypes.json` correction
that should apply everywhere.

## Field Names (camelCase metadata)

Most metadata field names are `SCREAMING_SNAKE_CASE` (the classic Stats
API convention, e.g. `"GAME_ID"`) and already valid, exported Go
identifiers - the template uses them as-is. NBA Live-Data-style endpoints
(`PlayByPlayV3`, `ScoreboardV3`, `LeagueStandings`, ...) instead use
camelCase field names starting with a lowercase letter (e.g.
`"gameId"`). Using that directly as a Go struct field name would produce
an **unexported** field - invisible to `encoding/json`, inaccessible to
any external caller. `goFieldName` in `generator.go` fixes this: it
capitalizes each camelCase word and fully uppercases recognized
initialisms (`ID`, `URL`, `UUID`, etc.), so `"gameId"` becomes `GameID`,
matching Go convention. The original field name is untouched as the
JSON tag - only the Go identifier changes.

This is deliberately conservative: it does not try to guess at
non-standard abbreviations globally. `VideoEvents`' `vl`/`vt`/`gc`/`surl`/
`durl`/`vurl`/`purl` fields are hand-committed as fully-uppercase
(`VL`/`VT`/.../`PURL`) despite not being recognized initialisms by any
standard convention - `goFieldName` alone would produce `Vl`/`Vt`/etc.
instead. Rather than add these to the global initialisms list (which
would silently affect any future field literally named `vl`/`gc`/etc. for
something else), `tools/generator/fieldname_overrides.json` holds them as
a narrow, `(endpoint, result set, field) -> Go identifier` exception,
mirroring `fieldtype_overrides.json`'s pattern:

```json
{
  "VideoEvents": {
    "Video": {
      "vl": "VL"
    }
  }
}
```

`resolveFieldGoType`'s type-override precedence has a name-override
equivalent: `goFieldName` checks `fieldname_overrides.json` before
falling back to the general capitalization rule.
`TestGoFieldNameOverridesReferenceRealMetadata` fails CI if an entry
doesn't match real committed metadata;
`TestGoFieldNameOverridesApplyOnlyWithinTheirEndpoint` proves an override
doesn't leak beyond its exact scope.

## Result Set Lookup and Header Validation

Generated code looks up each result set by its `name` field
(`findResultSet` in `pkg/stats/endpoints/types.go`), not by its position
in the response's `resultSets` array. It also validates the result set's
actual `headers` against the field order the generated struct assumes
before indexing any row positionally (`validateHeaders`), using the
struct's own `json` tags as the expected list (`jsonTags`) rather than a
second, separately-maintained copy of the field names.

This replaced `rawResp.ResultSets[0]`/`[1]`-style positional indexing
with no header check, which would have silently read the wrong result
set - or shifted every field after a column change into the wrong struct
field - if NBA.com ever reordered anything. A header mismatch now
returns an error instead of corrupting data silently.

**Not verified against live NBA.com responses.** If a field order has
drifted from what NBA.com currently returns, this surfaces as a new
error on upgrade rather than the old silent-wrong-data behavior - see
`CHANGELOG.md` for the full risk callout.

## Creating Metadata

### From Python nba_api

You can extract metadata from the Python nba_api library:

```python
from nba_api.stats.endpoints import playergamelog
import inspect
import json

endpoint = playergamelog.PlayerGameLog
metadata = {
    "name": "PlayerGameLog",
    "endpoint": endpoint.endpoint,
    "parameters": [],
    "result_sets": []
}

# Extract parameters from __init__ signature
sig = inspect.signature(endpoint.__init__)
for param_name, param in sig.parameters.items():
    if param_name not in ['self', 'proxy', 'headers', 'timeout', 'get_request']:
        metadata["parameters"].append({
            "name": param_name,
            "type": "string",
            "required": param.default == inspect.Parameter.empty
        })

# Extract expected data structure
metadata["result_sets"] = [
    {
        "name": key,
        "fields": fields
    }
    for key, fields in endpoint.expected_data.items()
]

print(json.dumps(metadata, indent=2))
```

### Manual Creation

For complex endpoints, manually create metadata:

1. Find the endpoint in Python nba_api
2. Note the endpoint name (lowercase)
3. List all parameters from `__init__`
4. List all result sets from `expected_data`
5. Create JSON following the format above

## Template Customization

Edit templates in `templates/` directory:
- `endpoint.tmpl` - SDK request/response structs, endpoint function, and parsing logic
- `handler.tmpl` - The endpoint's `cmd/nba-api-server` HTTP handler
- `dispatch.tmpl` - The server's route dispatch table (`generated_dispatch.go`), rendered once against every endpoint's metadata combined
- `endpoint_test.tmpl` - The endpoint's response-parsing test, with a fixture synthesized from its own result-set field names/types

## Development Workflow

1. **Get metadata** - extract from Python nba_api (see "Creating Metadata" above) or write by hand
2. **Review metadata** for accuracy
3. **Generate code** with `-dry-run` first
4. **Verify output** looks correct
5. **Generate files** without dry-run - this produces the SDK code, the HTTP handler, and a parsing test in one pass
6. **Run `-all-handlers`** so the new route is wired into the dispatch table
7. **Add each new field's verified type** to `fieldtypes.json` if `inferGoType` had to fall back (see "Field Types" above) - `TestAllMetadataFieldsHaveExplicitTypes` fails CI otherwise
8. **Create example** usage code under `examples/`

## Roadmap

### Done
- [x] Basic template structure
- [x] Single endpoint generation
- [x] Metadata-driven generation
- [x] Command-line interface
- [x] Automatic parsing function generation
- [x] Result set struct generation
- [x] Batch generation for all 141 endpoints (`-metadata`/`-all-handlers`)
- [x] HTTP handler generation (`cmd/nba-api-server`), including the route dispatch table
- [x] Response-parsing test generation (`generated_<name>_test.go`)

### Future Enhancements
- [ ] Example code generation
- [ ] Integration with Python analyzer (see "Creating Metadata" above for the current manual extraction script)
- [ ] Documentation generation

## Examples

### Generate Multiple Endpoints

```bash
# Create metadata file with multiple endpoints
cat > endpoints.json << EOF
[
  {"name": "TeamInfoCommon", "endpoint": "teaminfocommon", ...},
  {"name": "BoxScoreSummaryV2", "endpoint": "boxscoresummaryv2", ...},
  {"name": "ShotChartDetail", "endpoint": "shotchartdetail", ...}
]
EOF

# Generate all at once
go run . -metadata endpoints.json
```

### Custom Output Directory

```bash
go run . -endpoint PlayerStats -output /tmp/generated -server-output /tmp/generated-handlers
```

### Regenerate All Handlers + Dispatch Table

```bash
# After adding a new endpoint's metadata, wire its route into the server:
go run . -all-handlers
```

## Architecture

```
tools/generator/
├── main.go                        # CLI entry point
├── generator.go                   # Core generator logic (SDK, handler, and test generation)
├── fieldtypes.json                # Field name -> Go type dictionary (see "Field Types" above)
├── fieldtype_overrides.json       # Per-(endpoint, result set, field) type exceptions
├── fieldname_overrides.json       # Per-(endpoint, result set, field) Go-identifier exceptions
├── templates/
│   ├── endpoint.tmpl              # SDK request/response structs + endpoint function
│   ├── handler.tmpl               # cmd/nba-api-server HTTP handler
│   ├── dispatch.tmpl              # cmd/nba-api-server route dispatch table
│   └── endpoint_test.tmpl         # Generated response-parsing test
├── metadata/
│   └── *.json                     # Endpoint metadata, one entry per endpoint (batched or per-file)
└── README.md                      # This file

# Generation targets (outside this directory), for any endpoint with
# metadata - the 6 hand-written endpoints have none, so they're untouched:
#   pkg/stats/endpoints/<name>.go                 - SDK code
#   pkg/stats/endpoints/generated_<name>_test.go  - parsing test
#   cmd/nba-api-server/generated_<name>.go        - HTTP handler
#   cmd/nba-api-server/generated_dispatch.go      - route dispatch table (-all-handlers only)
```

## Contributing

When adding new templates or features:
1. Test with `-dry-run` first
2. Verify generated code compiles
3. Ensure code follows project style
4. Update this README

## Notes

- **Generated files are not safe to hand-edit and later regenerate over** - a future regeneration from updated metadata will silently discard hand edits. If a generated file needs to differ from what the generator would currently produce, fix it via the metadata/templates/overrides, not by editing the generated `.go` file directly. (The `gamerotation`/`leaguedashplayerstats`/`videoevents` endpoints are permanently hand-excluded from regeneration for exactly this reason - see `CHANGELOG.md`'s `[2.0.0]` section.)
- Verify field types against a live (or recorded) NBA.com response before adding them to `fieldtypes.json` - don't trust `inferGoType`'s fallback guess
- Run `go test ./pkg/stats/endpoints/... ./cmd/nba-api-server/...` after regenerating to confirm the generated parsing test and handler test still pass
- Update this README when the generator's own capabilities change (templates, flags, generation targets) - out of date documentation here is exactly the kind of drift `TestAllMetadataFieldsHaveExplicitTypes` and friends can't catch

## See Also

- [ADR 001](../../docs/adr/001-go-replication-strategy.md) - Architecture decisions
- [CONTRIBUTING.md](../../CONTRIBUTING.md) - Contribution guidelines
- [Python nba_api](https://github.com/swar/nba_api) - Source of endpoint information
