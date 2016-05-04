package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/danbondd/sainsburys-scraper-go"
)

func main() {
	url := os.Args[1]

	productScraper := scraper.NewScraper(http.DefaultClient)
	product := productScraper.Scrape(url)
	resp, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp))
}
