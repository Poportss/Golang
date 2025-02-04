package main

import "fmt"

//funcao ou subrotina  ou procendimento

func media(list []float64) float64 {

	total := 0.0

	for _, valor := range list {
		total += valor
	}

	return total / float64(len(list))

}

func main() {
	list := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println(media(list))
}
