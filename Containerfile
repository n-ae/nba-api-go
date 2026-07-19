# Multi-stage build for minimal final image
# Build stage
FROM golang:1.26-alpine AS builder

# Pass the git short hash and build time in from the host (via
# --build-arg) rather than computing them from a .git directory inside the
# image - keeps the build context free of repo history and works the same
# whether or not .git happens to be present in the context.
ARG GIT_COMMIT=unknown
ARG BUILD_TIME=unknown

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags="-w -s -X main.gitCommit=${GIT_COMMIT} -X main.buildTime=${BUILD_TIME}" \
    -o nba-api-server ./cmd/nba-api-server

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/nba-api-server .

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080
ENV LOG_LEVEL=info

# Run as non-root user
RUN adduser -D -u 1000 appuser
USER appuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./nba-api-server"]
