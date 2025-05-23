package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Home Page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	HandleRequest()
}
