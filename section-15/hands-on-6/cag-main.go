package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	func() {
		fmt.Printf("\nSerie para armar sto1\n")
		total := 0
		var sto1 []int
		var sto2 = make(map[int][]int)
		for i, v := range x {
			total += v
			sto1 = append(sto1, total)
			sto2[i] = sto1
			fmt.Println(i, total)
		}
		// Imprime sto1 ([]int)
		fmt.Printf("\nSlice st1\n")
		fmt.Println(sto1)

		// Imprime sto2 (MAP[int][]int) en desorden. A Map us an UNORDERED set of key:value pairs
		fmt.Printf("\nSlice sto2 sin ordenar\n")
		for k, v := range sto2 {
			fmt.Println(k, v)
		}

		// Crea un sto2keys ([]int) para almacenar los KEYS de sto2. Luego ordena sto2 keys. Finalmente imprime los valores de sto2 a partir de sto2keys ordenado
		fmt.Printf("\nSlice sto2 ordenado\n")
		sto2keys := make([]int, 0, len(sto2))
		for k := range sto2 {
			sto2keys = append(sto2keys, k)
		}
		sort.Ints(sto2keys)

		for _, k := range sto2keys {
			fmt.Println(k, sto2[k])
		}
	}()
	fmt.Println()
	fmt.Println("Fin de la funciòn")
}
