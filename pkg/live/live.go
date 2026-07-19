package live

import (
	"net/http"
	"time"

	"github.com/n-ae/nba-api-go/internal/middleware"
	"github.com/n-ae/nba-api-go/pkg/client"
)

const (
	LiveBaseURL = "https://cdn.nba.com/static/json/liveData"
)

type Client struct {
	client *client.Client
}

// Config configures a live Client. BaseURL overrides LiveBaseURL, mainly
// for pointing the client at a test server. Headers is a plain map for
// ergonomics; values are copied into an http.Header before being
// forwarded. Timeout is in seconds; zero means the client default (30s).
type Config struct {
	BaseURL     string
	Headers     map[string]string
	Timeout     int
	Middlewares []middleware.Middleware
}

func NewClient(config Config) *Client {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = LiveBaseURL
	}

	clientConfig := client.Config{
		BaseURL: baseURL,
	}

	if len(config.Headers) > 0 {
		headers := make(http.Header, len(config.Headers))
		for key, value := range config.Headers {
			headers.Set(key, value)
		}
		clientConfig.Headers = headers
	}

	if config.Timeout > 0 {
		clientConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	if len(config.Middlewares) > 0 {
		clientConfig.Middlewares = config.Middlewares
	} else {
		clientConfig.Middlewares = []middleware.Middleware{
			middleware.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
			middleware.WithPerHostRateLimit(5, 10),
		}
	}

	return &Client{
		client: client.NewClient(clientConfig),
	}
}

func NewDefaultClient() *Client {
	return NewClient(Config{})
}
