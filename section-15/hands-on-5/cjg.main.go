package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
	largo float64
	ancho float64
}

func (u cuadrado) area() float64 {
	return u.largo * u.ancho
}

type circulo struct {
	radio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

type forma interface {
	area() float64
	//perimetro() float64
}

func informe(f forma) {
	fmt.Println("El area del circulo es: ", c1.area())
	fmt.Println("El area del cuadrado es: ", s1.cuadrado)
}

func main() {
	s1 := cuadrado{
		largo: 4,
		ancho: 5,
	}

	c1 := circulo{
		radio: 10.5,
	}
	informe(c1)
	informe(s1)
}
