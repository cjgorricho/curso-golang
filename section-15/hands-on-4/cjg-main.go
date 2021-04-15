package main

import (
	"fmt"
)

type persona struct {
	first  string
	last   string
	age    int
	nacion string
}

func (p persona) speak() {
	fmt.Println("yo soy", p.first, p.last, "tengo", p.age, "años y soy de", p.nacion)
	fmt.Printf("Mi nombre es %s %s, tengo  %d años y soy de %s", p.first, p.last, p.age, p.nacion)
	fmt.Println()
}
func main() {
	p1 := persona{
		first:  "Carlos J.",
		last:   "Gorricho",
		age:    78,
		nacion: "Colombia",
	}

	p1.speak()

}
