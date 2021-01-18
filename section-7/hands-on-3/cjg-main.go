package main

import (
	"fmt"
)

const (
	x int = 1242
	y     = "Carlos"
)

func main() {

	fmt.Println(x, y)
	fmt.Printf("%T, %T", x, y)
}
