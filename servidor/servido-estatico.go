package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("servidor/static"))
	http.Handle("/", fs)
	log.Print("Servidor estatico on : 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
