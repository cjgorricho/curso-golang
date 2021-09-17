package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "wplay-live-data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("No se pudo crear archivo, error: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnHTML("#inplay-tab-FOOT", func(e *colly.HTMLElement) {

		text := e.ChildText(`div[class="ev-type-header"]`)
		fmt.Printf("Dato obtenido: %v, tipo: %T\n", text, text)
		writer.Write([]string{
			text,
		})
	})

	c.Visit("https://apuestas.wplay.co/es/s/FOOT/Futbol")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
