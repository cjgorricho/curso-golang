package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {

		if i%3 == 0 {

			fmt.Printf("%#U\n", i)
		}
	}
}
