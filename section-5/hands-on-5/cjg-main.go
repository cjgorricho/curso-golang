package main

import "fmt"

// crear tipo propio

type cjg int

// crear variable con identificador x y el nuevo tipo

var x cjg
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//fmt.Println(s)
	x = 137
	fmt.Println("x =", x)

	y = int(x)
	fmt.Println("y =", y)
	fmt.Printf("type = %T\n", y)
}
