package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {

		if i%3 == 0 {
			fmt.Printf("%#U\n", i)
			fmt.Println(i)
		} else if i%2 != 0 {
			fmt.Printf("%#U\t", i)
			fmt.Println(i)
		} else if i%2 == 1 {
			fmt.Println("Esta imprime")

		}
	}
}
