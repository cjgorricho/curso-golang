package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v dividido por 4, tiene un residuo de %v\t%.3f\n", i, i%4, float64(i)/4)
	}
}
