package main

import (
	"fmt"
)

func main() {

	g := f1 // asigna f1 a g

	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func f1(xi []int) int {
	if len(xi) == 0 {
		return 0
	}
	if len(xi) == 1 {
		return xi[0]
	}
	return xi[0] + xi[len(xi)-1]
}

// EJEMPLOS DE APLICACION ( A DESARROLLAR)
// 1. Se define una funcion para cada una de las estadìsticas que se le quieran calcular a una muestra (promedio, desv est, mediana, quartiles, etc). Se asigna cada una de estas funciones a una variable de tipo FUNC. Luego se utiliza la funciòn como call back sobre la muestra
