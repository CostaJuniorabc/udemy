package main

import (
	"fmt"
	"math"
)

func main() {

	var numero1 = 321
	numero2 := 123
	a := 3
	b := 2
	c := 3.0
	d := 2.0

	fmt.Println("Valor da Soma =>", numero1+numero2) // soma dois numeros

	fmt.Println("Valor da Subtraçõa =>", numero1-numero2) // subtrai 2 numeros

	fmt.Println("Valor da Multiplicação =>", numero1*numero2) // multiplica 2 numeros

	fmt.Println("Valor da Divisão =>", numero1/numero2) // divide 2 numeros

	fmt.Println("Valor do Modulo =>", numero1%numero2) // resto da divisão

	// operações bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("OU =>", a|b)  // 11 | 10 = 01
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01

	// outras operacoes usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
