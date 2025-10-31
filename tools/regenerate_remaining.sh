#!/bin/bash
# Batch Regeneration Script for Remaining Endpoints
# This script regenerates all endpoints that still have interface{} types

set -e

echo "🚀 Starting batch regeneration of remaining endpoints..."
echo ""

GENERATOR="./tools/generator/bin/generator"
METADATA_DIR="tools/generator/metadata"
OUTPUT_DIR="pkg/stats/endpoints"

# Build generator if needed
if [ ! -f "$GENERATOR" ]; then
    echo "📦 Building generator..."
    go build -o "$GENERATOR" ./tools/generator
    echo "✅ Generator built"
    echo ""
fi

# List of endpoints to regenerate
declare -a ENDPOINTS=(
    "boxscoresummaryv2"
    "shotchartdetail"
    "teamyearbyyearstats"
    "playerdashboardbygeneralsplits"
    "teamdashboardbygeneralsplits"
    "playbyplayv2"
    "teaminfocommon"
)

TOTAL=${#ENDPOINTS[@]}
CURRENT=0

echo "📋 Endpoints to regenerate: $TOTAL"
echo ""

for endpoint in "${ENDPOINTS[@]}"; do
    CURRENT=$((CURRENT + 1))
    echo "[$CURRENT/$TOTAL] Regenerating $endpoint..."

    METADATA_FILE="$METADATA_DIR/${endpoint}.json"

    if [ -f "$METADATA_FILE" ]; then
        $GENERATOR -metadata "$METADATA_FILE" -output "$OUTPUT_DIR"
        echo "  ✅ $endpoint regenerated"
    else
        echo "  ⚠️  Metadata file not found: $METADATA_FILE"
        echo "  Checking batch files..."

        # Try batch2_endpoints.json
        if grep -q "\"$endpoint\"" "$METADATA_DIR/batch2_endpoints.json" 2>/dev/null; then
            echo "  Found in batch2_endpoints.json - manual regeneration needed"
        fi
    fi

    echo ""
done

echo ""
echo "🎉 Batch regeneration complete!"
echo ""
echo "📝 Next steps:"
echo "  1. Run: go build ./pkg/stats/endpoints"
echo "  2. Run: go test ./pkg/stats/endpoints"
echo "  3. Review: git diff pkg/stats/endpoints/"
echo ""
echo "🔍 Verify no interface{} remains:"
echo "  grep -r 'interface{}' pkg/stats/endpoints/*.go | grep -v types.go | grep -v _test.go"
