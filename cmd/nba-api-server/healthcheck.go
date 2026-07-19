package main

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
)

const (
	nbaAPIStatusUnknown     = "unknown"
	nbaAPIStatusOperational = "operational"
	nbaAPIStatusDegraded    = "degraded"
)

// HealthChecker probes the upstream NBA Stats API on a background ticker
// and caches the result, so /health and /readyz answer instantly from
// memory instead of making a live upstream request on every probe. Without
// this, a load balancer or container runtime hitting /health every few
// seconds turns health checking into a steady stream of NBA.com traffic
// that competes with real requests for the same per-host rate limit.
type HealthChecker struct {
	client *stats.Client
	status atomic.Value // string
}

func NewHealthChecker(client *stats.Client) *HealthChecker {
	hc := &HealthChecker{client: client}
	hc.status.Store(nbaAPIStatusUnknown)
	return hc
}

// Status returns the most recently observed upstream status. It never
// blocks on a network call.
func (hc *HealthChecker) Status() string {
	if status, ok := hc.status.Load().(string); ok {
		return status
	}
	return nbaAPIStatusUnknown
}

// Start launches the background polling loop and returns immediately; the
// first probe runs asynchronously rather than blocking startup. Status()
// reports nbaAPIStatusUnknown until that first probe completes. Callers
// should not invoke Start from test setup - constructing a Server/
// StatsHandler must stay network-free so the unit test suite runs offline.
func (hc *HealthChecker) Start(ctx context.Context, interval time.Duration) {
	go func() {
		hc.check(ctx)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				hc.check(ctx)
			}
		}
	}()
}

// check derives its request deadline from parent so that canceling parent
// (e.g. on server shutdown) aborts an in-flight probe immediately instead
// of leaving it to run for up to its own 3s timeout after the caller has
// already stopped waiting on it.
func (hc *HealthChecker) check(parent context.Context) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()

	req := endpoints.CommonAllPlayersRequest{Season: "2023-24"}
	if _, err := endpoints.GetCommonAllPlayers(ctx, hc.client, req); err != nil {
		hc.status.Store(nbaAPIStatusDegraded)
		return
	}
	hc.status.Store(nbaAPIStatusOperational)
}
