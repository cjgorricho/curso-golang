package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Concurrencia perfecta - en este ejemplo se declara la funciòn concurrente de manera explìcita (func imprimir(...)) y se llama de forma explìcita en el cuerpo del programa (go imprimir(...))

var wg sync.WaitGroup
var aleatorio []int32

func main() {

	for i := 1; i <= 10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
		time.Sleep(5 * time.Millisecond)
		aleatorio = append(aleatorio, r)
	}
	fmt.Println(aleatorio)
	fmt.Println()

	start := time.Now()
	fmt.Println("Imprimiendo serie secuencial")
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(aleatorio[i-1])) // payload
		fmt.Printf("%v\ttiempo: %v\t\n", i, time.Duration(aleatorio[i-1]))
	}
	fmt.Printf("Tiempo total secuencial: %.3f s\n\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println("Imprimiendo serie concurrente")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go imprimir(i, aleatorio[i-1])
	}
	wg.Wait()
	fmt.Printf("Tiempo total concurrente: %.3f s\n\n", time.Since(start).Seconds())

}

func imprimir(ind int, ale int32) {
	time.Sleep(time.Duration(ale)) // payload
	fmt.Printf("%v\ttiempo: %v\t\n", ind, time.Duration(ale))
	wg.Done()
}
