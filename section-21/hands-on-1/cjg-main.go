package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		fmt.Println("CPU inicial :", runtime.NumCPU())
		fmt.Println("gs Iniciales:", runtime.NumGoroutine())

		for i := 0; i < 3; i++ {
			fmt.Println("Imprime Uno :", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 3; i < 6; i++ {
			fmt.Println("Imprime Dos", i)
		}
		wg.Done()
	}()

	fmt.Println("CPU medio :", runtime.NumCPU())
	fmt.Println("gs medio: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Salida")
	fmt.Println("CPU final :", runtime.NumCPU())
	fmt.Println("gs final:", runtime.NumGoroutine())
}
