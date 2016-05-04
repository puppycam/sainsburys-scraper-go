package scraper

type Product struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Size        int     `json:"size"`
	UnitPrice   float64 `json:"unitprice"`
}
