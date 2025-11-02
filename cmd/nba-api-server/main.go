package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "0.1.0"

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
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		logger:       logger,
		statsHandler: NewStatsHandler(),
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth())
	mux.Handle("/api/v1/stats/", s.statsHandler)

	return s.loggingMiddleware(s.corsMiddleware(mux))
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
		Status         string `json:"status"`
		Version        string `json:"version"`
		EndpointsCount int    `json:"endpoints_count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":                 "healthy",
			"version":                version,
			"sdk_endpoints_total":    139,
			"http_endpoints_exposed": 139,
			"timestamp":              time.Now().Unix(),
		})
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
