package main

import (
	"fmt"
	"sync"
)

// Concurrencia perfecta - en este ejemplo se declara la funciòn concurrente de manera explìcita (func imprimir(...)) y se llama de forma explìcita en el cuerpo del programa (go imprimir(...))

var wg sync.WaitGroup

func main() {
	wg.Add(10)

	for i := 1; i <= 10; i++ {

		go imprimir(i)

	}
	wg.Wait()

}

func imprimir(ind int) {
	fmt.Println(ind)
	wg.Done()
}
