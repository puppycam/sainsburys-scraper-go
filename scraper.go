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
func (scraper Scraper) Scrape(url string) (*Collection, error) {
	var products []Product

	resp, err := scraper.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting response from %s: %v", url, err)
	}

	productLinks, err := scraper.scrapeProductListBodyContent(resp)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	for _, link := range productLinks {
		res, _ := scraper.httpClient.Get(link)
		product, err := scraper.scrapeProductBodyContent(res)
		if err != nil {
			log.Printf("error getting product data: %v", err)
			continue
		}
		products = append(products, product)
	}

	return NewCollection(products), nil
}

func (scraper Scraper) scrapeProductListBodyContent(resp *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("error creating document from response: %v", err)
	}

	var links []string

	doc.Find(".product .productInner").Each(func(i int, div *goquery.Selection) {
		href, ok := div.Find("h3 a").Attr("href")
		if ok {
			links = append(links, href)
		}
	})

	return links, nil
}

func (scraper Scraper) scrapeProductBodyContent(resp *http.Response) (Product, error) {
	size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return Product{}, fmt.Errorf("error converting product size: %v", err)
	}
	size = (size / 1024)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return Product{}, fmt.Errorf("error creating document from response: %v", err)
	}

	productSummary := doc.Find(".productSummary").First()
	title := productSummary.Find("h1").Text()
	desc := strings.TrimSpace(doc.Find(".productText").First().Text())
	unitPrice := productSummary.Find("p.pricePerUnit").Text()

	reg, err := regexp.Compile("[^0-9.]")
	if err != nil {
		return Product{}, fmt.Errorf("error compiling regex: %v", err)
	}

	unitPrice = reg.ReplaceAllString(unitPrice, "")
	newUnitPrice, err := strconv.ParseFloat(unitPrice, 64)
	if err != nil {
		return Product{}, fmt.Errorf("error converting product price: %v", err)
	}

	return Product{title, desc, size, newUnitPrice}, nil
}
