package main

import "fmt"

//capacidade de uma funcao chamar ela mesmo

func fatorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * fatorial(x-1)
	}
}

func main() {
	fmt.Println(fatorial(5))
}
