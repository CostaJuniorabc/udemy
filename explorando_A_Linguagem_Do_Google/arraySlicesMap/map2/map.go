package main

import (
	"fmt"
)

func main() {

	fmt.Println("Lista de Funcionarios")

	//declarando e inicinado o map[type]type
	funcsESalarios := map[string]float64{
		"José João":    11325.45,
		"João José":    1556.78,
		"Pedro Junior": 1200.0,
	}

	fmt.Println(funcsESalarios)
	fmt.Println()

	// add novo Funcionarios
	funcsESalarios["Rafael Filho"] = 1350.0
	fmt.Println(funcsESalarios)
	fmt.Println()

	//remover um Funcionarios
	delete(funcsESalarios, "Inexistente")
	fmt.Println(funcsESalarios)
	fmt.Println()

	//percorrer a Lista
	fmt.Println("Usando o FOR")
	for name, salario := range funcsESalarios {
		fmt.Printf("%s (Salario:%v)\n", name, salario)
	}
	fmt.Println()

	fmt.Println("Usando o FOR1")
	for _, salario := range funcsESalarios {
		fmt.Printf("(Salario:%v)\n", salario)
	}
	fmt.Println()

	fmt.Println("Usando o FOR2")
	for name, _ := range funcsESalarios {
		fmt.Printf("%s \n", name)
	}
}
