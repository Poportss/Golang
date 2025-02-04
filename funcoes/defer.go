package main

import "fmt"

func dia1() {
	fmt.Println("dia1")
}

func dia2() {
	fmt.Println("dia2")
}

// faz na ordem inversa sempre
func main() {
	defer dia1()
	defer dia2()
}
