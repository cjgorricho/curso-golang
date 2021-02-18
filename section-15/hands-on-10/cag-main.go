package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println("x(f): ", f())
	fmt.Println("x(f): ", f())

	g := foo()
	fmt.Println("x(g): ", g())
	fmt.Println("x(g): ", g())
	fmt.Println("x(g): ", g())
	fmt.Println("x(g): ", g())
	fmt.Println("x(g): ", g())
	fmt.Println("x(g): ", g())

	fmt.Println("x(f): ", f())
	fmt.Println("x(f): ", f())

	fmt.Println("x(f): ", f())
	fmt.Println("x(f): ", f())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
