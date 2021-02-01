package main

import (
	"fmt"
)

func main() {
	favSport := "futbol"
	switch favSport {
	case "futbol":
		fmt.Println("ponte los guayos")
	case "natacion":
		fmt.Println("a la piscina")
	case "torear":
		fmt.Println("cuidado con el toro")
	case "tennis":
		fmt.Println("coge la raqueta")
	default:
		fmt.Println("Este es por defecto")
	}
}
