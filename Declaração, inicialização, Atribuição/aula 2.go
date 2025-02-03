package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Println("Ebulição ahgua em F", tempF)
	fmt.Println("Ebulição ahgua em C", tempC)
}
