package main

import "fmt"

func main() {
	a := 42

	fmt.Printf("Value of a: %v\n", a)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Address of a: %v\n", &a)
	fmt.Printf("Type of address of a: %T\n", &a)

	b := &a

	fmt.Println()
	fmt.Printf("Value of b: %v\n", b)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Address of b: %v\n", &b)
	fmt.Printf("Type of address of b: %T\n", &b)
	fmt.Printf("Value of b: %v\n", *b)

	c := &a
	d := &a
	e := &a

	a = 53 // con esta asignaciòn se cambia el valor de todas las variables que tienen un POINTER a la direcciòn de a. NOTA: aunque NO se hace una reasignaciòn del valor de b, esta vatiable adquiere inmediatamente el nuevo valor

	fmt.Println()
	fmt.Printf("Address of a is %v and value is %v\n", &a, a)
	fmt.Printf("Address of b is %v and value is %v\n", b, *b)
	fmt.Printf("Address of c is %v and value is %v\n", c, *c)
	fmt.Printf("Address of d is %v and value is %v\n", d, *d)
	fmt.Printf("Address of e is %v and value is %v\n", e, *e)

	a *= 2 // con esta asignaciòn se cambia el valor de todas las variables que tienen un POINTER a la direcciòn de a. NOTA: aunque NO se hace una reasignaciòn del valor de b, esta vatiable adquiere inmediatamente el nuevo valor

	fmt.Println()
	fmt.Printf("Ahora el valor de b es %v\n\n", *b)

	x := funcion
	x()
	y := &x
	z := *y
	z()
}

func funcion() {
	fmt.Println("Esto se imprime desde la función")
}
