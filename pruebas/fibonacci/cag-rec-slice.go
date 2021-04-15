package main

import (
	"fmt"
	"time"
)

// TO DO: econtrar la forma de hacer crecer el SLICE memo de forma flexible

var memo []uint64

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
	n := 100
	memo = make([]uint64, n+1)
	fmt.Println(fib(n))
	fmt.Printf("%.3fms elapsed\n", float64(time.Since(start).Microseconds())/1000)
}
