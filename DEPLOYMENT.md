# Deployment Guide

This guide covers deploying the NBA API Go HTTP server to production.

## Quick Start

### Build

```bash
# Standard build
go build -o nba-api-server ./cmd/nba-api-server

# With build info
go build \
  -ldflags="-X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) -X main.gitCommit=$(git rev-parse --short HEAD)" \
  -o nba-api-server \
  ./cmd/nba-api-server
```

### Run

```bash
# Default (port 8080)
./nba-api-server

# Custom port
PORT=3000 ./nba-api-server

# With logging
LOG_LEVEL=debug ./nba-api-server
```

## Deployment Options

### 1. Systemd Service (Recommended for VPS)

Create `/etc/systemd/system/nba-api.service`:

```ini
[Unit]
Description=NBA API Go Server
After=network.target

[Service]
Type=simple
User=nba-api
WorkingDirectory=/opt/nba-api
Environment="PORT=8080"
Environment="LOG_LEVEL=info"
ExecStart=/opt/nba-api/nba-api-server
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl enable nba-api
sudo systemctl start nba-api
sudo systemctl status nba-api
```

### 2. Docker

Create `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o nba-api-server ./cmd/nba-api-server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/nba-api-server .
EXPOSE 8080
CMD ["./nba-api-server"]
```

Build and run:

```bash
# Build
docker build -t nba-api-go .

# Run
docker run -d \
  --name nba-api \
  -p 8080:8080 \
  -e LOG_LEVEL=info \
  --restart unless-stopped \
  nba-api-go
```

### 3. Podman (Rootless Alternative to Docker)

```bash
# Build
podman build -t nba-api-go .

# Run
podman run -d \
  --name nba-api \
  -p 8080:8080 \
  -e LOG_LEVEL=info \
  --restart unless-stopped \
  nba-api-go
```

### 4. Cloud Platforms

#### Fly.io

Create `fly.toml`:

```toml
app = "nba-api-go"

[build]
  builder = "paketobuildpacks/builder:base"
  buildpacks = ["gcr.io/paketo-buildpacks/go"]

[env]
  PORT = "8080"
  LOG_LEVEL = "info"

[[services]]
  internal_port = 8080
  protocol = "tcp"

  [[services.ports]]
    handlers = ["http"]
    port = 80

  [[services.ports]]
    handlers = ["tls", "http"]
    port = 443
```

Deploy:

```bash
fly deploy
```

#### Railway

```bash
# Install Railway CLI
npm install -g @railway/cli

# Deploy
railway login
railway init
railway up
```

Set environment variables in Railway dashboard:
- `PORT=8080`
- `LOG_LEVEL=info`

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port |
| `LOG_LEVEL` | `info` | Logging verbosity (debug, info, warn, error) |

## Monitoring

### Health Check

```bash
curl http://localhost:8080/health
```

Response includes:
- Server status
- NBA API connectivity status
- Build information
- Endpoint counts
- Timestamp

### Metrics

```bash
curl http://localhost:8080/metrics
```

Response includes:
- Uptime
- Total requests
- Error count
- Requests by status code
- Requests by path
- Response time statistics (avg, min, max)

### Monitoring with External Tools

#### Prometheus

The `/metrics` endpoint can be scraped by Prometheus. Add to `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'nba-api'
    static_configs:
      - targets: ['localhost:8080']
    metrics_path: '/metrics'
    scrape_interval: 30s
```

#### Uptime Monitoring

Use services like:
- **UptimeRobot**: Free, monitors `/health` endpoint
- **Better Uptime**: Advanced monitoring with incident management
- **Healthchecks.io**: Simple HTTP endpoint monitoring

Example UptimeRobot configuration:
- Monitor Type: HTTP(S)
- URL: `https://your-domain.com/health`
- Interval: 5 minutes

## Performance Tuning

### Rate Limiting

Default: 100 requests/second per IP, burst of 200

To adjust, modify `main.go`:

```go
rateLimiter := NewRateLimiter(50, 100) // 50 req/s, burst 100
```

### Timeouts

Configured in `main.go`:

```go
srv := &http.Server{
    ReadTimeout:  15 * time.Second,  // Time to read request
    WriteTimeout: 30 * time.Second,  // Time to write response
    IdleTimeout:  60 * time.Second,  // Keep-alive timeout
}
```

### NBA API Health Check Timeout

