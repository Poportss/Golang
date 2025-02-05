package main

import "fmt"

func main() {
	a := soma(1, 2, 3)
	b := multiplica(10, 10)
	c := divide(100, 2, 5)
	d := subtrai(50, 10, 5)

	fmt.Println(a, b, c, d)
}

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func multiplica(numeros ...int) int {
	total := 1
	for _, num := range numeros {
		total *= num
	}
	return total
}

func divide(numeros ...int) int {
	total := numeros[0]
	for _, num := range numeros[1:] {
		total /= num
	}
	return total
}

func subtrai(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}

	total := numeros[0]
	for _, num := range numeros[1:] {
		total -= num
	}
	return total
}
