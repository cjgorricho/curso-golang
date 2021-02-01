package main

import (
	"fmt"
)

type person struct {
	nombre   string
	apellido string
	sabores  []string
}

func main() {
	p1 := person{
		nombre:   "James",
		apellido: "Bond",
		sabores: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		nombre:   "Miss",
		apellido: "Moneypenny",
		sabores: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
	for i, v := range p1.sabores {
		fmt.Println(i, v)
	}

	fmt.Println(p2.nombre)
	fmt.Println(p2.apellido)
	for i, v := range p2.sabores {
		fmt.Println(i, v)
	}
}
