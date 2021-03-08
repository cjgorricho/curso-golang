package main

import (
	"fmt"
)

//array
func main() {
	m := [8]string{
		"Carlos",
		"Juan",
		"Julia",
		"Ines",
		"Jesus",
		"antonio",
		"Federico",
		""}
	fmt.Println(len(m)) // da len del array = a cap
	fmt.Println(cap(m))
	fmt.Println((m))

	for i, v := range m {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", m)
}
