package main

import (
	"fmt"
)

// Tomado directamente de los ejercicios del curso.
// Este ejemplo es importante para entender manejo de matrices

func main() {
	// xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	// xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	// fmt.Println(xs1)
	// fmt.Println(xs2)

	// Una matriz de dos dimensiones se puede crear con el còdigo de arriba, o se puede crear directamente, como un conjunto de SLICES unidimensionales

	xxs := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
