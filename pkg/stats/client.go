package stats

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
)

func (c *Client) GetJSON(ctx context.Context, endpoint string, params url.Values, v interface{}) error {
	return c.client.GetJSON(ctx, endpoint, params, v)
}

func (c *Client) Get(ctx context.Context, endpoint string, params url.Values) (*models.RawResponse, error) {
	return c.client.Get(ctx, endpoint, params)
}
