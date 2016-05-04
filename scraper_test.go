package scraper_test

import (
	"fmt"
	"net/http"
	"testing"

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

func getMockClient() scraper.Scraper {
	client := scraper.NewScraper(mockHttpClient{})
	return client
}

func TestSeomthing(t *testing.T) {
	client := getMockClient()
	_ = client.Scrape("http://fake.url")
	// if err.Error() != notFound {
	// 	t.Error("Something")
	// }
}
