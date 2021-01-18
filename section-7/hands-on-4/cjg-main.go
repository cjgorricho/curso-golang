package main

import "fmt"

func main() {

	x := 42
	//w := 5
	fmt.Printf("%8[1]d\n%8[1]b\n%#8[1]x\n", x)
	//fmt.Printf("%[1]d\t%[1]b\t%[1]x\n%[2]d\t%[2]b\t%[2]x\n", x, w)
	fmt.Printf("%-8[1]d\n%-8[1]b\n%#-8[1]x\n", x)
	//y := x << 1
	//fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	//z := x >> 1
	//fmt.Printf("%d\t%b\t%#x", z, z, z)

}
