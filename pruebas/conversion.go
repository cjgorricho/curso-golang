package main

import (
	"fmt"
)

func main() {
	y := 17
	s := fmt.Sprintf("Decimal : %d \nBinario: %b \nHex: %X \n", y, y, y)
	fmt.Println(s)
}
