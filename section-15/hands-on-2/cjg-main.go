package main

import (
	"fmt"
)

func main() {
	ii := []int{2, 40, -60, 8, 10, 12}
	n := foo(ii...)
	fmt.Println("El total es: ", n)

	ii2 := []int{2, 100, -60, 8, 10, 12}
	n2 := bar(ii2)
	fmt.Println("Nuevo total es: ", n2)

}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
