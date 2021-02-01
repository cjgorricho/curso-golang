package main

import (
	"fmt"
)

func main() {
	i := 10

	for {
		if i > 100 {
			break
		}
		if i%4 == 0 {
			fmt.Println(i)
		}

		i++
	}
}
