package stats

import (
	"github.com/n-ae/nba-api-go/internal/middleware"
	"github.com/n-ae/nba-api-go/pkg/client"
)

const (
	StatsBaseURL = "https://stats.nba.com/stats"
)

type Client struct {
	client *client.Client
}

type Config struct {
	Headers     map[string]string
	Timeout     int
	Middlewares []middleware.Middleware
}

func NewClient(config Config) *Client {
	clientConfig := client.Config{
		BaseURL: StatsBaseURL,
	}

	if len(config.Middlewares) > 0 {
		clientConfig.Middlewares = config.Middlewares
	} else {
		clientConfig.Middlewares = []middleware.Middleware{
			middleware.WithRetry(middleware.DefaultRetryConfig()),
			middleware.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
			middleware.WithReferer("https://www.nba.com/"),
			middleware.WithAccept("application/json"),
			middleware.WithPerHostRateLimit(3, 5),
		}
	}

	return &Client{
		client: client.NewClient(clientConfig),
	}
}

func NewDefaultClient() *Client {
	return NewClient(Config{})
}
