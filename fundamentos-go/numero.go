package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("Title: %s, Author: %s", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func main() {
	//tipos
	nome := "rafael"
	idade := 22
	versao := 2.2
	fmt.Println("1 + 1 = ", nome, idade, versao)
	fmt.Println("1.1 + 1.1 = ", 1.1+1.1)

	//arrays
	var x = [10]int{}
	x[4] = 80
	fmt.Println(x)

	//interface
	a := Article{Title: "Hello World", Author: "John Doe"}
	fmt.Println(a)

	expressiom := (true && false) || (false && true) || !(false && false)
	fmt.Println(expressiom)

}

func Print(s Stringer) {
	fmt.Println(s.String())
}
