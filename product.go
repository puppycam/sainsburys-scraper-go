package scraper

// Product contains the relevant data fields for each fruit product.
type Product struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Size        int64   `json:"size"`
	UnitPrice   float64 `json:"unitprice"`
}
