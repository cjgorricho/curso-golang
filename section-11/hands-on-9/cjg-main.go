package main

import "fmt"

func main() {
	// map m:= map[type][]type{}
	// se reemplazo `` por ""
	// se reemplazo "" por '', no funciona
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	//estructura para agregar m[nueva clave]=[]tipo{"ad1","ad2","ad3"}
	m["Carlos"] = []string{"ajedrez", "cine", "ciencia"}

	//fmt.Println(m)
	//En el rango m leer registro k
	for k, v := range m {
		fmt.Println("Este es el registro para:", k)
		// para cada v2 en el rango v leer i, v2
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
		fmt.Println()
	}

}
