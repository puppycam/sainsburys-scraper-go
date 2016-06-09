package scraper_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/danbondd/sainsburys-scraper-go"
)

type urlResponse map[string]*http.Response

type mockHTTPClient struct {
	urls urlResponse
}

func (m mockHTTPClient) Get(url string) (*http.Response, error) {
	res, found := m.urls[url]
	if !found {
		return nil, fmt.Errorf("Not Found")
	}
	return res, nil
}

func newMockHTTPClient() mockHTTPClient {
	url := make(urlResponse)
	return mockHTTPClient{urls: url}
}

func newHTTPResponse(body string) *http.Response {
	headers := http.Header{}
	headers["Content-Length"] = []string{"45"}

	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		Header:     headers,
		Request:    &http.Request{},
		StatusCode: http.StatusOK,
	}
}

func TestScrapeWithBadResponse(t *testing.T) {
	mock := newMockHTTPClient()
	scraper := scraper.NewScraper(mock)

	collection, err := scraper.Scrape("http://fake.url")
	if err == nil {
		t.Error("expected error, got: nil")
	}

	if collection != nil {
		t.Error("expected collection to be nil")
	}
}

func TestScrapeReturnsCorrectCollection(t *testing.T) {
	mock := newMockHTTPClient()
	mock.urls["http://fake.product/1"] = newHTTPResponse("<div class='productSummary'><h1>Awesome Product 1</h1><p class='pricePerUnit'>£1.50</p><p class='productText'>This is a really awesome product!</p></div>")
	mock.urls["http://fake.product/2"] = newHTTPResponse("<div class='productSummary'><h1>Awesome Product 2</h1><p class='pricePerUnit'>£2.50</p><p class='productText'>This is a really awesome product!</p></div>")
	mock.urls["http://product.list"] = newHTTPResponse("<div class='product'><div class='productInner'><h3><a href='http://fake.product/1'></a></h3></div></div><div class='product'><div class='productInner'><h3><a href='http://fake.product/2'></a></h3></div></div>")
	scraper := scraper.NewScraper(mock)

	collection, err := scraper.Scrape("http://product.list")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(collection.Results) == 0 {
		t.Errorf("Results should contain 2 products, got: %d", len(collection.Results))
	}

	if collection.Total != 4.00 {
		t.Errorf("Total should be 4.00, got: %.2f", collection.Total)
	}
}
