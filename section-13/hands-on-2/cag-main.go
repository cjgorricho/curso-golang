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

	m := map[string]person{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	fmt.Println(m)
	fmt.Println()
	fmt.Println(m["Bond"])
	fmt.Println(m["Moneypenny"])
	fmt.Println()

	// Al asignar un STRUCT a un MAP, los paràmetros del STRUCT son "elevados" auotmàticamente al nivel del MAP. Poe eso se pueden llamar como se hace a continuaciòn, sin necesidad de invocar el STRUCT intermedio

	for k, v := range m {
		fmt.Println("Key :", k)
		fmt.Println(v.nombre)
		fmt.Println(v.apellido)
		for i, val := range v.sabores {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
