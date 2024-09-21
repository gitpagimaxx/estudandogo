package main

import "fmt"

func main() {

	func(texto string) {
		println(texto)
	}("Olá, mundo!")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebi: %s %d", texto, 10)
	}("Olá, mundo!!!")

	fmt.Println(retorno)
}
