package main

import "fmt"

var x int

const (
	a     = 42 // Untyped constant
	b int = 43 // Typed constant
)

func main() {
	fmt.Println(a, b)
}
