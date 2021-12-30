package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {

	fName := "FOOT-live-data.csv"
	file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("No se pudo crear archivo, error: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	liga := ""
	start := time.Now()
	cons := 0
	eventodds := make([]string, 3, 5)
	var data [][]string

	c := colly.NewCollector()

	c.OnHTML(`div#USInplay-tab-FOOT div.table-row.row-wrap`, func(e *colly.HTMLElement) {

		if e.DOM.Children().Length() == 1 {
			if liga != e.ChildText(`div.ev-type-header`) {
				liga = e.ChildText(`div.ev-type-header`)
			}
		} else {
			cons = cons + 1
			timestamp := time.Now().Round(time.Second)
			eventid := e.Attr("data-mkt_id")
			eventinfo := strings.Split(e.ChildText(`div[class="team-score"]`), "\n")
			eventtime := e.ChildText(`span.clock_mode_forward`)
			eventperiod := e.ChildText(`span.period`)

			e.ForEach(`div.inplay td.seln`, func(i int, elem *colly.HTMLElement) {
				eventodds[i] = elem.ChildText(`span.dec`)
			})

			data = append(data, []string{
				timestamp.Format(time.RFC822Z),
				"WPlay",
				liga,
				eventid,
				eventinfo[1] + " - " + eventinfo[4],
				eventinfo[0] + "-" + eventinfo[3],
				eventtime,
				eventperiod,
				eventodds[0],
				eventodds[1],
				eventodds[2],
			})

		}

	})

	c.Visit("https://apuestas.wplay.co/es")

	writer.WriteAll(data)

	fmt.Println("#", "TimeStamp", "BettingHouse", "League_Tournament", "EventID", "EventName", "EventScore", "EventTime", "EventPeriod", "OddsLocal", "OddsDraw", "OddsVisitor")
	for i, line := range data {
		fmt.Println(i, line)
	}

	log.Printf("Scraping finished, check file %q for results\n", fName)

	secs := time.Since(start).Seconds()
	fmt.Printf("Tiempo total: %v segs", secs)
}
