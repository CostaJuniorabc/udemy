package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não Sei"
	}
}

func main() {

	fmt.Println("Qual o tipo : ", tipo((1.3)))
	fmt.Println("Qual o tipo : ", tipo((1)))
	fmt.Println("Qual o tipo : ", tipo("Ops!!"))
	fmt.Println("Qual o tipo : ", tipo((func() {})))
	fmt.Println("Qual o tipo : ", tipo((time.Now())))

}
