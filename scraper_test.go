package scraper_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/danbondd/sainsburys-scraper-go"
)

type url map[string]*http.Response

type mockHttpClient struct {
	url url
}

func (m mockHttpClient) Get(url string) (*http.Response, error) {
	res, found := m.url[url]
	if !found {
		return nil, fmt.Errorf("Not Found")
	}
	return res, nil
}

func newMockHTTPClient() mockHttpClient {
	url := make(url)
	return mockHttpClient{url: url}
}

func TestBadResponse(t *testing.T) {
	mock := newMockHTTPClient()
	scraper := scraper.NewScraper(mock)

	_, err := scraper.Scrape("http://fake.url")
	if err == nil {
		t.Errorf("expected error")
	}
}

// func TestScrape(t *testing.T) {
// 	mock := newMockHTTPClient()
// 	mock.url["http://correct.url"] = &http.Response{
// 		Body:       ioutil.NopCloser(bytes.NewBufferString("Hello World")),
// 		StatusCode: 200,
// 	}
// 	scraper := scraper.NewScraper(mock)
//
// 	_, err := scraper.Scrape("http://correct.url")
// 	if err != nil {
// 		t.Errorf("expected error: %v", err)
// 	}
// }
