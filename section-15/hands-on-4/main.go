package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %s %s and my age is %d", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Carlos A",
		last:  "Gorricho",
		age:   53,
	}
	p1.speak()

}
