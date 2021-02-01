package main

import (
	"fmt"
)

func main() {
	// COMPOSITE LITERAL hace referencia a la definiciòn integral del SLICE: tipo + contenido: []tipo{contenido}
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // sabemos que es un SLICE porque el tamaño NO està definido
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
