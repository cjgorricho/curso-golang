package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5}
	x := foo(i...)
	fmt.Println(x)
	y := bar(i)
	fmt.Println(y)
}

func foo(x ...int) int {
	fmt.Println(x)
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	fmt.Println(x)
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
