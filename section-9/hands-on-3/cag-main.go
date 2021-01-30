package main

import (
	"fmt"
)

func main() {
	bd := 1968
	edad := 0
	for bd <= 2021 {
		fmt.Println(bd)
		bd++
		edad++
	}
	fmt.Println("Edad: ", edad-1)
}
