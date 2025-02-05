package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)

	fmt.Println(x, y)
}

func soma(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

func multiplica(numeros ...int) int {
	total := 1
	for _, num := range numeros {
		total = total * num
	}
	return total
}
