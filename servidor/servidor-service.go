package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v : %v /n", c, nome)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalhos", cabecalhos)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}
