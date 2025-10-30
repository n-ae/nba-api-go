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
- `-output <dir>` - Output directory (default: pkg/stats/endpoints)
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
