package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{

		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.79,
		},
		"J": {
			"Jose João": 11234.56,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	/*
		for name, salario := range funcsPorLetra {
			fmt.Printf("%s (Salario:%v)\n", name, salario)
		}*/

}
