package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

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
