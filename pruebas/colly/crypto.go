package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	fName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML(`tbody tr`, func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(`div[class*="__column-name"]`),
			e.ChildText(`td[class*="symbol"]`),
			e.ChildText(`td[class*="price"]`),
			e.ChildText(`td[class*="volume-24-h"]`),
			e.ChildText(`td[class*="market-cap"]`),
			e.ChildText(`td[class*="percent-change-1-h"]`),
			e.ChildText(`td[class*="percent-change-24-h"]`),
			e.ChildText(`td[class*="percent-change-7-d"]`),
		})
	})

	c.Limit(&colly.LimitRule{
		Delay: 5 * time.Second,
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
