package main

import (
	"fmt"
)

// Concurrencia incompleta - las funciones se lanzan al àmbito concurrente anteponiendo la palabra "go", pero el programa principal no espera a su culminaciòn

func main() {

	for i := 1; i <= 10; i++ {
		if i < 10 {

			go fmt.Println(i)

		} else {
			fmt.Println(i)
		}
	}

}
