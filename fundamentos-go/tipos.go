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

	//mapas
	//maps := make(map[string]int)
	//
	//maps["chave"] = 10
	//fmt.Println(maps["chave"])

	//mapas
	//maps := make(map[int]int)
	//
	//maps[1] = 20
	//fmt.Println(maps[1])

	elemento := make(map[string]string)

	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "L;itio"
	fmt.Println(elemento)

	slice := make([]int, 3, 9)
	fmt.Println(len(slice), cap(slice))

	arrays := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(arrays[2:5])

}

func Print(s Stringer) {
	fmt.Println(s.String())
}
