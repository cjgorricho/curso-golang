package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 4; x++ {
		switch {
		case false:
			fmt.Println("No ingreso un numero")
		case (x == 1):
			fmt.Println("uno")
		case (x == 2):
			fmt.Println("dos")
		case (x == 3):
			fmt.Println("tres")
		default:
			fmt.Println("Este es por defecto")
		}
	}

}
