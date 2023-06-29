package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5, 6} // compilador conta !

	fmt.Println("----------pegando o indice e o numero----------")
	for i, numero := range numeros {
		//fmt.Println(i+1, ")", numero)
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	fmt.Println("----------Ignorando o indice e pegando o numero----------")
	for _, num := range numeros {
		fmt.Println(num)
	}

	fmt.Println("----------Ignorando o numero e pegando o indice----------")
	for i, _ := range numeros {
		fmt.Println(i)
	}

}
