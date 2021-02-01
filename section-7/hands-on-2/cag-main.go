package main

import "fmt"

var x int

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Printf("%-5[1]v - Type: %[1]T\n", a)
	fmt.Printf("%-5[1]v - Type: %[1]T\n", b)
	fmt.Printf("%-5[1]v - Type: %[1]T\n", c)
	fmt.Printf("%-5[1]v - Type: %[1]T\n", d)
	fmt.Printf("%-5[1]v - Type: %[1]T\n", e)
	fmt.Printf("%-5[1]v - Type: %[1]T\n", f)
}
