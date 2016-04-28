package scraper

import (
	"io/ioutil"
	"net/http"
)

type Scraper struct {
	httpClient HttpClient
}

func NewScraper(httpClient HttpClient) Scraper {
	return Scraper{httpClient: httpClient}
}

func (scraper Scraper) GetContent(url string) (body string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return body, err
	}

	return string(b), nil
}
