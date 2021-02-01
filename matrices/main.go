package main

import (
	"fmt"
	"time"
)

func ones(mat [][]int, rows, cols int) [][]int {
	for i := 0; i < rows; i++ {
		var row []int
		for j := 0; j < cols; j++ {
			row = append(row, 1)
		}
		mat = append(mat, row)
	}
	return mat
}

func mult(mat, mat1, mat2 [][]int) [][]int {

	for _, mt1 := range mat1 {
		var row []int
		for k := 0; k < len(mat2[0]); k++ {
			pos := 0
			for j, mt2 := range mat2 {
				pos = pos + mt1[j]*mt2[k]
			}
			row = append(row, pos)
		}
		mat = append(mat, row)
	}
	return mat
}

func main() {
	start := time.Now()
	var m1, m2 [][]int
	var m3 [][]int
	m1 = ones(m1, 20, 300)
	m2 = ones(m2, 300, 10)
	m3 = mult(m3, m1, m2)
	fmt.Printf("Matriz 1\nTipo: %[1]T\nTamaño :%v\nCapacidad :%v\n", m1, len(m1), cap(m1))
	fmt.Println()
	fmt.Printf("Matriz 2\nTipo: %[1]T\nTamaño :%v\nCapacidad :%v\n", m2, len(m2), cap(m2))
	fmt.Println()
	fmt.Printf("Matriz 3\nTipo: %[1]T\nTamaño :%v\nCapacidad :%v\n", m3, len(m3), cap(m3))
	fmt.Printf("\nMatriz resultado\n")
	for i, mat := range m3 {
		fmt.Printf("%d\t%v\n", i, mat)
	}
	fmt.Printf("\n%.3fms elapsed\n", float64(time.Since(start).Microseconds())/1000)

}
