package main

import (
	"fmt"
	"math/rand"
	"time"
)

func escrever(msg string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", msg)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}

func multiplexador(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func main() {

	fmt.Println("Padrão Multiplexador")

	canal := multiplexador(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
