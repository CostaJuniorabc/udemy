package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("================== FOR ==================")
	i := 0
	for i <= 200 { // não tem while em go
		fmt.Println(i)
		i++
	}

	fmt.Println("================== FOR 2 ==================")
	for i := 0; i <= 400; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("================== FOR 3 ==================")
	fmt.Println("\nMisturando...")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Par : ", i)
		} else {
			fmt.Println("Impar : ", i)
		}
	}

	fmt.Println("================== FOR 4 ==================")
	for {
		// laço infinito
		fmt.Println("Para Sempre... :(")
		time.Sleep(time.Second * 3)
	}

	// Veremos o foreach no capítulo de array
}
