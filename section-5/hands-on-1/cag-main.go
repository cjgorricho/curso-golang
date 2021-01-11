package main

import "fmt"

func main() {

	//Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
	//	a. 42
	//	b. “James Bond”
	//	c. true

	x := 42
	y := "James Bond"
	z := true

	// 2. Now print the values stored in those variables using
	//		a. a single print statement

	fmt.Println("Single print statement 1 - fmt.Printf")
	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)
	fmt.Println()

	fmt.Println("Single print statement 2 - fmt.Sprintln")
	s := fmt.Sprintf("x: %v\ny: %v\nz: %v\n", x, y, z)
	fmt.Println(s)

	//		b. multiple print statements

	fmt.Println("Multiple print statement - fmt.Println")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

}
