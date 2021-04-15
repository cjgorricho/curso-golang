package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/*


 */

var numeros = map[int]string{
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
	24:   "veinticuatro",
	25:   "veinticinco",
	26:   "veintiseis",
	27:   "veintisiete",
	28:   "veintiocho",
	29:   "veintinueve",
	30:   "treinta",
	40:   "cuarenta",
	50:   "cincuenta",
	60:   "sesenta",
	70:   "setenta",
	80:   "ochenta",
	90:   "noventa",
	100:  "cien",
	101:  "ciento",
	200:  "doscientos",
	300:  "trescientos",
	400:  "cuatrocientos",
	500:  "quinientos",
	600:  "seiscientos",
	700:  "setecientos",
	800:  "ochocientos",
	900:  "novecientos",
	1e3:  "mil",
	1e6:  "millón",
	2e6:  "millones",
	1e12: "billón",
	//2e12: "billones",
	//1e18: "trillón",
	//2e18: "trillones",
}

var num int
var bi int

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt64)
	time.Sleep(1 * time.Nanosecond)
	num = r

	fmt.Println("num: ", num)
	snum := strconv.Itoa(num)

	//fmt.Printf("%v\t%T\n", snum, snum)
	tri := snum[len(snum)-3 : len(snum)]
	fmt.Printf("%v\t%T\n ", tri, tri)
	rest := snum[:len(snum)-3]
	fmt.Printf("%v\t%T\n", rest, rest)

	tri1 := tri[0:1] + "00"
	tri1int, err := strconv.Atoi(tri1)
	if err != nil {
		fmt.Println(err)
	}

	tri2 := tri[1:2] + "0"
	tri2int, err := strconv.Atoi(tri2)
	if err != nil {
		fmt.Println(err)
	}
	tri3 := string(tri[2])
	tri3int, err := strconv.Atoi(tri3)
	if err != nil {
		fmt.Println(err)
	}

	lit1 := numeros[tri1int] + " " + numeros[tri2int] + " y " + numeros[tri3int]
	//fmt.Println("tri1: ", tri1int)
	//fmt.Println("tri2: ", tri2int)
	//fmt.Println("tri3: ", tri3int)
	fmt.Println("lit1: ", lit1)

}
