package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/danbondd/sainsburys-scraper-go"
)

func main() {
	url := os.Args[1]

	productScraper := scraper.NewScraper(http.DefaultClient)
	collection, err := productScraper.Scrape(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(collection, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshalling JSON: %v", err)
		os.Exit(1)
	}

	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

	fmt.Println(string(b))
}
