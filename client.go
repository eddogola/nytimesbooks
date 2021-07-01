package books

import "net/http"

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
