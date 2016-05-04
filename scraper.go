package scraper

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	httpClient HttpClient
}

func NewScraper(httpClient HttpClient) Scraper {
	return Scraper{httpClient: httpClient}
}

func (scraper Scraper) Scrape(url string) Product {
	resp, err := scraper.getBodyContent(url)
	if err != nil {
		fmt.Println(err)
		return Product{}
	}

	product := scraper.scrapeBodyContent(resp)

	return product
}

func (scraper Scraper) getBodyContent(url string) (resp *http.Response, err error) {
	resp, err = scraper.httpClient.Get(url)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (scraper Scraper) scrapeBodyContent(resp *http.Response) Product {
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find(".productSummary h1").Text()
	desc := strings.TrimSpace(doc.Find(".productText").First().Text())
	unitPrice := doc.Find("p.pricePerUnit").Text()

	reg, err := regexp.Compile("[^0-9.]")
	if err != nil {
		log.Fatal(err)
	}

	unitPrice = reg.ReplaceAllString(unitPrice, "")
	newUnitPrice, _ := strconv.ParseFloat(unitPrice, 64)

	return Product{
		Title:       title,
		Description: desc,
		Size:        size,
		UnitPrice:   newUnitPrice,
	}
}
