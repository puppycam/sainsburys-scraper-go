package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danbondd/sainsburys-scraper-go"
)

func main() {
	url := os.Args[1]

	productScraper := scraper.NewScraper(http.DefaultClient)
	body, err := productScraper.GetContent(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(body)
}
