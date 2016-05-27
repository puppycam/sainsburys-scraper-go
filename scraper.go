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
func (scraper Scraper) Scrape(url string) *Collection {
	var products []Product

	resp, err := scraper.httpClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	productLinks := scraper.scrapeProductListBodyContent(resp)

	for _, link := range productLinks {
		res, _ := scraper.httpClient.Get(link)
		product := scraper.scrapeProductBodyContent(res)
		products = append(products, product)
	}

	return NewCollection(products)
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
	size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	size = (size / 1024)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	productSummary := doc.Find(".productSummary").First()
	title := productSummary.Find("h1").Text()
	desc := strings.TrimSpace(doc.Find(".productText").First().Text())
	unitPrice := productSummary.Find("p.pricePerUnit").Text()

	reg, err := regexp.Compile("[^0-9.]")
	if err != nil {
		log.Fatal(err)
	}

	unitPrice = reg.ReplaceAllString(unitPrice, "")
	newUnitPrice, err := strconv.ParseFloat(unitPrice, 64)
	if err != nil {
		log.Fatal(err)
	}

	return Product{title, desc, size, newUnitPrice}
}
