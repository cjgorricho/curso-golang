package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5}
	defer fmt.Println("foo DEFERRED", foo(i...))
	fmt.Println("foo", foo(i...))
	fmt.Println("bar", bar(i))
}

func foo(x ...int) int {
	// fmt.Println(x)
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	// fmt.Println(x)
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
