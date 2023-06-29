package main

import "fmt"

// como declarar uma função || essa função nao recebe nada e não retorna nada
func f1() {
	fmt.Println("F1")
}

// vai receber 2 parametros e nao vai retornar nada
func f2(p1, p2 string) {
	fmt.Println("F2, parametro1 :", p1)
	fmt.Println("F2, parametro2 :", p2)
}

// não vai receber parametro mas retorna uma string
func f3() string {
	return "F3"
}

// recebe 2 parametros tipo string e retorna uma string ||usamosSprintf devido a ser um retorno
func f4(p1 string, p2 string) string {
	return fmt.Sprintf("F4 : Parametro1 : %s , Parametro2 : %s", p1, p2)
}

// essa função nao recebe nada porem retorna 2 valores tipo string
func f5() (string, string) {
	return "Parametro1", "Parametro2"
}

func main() {
	fmt.Println("Retorno das funçoes")
	fmt.Println()

	fmt.Println("Função F1")
	f1()
	fmt.Println()

	fmt.Println("Função F2")
	f2("Parametro1", "Parametro2")
	fmt.Println()

	fmt.Println("Função F3")
	r3 := f3()
	fmt.Println(r3)
	fmt.Println()

	fmt.Println("Função F4")
	r4 := f4("Retorno 1", "Retorno 2")
	fmt.Println(r4)
	fmt.Println()

	fmt.Println("Função F5")
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
	fmt.Println()

	// iginorando retorno 1
	fmt.Println("Função F5 ignorando um retorno")
	_, i52 := f5()
	fmt.Println("F5:", i52)
	fmt.Println()

	// iginorando retorno 2
	fmt.Println("Função F5 ignorando um retorno")
	i51, _ := f5()
	fmt.Println("F5:", i51)
	fmt.Println()

}
