package scraper

import "io/ioutil"

type Scraper struct {
	httpClient HttpClient
}

func NewScraper(httpClient HttpClient) Scraper {
	return Scraper{httpClient: httpClient}
}

func (scraper Scraper) GetContent(url string) (body string, err error) {
	resp, err := scraper.httpClient.Get(url)
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
