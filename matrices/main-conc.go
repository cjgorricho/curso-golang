package main

import (
	"fmt"
	"time"
)

// TO DO: incorporar màs funciones de creaciòn y transformaciòn de matrices
// TO DO: incorporar concurrencia tanto en la creaciòn como en la transfomaciòn de matrices

func zeros(mat [][]int, rows, cols int) [][]int {
	c := make(chan []int, rows)
	fmt.Println(cap(c))
	for i := 0; i < rows; i++ {

		go func() {
			var row []int
			for j := 0; j < cols; j++ {
				row = append(row, 0)
			}
			c <- row
		}()
		mat = append(mat, <-c)
	}

	return mat
}

func ones(mat [][]int, rows, cols int) [][]int {
	c := make(chan []int, rows)
	fmt.Println(cap(c))
	for i := 0; i < rows; i++ {

		go func() {
			var row []int
			for j := 0; j < cols; j++ {
				row = append(row, 1)
			}
			c <- row
		}()
		mat = append(mat, <-c)

	}

	return mat
}

func mult(mat, mat1, mat2 [][]int) [][]int {
	c := make(chan []int, len(mat1))
	var sto = make(map[int][]int)
	for i, mt1 := range mat1 {

		go func() {
			var row []int
			row = append(row, i)
			for k := 0; k < len(mat2[0]); k++ {
				pos := 0
				for j, mt2 := range mat2 {
					pos = pos + mt1[j]*mt2[k]
				}
				row = append(row, pos)
			}
			c <- row
		}()
		st := <-c
		sto[st[0]] = st[1:]
	}
	mat = zeros(mat, len(mat1), len(mat2[0])) // inicializa la matriz resultados con 0 en las dimensiones n * m necesarias
	for k, v := range sto {
		mat[k] = v
	}

	return mat
}

func main() {
	start := time.Now()
	var m1, m2 [][]int
	var m3 [][]int
	m1 = ones(m1, 50, 30)
	m2 = ones(m2, 30, 50)
	m3 = mult(m3, m1, m2)
	fmt.Printf("\n%.3fms elapsed\n\n", float64(time.Since(start).Microseconds())/1000)
	//fmt.Printf("Matriz 1\nTipo: %[1]T\nTamaño :%v\nCapacidad :%v\n", m1, len(m1), cap(m1))
	//fmt.Printf("\nMatriz\n")
	//for i, mat := range m1 {
	//	fmt.Printf("%d\t%v\n", i, mat)
	//}
	//fmt.Printf("\nMatriz 2\nTipo: %[1]T\nTamaño :%v\nCapacidad :%v\n", m2, len(m2), cap(m2))
	//fmt.Printf("\nMatriz\n")
	//for i, mat := range m2 {
	//	fmt.Printf("%d\t%v\n", i, mat)
	//}
	//fmt.Printf("\nMatriz 3\nTipo: %[1]T\nTamaño :%v\nCapacidad :%v\n", m3, len(m3), cap(m3))
	//fmt.Printf("\nMatriz\n")
	//for i, mat := range m3 {
	//	fmt.Printf("%d\t%v\n", i, mat)
	//}

}
