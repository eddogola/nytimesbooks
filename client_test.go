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
