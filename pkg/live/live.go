package live

import (
	"net/http"
	"time"

	"github.com/n-ae/nba-api-go/pkg/client"
	"github.com/n-ae/nba-api-go/pkg/client/middleware"
)

const (
	LiveBaseURL = "https://cdn.nba.com/static/json/liveData"
)

type Client struct {
	client *client.Client
}

// Config configures a live Client. BaseURL overrides LiveBaseURL, mainly
// for pointing the client at a test server. Headers, Timeout, and
// MaxResponseBytes forward directly to the underlying client.Config; a
// zero Timeout means the client default (30s), and a zero
// MaxResponseBytes means client.DefaultMaxResponseBytes.
//
// A caller-supplied Middlewares replaces the default chain entirely
// rather than extending it - use append(live.DefaultMiddlewares(),
// yourMiddleware...) if you want the defaults plus your own additions
// instead of replacing them outright. AdditionalMiddlewares is the same
// idea as an explicit Config field instead of a manual append: it's
// appended after whichever chain Middlewares resolves to (the defaults
// if Middlewares is empty, or your override if it isn't). The built-in
// constructors (WithUserAgent, WithPerHostRateLimit, ...) live in the
// importable pkg/client/middleware package if you want to reconfigure
// rate limits rather than only add to them.
type Config struct {
	BaseURL               string
	Headers               http.Header
	Timeout               time.Duration
	MaxResponseBytes      int64
	Middlewares           []client.Middleware
	AdditionalMiddlewares []client.Middleware
}

// DefaultMiddlewares returns the middleware chain NewClient uses when
// Config.Middlewares is empty: a User-Agent header and a per-host rate
// limit. It's exported so a custom chain can extend these defaults - see
// Config's doc comment - instead of silently replacing them.
func DefaultMiddlewares() []client.Middleware {
	return []client.Middleware{
		middleware.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
		middleware.WithPerHostRateLimit(5, 10),
	}
}

func NewClient(config Config) *Client {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = LiveBaseURL
	}

	clientConfig := client.Config{
		BaseURL:          baseURL,
		Headers:          config.Headers,
		Timeout:          config.Timeout,
		MaxResponseBytes: config.MaxResponseBytes,
	}

	if len(config.Middlewares) > 0 {
		clientConfig.Middlewares = config.Middlewares
	} else {
		clientConfig.Middlewares = DefaultMiddlewares()
	}
	if len(config.AdditionalMiddlewares) > 0 {
		clientConfig.Middlewares = append(clientConfig.Middlewares, config.AdditionalMiddlewares...)
	}

	return &Client{
		client: client.NewClient(clientConfig),
	}
}

func NewDefaultClient() *Client {
	return NewClient(Config{})
}
