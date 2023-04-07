package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func subtrair(c, d int) int {
	return c - d
}

func multiplica(e, f int) int {
	return e * f
}

func dividir(g, h float64) float64 {
	return g / h
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	somar := somar(2, 2)
	subtrair := subtrair(2, 2)
	multiplicar := multiplica(2, 6)
	dividir := dividir(2, 5)

	imprimir(somar)
	imprimir(subtrair)
	imprimir(multiplicar)
	imprimir(int(dividir))

}
