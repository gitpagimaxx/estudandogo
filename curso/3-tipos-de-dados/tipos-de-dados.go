package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 10
	fmt.Println(numero)

	var numeroNegativo int16 = -10
	fmt.Println(numeroNegativo)

	var numero2 uint32 = 100
	fmt.Println(numero2)

	// ALIAS
	var numero3 rune = 1000
	fmt.Println(numero3)

	var numero4 byte = 100
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1230001.45
	fmt.Println(numeroReal3)

	// STRING
	var texto string = "Texto"
	fmt.Println(texto)

	texto2 := "Texto 2"
	fmt.Println(texto2)

	char := 'B'
	fmt.Println(char)

	// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := false
	fmt.Println(booleano2)

	// ERRO
	var erro error = fmt.Errorf("Erro")
	fmt.Println(erro)

	var erro2 error
	fmt.Println(erro2)

	var error3 error = errors.New("Erro interno")
	fmt.Println(error3)

	// PONTEIRO
	var ponteiro *int = new(int)
	*ponteiro = 10
}
