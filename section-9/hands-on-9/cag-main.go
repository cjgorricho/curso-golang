package main

import (
	"fmt"
)

func main() {
	texto := "Escobar"
	switch texto {
	case "Carlos":
		fmt.Println(texto, "es el primer nombre")
	case "Alejandro":
		fmt.Println(texto, "es el segundo nombre")
	case "Gorricho":
		fmt.Println(texto, "es el primer apellido")
	case "Escobar":
		fmt.Println(texto, "es el segundo apellido")
	}
}
