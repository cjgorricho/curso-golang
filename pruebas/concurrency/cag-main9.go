package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Concurrencia perfecta - en este ejemplo se declara la funciòn concurrente de manera explìcita (func imprimir(...)) y se llama de forma explìcita en el cuerpo del programa (go worker(...))

var aleatorio []int32

type data struct {
	ind      int
	duracion time.Duration
}

var comms = make(chan data)

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

	// este bloque del programa imprime el resultado de las rutinas de trabajo (workers) de forma concurrente

	start1 := time.Now()
	fmt.Println("Imprimiendo serie concurrente")
	for i := 1; i <= limite; i++ {
		go worker(i, aleatorio[i-1], comms)
	}

	for i := 1; i <= limite; i++ {
		datosrecv := <-comms
		fmt.Printf("%v\ttiempo: %v\t\n", datosrecv.ind, datosrecv.duracion)
	}

	fmt.Printf("Tiempo total concurrente: %.3f s\n\n", time.Since(start1).Seconds())

}

// La función worker simula una subrutina con una tarea especifica. Utilica un canal de comunicaciones llamado comms

func worker(ind int, ale int32, comms chan data) {
	time.Sleep(time.Duration(ale)) // payload
	datos := data{
		ind:      ind,
		duracion: time.Duration(ale),
	}
	comms <- datos
}
