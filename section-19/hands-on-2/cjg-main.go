package main

import (
	"fmt"
)

type person struct {
	name string
}
type human interface {
	speak()
}

func (p *person) speak() {
	//saySomething()
	fmt.Println("Hola amigos")

}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Carlos",
	}
	saySomething(&p1)

}
