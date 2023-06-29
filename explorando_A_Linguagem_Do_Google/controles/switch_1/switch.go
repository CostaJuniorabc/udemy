package main

import "fmt"

func notaParaConceito(n float64) string {

	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida"
	}
}

func main() {
	fmt.Println("Nota Daniele :", notaParaConceito(10))
	fmt.Println("Nota Marcelo :", notaParaConceito(9.8))
	fmt.Println("Nota Alaisa :", notaParaConceito(8))
	fmt.Println("Nota Junior :", notaParaConceito(6))
	fmt.Println("Nota Rosana :", notaParaConceito(4))
	fmt.Println("Nota Clara :", notaParaConceito(2))
	fmt.Println("Nota Lucas :", notaParaConceito(-1))
}
