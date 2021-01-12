package main

import "fmt"

func main() {

	// Usando el operador de declaraci'on corta, ASIGNE estos Valores a la correspondientes VARIABLES

	x := 42
	y := "James Bond"
	z := true

	//Imprimir en un solo statement

	fmt.Println("Imprimir en un solo statement, usando Printf")
	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)
	fmt.Println()

	// Imprimir en varias lineas
	fmt.Println("Imprimir en un solo statement")
	s := fmt.Sprintf("x: %v\ny: %v\nz: %v\n", x, y, z)
	fmt.Println(s)

	//Imprimir en varias lineas, separadas
	fmt.Println("Imprimir en varias lineas separadas")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)

}
