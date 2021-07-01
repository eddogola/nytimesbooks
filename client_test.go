package books

import (
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