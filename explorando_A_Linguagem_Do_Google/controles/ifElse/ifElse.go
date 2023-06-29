package main

import "fmt"

// função que devolve um parametro nota do tipo float64
func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aluno APROVADO com nota =", nota)
	} else {
		fmt.Println("Aluno REPROVADO com nota =", nota)
	}
}

func main() {
	joao := 7.3
	maria := 5.1

	imprimirResultado(joao)
	imprimirResultado(maria)
}
