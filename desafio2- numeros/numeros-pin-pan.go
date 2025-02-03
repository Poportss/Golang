package main

import "fmt"

func main() {

	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Println("Pin")
		} else if x%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(x)
		}
	}

}
