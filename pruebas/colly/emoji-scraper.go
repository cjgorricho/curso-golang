package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("emojipedia.org"),
	)

	// Un programa que usa Colly se desarrolla de ABAJO para ARRIBA
	// Lo que está arriba es lo último que sucede en la cadena de eventos

	// Este bloque de código contiene las acciones principales de SCRAPING
	// Se procesa justo después de CADA REQUEST de un nuevo link
	c.OnHTML("article", func(e *colly.HTMLElement) {
		isEmojiPage := false

		// Extract meta tags from the document
		metaTags := e.DOM.ParentsUntil("~").Find("meta")
		metaTags.Each(func(_ int, s *goquery.Selection) {
			// Search for og:type meta tags
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				// Emoji pages have "article" as their og:type
				isEmojiPage = strings.EqualFold(content, "article")
			}
		})

		if isEmojiPage {
			// Find the emoji page title
			fmt.Println("Emoji: ", e.DOM.Find("h1").Text())
			// Grab all the text from the emoji's description
			fmt.Println(
				"Description: ",
				e.DOM.Find(".description").Find("p").Text())
		}
	})

	// Este bloque de código funciona de manera recursiva, aunque no es aparente.
	// Busca todos los links que estén en elementos a con etiqueta href (a[href])
	// Visita CADA UNO de los links. Llama la atención que no hay un FOR o EACH statement
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the linked URL from the anchor tag
		link := e.Attr("href")
		// Have our crawler visit the linked URL
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Reglas de límites, VS Code me sugiriò completarlas
	c.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*",
		Delay:        1,
		RandomDelay:  1 * time.Second,
		Parallelism:  4, // No percibì mucha diferencia entre 2 y 4 instancias. El CRAWLING se hacìa todas las veces en el mismo orden
	})

	//Cada vez que hay un request (visit web site u otro) ejeuta una acción
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Visita el web site original en donde comienza el crawling / scraping
	c.Visit("https://emojipedia.org")

	// Espera wait groups que se crean por Colly. La primera vez que activè paralelismo no tenìa esta instrucciòn y el programa funcionó
	c.Wait()
}
