package main

import (
	"fmt"
)

func main() {
	ii := []int{2, 40, -60, 8, 10, 12}
	defer fmt.Println("El total diferido es: ", foo(ii...))

	ii2 := []int{2, 100, -60, 8, 10, 12}
	n2 := bar(ii2)
	fmt.Println("Nuevo total es: ", n2)

	ii3 := []int{2, 1, -60, 8, 10, 12}
	n3 := bar(ii3)
	fmt.Println("Nuevo total no diferido es: ", n3)

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
