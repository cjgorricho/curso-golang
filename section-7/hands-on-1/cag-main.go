package main

import "fmt"

var x int

func main() {
	x = 42

	// Para resolver este problema estoy utilizando funciones avanzadas de formato de los VERBOS
	// - el ancho del espacio para imprimir (por ejemplo 8 despuès de %)
	// - hacer referencia al primer paràmetro a imprimir ([1]), lo que permite no tener que repetir la misma variable muchas veces
	// - usar el signo "-" so sequiere alinear el valor impreso dentro del ancho asignado a la derecha o a la izquierda.
	//Todos estos paràmetros deben ir antes de el tipo de variable que se quiere imprimir

	fmt.Printf("Dec: %8[1]d\nBin: %8[1]b\nHex: %8[1]X\n", x)
	fmt.Println()
	fmt.Printf("Dec: %-8[1]d\nBin: %-8[1]b\nHex: %-8[1]X\n", x)

}
