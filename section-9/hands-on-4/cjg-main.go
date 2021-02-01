package main

import (
	"fmt"
)

func main() {
	i := 1942

	for {
		if i > 2020 {
			break
		}
		fmt.Println(i, i-1942)
		i++
	}
}
