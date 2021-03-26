package main

import (
	"fmt"
)

func main() {
	m := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 98, 3000}

	for i, v := range m {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", m)
	fmt.Println(len(m), cap(m))
}
