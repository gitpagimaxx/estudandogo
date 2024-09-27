package main

import (
	"fmt"
)

func main() {

	fmt.Println("Canais com buffer")

	canal := make(chan string, 2)
	canal <- "Olá Mundo"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
