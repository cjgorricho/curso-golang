package main

import (
	"fmt"
	"sync"
)

// Concurrencia perfecta - se incorpora la impresiòn del i=10 dentro de las funciones concurrentes. Se debe incrementar el número de grupos de espera a 10, porque ahora se lanzan 10 funciones concurrentes.  La funciòn concurrente se declara de forma anònima (go func() {}())

var wg sync.WaitGroup

func main() {
	wg.Add(10)

	for i := 1; i <= 10; i++ {

		go func(ind int) {
			fmt.Println(ind)
			wg.Done()
		}(i)

	}
	wg.Wait()

}
