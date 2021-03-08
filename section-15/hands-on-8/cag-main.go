package main

import "fmt"

func main() {
	f := funcion1()  // el argumento se define en el cuerpo de la funcion. NO CUMPLE CON EL PROPOSITO DE CREAR UNA FUNCION QUE DEVUELVA UN TIPO FUNCION
	g := funcion2()  // el argumento se define al llamar la funcion desde g()
	h := funcion3(1) // el argumento se define al asignar la funcion

	fmt.Println(f)
	fmt.Printf("%T\n", f)

	fmt.Println(g(1))
	fmt.Printf("%T\n", g)

	fmt.Println(h())
	fmt.Printf("%T\n", h)
}

func funcion1() int {
	f := func(x int) int {
		return 51 + x
	}
	return f(1)
}

func funcion2() func(x int) int {
	f := func(x int) int {
		return 52 + x
	}
	return f
}

func funcion3(x int) func() int {
	return func() int {
		return 53 + x
	}
}
