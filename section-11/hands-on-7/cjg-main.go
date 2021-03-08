package main

import (
	"fmt"
)

func main() {
	dato1 := []string{"James", "Bond", "Shaken, not stirred"}
	dato2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(dato1, dato2)
	datos := [][]string{dato1, dato2}
	fmt.Println(datos)
	fmt.Println("Valores", len(datos), cap(datos))
	//corre i dos veces porque hay 2 strings en datos
	for i, x := range datos {
		fmt.Println("registros: ", i)
		//por cada corrida de i, corre j tres veces porque cada slice tiene tres datos
		fmt.Println(len(dato1), len(dato2))
		for j, val := range x {
			fmt.Printf("\t posicion: %v \t valor: %v \n", j, val)
		}
	}
}
