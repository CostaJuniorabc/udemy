package main

import (
	"fmt"
)

func main() {

	array1 := [3]int{1, 2, 3}
	fmt.Println("Valor ARRAY1", array1)
	fmt.Println()

	var slice1 []int
	fmt.Println("Valor SLICE1", slice1)
	fmt.Println()

	// array1 = append(array1, 4, 5, 6) undefined: array1compilerUndeclaredName ( não funciona pq o primeiro argumento sempre tem que ser um slice)

	fmt.Println("Com append")
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println("Valor ARRAY1:", array1, "||", "Valor SLICE1:", slice1)
	fmt.Println()

	fmt.Println("Com make e copy") // copy não aumenta não expande o tamanho para o outro
	slice2 := make([]int, 2)
	fmt.Println("Valor SLICE2:", slice2)

	copy(slice2, slice1) // nao podemos passsar um array como fonte de dados nem como origem
	fmt.Println("Com copy, Valor SLICE2 :", slice2)

}
