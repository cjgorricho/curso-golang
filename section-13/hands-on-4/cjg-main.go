package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		frutas   string
		bebidas  string
		cantidad int
	}{
		frutas:   "Banano",
		bebidas:  "Coca Cola",
		cantidad: 3,
	}

	fmt.Println(p1)

}
