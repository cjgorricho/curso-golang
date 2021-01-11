package main

import "fmt"

// 1. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

type tipo_cag int

var x tipo_cag
var y int

func main() {
	//		a.i. print out the value of the variable “x”
	fmt.Println("Imprimiendo x...")
	fmt.Println(x)

	//		a.ii. print out the type of the variable “x”
	fmt.Printf("%T\n", x)

	//		a.iii assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x = 42

	// 		a.iv print out the value of the variable “x”
	fmt.Println(x)

	// 		b. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	//			1. then use the “=” operator to ASSIGN that value to “y”
	y = int(x)

	//			2. print out the value stored in “y”
	fmt.Println()
	fmt.Println("Imprimiendo y...")
	fmt.Println(y)

	//			3. print out the type of “y”
	fmt.Printf("%T\n", y)

}
