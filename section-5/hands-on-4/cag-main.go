package main

import "fmt"

// 1. Create your own type. Have the underlying type be an int.

type tipo_cag int

// 2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword

var x tipo_cag

func main() {
	//		a. print out the value of the variable “x”
	fmt.Println(x)

	//		b. print out the type of the variable “x”
	fmt.Printf("%T\n", x)

	//		c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x = 42

	// 		d. print out the value of the variable “x”
	fmt.Println(x)

}
