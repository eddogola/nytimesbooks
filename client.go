package books

import (
	"context"
	"net/http"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	base       string
	apiKey     string
	HTTPClient *http.Client
}

type OptionFunc func(*Client)

func NewClient(apiKey string, options ...OptionFunc) *Client {
	c := &Client{
		base:       "https://api.nytimes.com/svc/books/v3",
		apiKey:     apiKey,
		HTTPClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func WithHTTPClient(httpc *http.Client) OptionFunc {
	return func(c *Client) {
		c.HTTPClient = httpc
	}
}

func (c *Client) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
