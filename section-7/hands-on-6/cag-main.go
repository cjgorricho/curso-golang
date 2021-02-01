package main

import "fmt"

const (
	a = 2017 + iota
	b
	c
	d
	v = 1
	w = v << iota // si se declara iota nuevamente, con al menos una lìnea de diferencia a la ùltima declaraciòn, se reinicia el ìndice desde 0
	x
	y
	z
	aa
	ab
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println()
	fmt.Println(v)
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(aa)
	fmt.Println(ab)
}
