package main

import (
	"fmt"
	"time"
)

var memo map[int]uint64

func fib(n int) uint64 {

	if v, ok := memo[n]; ok {
		return v
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
