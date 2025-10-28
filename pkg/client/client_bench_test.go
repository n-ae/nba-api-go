package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func BenchmarkClientGet(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {"value": 42}}`))
	}))
	defer server.Close()

	client := NewClient(Config{
		BaseURL: server.URL,
	})

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Get(ctx, "/test", nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClientGetWithParams(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := NewClient(Config{
		BaseURL: server.URL,
	})

	params := url.Values{
		"param1": []string{"value1"},
		"param2": []string{"value2"},
		"param3": []string{"value3"},
	}

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Get(ctx, "/test", params)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildURL(b *testing.B) {
	client := NewClient(Config{
		BaseURL: "https://api.example.com",
	})

	params := url.Values{
		"param1": []string{"value1"},
		"param2": []string{"value2"},
		"param3": []string{"value3"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.buildURL("/test", params)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSortParams(b *testing.B) {
	client := NewClient(Config{
		BaseURL: "https://api.example.com",
	})

	params := url.Values{
		"z": []string{"last"},
		"a": []string{"first"},
		"m": []string{"middle"},
		"d": []string{"fourth"},
		"p": []string{"fifth"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.sortParams(params)
	}
}
