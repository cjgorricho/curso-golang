package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Mi nombre es %s %s y tengo  %d años", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Carlos A",
		last:  "Gorricho",
		age:   53,
	}
	p1.speak()

}
