package main

import "fmt"

type bar struct {
	text string
	age  int
}

func main() {
	x := foo()
	fmt.Println(x)
	y1, y2 := bar1()
	fmt.Println(y1, y2)
	z1, z2 := bar2()
	fmt.Println(z1, z2)

}

func foo() int {
	x := 53
	return x
}

// en la siguiente funciòn se definen los TYPE de los paràmetros que revuelve la funciòn. Las variables correspondientes deben ser inicializadas con el mètodo corto (con :=). El RETURN debe contener el nombre de las variables inicializadas

func bar1() (string, int) {
	texto := "carlos tiene"
	edad := 53
	return texto, edad
}

// en la siguiente funciòn se definen los NOMBRES y los TYPE de los paràmetros que revuelve la funciòn. Las variables correspondientes son INICIALIZDAS por el encabezado de la funciòn, en consecuencia solo se define su valor (con =). El RETURN se escribe solo, sin el nombre de las variables que devuelve la funciòn

func bar2() (texto string, edad int) {
	texto = "ahora carlos tiene"
	edad = 53
	return
}
