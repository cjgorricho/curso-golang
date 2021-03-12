package main

import (
	"fmt"
	"sync"
)

// Concurrencia perfecta - se incorporan los grupos de espera para cada funciòn concurrente. Son nueve funciones concurrentes entonces solo necesito 9 grupos de espera. Adicionalmente se pasa la variable i como un argumento de la funciòn concurrente, de manera que se "fija" para cada funciòn.  La funciòn concurrente se declara de forma anònima (go func() {}())

var wg sync.WaitGroup

func main() {
	wg.Add(9)
	for i := 1; i <= 10; i++ {
		if i < 10 {

			go func(ind int) {
				fmt.Println(ind)
				wg.Done()
			}(i)
		} else {
			fmt.Println(i)
		}
	}
	wg.Wait()

}
