package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	fmt.Println()
	fmt.Println("Seção 4.29 - If com Init")
	fmt.Println()
	fmt.Println("Link = https://www.udemy.com/course/curso-go/learn/lecture/8603234#overview")
	fmt.Println()

	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou !!! -", i)
	} else {
		fmt.Println("Perdeu !!! - ", i)
	}

	fmt.Println()
	fmt.Println("__________FIM . . .__________ ")
	fmt.Println()
}
