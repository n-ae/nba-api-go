package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"
)

type httpClientFunc func(*http.Request) (*http.Response, error)

func (f httpClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestClient_Get(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       string
		params         url.Values
		serverResponse string
		serverStatus   int
		wantErr        bool
	}{
		{
			name:           "successful request",
			endpoint:       "/test",
			params:         nil,
			serverResponse: `{"success": true}`,
			serverStatus:   http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "404 error",
			endpoint:       "/notfound",
			params:         nil,
			serverResponse: `{"error": "not found"}`,
			serverStatus:   http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:     "with query parameters",
			endpoint: "/test",
			params: url.Values{
				"param1": []string{"value1"},
				"param2": []string{"value2"},
			},
			serverResponse: `{"success": true}`,
			serverStatus:   http.StatusOK,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverStatus)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := NewClient(Config{
				BaseURL: server.URL,
			})

			resp, err := client.Get(context.Background(), tt.endpoint, tt.params)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("Client.Get() returned nil response")
			}

			if !tt.wantErr && resp != nil {
				if resp.StatusCode != tt.serverStatus {
					t.Errorf("Client.Get() status = %d, want %d", resp.StatusCode, tt.serverStatus)
				}
				if string(resp.Body) != tt.serverResponse {
					t.Errorf("Client.Get() body = %s, want %s", string(resp.Body), tt.serverResponse)
				}
			}
		})
	}
}

func TestClient_buildURL(t *testing.T) {
	client := NewClient(Config{
		BaseURL: "https://api.example.com",
	})

	tests := []struct {
		name     string
		endpoint string
		params   url.Values
		want     string
	}{
		{
			name:     "simple endpoint",
			endpoint: "/test",
			params:   nil,
			want:     "https://api.example.com/test",
		},
		{
			name:     "with parameters",
			endpoint: "/test",
			params: url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			want: "https://api.example.com/test?a=1&b=2",
		},
		{
			name:     "sorted parameters",
			endpoint: "/test",
			params: url.Values{
				"z": []string{"last"},
				"a": []string{"first"},
				"m": []string{"middle"},
			},
			want: "https://api.example.com/test?a=first&m=middle&z=last",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.buildURL(tt.endpoint, tt.params)
			if err != nil {
				t.Errorf("Client.buildURL() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Client.buildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientBuildURLRejectsInvalidBaseURL(t *testing.T) {
	client := NewClient(Config{BaseURL: "://invalid"})
	if _, err := client.buildURL("health", nil); err == nil {
		t.Fatal("buildURL() succeeded with an invalid base URL")
	}
}

func TestClientHeaderMutationsAreSafeDuringRequests(t *testing.T) {
	client := NewClient(Config{
		BaseURL: "https://api.example.com",
		HTTPClient: httpClientFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(http.NoBody),
			}, nil
		}),
	})

	const iterations = 200
	errs := make(chan error, iterations)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			client.SetHeader("X-Request-ID", "set")
			client.AddHeader("X-Request-ID", "add")
			client.SetHeaders(http.Header{"X-Request-ID": {"replace"}})
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			if _, err := client.Get(context.Background(), "health", nil); err != nil {
				errs <- err
			}
		}
	}()
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Errorf("Get() returned error: %v", err)
	}
}

// TestClientTimeoutAppliesWithCustomHTTPClient guards the fix for Config.Timeout
// being silently ignored whenever a caller supplies their own HTTPClient (which
// the SDK can't stamp with an http.Client.Timeout). Get now imposes the timeout
// as a per-request context deadline, so a custom client that respects the
// request context is still bounded by it.
func TestClientTimeoutAppliesWithCustomHTTPClient(t *testing.T) {
	client := NewClient(Config{
		BaseURL: "https://api.example.com",
		Timeout: 20 * time.Millisecond,
		// A custom client with no timeout of its own that only returns once
		// the request context is done - i.e. it would hang forever if Get
		// didn't impose a deadline.
		HTTPClient: httpClientFunc(func(req *http.Request) (*http.Response, error) {
			<-req.Context().Done()
			return nil, req.Context().Err()
		}),
	})

	done := make(chan error, 1)
	go func() {
		_, err := client.Get(context.Background(), "health", nil)
		done <- err
	}()

	select {
	case err := <-done:
		if err == nil {
			t.Fatal("Get() returned nil error; expected a deadline-exceeded failure")
		}
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Get() error = %v; want it to wrap context.DeadlineExceeded", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Get() did not return within 2s; Config.Timeout was not enforced for the custom HTTPClient")
	}
}
