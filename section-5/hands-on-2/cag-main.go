package main

import "fmt"

// 1. Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).

var x int
var y string
var z bool

func main() {
	// 2. in func main
	//		a. print out the values for each identifier

	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)

	// 		b. The compiler assigned values to the variables. What are these values called? -- ZERO VALUES

}
