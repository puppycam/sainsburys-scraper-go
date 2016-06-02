package scraper

import (
	"net/http"
)

// HTTPClient is the interface that wraps the basic Get method.
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// HTTPResponse contains specific information about a GET request
type HTTPResponse struct {
	url      string
	response *http.Response
	err      error
}
