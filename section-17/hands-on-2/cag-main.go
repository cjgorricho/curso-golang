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
	cambio(&p1) // se pasa la drieccion a la variable
	fmt.Println(p1)

}

func cambio(p *person) {
	p.nombre = "Carlos Alejandro"
}
