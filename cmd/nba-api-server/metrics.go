package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type Metrics struct {
	mu                sync.RWMutex
	startTime         time.Time
	totalRequests     atomic.Int64
	totalErrors       atomic.Int64
	requestsByStatus  map[int]*atomic.Int64
	requestsByPath    map[string]*atomic.Int64
	responseTimes     []time.Duration
	maxResponseTimes  int
}

func NewMetrics() *Metrics {
	m := &Metrics{
		startTime:        time.Now(),
		requestsByStatus: make(map[int]*atomic.Int64),
		requestsByPath:   make(map[string]*atomic.Int64),
		responseTimes:    make([]time.Duration, 0, 1000),
		maxResponseTimes: 1000,
	}
	return m
}

func (m *Metrics) RecordRequest(path string, status int, duration time.Duration) {
	m.totalRequests.Add(1)

	if status >= 400 {
		m.totalErrors.Add(1)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.requestsByStatus[status]; !exists {
		m.requestsByStatus[status] = &atomic.Int64{}
	}
	m.requestsByStatus[status].Add(1)

	if _, exists := m.requestsByPath[path]; !exists {
		m.requestsByPath[path] = &atomic.Int64{}
	}
	m.requestsByPath[path].Add(1)

	if len(m.responseTimes) < m.maxResponseTimes {
		m.responseTimes = append(m.responseTimes, duration)
	}
}

func (m *Metrics) GetSnapshot() MetricsSnapshot {
	m.mu.RLock()
	defer m.mu.RUnlock()

	snapshot := MetricsSnapshot{
		Uptime:           time.Since(m.startTime).Seconds(),
		TotalRequests:    m.totalRequests.Load(),
		TotalErrors:      m.totalErrors.Load(),
		RequestsByStatus: make(map[int]int64),
		RequestsByPath:   make(map[string]int64),
	}

	for status, count := range m.requestsByStatus {
		snapshot.RequestsByStatus[status] = count.Load()
	}

	for path, count := range m.requestsByPath {
		snapshot.RequestsByPath[path] = count.Load()
	}

	if len(m.responseTimes) > 0 {
		var total time.Duration
		min := m.responseTimes[0]
		max := m.responseTimes[0]

		for _, d := range m.responseTimes {
			total += d
			if d < min {
				min = d
			}
			if d > max {
				max = d
			}
		}

		snapshot.AvgResponseTime = total / time.Duration(len(m.responseTimes))
		snapshot.MinResponseTime = min
		snapshot.MaxResponseTime = max
	}

	return snapshot
}

type MetricsSnapshot struct {
	Uptime           float64           `json:"uptime_seconds"`
	TotalRequests    int64             `json:"total_requests"`
	TotalErrors      int64             `json:"total_errors"`
	RequestsByStatus map[int]int64     `json:"requests_by_status"`
	RequestsByPath   map[string]int64  `json:"requests_by_path"`
	AvgResponseTime  time.Duration     `json:"avg_response_time_ns"`
	MinResponseTime  time.Duration     `json:"min_response_time_ns"`
	MaxResponseTime  time.Duration     `json:"max_response_time_ns"`
}
