// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	urls := []string{
		"https://apuestas.wplay.co/es",
		"https://www.rivalo.co/es/apuestas/",
		"https://sports.yajuego.co/",
		"https://www.zamba.co/es",
		"https://m.codere.com.co/deportescolombia/#/HomePage",
		"https://www.rushbet.co/?page=sports#filter/all/",
		"https://www.luckia.co/apuestas",
		"https://mozzartbet.com.co/es/apuestas",
		"https://betplay.com.co/",
		"https://sports.sportium.com.co/home",
	}

	start := time.Now()
	ch := make(chan string)
	for i, url := range urls {
		go fetch(i, url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}

func fetch(ind int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	body, errBody := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if errBody != nil {
		ch <- fmt.Sprintf("Body while reading %s: %v", url, err)
		return
	}

	bodyText := string(body)
	numLinks := strings.Count(bodyText, "<a")

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("Subroutine %d: %.4fs %10d  %-50s\t%d links", ind, secs, len(bodyText), url, numLinks)

}

//!-
