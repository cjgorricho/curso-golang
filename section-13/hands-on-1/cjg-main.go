package main

import "fmt"

type person struct {
	nombre   string
	apellido string
	sabores  []string
}

func main() {
	p1 := person{
		nombre:   "Carlos",
		apellido: "Gorricho",
		sabores: []string{
			"cacao",
			"cafe",
			"mani",
		},
	}
	p2 := person{
		nombre:   "Julia Ines",
		apellido: "Escobar",
		sabores: []string{
			"te",
			"limonada",
			"leche",
		},
	}
	fmt.Println(p1.nombre, p1.apellido)
	for i, v := range p1.sabores {
		fmt.Println(i, v)
	}
	fmt.Println()

	fmt.Println(p2.nombre, p2.apellido)
	for i, v := range p2.sabores {
		fmt.Println(i, v)
	}
}
