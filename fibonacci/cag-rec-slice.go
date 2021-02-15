package main

import (
	"fmt"
	"time"
)

// TO DO: econtrar la forma de hacer crecer el SLICE memo de forma flexible

var memo = make([]uint64, 101)

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
	start := time.Now()
	fmt.Println(fib(100))
	fmt.Printf("%.3fms elapsed\n", float64(time.Since(start).Microseconds())/1000)
}
