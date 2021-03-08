package main

import "fmt"

type person struct {
	nombre string
}

func main() {
	p1 := person{
		nombre: "Carlos Gorricho",
	}

	fmt.Println(p1)
	cambio1(&p1) // se pasa la drieccion a la variable
	fmt.Println(p1)

	p2 := person{
		nombre: "Maria Gorricho",
	}

	fmt.Println(p2)
	p2 = cambio2(p2) // se pasa la variable
	fmt.Println(p2)

}

func cambio1(p *person) {
	p.nombre = "Carlos Alejandro"
}

func cambio2(p person) person {
	p.nombre = "Maria Andrea"
	return p
}
