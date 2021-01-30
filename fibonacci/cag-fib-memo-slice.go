package main

import (
	"fmt"
)

var memo = make([]uint64, 1001)

func fib(n int) uint64 {
	if memo[n] != 0 {
		return memo[n]
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

func main() {

	fmt.Println(fib(10))
	fmt.Println(fib(100))
	fmt.Println(fib(200))

}
