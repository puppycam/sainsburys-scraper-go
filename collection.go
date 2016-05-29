package scraper

import (
	"encoding/json"
	"fmt"
)

// Collection contains a list of products along with the accumulated total product prices.
type Collection struct {
	Results []Product
	Total   float64
}

// NewCollection returns an aggregate collection of products along with the accumulated total product prices.
func NewCollection(products []Product) *Collection {
	var total float64

	for _, product := range products {
		total += product.UnitPrice
	}

	return &Collection{products, total}
}

// MarshalJSON takes a collection of products and returns a correctly formatted list of results and total price.
func (c *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"results": c.Results,
		"total":   fmt.Sprintf("%.2f", c.Total),
	})
}
