package main

import (
	"fmt"
)

func main() {
	// COMPOSITE LITERAL hace referencia a la definiciòn integral del ARRAY: tamaño tipo + contenido: [tamaño]tipo{contenido}
	x := [5]int{10, 20, 30, 40, 50} // sabemos que es un ARRAY porque el tamaño està definido dentro de []
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x) // el TYPE incluye el tamaño del ARRAY
}
