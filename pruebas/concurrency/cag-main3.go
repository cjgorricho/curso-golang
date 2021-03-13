package main

import (
	"fmt"
	"sync"
)

// Concurrencia imperfecta - se incorpora el concepto de los grupos de espera para cada funciòn concurrente,  pero no se fija el valor de i. Cuando cada una de las funciones concurrentes va a buscar el valor de i para imprimirlo, se crea una inconsistencia de lectura i. La funciòn concurrente se declara de forma anònima (go func() {}())

var wg sync.WaitGroup

func main() {
	wg.Add(9)
	for i := 1; i <= 10; i++ {
		if i < 10 {

			go func() {
				fmt.Println(i)
				wg.Done()
			}()
		} else {
			fmt.Println(i)
		}
	}
	wg.Wait()

}
