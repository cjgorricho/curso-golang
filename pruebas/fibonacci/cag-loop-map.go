package main

import (
	"fmt"
	"time"
)

func fib(n int) uint64 {
	var memo = make(map[int]uint64)
	for i := 0; i <= n; i++ {
		if i <= 1 {
			memo[i] = uint64(i)
		} else {
			memo[i] = memo[i-1] + memo[i-2]
		}
	}
	return memo[n]
}

func main() {
	start := time.Now()
	fmt.Println(fib(100))
	fmt.Printf("%.3fms elapsed\n", float64(time.Since(start).Microseconds())/1000)
}
