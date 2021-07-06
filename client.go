package books

import (
	"context"
	"encoding/json"
	"fmt"
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

func (c *Client) get(ctx context.Context, url string) (*http.Response, error) {
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

// GetBestSellersList Gets Best Sellers list. If no date is provided returns the latest list.
func (c *Client) GetBestSellersList(qp QueryParam) (*List, error) {
	URL, err := c.makeLink(ListsEndpoint, qp)
	if err != nil {
		return nil, err
	}
	resp, err := c.get(context.Background(), URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list List
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, err
}

// GetBestSellersListByDate Gets Best Sellers list by date.
func (c *Client) GetBestSellersListByDate(date, listName string, qp QueryParam) (*ListByDate, error) {
	endpoint := fmt.Sprintf(ListsByDateEndpoint, date, listName)
	URL, err := c.makeLink(endpoint, qp)
	if err != nil {
		return nil, err
	}
	resp, err := c.get(context.Background(), URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list ListByDate
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, err
}

// GetBestSellersListHistory Gets Best Sellers list history.
func (c *Client) GetBestSellersListHistory(qp QueryParam) (*ListHistory, error) {
	URL, err := c.makeLink(HistoryEndpoint, qp)
	if err != nil {
		return nil, err
	}
	resp, err := c.get(context.Background(), URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hist ListHistory
	err = json.NewDecoder(resp.Body).Decode(&hist)
	if err != nil {
		return nil, err
	}

	return &hist, err
}

// GetBestSellersListNames Gets Best Sellers list names.
func (c *Client) GetBestSellersListNames() (*Names, error) {
	URL, err := c.makeLink(NamesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.get(context.Background(), URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var names Names
	err = json.NewDecoder(resp.Body).Decode(&names)
	if err != nil {
		return nil, err
	}

	return &names, err
}

// GetOverview Gets top 5 books for all the Best Sellers lists for specified date.
func (c *Client) GetOverview(qp QueryParam) (*Overview, error) {
	URL, err := c.makeLink(OverviewEndpoint, qp)
	if err != nil {
		return nil, err
	}
	resp, err := c.get(context.Background(), URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var overview Overview
	err = json.NewDecoder(resp.Body).Decode(&overview)
	if err != nil {
		return nil, err
	}

	return &overview, err
}

// GetReviews Gets book reviews.
func (c *Client) GetReviews(qp QueryParam) (*Reviews, error) {
	URL, err := c.makeLink(ReviewsEndpoint, qp)
	if err != nil {
		return nil, err
	}
	resp, err := c.get(context.Background(), URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reviews Reviews
	err = json.NewDecoder(resp.Body).Decode(&reviews)
	if err != nil {
		return nil, err
	}

	return &reviews, err
}