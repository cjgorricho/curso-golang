package main

import (
	"fmt"
)

type person struct {
	nombre string
}

func main() {
	p1 := person{
		nombre: "Carlos Gorricho",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
func changeMe(p *person) {
	p.nombre = "Julia Ines"
	(*p).nombre = "Carlos Alejandro"
}
