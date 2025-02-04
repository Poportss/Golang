package main

import (
	"io"
	"log"
	"os"
)

func main() {

	if _, err := io.WriteString(os.Stdout, "some io.Reader stream to be read"); err != nil {
		log.Fatal(err)
	}

}
