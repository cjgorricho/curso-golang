package main

import "fmt"

// crear tipo propio

type cjg int

// crear variable con identificador x y el nuevo tipo

var x cjg

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//fmt.Println(s)
	x = 42
	fmt.Println(x)

}
