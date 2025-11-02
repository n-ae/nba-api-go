package live

import (
	"github.com/n-ae/nba-api-go/internal/middleware"
	"github.com/n-ae/nba-api-go/pkg/client"
)

const (
	LiveBaseURL = "https://cdn.nba.com/static/json/liveData"
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
		BaseURL: LiveBaseURL,
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
