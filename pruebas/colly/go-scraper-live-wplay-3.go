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

	fName := "wplay-live-data-3.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("No se pudo crear archivo, error: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	liga := ""

	c.OnHTML(`div#USInplay-tab-FOOT div.table-row.row-wrap`, func(e *colly.HTMLElement) {

		if e.DOM.Children().Length() == 1 {
			if liga != e.ChildText(`div.ev-type-header`) {
				liga = e.ChildText(`div.ev-type-header`)
			}
		} else {
			eventid := e.Attr("data-mkt_id")
			eventinfo := strings.Split(e.ChildText(`div[class="team-score"]`), "\n")
			fmt.Printf(
				"%s, %s, %s - %s\n",
				liga,
				eventid,
				eventinfo[1],
				eventinfo[4],
			)

		}

	})

	c.Visit("https://apuestas.wplay.co/es")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
