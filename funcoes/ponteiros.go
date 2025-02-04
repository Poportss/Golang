package main

import "fmt"

//ponteiro armazena um valor na memória, mas não é um valor proprimamente escrito

func inical(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	inical(&x)
	fmt.Println(x)
}
