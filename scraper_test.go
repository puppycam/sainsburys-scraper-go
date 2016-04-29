package scraper_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/danbondd/sainsburys-scraper-go"
)

type mockHttpClient struct {
	url map[string] *http.Response
}

func (m mockHttpClient) Get(url string) (resp *http.Response, err error) {
	resp, found := m.url[url]
	if !found {
		return nil, fmt.Errorf("lol")
	}
	return resp, nil
}

func getMockClient() scraper.Scraper {
	client := scraper.NewScraper(mockHttpClient{})
	return client
}

func TestSeomthing(t *testing.T) {
	client := getMockClient()
	_, err := client.GetContent("")
	if err.Error() != "lol" {
		t.Error("Something")
	}
}