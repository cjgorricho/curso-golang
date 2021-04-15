package main

import (
	"fmt"
	"time"
)

func fib(n int) int {

	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

func main() {
	start := time.Now()
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(40))
	fmt.Printf("%dms elapsed\n", time.Since(start).Milliseconds())

}
