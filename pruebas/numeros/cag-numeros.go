package main

import (
	"fmt"
	"math/rand"
	"time"
)

var aleatorio []int64

/*
Todo lo que està en este bloque queda como un comenario de varias lìneas
*/

var numeros = map[int64]string{
	1:    "uno",
	2:    "dos",
	3:    "tres",
	4:    "cuatro",
	5:    "cinco",
	6:    "seis",
	7:    "siete",
	8:    "ocho",
	9:    "nueve",
	10:   "diez",
	11:   "once",
	12:   "doce",
	13:   "trece",
	14:   "catorce",
	15:   "quince",
	16:   "dieciséis",
	17:   "diecisiete",
	18:   "dieciocho",
	19:   "diecinueve",
	20:   "veinte",
	21:   "veintiuno",
	22:   "veintidós",
	23:   "veintitrés",
	24: 		"veinticuatro",
	25: 		"veinticinco",
	26: "veintiseis",
	27: "veintisiete",
	28: ""

	30:   "treinta",
	40:   "cuarenta",
	50:   "cincuenta",
	60:   "sesenta",
	70:   "setenta",
	80:   "ochenta",
	90:   "noventa",
	100:  "cien",
	200:  "doscientos",
	300:  "trescientos",
	400:  "cuatrocientos",
	500:  "quinientos",
	600:  "seiscientos",
	700:  "setecientos",
	800:  "ochocientos",
	900:  "novecientos",
	1000: "mil",
	1e6:  "millón",
	1e12: "billón",
}

func main() {
	limite := 10

	for i := 1; i <= limite; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
		time.Sleep(5 * time.Millisecond)
		aleatorio = append(aleatorio, r)
	}
	fmt.Println(aleatorio)
	fmt.Println()

	fmt.Println()

}
