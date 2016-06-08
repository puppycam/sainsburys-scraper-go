package scraper

import (
	"encoding/json"
	"fmt"
)

// Product contains the relevant data fields for each fruit product.
type Product struct {
	Title       string
	Description string
	Size        int64
	UnitPrice   float64
}

// MarshalJSON takes a product and returns a correctly formatted list attributes.
func (p *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"title":       p.Title,
		"description": p.Description,
		"size":        fmt.Sprintf("%dkb", p.Size),
		"unitprice":   fmt.Sprintf("%.2f", p.UnitPrice),
	})
}
