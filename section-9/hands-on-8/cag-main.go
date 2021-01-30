package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("no imprime")
	case true:
		fmt.Println("si imprime")
	}
	texto := "Gorricho"
	switch {
	case texto == "Carlos":
		fmt.Println(texto, "es el nombre")
	case texto == "Gorricho":
		fmt.Println(texto, "es el apellido")
	}
}
