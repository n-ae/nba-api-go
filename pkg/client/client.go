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
	"time"

	"github.com/n-ae/nba-api-go/internal/middleware"
	"github.com/n-ae/nba-api-go/pkg/models"
)

const (
	DefaultUserAgent = "nba-api-go/1.0"
	DefaultTimeout   = 30 * time.Second
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL    string
	httpClient HTTPClient
	headers    http.Header
	timeout    time.Duration
	transport  middleware.RoundTripper
}

type Config struct {
	BaseURL     string
	HTTPClient  HTTPClient
	Headers     http.Header
	Timeout     time.Duration
	Middlewares []middleware.Middleware
}

func NewClient(config Config) *Client {
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	if config.HTTPClient == nil {
		transport := &http.Transport{
			DisableKeepAlives:     true,
			MaxIdleConns:          1,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   30 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
		}

		config.HTTPClient = &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		}
	}

	if config.Headers == nil {
		config.Headers = make(http.Header)
	}

	if config.Headers.Get("User-Agent") == "" {
		config.Headers.Set("User-Agent", DefaultUserAgent)
	}

	baseTransport := &baseRoundTripper{client: config.HTTPClient}

	var transport middleware.RoundTripper = baseTransport
	if len(config.Middlewares) > 0 {
		chained := middleware.Chain(config.Middlewares...)
		transport = chained(baseTransport)
	}

	return &Client{
		baseURL:    config.BaseURL,
		httpClient: config.HTTPClient,
		headers:    config.Headers,
		timeout:    config.Timeout,
		transport:  transport,
	}
}

type baseRoundTripper struct {
	client HTTPClient
}

func (b *baseRoundTripper) RoundTrip(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return b.client.Do(req)
}

func (c *Client) Get(ctx context.Context, endpoint string, params url.Values) (*models.RawResponse, error) {
	reqURL, err := c.buildURL(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, values := range c.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := c.transport.RoundTrip(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		if apiErr := models.HTTPStatusToError(resp.StatusCode, reqURL); apiErr != nil {
			return nil, apiErr
		}
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

func (c *Client) buildURL(endpoint string, params url.Values) (string, error) {
	baseURL, err := url.Parse(c.baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	baseURL.Path = path.Join(baseURL.Path, endpoint)

	if params != nil {
		sortedParams := c.sortParams(params)
		baseURL.RawQuery = sortedParams.Encode()
	}

	return baseURL.String(), nil
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
	c.headers.Set(key, value)
}

func (c *Client) AddHeader(key, value string) {
	c.headers.Add(key, value)
}

func (c *Client) SetHeaders(headers http.Header) {
	c.headers = headers
}
