package stats

import (
	"fmt"
	"net/http"
	"time"

	"github.com/n-ae/nba-api-go/v2/pkg/client"
	"github.com/n-ae/nba-api-go/v2/pkg/client/middleware"
)

const (
	StatsBaseURL = "https://stats.nba.com/stats"
)

type Client struct {
	client *client.Client
}

// Config configures a stats Client. Headers, Timeout, and MaxResponseBytes
// forward directly to the underlying client.Config; a zero Timeout means
// the client default (30s), and a zero MaxResponseBytes means
// client.DefaultMaxResponseBytes.
//
// A caller-supplied Middlewares replaces the default chain entirely
// rather than extending it - use append(stats.DefaultMiddlewares(),
// yourMiddleware...) if you want the defaults (retry, NBA-required
// headers, rate limiting) plus your own additions instead of replacing
// them outright. AdditionalMiddlewares is the same idea as an explicit
// Config field instead of a manual append: it's appended after whichever
// chain Middlewares resolves to (the defaults if Middlewares is empty,
// or your override if it isn't), so
//
//	stats.Config{AdditionalMiddlewares: []client.Middleware{yourMiddleware}}
//
// is equivalent to the append(...) form above but doesn't require calling
// DefaultMiddlewares() yourself, and still layers on top of an explicit
// Middlewares override if you supply one too. The built-in constructors
// (WithRetry, WithPerHostRateLimit, WithHeaders, WithUserAgent, ...) live
// in the importable pkg/client/middleware package if you want to
// reconfigure retry/backoff or rate limits rather than only add to them.
type Config struct {
	BaseURL               string
	Headers               http.Header
	Timeout               time.Duration
	MaxResponseBytes      int64
	Middlewares           []client.Middleware
	AdditionalMiddlewares []client.Middleware
}

// DefaultMiddlewares returns the middleware chain NewClient uses when
// Config.Middlewares is empty: retry with backoff, the headers NBA.com's
// API expects, and a per-host rate limit. It's exported so a custom chain
// can extend these defaults - see Config's doc comment - instead of
// silently replacing them.
func DefaultMiddlewares() []client.Middleware {
	return []client.Middleware{
		middleware.WithRetry(middleware.DefaultRetryConfig()),
		middleware.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
		middleware.WithReferer("https://www.nba.com/"),
		middleware.WithAccept("application/json"),
		middleware.WithPerHostRateLimit(3, 5),
	}
}

// NewClient validates config and constructs a Client. It returns an error
// if config.BaseURL doesn't parse - see client.NewClient.
func NewClient(config Config) (*Client, error) {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = StatsBaseURL
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

	c, err := client.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}

	return &Client{client: c}, nil
}

// NewDefaultClient constructs a Client against StatsBaseURL, a
// compile-time-valid constant, so construction can't fail.
func NewDefaultClient() *Client {
	c, err := NewClient(Config{})
	if err != nil {
		// Unreachable: StatsBaseURL is a valid constant and NewDefaultClient
		// never overrides it, so NewClient's only failure mode (an invalid
		// BaseURL) can't occur here.
		panic(fmt.Sprintf("stats: NewDefaultClient: %v", err))
	}
	return c
}
