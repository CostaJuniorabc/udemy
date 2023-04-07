package main

import (
	"fmt"
	m "math" //m foi usado para criar um atalho para o import
)

func main() {

	const PI float64 = 3.1415 // para criar uma constante
	var raio = 3.2            // tipo float64

	//forma reduzida para criar uma variavel
	area := PI * m.Pow(raio, 2)

	fmt.Println("Area =", area)

	//outra forma de declarar constantes
	const (
		a = 1
		b = 2
		c = 3
	)

	// outra forma de declarar variaveis
	var (
		d = 4
		e = 5
		f = 6
	)

	fmt.Println("constantes =", a, b, c)
	fmt.Println("Variaveis =", d, e, f)

	//outro modo de declarar variavel
	var g, h bool = true, false
	fmt.Println(g, h)

	// outro modo de declarar
	i, j, k := 2, false, "epa!"
	fmt.Println(i, j, k)

}
