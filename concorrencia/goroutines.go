package main

import "fmt"

// executar outras funções ao mesmo tempo

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var escreva string

	fmt.Scanln(&escreva)
}
