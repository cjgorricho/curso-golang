package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "rojo",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "azul",
		},
		luxury: true,
	}
	fmt.Print("Camion\n")
	fmt.Println("puertas: ", t.doors)
	fmt.Println("color: ", t.color)
	if t.fourWheel == true {
		fmt.Println("Es de 4 ruedas")
	}
	fmt.Println("-------------")
	fmt.Println("Sedan")
	fmt.Println("puertas: ", s.doors)
	fmt.Println("color: ", s.color)
	if s.luxury == true {
		fmt.Println("Es de lujo")
	}
	//fmt.Println(s)

}
