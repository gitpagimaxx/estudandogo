package main

import "fmt"

func main() {
	// ARITMÉTICOS
	// Soma
	var soma int = 10 + 10
	println(soma)

	// Subtração
	var subtracao int = 10 - 5
	println(subtracao)

	// Multiplicação
	var multiplicacao int = 10 * 5
	println(multiplicacao)

	// Divisão
	var divisao int = 10 / 2
	println(divisao)

	// Módulo
	var modulo int = 15 % 4
	println("Modulo", modulo)

	// COMPARAÇÃO
	// Igual
	var igual bool = 10 == 10
	fmt.Println("Igual:", igual)

	// Diferente
	var diferente bool = 10 != 10
	println(diferente)

	// Maior
	var maior bool = 10 > 5
	println(maior)

	// Menor
	var menor bool = 10 < 5
	println(menor)

	// como eu faaço uma concatenação de strings?
	var texto string = "Texto" + " " + "Texto 2"
	fmt.Println(texto)

}
