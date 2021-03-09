package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TO DO: incorporar màs funciones de creaciòn y transformaciòn de matrices
// TO DO: incorporar concurrencia tanto en la creaciòn como en la transfomaciòn de matrices

func randmat(mat [][]int, rows, cols int) [][]int {
	for i := 0; i < rows; i++ {
		var row []int
		for j := 0; j < cols; j++ {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(5 * time.Millisecond)
			value := rand.Intn(5) + 1
			row = append(row, value)
		}
		mat = append(mat, row)
	}
	return mat
}

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
				//fmt.Println(i, j, pos)
			}
			row = append(row, pos)
		}
		mat = append(mat, row)
	}
	//fmt.Println(mat)
	return mat
}

func printmat(m ...[][]int) {
	for m, mat := range m {
		fmt.Printf("\nMatriz %v\nTipo: %T\nTamaño :%v\nCapacidad :%v\n\n", m+1, mat, len(mat), cap(mat))

		if len(mat) <= 50 && len(mat[0]) <= 50 {
			fmt.Printf("Matriz\n")
			for i, mt := range mat {
				fmt.Printf("%d\t%v\n", i, mt)
			}
		}
	}
}

func main() {
	start := time.Now()
	var m1, m2 [][]int
	var m3 [][]int
	m1 = ones(m1, 500, 30000)
	tm1 := time.Since(start).Milliseconds()
	m2 = ones(m2, 30000, 100)
	tm2 := time.Since(start).Milliseconds()
	m3 = mult(m3, m1, m2)
	tm3 := time.Since(start).Milliseconds()
	fmt.Printf("\nTiempo matriz 1: %v ms\nTiempo matriz 2: %v ms\nTiempo matriz 3: %v ms\nTiempo total: %v ms\n\n", tm1, tm2-tm1, tm3-tm2, tm3)
	printmat(m1, m2, m3)
}
