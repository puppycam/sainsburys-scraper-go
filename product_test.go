package scraper_test

import (
	"encoding/json"
	"testing"

	"github.com/danbondd/sainsburys-scraper-go"
)

func TestProductJSONOutput(t *testing.T) {
	product := scraper.Product{
		Title:       "Test Product",
		Description: "This is an awesome product!",
		Size:        40,
		UnitPrice:   10,
	}

	expectedJSON := "{\"description\":\"This is an awesome product!\",\"size\":\"40kb\",\"title\":\"Test Product\",\"unitprice\":10}"

	json, err := json.Marshal(&product)
	if err != nil {
		t.Errorf(err.Error())
	}

	if string(json) != expectedJSON {
		t.Errorf("Incorrect JSON output")
	}
}
