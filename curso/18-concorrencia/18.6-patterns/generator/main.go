package main

import (
	"fmt"
	"time"
)

func main() {

	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(msg string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", msg)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