Health endpoint checks NBA API with 3-second timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
```

## Reverse Proxy

### Nginx

```nginx
server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Timeouts
        proxy_connect_timeout 30s;
        proxy_send_timeout 30s;
        proxy_read_timeout 30s;
    }

    # Health check endpoint (bypass rate limiting)
    location /health {
        proxy_pass http://localhost:8080/health;
        access_log off;
    }
}
```

### Caddy

```
api.yourdomain.com {
    reverse_proxy localhost:8080
}
```

## SSL/TLS

### Let's Encrypt with Certbot

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Obtain certificate
sudo certbot --nginx -d api.yourdomain.com

# Auto-renewal (already configured by certbot)
sudo certbot renew --dry-run
```

### Cloudflare

1. Add domain to Cloudflare
2. Set DNS record: A record pointing to server IP
3. Enable "Full (strict)" SSL/TLS mode
4. Enable "Always Use HTTPS"

## Security

### Firewall

```bash
# UFW (Ubuntu)
sudo ufw allow 22/tcp  # SSH
sudo ufw allow 80/tcp  # HTTP
sudo ufw allow 443/tcp # HTTPS
sudo ufw enable
```

### Best Practices

1. **Run as non-root user**
   ```bash
   sudo useradd -r -s /bin/false nba-api
   ```

2. **Limit file permissions**
   ```bash
   sudo chown nba-api:nba-api /opt/nba-api/nba-api-server
   sudo chmod 550 /opt/nba-api/nba-api-server
   ```

3. **Enable CORS only for trusted domains**
   - Modify `corsMiddleware` in `main.go`
   - Replace `"*"` with specific domains

4. **Monitor logs**
   ```bash
   sudo journalctl -u nba-api -f
   ```

## Cost Estimates

### VPS (Recommended for Solo Projects)

| Provider | Specs | Cost/Month | Notes |
|----------|-------|------------|-------|
| Hetzner | 1 vCPU, 2GB RAM | €4.51 (~$5) | Best value |
| DigitalOcean | 1 vCPU, 1GB RAM | $6 | Easy setup |
| Linode | 1GB RAM | $5 | Reliable |
| Vultr | 1GB RAM | $6 | Good network |

### Platform-as-a-Service

| Provider | Free Tier | Paid |
|----------|-----------|------|
| Fly.io | 3 VMs, 160GB transfer | $1.94/mo for always-on |
| Railway | 500 hours/mo | $5/mo + usage |
| Render | 750 hours/mo | $7/mo per instance |

**Recommendation**: Hetzner VPS ($5/mo) or Fly.io for minimal cost.

## Backup & Recovery

### State

This server is **stateless** - no database, no persistent storage needed.

### Configuration

Back up only:
- Environment variables
- Reverse proxy configuration
- Systemd service file (if applicable)

### Disaster Recovery

1. Rebuild binary: `go build ./cmd/nba-api-server`
2. Deploy to server
3. Restart service

Total recovery time: < 5 minutes

## Scaling

### Single Instance Performance

- **Handles**: ~10,000 req/min easily on 1 vCPU
- **Bottleneck**: NBA API rate limits, not this server
- **Memory**: < 100MB typically

### When to Scale

You **probably don't need to scale** because:
- NBA API has rate limits
- This server is very efficient
- Single instance handles thousands of req/min

If you must scale:

1. **Horizontal**: Multiple instances behind load balancer
2. **Vertical**: Upgrade to 2+ vCPU first (cheaper)

## Troubleshooting

### Server won't start

```bash
# Check port availability
sudo lsof -i :8080

# Check logs
sudo journalctl -u nba-api -n 50
```

### High error rate

```bash
# Check metrics
curl http://localhost:8080/metrics | jq '.total_errors'

# Check NBA API status
curl http://localhost:8080/health | jq '.nba_api_status'
```

### Memory issues

```bash
# Check memory usage
ps aux | grep nba-api-server

# Monitor over time
watch -n 1 'ps aux | grep nba-api-server'
```

## Recommended Setup (Solo Maintainer)

For minimal cost and maximum reliability:

1. **Hetzner VPS** ($5/mo)
   - 1 vCPU, 2GB RAM
   - Ubuntu 22.04 LTS

2. **Systemd service** (simple, reliable)

3. **Caddy reverse proxy** (auto SSL)

4. **UptimeRobot** (free monitoring)

5. **Total cost**: $5/month
6. **Maintenance time**: ~1 hour/month

This setup has run production apps for years with minimal intervention.

## Support

- Issues: https://github.com/n-ae/nba-api-go/issues
- Documentation: https://github.com/n-ae/nba-api-go
