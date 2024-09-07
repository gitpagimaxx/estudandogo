package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá, mundo!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("diego@gmail.com")
	fmt.Println(erro)
}
