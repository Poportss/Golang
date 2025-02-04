package main

import "fmt"

//capacidade de uma funcao de chamar um variavel que está em outra funcão

func main() {
	x := 0

	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())

}
