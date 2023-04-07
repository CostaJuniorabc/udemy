package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	//	fmt.Println("Resultado de X / Y :", x / y) // não funciona assim
	fmt.Println("Resultado de X / Y :", x/float64(y)) // tem que ser assim informar o tipo da variavel para poder fazer o calculo
	println()

	// conversão de float para inteiro
	nota := 6.9
	notaFinal := int(nota) // converte a variavel nota para um tipo inteiro
	fmt.Println("Qual a nota arredondada :", notaFinal)

	//cuidado para concatenar uma string com string, nesse caso concatenamos umas string com outra stringo porem o valor da segunda estring acaba se referindo ao valor do codigo da tabela ASCII - tabela = https://www.ime.usp.br/~kellyrb/mac2166_2015/tabela_ascii.html
	fmt.Println("String 1 + " + string(rune(97)))

	// int para string
	fmt.Println("Teste convert" + strconv.Itoa(123))

	// string para int nessa caso ele retorna 2 valores um e o valor e o outro e erro para ignorar um dos metodos so colocar um _
	num, _ := strconv.Atoi("1234")
	fmt.Println("Resultado : ", num-1230)

	// string para boolean esse metodo tbm retorna 2 valores
	boo, _ := strconv.ParseBool("1")
	if boo == true {
		fmt.Println("Verdadeiro")

	} else {
		fmt.Println("Falso")
	}

}
