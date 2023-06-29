package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2    // se 1 e 2 for correto
	comprarTv32 := trab1 != trab2    // ou exclusivo
	comprarSorvete := trab1 || trab2 // ou

	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("\n TV50: %t", tv50)
	fmt.Printf("\n TV32: %t", tv32)
	fmt.Printf("\n Sorvete: %t", sorvete)
	fmt.Printf("\n Saudável: %t", !sorvete)
	fmt.Println()

	tv501, tv321, sorvete1 := compras(true, false)
	fmt.Printf("\n TV50: %t", tv501)
	fmt.Printf("\n TV32: %t", tv321)
	fmt.Printf("\n Sorvete: %t", sorvete1)
	fmt.Printf("\n Saudável: %t", !sorvete1)
	fmt.Println()

	tv502, tv322, sorvete2 := compras(false, false)
	fmt.Printf("\n TV50: %t", tv502)
	fmt.Printf("\n TV32: %t", tv322)
	fmt.Printf("\n Sorvete: %t", sorvete2)
	fmt.Printf("\n Saudável: %t", !sorvete2)
	fmt.Println()

	tv503, tv323, sorvete3 := compras(false, true)
	fmt.Printf("\n TV50: %t", tv503)
	fmt.Printf("\n TV32: %t", tv323)
	fmt.Printf("\n Sorvete: %t", sorvete3)
	fmt.Printf("\n Saudável: %t", !sorvete3)

}
