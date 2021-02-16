package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type exchanges struct {
	Exchanges []string `json:"exchanges"`
}

func main() {
	// fmt.Println("Calling API...")
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://cryptocandledata.com/api/exchanges", strings.NewReader(""))

	if err != nil {
		fmt.Print(err.Error())
	}

	// req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject exchanges
	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Printf("API Response as struct %v\n\n", responseObject)
}
