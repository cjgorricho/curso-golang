package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		frutas   string
		bebidas  map[string]string
		cantidad []int
	}{
		frutas: "Banano",
		bebidas: map[string]string{
			"gaseosa":  "Coca Cola",
			"malteada": "Chocolate",
			"caliente": "cafe",
		},
		//cantidad: 3,
	}

	fmt.Println(p1)

}
