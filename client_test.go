package books

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
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
	resp, err := c.get(context.Background(), "someplace.com")
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

func TestGetBestSellersList(t *testing.T) {
	jsonData := `{"status": "OK",
				  "copyright": "Copyright (c) 2019 The New York Times Company.  All Rights Reserved.",
				  "num_results": 1,
				  "last_modified": "2016-03-11T13:09:01-05:00"
				}`

	// setup mock http client
	mc := &MockClient{
		func(r *http.Request) (*http.Response, error) {
			body := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))

			return &http.Response{Body: body}, nil
		},
	}

	c := NewClient("apikey", WithHTTPClient(mc))
	got, err := c.GetBestSellersList(nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var want List
	err = json.Unmarshal([]byte(jsonData), &want)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetBestSellersListByDate(t *testing.T) {
	jsonData := `{  "status": "OK",  "copyright": "Copyright (c) 2019 The New York Times Company.  All Rights Reserved.",  "num_results": 15,  "last_modified": "2015-12-25T13:05:20-05:00",  "results": {    "list_name": "Trade Fiction Paperback",    "bestsellers_date": "2015-12-19",    "published_date": "2016-01-03",    "display_name": "Paperback Trade Fiction",    "normal_list_ends_at": 10,    "updated": "WEEKLY",    "books": [      {        "rank": 1,        "rank_last_week": 0,        "weeks_on_list": 60,        "asterisk": 0,        "dagger": 0,        "primary_isbn10": "0553418025",        "primary_isbn13": "9780553418026",        "publisher": "Broadway",        "description": "Separated from his crew, an astronaut embarks on a quest to stay alive on Mars. The basis of the movie.",        "price": 0,        "title": "THE MARTIAN",        "author": "Andy Weir",        "contributor": "by Andy Weir",        "contributor_note": "",        "book_image": "http://du.ec2.nytimes.com.s3.amazonaws.com/prd/books/9780804139038.jpg",        "amazon_product_url": "http://www.amazon.com/The-Martian-Novel-Andy-Weir-ebook/dp/B00EMXBDMA?tag=thenewyorktim-20",        "age_group": "",        "book_review_link": "",        "first_chapter_link": "",        "sunday_review_link": "",        "article_chapter_link": "",        "isbns": [          {            "isbn10": "0804139024",            "isbn13": "9780804139021"          }        ]      }    ],    "corrections": []  }}`

	// setup mock http client
	mc := &MockClient{
		func(r *http.Request) (*http.Response, error) {
			body := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))

			return &http.Response{Body: body}, nil
		},
	}

	c := NewClient("apikey", WithHTTPClient(mc))
	got, err := c.GetBestSellersListByDate("2020-07-06", "hardcover-fiction", nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var want ListByDate
	err = json.Unmarshal([]byte(jsonData), &want)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetBestSellersListHistory(t *testing.T) {
	jsonData := `{  "status": "OK",  "copyright": "Copyright (c) 2019 The New York Times Company.  All Rights Reserved.",  "num_results": 28970,  "results": [    {      "title": "#GIRLBOSS",      "description": "An online fashion retailer traces her path to success.",      "contributor": "by Sophia Amoruso",      "author": "Sophia Amoruso",      "contributor_note": "",      "price": 0,      "age_group": "",      "publisher": "Portfolio/Penguin/Putnam",      "isbns": [        {          "isbn10": "039916927X",          "isbn13": "9780399169274"        }      ],      "ranks_history": [        {          "primary_isbn10": "1591847931",          "primary_isbn13": "9781591847939",          "rank": 8,          "list_name": "Business Books",          "display_name": "Business",          "published_date": "2016-03-13",          "bestsellers_date": "2016-02-27",          "weeks_on_list": 0,          "ranks_last_week": null,          "asterisk": 0,          "dagger": 0        }      ],      "reviews": [        {          "book_review_link": "",          "first_chapter_link": "",          "sunday_review_link": "",          "article_chapter_link": ""        }      ]    }  ]}`

	// setup mock http client
	mc := &MockClient{
		func(r *http.Request) (*http.Response, error) {
			body := ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))

			return &http.Response{Body: body}, nil
		},
	}

	c := NewClient("apikey", WithHTTPClient(mc))
	got, err := c.GetBestSellersListHistory(nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var want ListHistory
	err = json.Unmarshal([]byte(jsonData), &want)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v want %v", got, want)
	}
}
