# NBA API Go - Code Generator

This tool generates Go endpoint code from metadata about NBA.com API endpoints.

## Overview

The code generator helps automate the creation of endpoint wrappers for the 139+ stats API endpoints. It generates:
- Request structs with typed parameters
- Response structs for result sets
- Endpoint functions with validation
- Parsing functions for data conversion

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

- `-endpoint <name>` - Generate a single endpoint
- `-metadata <file>` - Generate from metadata JSON file
- `-output <dir>` - Output directory (default: `<repo-root>/pkg/stats/endpoints`, resolved from this file's own location so it's correct regardless of the working directory you run `go run .` from)
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
non-standard abbreviations. `VideoEvents`' `vl`/`vt`/`gc`/`surl`/`durl`/
`vurl`/`purl` fields are hand-committed as fully-uppercase
(`VL`/`VT`/.../`PURL`) despite not being recognized initialisms by any
standard convention - regenerating that file today would produce
`Vl`/`Vt`/etc. instead, so it stays excluded from regeneration rather
than force a guess into the initialisms list that wouldn't generalize.

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
- `endpoint.tmpl` - Main endpoint template
- `request.tmpl` - Request struct template (future)
- `response.tmpl` - Response struct template (future)
- `example.tmpl` - Example code template (future)

## Development Workflow

1. **Extract metadata** from Python nba_api
2. **Review metadata** for accuracy
3. **Generate code** with `-dry-run` first
4. **Verify output** looks correct
5. **Generate files** without dry-run
6. **Add parsing logic** for result sets
7. **Add tests** for the endpoint
8. **Create example** usage code

## Roadmap

### Current (v0.1)
- [x] Basic template structure
- [x] Single endpoint generation
- [x] Metadata-driven generation
- [x] Command-line interface

### Future Enhancements
- [ ] Automatic parsing function generation
- [ ] Result set struct generation
- [ ] Test skeleton generation
- [ ] Example code generation
- [ ] Integration with Python analyzer
- [ ] Batch generation for all 139 endpoints
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
go run . -endpoint PlayerStats -output /tmp/generated
```

## Architecture

```
tools/generator/
├── main.go              # CLI entry point
├── generator.go         # Core generator logic
├── analyzer.go          # Python endpoint analyzer (future)
├── templates/
│   ├── endpoint.tmpl    # Endpoint function template
│   ├── request.tmpl     # Request struct template (future)
│   └── response.tmpl    # Response struct template (future)
├── metadata/
│   └── endpoints.json   # Endpoint metadata database
└── README.md            # This file
```

## Contributing

When adding new templates or features:
1. Test with `-dry-run` first
2. Verify generated code compiles
3. Ensure code follows project style
4. Update this README

## Notes

- Generated code is a **starting point**, not production-ready
- Always review and customize generated code
- Add proper error handling and validation
- Write tests for generated endpoints
- Update documentation

## See Also

- [ADR 001](../../docs/adr/001-go-replication-strategy.md) - Architecture decisions
- [CONTRIBUTING.md](../../CONTRIBUTING.md) - Contribution guidelines
- [Python nba_api](https://github.com/swar/nba_api) - Source of endpoint information
