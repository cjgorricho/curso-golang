package main

import (
	"fmt"
	"time"
)

var memo = make(map[int]uint64)

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
	// fmt.Println(fib(10))
	// fmt.Println(fib(100))
	fmt.Println(memo)
	fmt.Println(fib(10))
	for k, v := range memo {
		fmt.Println(k, v)
	}
	fmt.Printf("%.3fms elapsed\n", float64(time.Since(start).Microseconds())/1000)
}
