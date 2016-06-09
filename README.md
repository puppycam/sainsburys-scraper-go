# Sainsbury's Grocery Scraper [![Build Status](https://travis-ci.org/danbondd/sainsburys-scraper-go.svg?branch=master)](https://travis-ci.org/danbondd/sainsburys-scraper-go)

A console application, written in Go, that scrapes the Ripe Fruits page of the Sainsbury's grocery website and returns a JSON encoded collection of products along with their running totals.

## Usage

- To run the application `go run console/console.go "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html"`
- To test the application `go test ./... --cover`

## Vendoring

- [FiloSottile/gvt](https://github.com/FiloSottile/gvt)

## Dependencies

- [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
