package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}

	return resultado
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// funções variáticas
	// funções que aceitam um número variável de argumentos
	// é possível passar zero ou mais argumentos
	// os argumentos são acessíveis dentro da função como um slice
	totalSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)
	soma(10, 20, 30)
	soma(100, 200)
	soma(0)
	soma()

	escrever("Olá mundo", 1, 2, 3, 4, 5)

}
