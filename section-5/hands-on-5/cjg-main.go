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
	fmt.Println(x)

	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
