package main

import (
	"fmt"
)

func main() {
	g := 42 == 42
	h := 42 <= 40
	i := 42 >= 40
	j := 42 != 40
	k := 42 < 40
	l := 42 > 40

	fmt.Println(g, h, i, j, k, l)

}
