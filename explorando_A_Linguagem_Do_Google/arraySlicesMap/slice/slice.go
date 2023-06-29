package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("---------- Diferençã entre Array e Slice ----------")
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println("Array", a1, "||", "Slice :", s1)
	fmt.Println("Tipo de Array :", reflect.TypeOf(a1), "||", "Tipo de Slice :", reflect.TypeOf(s1))

	//array base para trabalhar com slice
	a2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println()
	fmt.Println("---------- Slice s2 ----------")
	// Slice não é um array ! Slice define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println("Arrai s2 :", a2, "||", "Slace s2 :", s2)

	fmt.Println()
	fmt.Println("---------- Slice s3 ----------")
	//novo slice, mas aponta para o mesmo array
	s3 := a2[:2]
	fmt.Println("Array s2 :", a2, "||", "Slice s3 :", s3)

	fmt.Println()
	fmt.Println("---------- Slice s4 | slice de slice----------")
	// vc pode imaginar um slice como : uma strutura que tem tamanho e ele tambem tem um ponteiro para um elemento de array
	s4 := s2[:1]
	fmt.Println("Slice s2 :", s2, "||", "Slice s4 :", s4)

}
