package main

import (
	"fmt"
)

func main() {

	x := 1
	y := 2

	fmt.Println("Valor de X =>", x)
	fmt.Println("Valor de Y =>", y)
	fmt.Println()

	// apenas posfix
	x++ // add +1 em x | x += 1 oux + 1
	fmt.Println("Posfix de X =>", x)
	fmt.Println()
	x++ // add +1 em x |x += 1 ou x + 1
	fmt.Println("Posfix de X =>", x)
	fmt.Println()

	// apenas posfix
	y-- // subtrai -1 em y | y -= 1 ou x - 1
	fmt.Println("Posfix de y =>", y)
	fmt.Println()
	y-- // subtrai -1 de y
	fmt.Println("Posfix de Y =>", y)
	fmt.Println()
	x++ // add +1 em x |x += 1 oux + 1
	fmt.Println("Posfix de Y =>", y)
	fmt.Println()
}
