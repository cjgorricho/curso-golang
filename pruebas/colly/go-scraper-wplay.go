package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "wplay-data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("No se pudo crear archivo, error: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnHTML(`div[data-src_code="COUPON_TYPE"]`, func(e *colly.HTMLElement) {

		text := e.ChildText(`h4`)
		fmt.Printf("Dato obtenido: %v, tipo: %T\n", text, text)
		writer.Write([]string{
			text,
		})

	})

	c.Visit("https://apuestas.wplay.co/es/s/FOOT/Futbol")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
