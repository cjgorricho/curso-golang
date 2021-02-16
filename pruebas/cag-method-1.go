package main

import (
	"fmt"
	"strings"
)

type txttransform struct {
	text         string
	texttoremove string
}

func (s txttransform) trimany() string {
	tx := strings.Trim(s.text, s.texttoremove)
	return tx
}

type text string

func (s text) trim() string {
	tx := strings.Trim(string(s), " ")
	return tx
}

func main() {
	txt1 := txttransform{
		text:         `  *+*+* este es el texto limpio *+*+*+*  `,
		texttoremove: " *+",
	}
	txt2 := text(`    este es otro texto limpio     `)

	// fmt.Println(txt1)
	// fmt.Printf("Type %T:\n", txt1)
	fmt.Println(txt1.trimany())
	fmt.Printf("%T\n", txt1.trimany())
	fmt.Println(txt2.trim())
	fmt.Printf("%T\n", txt2.trim())

}
