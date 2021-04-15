package main

import (
	"fmt"
)

func main() {

	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int {
	x := 1212
	return x
}
func bar() (string, int) {
	nombre := "Carlos"
	cuenta := 1234
	return nombre, cuenta
}
