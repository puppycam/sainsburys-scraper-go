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

// Scraper contains the httpClient field which has the custom HTTPClient type.
type Scraper struct {
	httpClient HTTPClient
}

// NewScraper returns a new Scraper with all of its required fields.
func NewScraper(httpClient HTTPClient) Scraper {
	return Scraper{httpClient: httpClient}
}

// Scrape takes a URL, gets its HTML content and returns a new Product with the required data.
func (scraper Scraper) Scrape(url string) (products []Product) {
	resp, err := scraper.getBodyContent(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	productLinks := scraper.scrapeProductListBodyContent(resp)

	for _, link := range productLinks {
		body, _ := scraper.getBodyContent(link)
		product := scraper.scrapeProductBodyContent(body)
		products = append(products, product)
	}

	return products
}

func (scraper Scraper) getBodyContent(url string) (resp *http.Response, err error) {
	resp, err = scraper.httpClient.Get(url)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (scraper Scraper) scrapeProductListBodyContent(resp *http.Response) []string {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	links := doc.Find(".product .productInner").Map(func(i int, div *goquery.Selection) string {
		href, _ := div.Find("h3 a").Attr("href")

		return href
	})

	return links
}

func (scraper Scraper) scrapeProductBodyContent(resp *http.Response) Product {
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
