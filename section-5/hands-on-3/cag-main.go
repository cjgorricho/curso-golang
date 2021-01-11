package main

import "fmt"

// 1. At the package level scope, assign the following values to the three variables

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// 2. in func main
	//		a. print out the values for each identifier

	s := fmt.Sprintf("x: %v\ny: %v\nz: %v\n", x, y, z)

	//		b. print out the value stored by variable “s”

	fmt.Println(s)

}
