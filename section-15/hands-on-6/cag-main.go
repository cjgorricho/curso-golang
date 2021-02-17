package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	func() {
		total := 0
		var sto1 []int
		var sto2 = make(map[int][]int)
		for i, v := range x {
			total += v
			sto1 = append(sto1, total)
			sto2[i] = sto1
			fmt.Println(i, total)
		}
		fmt.Println(sto1)

		sto2keys := make([]int, 0, len(sto2))
		for k := range sto2 {
			sto2keys = append(sto2keys, k)
		}
		sort.Ints(sto2keys)

		for _, k := range sto2keys {
			fmt.Println(k, sto2[k])
		}
	}()
	fmt.Println("Fin de la funciòn")
}
