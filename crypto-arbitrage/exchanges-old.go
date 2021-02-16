package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	client := &http.Client{}

	url := "https://cryptocandledata.com/api/exchanges"
	method := "GET"

	payload := strings.NewReader("")

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
