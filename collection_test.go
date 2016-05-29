package scraper_test

import (
	"encoding/json"
	"testing"

	"github.com/danbondd/sainsburys-scraper-go"
)

func TestCollectionJSONOutput(t *testing.T) {
	var products []scraper.Product

	product := scraper.Product{
		Title:       "Test Product",
		Description: "This is an awesome product!",
		Size:        40,
		UnitPrice:   10,
	}

	products = append(products, product)

	collection := scraper.NewCollection(products)
	expectedJSON := "{\"results\":[{\"description\":\"This is an awesome product!\",\"size\":\"40kb\",\"title\":\"Test Product\",\"unitprice\":10}],\"total\":\"10.00\"}"

	json, err := json.Marshal(&collection)
	if err != nil {
		t.Errorf(err.Error())
	}

	if string(json) != expectedJSON {
		t.Errorf("Incorrect JSON output")
	}
}
