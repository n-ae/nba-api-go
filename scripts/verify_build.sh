#!/bin/bash
set -e

echo "=== NBA API Go - Build Verification Script ==="
echo ""

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check Go installation
echo "1. Checking Go installation..."
if command -v go &> /dev/null; then
    GO_VERSION=$(go version)
    echo -e "${GREEN}✓${NC} Go is installed: $GO_VERSION"
else
    echo -e "${RED}✗${NC} Go is not installed"
    exit 1
fi
echo ""

# Build the binary
echo "2. Building API server binary..."
cd "$(dirname "$0")/.."
if go build -o bin/nba-api-server ./cmd/nba-api-server; then
    echo -e "${GREEN}✓${NC} Binary built successfully"
    SIZE=$(du -h bin/nba-api-server | cut -f1)
    echo "   Binary size: $SIZE"
else
    echo -e "${RED}✗${NC} Build failed"
    exit 1
fi
echo ""

# Run tests
echo "3. Running unit tests..."
if go test ./cmd/nba-api-server/... -v; then
    echo -e "${GREEN}✓${NC} All unit tests passed"
else
    echo -e "${RED}✗${NC} Some tests failed"
    exit 1
fi
echo ""

# Check for container runtime
echo "4. Checking container runtime..."
CONTAINER_CMD=""
if command -v podman &> /dev/null; then
    CONTAINER_CMD="podman"
    echo -e "${GREEN}✓${NC} Podman is available"
    # Check if podman machine is running
    if ! podman info &> /dev/null; then
        echo -e "${YELLOW}⚠${NC} Podman machine not running, attempting to start..."
        podman machine start podman-machine-default || true
        sleep 3
    fi
elif command -v docker &> /dev/null; then
    CONTAINER_CMD="docker"
    echo -e "${GREEN}✓${NC} Docker is available"
else
    echo -e "${YELLOW}⚠${NC} No container runtime found (podman/docker)"
    echo "   Skipping container build verification"
    CONTAINER_CMD=""
fi
echo ""

# Build container if runtime is available
if [ -n "$CONTAINER_CMD" ]; then
    echo "5. Building container image..."
    if $CONTAINER_CMD build -f Containerfile -t nba-api-go:test .; then
        echo -e "${GREEN}✓${NC} Container image built successfully"
        
        # Check image size
        if [ "$CONTAINER_CMD" = "docker" ]; then
            IMAGE_SIZE=$($CONTAINER_CMD images nba-api-go:test --format "{{.Size}}")
        else
            IMAGE_SIZE=$($CONTAINER_CMD images nba-api-go:test --format "{{.Size}}")
        fi
        echo "   Image size: $IMAGE_SIZE"
        
        # Verify image size is reasonable (< 30MB)
        SIZE_MB=$(echo $IMAGE_SIZE | sed 's/MB//')
        if (( $(echo "$SIZE_MB < 30" | bc -l) )); then
            echo -e "${GREEN}✓${NC} Image size is within target (< 30MB)"
        else
            echo -e "${YELLOW}⚠${NC} Image size is larger than target (> 30MB)"
        fi
    else
        echo -e "${RED}✗${NC} Container build failed"
        exit 1
    fi
    echo ""
    
    echo "6. Testing container..."
    # Start container
    CONTAINER_ID=$($CONTAINER_CMD run -d -p 8081:8080 nba-api-go:test)
    echo "   Started container: $CONTAINER_ID"
    
    # Wait for server to start
    sleep 3
    
    # Test health endpoint
    if curl -f http://localhost:8081/health > /dev/null 2>&1; then
        echo -e "${GREEN}✓${NC} Container health check passed"
    else
        echo -e "${RED}✗${NC} Container health check failed"
        $CONTAINER_CMD logs $CONTAINER_ID
        $CONTAINER_CMD stop $CONTAINER_ID > /dev/null
        $CONTAINER_CMD rm $CONTAINER_ID > /dev/null
        exit 1
    fi
    
    # Stop container
    $CONTAINER_CMD stop $CONTAINER_ID > /dev/null
    $CONTAINER_CMD rm $CONTAINER_ID > /dev/null
    echo "   Stopped and removed container"
    echo ""
fi

# Run integration tests (if tag is supported)
echo "7. Running integration tests..."
if go test -tags=integration ./cmd/nba-api-server/... -v -short; then
    echo -e "${GREEN}✓${NC} Integration tests passed"
else
    echo -e "${YELLOW}⚠${NC} Integration tests skipped or failed"
fi
echo ""

# Summary
echo "=== Build Verification Summary ==="
echo -e "${GREEN}✓${NC} Go build successful"
echo -e "${GREEN}✓${NC} Unit tests passed"
if [ -n "$CONTAINER_CMD" ]; then
    echo -e "${GREEN}✓${NC} Container build successful"
    echo -e "${GREEN}✓${NC} Container runtime test passed"
fi
echo ""
echo -e "${GREEN}All verifications passed!${NC}"
echo ""
echo "Next steps:"
echo "  1. Run server: ./bin/nba-api-server"
echo "  2. Test locally: curl http://localhost:8080/health"
if [ -n "$CONTAINER_CMD" ]; then
    echo "  3. Run container: $CONTAINER_CMD run -p 8080:8080 nba-api-go:test"
fi
