package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
)

const version = "1.1.1"

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

	server := NewServer(logger)

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

type Server struct {
	logger       *log.Logger
	statsHandler *StatsHandler
	metrics      *Metrics
	rateLimiter  *RateLimiter
}

func NewServer(logger *log.Logger) *Server {
	rateLimiter := NewRateLimiter(100, 200)
	rateLimiter.CleanupOldLimiters(5 * time.Minute)

	return &Server{
		logger:       logger,
		statsHandler: NewStatsHandler(),
		metrics:      NewMetrics(),
		rateLimiter:  rateLimiter,
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth())
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		nbaAPIStatus := s.checkNBAAPI()

		resp := healthResponse{
			Status:  "healthy",
			Version: version,
			BuildInfo: map[string]string{
				"go_version": runtime.Version(),
				"build_time": buildTime,
				"git_commit": gitCommit,
			},
			EndpointsCount: map[string]int{
				"sdk_total":    140,
				"http_exposed": 149,
			},
			Dependencies: map[string]string{
				"nba_api": "stats.nba.com",
			},
			NBAAPIStatus: nbaAPIStatus,
			Timestamp:    time.Now().Unix(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func (s *Server) handleMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snapshot := s.metrics.GetSnapshot()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(snapshot)
	}
}

func (s *Server) checkNBAAPI() string {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client := stats.NewDefaultClient()

	req := endpoints.CommonAllPlayersRequest{
		Season: "2023-24",
	}

	_, err := endpoints.GetCommonAllPlayers(ctx, client, req)
	if err != nil {
		return "degraded"
	}

	return "operational"
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
