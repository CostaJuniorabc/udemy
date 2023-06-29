package main

import (
	"fmt"
)

func main() {

	s := make([]int, 10) // slice criado atraves do make
	s[9] = 12
	fmt.Println("Slice :", s)

	s = make([]int, 10, 20)                                                             // slice com 2 medidas
	fmt.Println("Slice:", s, "||", "Tamanho:", len(s), "||", "Capacidade Max:", cap(s)) // len(pega o tamanho) || cap(pega a capacidade maxima)

	//função apend
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Slice:", s, "||", "Tamanho:", len(s), "||", "Capacidade Max:", cap(s)) // len(pega o tamanho) || cap(pega a capacidade maxima)

	// qundo ultrapassa o tamanho maximo, ele aumenta o tamanho altomaticamente
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println("Slice:", s, "||", "Tamanho:", len(s), "||", "Capacidade Max:", cap(s)) // len(pega o tamanho) || cap(pega a capacidade maxima)

}
