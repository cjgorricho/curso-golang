package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TO DO: incorporar màs funciones de creaciòn y transformaciòn de matrices
// TO DO: incorporar concurrencia tanto en la creaciòn como en la transfomaciòn de matrices

func randmat(mat [][]int, rows, cols int) [][]int {
	mat = make([][]int, rows)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(rows)
	for i := 0; i < rows; i++ {
		go func(ind int) {
			var row []int
			for j := 0; j < cols; j++ {
				rand.Seed(time.Now().UnixNano())
				time.Sleep(5 * time.Millisecond)
				value := rand.Intn(5) + 1
				row = append(row, value)
			}
			mu.Lock()
			mat[ind] = row
			mu.Unlock()
			wg.Done()
		}(i)
		time.Sleep(5 * time.Millisecond)
	}
	wg.Wait()
	return mat
}

func ones(mat [][]int, rows, cols int) [][]int {
	mat = make([][]int, rows)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(rows)
	for i := 0; i < rows; i++ {
		//fmt.Printf("wg %d is added for mat of %v rows contains %v tipo %T\n", i, rows, wg, wg)
		go func(ind int) {
			var row []int
			for j := 0; j < cols; j++ {
				row = append(row, 1)
			}
			mu.Lock()
			mat[ind] = row
			mu.Unlock()
			wg.Done()
			//fmt.Printf("wg %d is done for mat of %v rows contains %v and the row %v\n", i-1, rows, wg, row)

		}(i)

	}

	wg.Wait()
	//fmt.Printf("matriz de %v filas es %v\n", rows, mat)
	return mat
}

func mult1(mat, mat1, mat2 [][]int) [][]int {
	var sto = make(map[int][]int) //variable global a todas las go routines
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(mat1))
	for i, m1 := range mat1 {

		go func(ind int, m1 []int, mt2 [][]int) {
			var row []int
			row = append(row, ind)
			for k := 0; k < len(mt2[0]); k++ {
				pos := 0
				comp := 0
				for j, m2 := range mt2 {
					comp = m1[j] * m2[k]
					pos = pos + comp
					//fmt.Println(ind+1, k+1, comp)
					//if j+1 >= len(mt2) {
					//	fmt.Println(ind+1, k+1, pos)
					//}
				}
				row = append(row, pos)
				//fmt.Println(row, len(row))
			}
			mu.Lock()
			sto[row[0]] = row[1:]
			mu.Unlock()
			wg.Done()
		}(i, m1, mat2)
	}
	wg.Wait()
	//fmt.Println(sto)
	mat = make([][]int, len(mat1)) // inicializa la matriz resultados con 0 en las dimensiones n * m necesarias
	for k, v := range sto {
		mat[k] = v
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
				//if j+1 >= len(mat2) {
				//	fmt.Println(i+1, k+1, pos)
				//}
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
	m1 = ones(m1, 50, 30)
	tm1 := time.Since(start).Milliseconds()
	m2 = ones(m2, 30, 10)
	tm2 := time.Since(start).Milliseconds()
	m3 = mult1(m3, m1, m2)
	tm3 := time.Since(start).Milliseconds()
	fmt.Printf("\nTiempo matriz 1: %v ms\nTiempo matriz 2: %v ms\nTiempo matriz 3: %v ms\nTiempo total: %v ms\n\n", tm1, tm2-tm1, tm3-tm2, tm3)
	printmat(m1, m2, m3)
}
