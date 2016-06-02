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
	product, err := productScraper.Scrape(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	resp, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp))
}
