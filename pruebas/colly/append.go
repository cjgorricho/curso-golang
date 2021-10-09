package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	var path = "sample.csv"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	var data [][]string
	data = append(data, []string{"bruce", "wayne"})
	data = append(data, []string{"clark", "kent"})
	data = append(data, []string{"hal", "jordan"})

	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appending succed")
}
