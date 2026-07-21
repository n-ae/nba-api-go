package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"sort"
	"sync"
	"time"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
)

const (
	// DefaultUserAgent is not applied automatically. The core client is
	// generic and NBA-agnostic; it's each facade's DefaultMiddlewares
	// (see pkg/stats, pkg/live) that installs the User-Agent NBA.com
	// expects, via middleware.WithUserAgent. DefaultUserAgent remains
	// exported for callers who construct client.Client directly and want
	// a reasonable fallback. Major-version-only (not "2.1.2") so it needn't
	// change on every patch release.
	DefaultUserAgent = "nba-api-go/2"
	DefaultTimeout   = 30 * time.Second

	// DefaultMaxResponseBytes bounds how much of a response body Get reads
	// into memory. NBA.com responses are ordinarily well under this, but
	// an upstream failure, proxy error page, or an unexpectedly large
	// endpoint could otherwise consume unbounded memory.
	DefaultMaxResponseBytes int64 = 50 * 1024 * 1024 // 50 MiB
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL          *url.URL
	httpClient       HTTPClient
	headersMu        sync.RWMutex
	headers          http.Header
	timeout          time.Duration
	transport        RoundTripper
	maxResponseBytes int64
}

type Config struct {
	BaseURL    string
	HTTPClient HTTPClient
	Headers    http.Header
	// Timeout bounds each Get: the SDK-built http.Client uses it, and Get
	// also imposes it as a per-request context deadline, so it applies
	// uniformly even when a custom HTTPClient (which the SDK can't
	// configure) is supplied. A caller-provided ctx with an earlier
	// deadline still wins. A value of 0 means DefaultTimeout. A negative
	// value disables timeout enforcement entirely (both the SDK-built
	// client and the per-request context deadline) - unlike 0, it is not
	// normalized to DefaultTimeout.
	Timeout     time.Duration
	Middlewares []Middleware
	// MaxResponseBytes bounds how much of a response body Get reads into
	// memory; a value <= 0 means DefaultMaxResponseBytes.
	MaxResponseBytes int64
}

// NewClient validates config and constructs a Client. It returns an error
// if config.BaseURL doesn't parse, rather than deferring that failure to
// the first call to Get - a caller who mistypes or misconfigures BaseURL
// finds out at startup, not under traffic.
func NewClient(config Config) (*Client, error) {
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	if config.MaxResponseBytes <= 0 {
		config.MaxResponseBytes = DefaultMaxResponseBytes
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	if config.HTTPClient == nil {
		// Clone http.DefaultTransport rather than building one from a
		// mostly zero-valued struct, so proxy support
		// (Proxy: ProxyFromEnvironment), HTTP/2 negotiation, and sane
		// connect timeouts aren't silently lost. This also re-enables
		// keep-alives, which a previous version of this transport
		// unconditionally disabled with no recorded rationale - see
		// ADR 003 for the full analysis and the conditions under which
		// that decision should be revisited.
		defaultTransport, ok := http.DefaultTransport.(*http.Transport)
		if !ok {
			// Practically unreachable: http.DefaultTransport is always
			// *http.Transport unless something replaced the package-level
			// var. Fail loudly rather than risk Clone on a nil transport.
			panic("client: http.DefaultTransport is not *http.Transport")
		}
		transport := defaultTransport.Clone()
		// NBA.com/Akamai's handshake and response times run slower than
		// most sites; give both more headroom than the stdlib defaults
		// (10s and unset, respectively).
		transport.TLSHandshakeTimeout = 30 * time.Second
		transport.ResponseHeaderTimeout = 60 * time.Second

		config.HTTPClient = &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		}
	}

	// Clone rather than alias config.Headers: without this, constructing a
	// Client would mutate the caller's map, later caller-side mutations
	// would silently reach into the client, and concurrent access to
	// either would race.
	headers := config.Headers.Clone()
	if headers == nil {
		headers = make(http.Header)
	}

	baseTransport := &baseRoundTripper{client: config.HTTPClient}

	var transport RoundTripper = baseTransport
	if len(config.Middlewares) > 0 {
		chained := Chain(config.Middlewares...)
		transport = chained(baseTransport)
	}

	return &Client{
		baseURL:          baseURL,
		httpClient:       config.HTTPClient,
		headers:          headers,
		timeout:          config.Timeout,
		transport:        transport,
		maxResponseBytes: config.MaxResponseBytes,
	}, nil
}

type baseRoundTripper struct {
	client HTTPClient
}

func (b *baseRoundTripper) RoundTrip(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return b.client.Do(req)
}

func (c *Client) Get(ctx context.Context, endpoint string, params url.Values) (*models.RawResponse, error) {
	// Impose the configured timeout as a per-request deadline so it applies
	// even when a custom HTTPClient is used (which the SDK can't configure
	// with an http.Client.Timeout of its own). context.WithTimeout keeps a
	// caller's earlier deadline, so this only ever tightens, never extends.
	if c.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.timeout)
		defer cancel()
	}

	reqURL := c.buildURL(endpoint, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.headersMu.RLock()
	headers := c.headers.Clone()
	c.headersMu.RUnlock()
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := c.transport.RoundTrip(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	//nolint:errcheck
	defer resp.Body.Close()

	// Read one byte past the limit so a body that's exactly at the limit
	// (allowed) can be told apart from one that's larger (rejected)
	// without needing a second read.
	limited := io.LimitReader(resp.Body, c.maxResponseBytes+1)
	body, err := io.ReadAll(limited)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode >= 400 {
		if apiErr := models.HTTPStatusToError(resp.StatusCode, reqURL, body); apiErr != nil {
			return nil, apiErr
		}
	}
	if int64(len(body)) > c.maxResponseBytes {
		return nil, fmt.Errorf("%w: %d bytes", models.ErrResponseTooLarge, c.maxResponseBytes)
	}

	return models.NewRawResponse(body, resp.StatusCode, reqURL, resp.Header), nil
}

func (c *Client) GetJSON(ctx context.Context, endpoint string, params url.Values, v interface{}) error {
	rawResp, err := c.Get(ctx, endpoint, params)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rawResp.Body, v); err != nil {
		return fmt.Errorf("%w: %v", models.ErrInvalidResponse, err)
	}

	return nil
}

func (c *Client) buildURL(endpoint string, params url.Values) string {
	baseURL := *c.baseURL

	baseURL.Path = path.Join(baseURL.Path, endpoint)

	if params != nil {
		sortedParams := c.sortParams(params)
		baseURL.RawQuery = sortedParams.Encode()
	}

	return baseURL.String()
}

func (c *Client) sortParams(params url.Values) url.Values {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	sorted := make(url.Values)
	for _, key := range keys {
		sorted[key] = params[key]
	}

	return sorted
}

func (c *Client) SetHeader(key, value string) {
	c.headersMu.Lock()
	defer c.headersMu.Unlock()
	c.headers.Set(key, value)
}

func (c *Client) AddHeader(key, value string) {
	c.headersMu.Lock()
	defer c.headersMu.Unlock()
	c.headers.Add(key, value)
}

// SetHeaders replaces the client's headers with a copy of headers. It
// clones rather than aliasing the caller's map so that later mutations to
// the map the caller passed in (or to c.headers via SetHeader/AddHeader)
// can't reach into each other unexpectedly.
func (c *Client) SetHeaders(headers http.Header) {
	c.headersMu.Lock()
	defer c.headersMu.Unlock()
	c.headers = headers.Clone()
	if c.headers == nil {
		c.headers = make(http.Header)
	}
}
