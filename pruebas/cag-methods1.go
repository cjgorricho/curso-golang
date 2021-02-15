package main

import (
	"fmt"
	"strings"
)

type txttransform struct {
	text         string
	texttoremove string
}

func (s txttransform) trimany(tx, txtr string) string {
	tx = strings.Trim(tx, txtr)
	return tx
}

func (s txttransform) trimspace(tx string) string {
	tx = strings.Trim(tx, " ")
	return tx
}

func main() {
	txt1 := txttransform{
		text:         `  este es el texto  `,
		texttoremove: " *+",
	}
	txt2 := txttransform{
		text: `  este es otro texto  `,
	}

	fmt.Println(txt1)
	fmt.Printf("Type %T:\n", txt1)
	fmt.Println(txt1.trimany(txt1.text, txt1.texttoremove))
	fmt.Println(txt2.trimspace(txt2.text))

}
