package scraper_test

import (
	"fmt"
	"net/http"

	"github.com/danbondd/sainsburys-scraper-go"
)

const notFound = "Not found"

type mockHttpClient struct {
	url map[string]*http.Response
}

func (m mockHttpClient) Get(url string) (resp *http.Response, err error) {
	resp, found := m.url[url]
	if !found {
		return nil, fmt.Errorf(notFound)
	}
	return resp, nil
}

func getScraperWithMockHTTPClient() scraper.Scraper {
	scraperMock := scraper.NewScraper(mockHttpClient{})
	return scraperMock
}
