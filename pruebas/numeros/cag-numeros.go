package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*


 */

var numeros = map[string]string{
	"1":   "uno",
	"2":   "dos",
	"3":   "tres",
	"4":   "cuatro",
	"5":   "cinco",
	"6":   "seis",
	"7":   "siete",
	"8":   "ocho",
	"9":   "nueve",
	"10":  "diez",
	"11":  "once",
	"12":  "doce",
	"13":  "trece",
	"14":  "catorce",
	"15":  "quince",
	"16":  "dieciséis",
	"17":  "diecisiete",
	"18":  "dieciocho",
	"19":  "diecinueve",
	"20":  "veinte",
	"21":  "veintiuno",
	"22":  "veintidós",
	"23":  "veintitrés",
	"24":  "veinticuatro",
	"25":  "veinticinco",
	"26":  "veintiseis",
	"27":  "veintisiete",
	"28":  "veintiocho",
	"29":  "veintinueve",
	"30":  "treinta",
	"40":  "cuarenta",
	"50":  "cincuenta",
	"60":  "sesenta",
	"70":  "setenta",
	"80":  "ochenta",
	"90":  "noventa",
	"100": "cien",
	"101": "ciento",
	"200": "doscientos",
	"300": "trescientos",
	"400": "cuatrocientos",
	"500": "quinientos",
	"600": "seiscientos",
	"700": "setecientos",
	"800": "ochocientos",
	"900": "novecientos",
}

var posicion = map[int]string{
	//1e0: "unidadaes"
	1e3:  "mil",
	1e6:  "millón",
	2e6:  "millones",
	1e12: "billón",
	2e12: "billones",
	1e18: "trillón",
	2e18: "trillones",
}

var aleatorio []int

func main() {
	limite := 5
	for i := 0; i < limite; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt64)
		time.Sleep(5 * time.Microsecond)
		aleatorio = append(aleatorio, r)
	}

	fmt.Println(aleatorio)
	fmt.Println()

	for j, num := range aleatorio {

		fmt.Printf("Aleatorio %d:\n", j)
		fmt.Printf("%v\t\t%T\n", num, num)

		snum := strconv.Itoa(num)

		fmt.Println("Numero de dígitos:\t", len(snum))
		fmt.Println("Grupos de 3:\t\t", len(snum)/3)
		fmt.Println("Residuo:\t\t", len(snum)%3)

		fmt.Println()
		fmt.Println("Ascendente:")
		fmt.Print("Ind: ")
		for i := 0; i < len(snum); i++ {
			fmt.Printf("%2v|", i)
		}
		fmt.Println()
		fmt.Print("     ")
		for i := 0; i < len(snum); i++ {
			fmt.Printf("---")
		}
		fmt.Println()
		fmt.Print("Val: ")
		for i := 0; i < len(snum); i++ {
			fmt.Printf("%2s|", string(snum[i]))
		}
		fmt.Println()
		fmt.Println()

		//gdt := len(snum) / 3
		residuo := len(snum) % 3
		var b strings.Builder

		switch residuo {
		case 2:
			fmt.Println("Residuo 2")
			fmt.Fprintf(&b, "0")
		case 1:
			fmt.Println("Residuo 1")
			fmt.Fprintf(&b, "00")

		default:
			fmt.Println("Residuo 0")
		}

		b.WriteString(snum)

		fmt.Println("String ajustado: ", b.String())
		fmt.Println("Numero dígitos ajust:\t", len(b.String()))
		fmt.Println("Grupos de 3 ajustado:\t", len(b.String())/3)
		fmt.Println("Residuo ajustado:\t", len(b.String())%3)

		fmt.Println()

		astr := b.String()
		for i := 0; i < len(astr)/3; i++ {
			fmt.Println(astr[3*i : 3*i+3])
		}

		fmt.Println()

	}
}
