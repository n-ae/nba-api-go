//go:build integration
// +build integration

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// Integration tests for the HTTP API server
// Run with: go test -tags=integration ./cmd/nba-api-server/...

func TestAPIServerIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	logger := log.New(os.Stdout, "[integration-test] ", log.LstdFlags)
	server := NewServer(logger)

	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	t.Run("HealthCheck", func(t *testing.T) {
		resp, err := client.Get(ts.URL + "/health")
		if err != nil {
			t.Fatalf("failed to call health endpoint: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		var health map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if status := health["status"]; status != "healthy" {
			t.Errorf("expected healthy status, got %v", status)
		}

		t.Logf("Health check: %+v", health)
	})

	t.Run("CORSHeaders", func(t *testing.T) {
		resp, err := client.Get(ts.URL + "/health")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if origin := resp.Header.Get("Access-Control-Allow-Origin"); origin != "*" {
			t.Errorf("expected CORS origin *, got %s", origin)
		}

		t.Logf("CORS headers: %v", resp.Header)
	})

	t.Run("UnknownEndpoint", func(t *testing.T) {
		resp, err := client.Get(ts.URL + "/unknown")
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", resp.StatusCode)
		}
	})
}

func TestPlayerGameLogEndpoint(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	logger := log.New(os.Stdout, "[integration-test] ", log.LstdFlags)
	server := NewServer(logger)

	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	t.Run("ValidRequest", func(t *testing.T) {
		url := fmt.Sprintf("%s/api/v1/stats/playergamelog?PlayerID=203999&Season=2023-24&SeasonType=Regular%%20Season", ts.URL)

		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("failed to call endpoint: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Logf("Response body: %s", body)
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if !result["success"].(bool) {
			t.Error("expected success=true")
		}

		t.Logf("Response keys: %v", getKeys(result))
	})

	t.Run("MissingPlayerID", func(t *testing.T) {
		url := fmt.Sprintf("%s/api/v1/stats/playergamelog?Season=2023-24", ts.URL)

		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("failed to call endpoint: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if result["success"].(bool) {
			t.Error("expected success=false for missing parameter")
		}
	})
}

func TestConcurrentRequests(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	logger := log.New(os.Stdout, "[integration-test] ", log.LstdFlags)
	server := NewServer(logger)

	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	const concurrency = 10
	errors := make(chan error, concurrency)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	for i := 0; i < concurrency; i++ {
		go func(id int) {
			resp, err := client.Get(ts.URL + "/health")
			if err != nil {
				errors <- fmt.Errorf("request %d failed: %v", id, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				errors <- fmt.Errorf("request %d got status %d", id, resp.StatusCode)
				return
			}

			errors <- nil
		}(i)
	}

	for i := 0; i < concurrency; i++ {
		select {
		case err := <-errors:
			if err != nil {
				t.Error(err)
			}
		case <-ctx.Done():
			t.Fatal("test timeout")
		}
	}

	t.Logf("Successfully handled %d concurrent requests", concurrency)
}

func TestServerTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	logger := log.New(os.Stdout, "[integration-test] ", log.LstdFlags)
	server := NewServer(logger)

	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	// Test with very short timeout
	client := &http.Client{
		Timeout: 1 * time.Millisecond,
	}

	_, err := client.Get(ts.URL + "/health")
	if err == nil {
		// Sometimes the request completes within 1ms, that's OK
		t.Log("Request completed within timeout (fast server)")
	} else {
		// Expected timeout error
		t.Logf("Got expected timeout: %v", err)
	}
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
