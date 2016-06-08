package scraper

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

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

// Scrape takes a URL, gets its HTML content and returns a new Collection or Products.
func (s Scraper) Scrape(url string) (*Collection, error) {
	var products []Product

	res, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting response from %s: %v", url, err)
	}

	productLinks, err := s.scrapeProductListBodyContent(res)
	if err != nil {
		return nil, fmt.Errorf("error getting product links: %v", err)
	}

	productResponses := s.getProductPageResponses(productLinks)

	for res := range productResponses {
		if res.err != nil || res.response == nil || res.response.StatusCode != http.StatusOK {
			res.response.Body.Close()
			continue
		}
		product, err := s.scrapeProductBodyContent(res.response)
		if err != nil {
			log.Printf("error getting product data: %v", err)
			continue
		}
		products = append(products, product)
	}

	return NewCollection(products), nil
}

func (s Scraper) getProductPageResponses(urls []string) (ch chan *HTTPResponse) {
	ch = make(chan *HTTPResponse)
	go func() {
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				res, err := s.httpClient.Get(url)
				ch <- &HTTPResponse{url, res, err}
			}(url)
		}
		wg.Wait()
		close(ch)
	}()

	return
}

func (s Scraper) scrapeProductListBodyContent(res *http.Response) (links []string, err error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, fmt.Errorf("error creating document from response: %v", err)
	}

	doc.Find(".product .productInner").Each(func(i int, div *goquery.Selection) {
		href, ok := div.Find("h3 a").Attr("href")
		if ok {
			links = append(links, href)
		}
	})

	return links, nil
}

func (s Scraper) scrapeProductBodyContent(res *http.Response) (Product, error) {
	size, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return Product{}, fmt.Errorf("error converting product size: %v", err)
	}
	size = (size / 1024)

	doc, err := goquery.NewDocumentFromResponse(res)
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
