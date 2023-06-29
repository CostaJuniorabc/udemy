package main

import "fmt"

func main() {

	i := 1

	//  não tem aritimetica de ponteiros em GO
	var p *int = nil

	p = &i // & serve para pegar o endereço armazenado no ponteiro || nesse caso pegando o endereço da variavel
	fmt.Println("Endereço da variável p :", p)

	// p++ da erro || não pode ser feito assim

	*p++ // desreferenciando o ponteiro (pegando o valor da variável)
	fmt.Println("Valor da Variavel *p++ = ", *p)

	i++
	fmt.Println("Valor da variável i =", i)
}
