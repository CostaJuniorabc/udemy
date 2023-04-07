package main

import "fmt"

/*
fmt. (serve para apresentar varias funçoes de saidas formatadas)
fmt.Sprint(x) (serve para transformar um valor em string)
Print (ele escreve tudo na mesma linha)
Println (ele quebra a linha no caso cria uma nova linha para baixo)
Printf (e o mesmo que print formatado ex: %f imprime um float, %s imprime uma string, %d imprime um inteiro, %t mostra um valor tipo boolean, %v serve para varios tipos de dados \n serve para pular vinha)

*/

func main() {

	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
