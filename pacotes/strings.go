package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra := "terra"

	if strings.Contains(palavra, "rr") {
		fmt.Println("palavra contains rr")
	}
}
