package books

import (
	"context"
	"net/http"
	"net/url"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	base       string
	apiKey     string
	HTTPClient Doer
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

func WithHTTPClient(doer Doer) OptionFunc {
	return func(c *Client) {
		c.HTTPClient = doer
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

type QueryParam map[string]string

func (qp QueryParam) String() string {
	vals := url.Values{}
	for key, val := range qp {
		vals.Add(key, val)
	}

	return vals.Encode()
}

func (c *Client) makeLink(endpoint string, qp QueryParam) (string, error) {
	v := url.Values{}
	v.Add("api-key", c.apiKey)
	queryParams := v.Encode()
	if qp != nil {
		queryParams += "&" + qp.String()
	}

	link := c.base + endpoint
	URL, err := url.ParseRequestURI(link)
	if err != nil {
		return "", nil
	}
	URL.RawQuery = queryParams

	return URL.String(), nil
}