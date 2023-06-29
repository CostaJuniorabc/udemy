package main

import (
	"fmt"
)

func main() {

	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println("Valor com Make", s1, "||", "Valor com Append", s2)

	s1[3] = 7
	fmt.Println("Valor com Make", s1, "||", "Valor com Append", s2)
}
