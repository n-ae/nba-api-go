package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/n-ae/nba-api-go/v3/pkg/client"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

const version = "3.1.1"

var (
	buildTime = "unknown"
	gitCommit = "unknown"
)

func main() {
	port := getEnv("PORT", "8080")
	logLevel := getEnv("LOG_LEVEL", "info")

	logger := log.New(os.Stdout, "[nba-api] ", log.LstdFlags)
	logger.Printf("Starting NBA API Server v%s", version)
	logger.Printf("Log level: %s", logLevel)

	upstreamTimeout, err := getDurationEnv("NBA_API_TIMEOUT", client.DefaultTimeout)
	if err != nil {
		logger.Fatalf("Invalid NBA_API_TIMEOUT: %v", err)
	}
	server := NewServerWithOptions(logger, ServerOptions{
		NBAAPITimeout:   upstreamTimeout,
		CORSAllowOrigin: getEnv("CORS_ALLOW_ORIGIN", "*"),
	})

	healthCtx, cancelHealthCheck := context.WithCancel(context.Background())
	server.healthChecker.Start(healthCtx, time.Minute)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      server.Routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Printf("Server listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("Shutting down server...")
	cancelHealthCheck()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server stopped gracefully")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}
	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		return 0, fmt.Errorf("must be a positive Go duration, got %q", value)
	}
	return duration, nil
}

type Server struct {
	logger          *log.Logger
	statsHandler    *StatsHandler
	metrics         *Metrics
	rateLimiter     *RateLimiter
	healthChecker   *HealthChecker
	corsAllowOrigin string
}

func NewServer(logger *log.Logger) *Server {
	return NewServerWithOptions(logger, ServerOptions{})
}

type ServerOptions struct {
	NBAAPITimeout   time.Duration
	CORSAllowOrigin string
}

func NewServerWithOptions(logger *log.Logger, options ServerOptions) *Server {
	if options.NBAAPITimeout <= 0 {
		options.NBAAPITimeout = client.DefaultTimeout
	}
	if options.CORSAllowOrigin == "" {
		options.CORSAllowOrigin = "*"
	}
	rateLimiter := NewRateLimiter(100, 200)
	rateLimiter.CleanupOldLimiters(5 * time.Minute)

	// ServerOptions has no BaseURL field, so stats.NewClient always
	// constructs against the valid StatsBaseURL constant here - this can't
	// fail in practice, but propagating the error would require every
	// caller of NewServer/NewServerWithOptions to handle a startup error
	// that can't occur; panic instead, matching the "practically
	// unreachable" pattern already used in pkg/client for exactly this
	// kind of invariant.
	statsClient, err := stats.NewClient(stats.Config{Timeout: options.NBAAPITimeout})
	if err != nil {
		panic(fmt.Sprintf("nba-api-server: %v", err))
	}

	return &Server{
		logger:          logger,
		statsHandler:    NewStatsHandler(statsClient),
		metrics:         NewMetrics(),
		rateLimiter:     rateLimiter,
		healthChecker:   NewHealthChecker(statsClient),
		corsAllowOrigin: options.CORSAllowOrigin,
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth())
	mux.HandleFunc("/readyz", s.handleReadyz())
	mux.HandleFunc("/metrics", s.handleMetrics())
	mux.Handle("/api/v1/stats/", s.statsHandler)

	return s.metricsMiddleware(s.loggingMiddleware(s.rateLimiter.Middleware(s.corsMiddleware(mux))))
}

func (s *Server) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		s.metrics.RecordRequest(r.URL.Path, rec.statusCode, duration)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", s.corsAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleHealth() http.HandlerFunc {
	type healthResponse struct {
		Status         string            `json:"status"`
		Version        string            `json:"version"`
		BuildInfo      map[string]string `json:"build_info"`
		EndpointsCount map[string]int    `json:"endpoints_count"`
		Dependencies   map[string]string `json:"dependencies"`
		NBAAPIStatus   string            `json:"nba_api_status"`
		Timestamp      int64             `json:"timestamp"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// /health is a local liveness check: it reports whether this
		// process is up and always answers 200, reading the upstream
		// status from the HealthChecker's cache rather than making a
		// live NBA.com request per probe. Use /readyz if you need a
		// probe that fails when the upstream dependency is degraded.
		nbaAPIStatus := s.healthChecker.Status()

		resp := healthResponse{
			Status:  "healthy",
			Version: version,
			BuildInfo: map[string]string{
				"go_version": runtime.Version(),
				"build_time": buildTime,
				"git_commit": gitCommit,
			},
			// sdk_total: distinct endpoint functions in pkg/stats/endpoints
			// (143 files minus the 2 shared helpers, dates.go and types.go).
			// http_exposed: distinct case labels in the StatsHandler switch
			// below; one of them ("playertrackingshotdashboard") is a legacy
			// alias that reaches the same SDK endpoint as
			// "playertrackingshootingefficiency", so http_exposed is one
			// higher than sdk_total rather than equal to it. Verified
			// 2026-07-19 - see docs/MAINTAINABILITY_ASSESSMENT_2026-07-19.md.
			EndpointsCount: map[string]int{
				"sdk_total":    141,
				"http_exposed": 142,
			},
			Dependencies: map[string]string{
				"nba_api": "stats.nba.com",
			},
			NBAAPIStatus: nbaAPIStatus,
			Timestamp:    time.Now().Unix(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding health response: %v", err)
		}
	}
}

func (s *Server) handleMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snapshot := s.metrics.GetSnapshot()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(snapshot); err != nil {
			log.Printf("Error encoding metrics response: %v", err)
		}
	}
}

// handleReadyz reports whether this instance can currently serve NBA data.
// Unlike /health it can return a non-200 status: 503 when the cached
// upstream status is degraded, so a load balancer or orchestrator can stop
// routing traffic here without restarting the process. A status of
// nbaAPIStatusUnknown (no background probe has completed yet) counts as
// ready, so readiness doesn't fail during the brief window after startup.
func (s *Server) handleReadyz() http.HandlerFunc {
	type readyzResponse struct {
		Ready        bool   `json:"ready"`
		NBAAPIStatus string `json:"nba_api_status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		nbaAPIStatus := s.healthChecker.Status()
		ready := nbaAPIStatus != nbaAPIStatusDegraded

		resp := readyzResponse{
			Ready:        ready,
			NBAAPIStatus: nbaAPIStatus,
		}

		w.Header().Set("Content-Type", "application/json")
		if ready {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding readyz response: %v", err)
		}
	}
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	type errorResponse struct {
		Success bool `json:"success"`
		Error   struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	resp := errorResponse{Success: false}
	resp.Error.Code = code
	resp.Error.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
