package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Concurrencia perfecta - en este ejemplo se declara la funciòn concurrente de manera explìcita (func imprimir(...)) y se llama de forma explìcita en el cuerpo del programa (go imprimir(...))

var wg sync.WaitGroup
var mu sync.Mutex
var aleatorio []int32

func main() {

	limite := 10

	// Creación de números aleatorios

	for i := 1; i <= limite; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(math.MaxInt32)
		time.Sleep(5 * time.Nanosecond)
		aleatorio = append(aleatorio, r)
	}
	fmt.Println(aleatorio) // Imprime el slice con los número aleatorios
	fmt.Println()

	// Imprime las rutinas de trabajo de forma secuencial

	start := time.Now()
	fmt.Println("Imprimiendo serie secuencial")
	for i := 1; i <= limite; i++ {
		time.Sleep(time.Duration(aleatorio[i-1])) // payload
		fmt.Printf("%v\ttiempo: %v\t\n", i, time.Duration(aleatorio[i-1]))
	}
	fmt.Printf("Tiempo total secuencial: %.3f s\n\n", time.Since(start).Seconds())

	// Imprime las rutinas de trabajo de forma concurrente usando go funcs

	start1 := time.Now()
	fmt.Println("Imprimiendo serie concurrente")
	for i := 1; i <= limite; i++ {
		wg.Add(1)
		go imprimir(i, aleatorio[i-1])
	}
	wg.Wait()
	fmt.Printf("Tiempo total concurrente: %.3f s\n\n", time.Since(start1).Seconds())

}

// La función imprimer controla la concurrencia con Wait groups y con Mutexes

func imprimir(ind int, ale int32) {
	time.Sleep(time.Duration(ale)) // payload
	mu.Lock()
	fmt.Printf("%v\ttiempo: %v\t\n", ind, time.Duration(ale))
	mu.Unlock()
	wg.Done()
}
