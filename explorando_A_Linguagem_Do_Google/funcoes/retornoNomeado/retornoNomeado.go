package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	primeiro = p1
	segundo = p2
	return // retorno limpo ps ja infomamos o valor
}

func main() {

	fmt.Println("Retorno Nomeado")

	// vai trocar os valoares so primeiro parametro para o valor do segundo e o segundo parametro vai ficar com o valor do primeiro
	x, y := trocar(2, 3)
	fmt.Println("Valor P1 :", x, "|| Valor P2 :", y)
	fmt.Println()

	//usando o mesmo nome de variavel da função
	segundo, primeiro := trocar(7, 1)
	fmt.Println("Valor P1 :", primeiro, "|| Valor P2 :", segundo)
	fmt.Println()

}
