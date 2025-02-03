package main

import "fmt"

// são colecoes de campos
type Pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(Pessoa{
		nome:  "Rafael",
		idade: 10,
	})

	fmt.Println(Pessoa{
		nome:  "Rafael 2 ",
		idade: 20,
	})

	fmt.Println(Pessoa{
		nome:  "Rafael 3",
		idade: 30,
	})
}
