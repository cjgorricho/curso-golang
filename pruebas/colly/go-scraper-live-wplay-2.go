package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	fName := "wplay-live-data-2.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("No se pudo crear archivo, error: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnHTML(`div[id="inplay-tab-FOOT"]`, func(e *colly.HTMLElement) {

		fmt.Printf("Numero de hijos %s[id=%s]: %d\n\n", e.Name, e.Attr("id"), e.DOM.Children().Length())

		e.ForEach(`div[class="ev-type-header"]`, func(i int, elem *colly.HTMLElement) {
			text := strings.Trim(elem.Text, "\n")
			fmt.Printf("Dato obtenido %v: %v\n", i, text)
			writer.Write([]string{
				text,
			})
		})

	})

	c.Visit("https://apuestas.wplay.co/es/s/FOOT/Futbol")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
