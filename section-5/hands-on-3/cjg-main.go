package main

import "fmt"

//Declarar 3 variables
var x int
var y string
var z bool

func main() {
	x = 42
	y = "Carlos"
	z = true

	// Imprimir valores para cada identificador
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

}
