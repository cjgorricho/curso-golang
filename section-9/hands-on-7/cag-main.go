package main

import (
	"fmt"
)

func main() {
	x := "Carlos A Gorricho"

	if x == "Juan" {
		fmt.Println(x)
	} else if x == "Carlos A Gorricho" {
		fmt.Println("(else if)", x)
	}
}
