package main

import "fmt"

var x int

func main() {
	a := 42
	fmt.Printf("%[1]d\t%[1]b\t%#[1]X\n", a)
	b := a << 1
	fmt.Printf("%[1]d\t%[1]b\t%#[1]X\n", b)
	c := a >> 1
	fmt.Printf("%[1]d\t%[1]b\t%#[1]X\n", c)

	fmt.Println()

	x := 1
	var bins []int
	for i := 0; i <= 10; i++ {
		bins = append(bins, x<<i)
		fmt.Printf("%5[1]d\tBin: %-16[1]b\tHex: %#[1]X\n", bins[i])
	}
}
