package books

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {

	t.Run("default http client", func(t *testing.T) {
		c := NewClient("apikey")
		if c.HTTPClient != http.DefaultClient {
			t.Errorf("NewClient(\"apikey\").HTTPClient == %v, want %v", c.HTTPClient, http.DefaultClient)
		}
	})

	t.Run("provided http client", func(t *testing.T) {
		httpc := &http.Client{Timeout: 45}
		c := NewClient("apikey", WithHTTPClient(httpc))
		if c.HTTPClient != httpc {
			t.Errorf("NewClient(\"apikey\").HTTPClient == %v, want %v", c.HTTPClient, httpc)
		}
	})
}

// mock http client
type MockClient struct {
	MockDo func(*http.Request) (*http.Response, error)
}

func (mc *MockClient) Do(req *http.Request) (*http.Response, error) {
	return mc.MockDo(req)
}

func TestGet(t *testing.T) {
	// setup mock client
	mc := &MockClient{
		MockDo: func(r *http.Request) (*http.Response, error) {
			content := "jetfuel"
			respBody := ioutil.NopCloser(bytes.NewReader([]byte(content)))
			return &http.Response{
				Body: respBody,
			}, nil
		},
	}
	c := NewClient("apikey", WithHTTPClient(mc))
	resp, err := c.Get(context.Background(), "someplace.com")
	if err != nil {
		t.Errorf("Got unexpected error %v", err)
	}
	defer resp.Body.Close()

	want := "jetfuel"
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	got := buf.String()

	if want != got {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestQueryParam(t *testing.T) {
	qp := make(QueryParam)
	qp["first"] = "ffs"
	qp["page"] = "5"

	got := fmt.Sprint(qp)
	want := "first=ffs&page=5"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMakeLink(t *testing.T) {
	c := NewClient("apikey")

	tests := []struct {
		name       string
		endpoint   string
		queryParam QueryParam
		want       string
	}{
		{
			name:       "no query parameters",
			endpoint:   ListsEndpoint,
			queryParam: nil,
			want:       c.base + ListsEndpoint + "?api-key=apikey",
		},
		{
			name:       "two query parameters",
			endpoint:   ListsEndpoint,
			queryParam: QueryParam{"list": "hardcover-fiction", "offset": "40"},
			want:       c.base + ListsEndpoint + "?api-key=apikey&list=hardcover-fiction&offset=40",
		},
		{
			name:       "two query parameters + endpoint with placeholders",
			endpoint:   fmt.Sprintf(ListsByDateEndpoint, "2021-07-06", "hardcover-fiction"),
			queryParam: QueryParam{"list": "hardcover-fiction", "offset": "40"},
			want:       c.base + "/lists/2021-07-06/hardcover-fiction.json" + "?api-key=apikey&list=hardcover-fiction&offset=40",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.makeLink(tt.endpoint, tt.queryParam)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
