package main

import (
	"fmt"
)

func main() {
	m := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(m[:5]) // no incluye el 5 (que en realidad es 6)

	fmt.Println(m[5:])  //incluye el 5 (que en realidad es 6)
	fmt.Println(m[2:7]) // comienza en el 3 que es el indice 2
	fmt.Println(m[1:6]) // desde indice 1 hasta el 5 excluye el 6

}
