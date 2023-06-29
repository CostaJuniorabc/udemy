package main

import "fmt"

func notaParaConceito(n float64) string {

	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println()
	fmt.Println("Seção 4.28 - If/Else If")
	fmt.Println()
	fmt.Println("Link = https://www.udemy.com/course/curso-go/learn/lecture/8603232#overview")
	fmt.Println()

	fmt.Println("Nota final João :", notaParaConceito(9.8))
	fmt.Println("Nota final Maria :", notaParaConceito(8.9))
	fmt.Println("Nota final Luiz :", notaParaConceito(7.9))
	fmt.Println("Nota final Luiza :", notaParaConceito(4.9))
	fmt.Println("Nota final Diogo :", notaParaConceito(2.9))

	fmt.Println()
	fmt.Println("__________FIM . . .__________ ")
	fmt.Println()
}
