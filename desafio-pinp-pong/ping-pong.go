package main

import (
	"fmt"
	"time"
)

//canal sincronizar a execuçao de goroutines

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" //usado para enviar e recber mensagem pelo canal
	}
}

func pongar(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" //usado para enviar e recber mensagem pelo canal
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)
	go pingar(c)
	go pongar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
