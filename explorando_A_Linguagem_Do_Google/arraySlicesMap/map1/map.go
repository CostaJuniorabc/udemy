package main

import (
	"fmt"
)

func main() {

	// var aprovados map[int] string (não funciona)
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[34512312345] = "Maria"
	aprovados[23456789098] = "Pedro"
	aprovados[83747485763] = "Ana"
	aprovados[00000000001] = "Junior"
	fmt.Println(aprovados)
	fmt.Println()

	fmt.Println("Usando um for no map")
	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF:%v)\n", name, cpf)
	}

	fmt.Println() // nesse caso foi deletado o junior
	delete(aprovados, 00000000001)
	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF:%v)\n", name, cpf)
	}
}
