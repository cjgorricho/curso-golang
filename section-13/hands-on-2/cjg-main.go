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

	//fmt.Println(p1.nombre, p1.apellido)
	//for i, v := range p1.sabores {
	//	fmt.Println(i, v)
	//}
	//fmt.Println()

	//fmt.Println(p2.nombre, p2.apellido)
	//for i, v := range p2.sabores {
	//	fmt.Println(i, v)
	//}

	fmt.Println("~~~~~~~~~~~~")
	m := map[string]person{
		p1.apellido: p1,
		p2.apellido: p2,
	}
	for k, v := range m {
		fmt.Println(v.nombre, k)
		//fmt.Println(v.apellido)
		for i, val := range v.sabores {
			fmt.Println(i, val)
		}
		fmt.Println("-------------")
	}
	fmt.Println(m)

}
