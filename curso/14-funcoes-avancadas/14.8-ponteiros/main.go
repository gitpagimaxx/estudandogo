package main

import "fmt"

func invertesinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	fmt.Println("Ponteiros")

	numero := 20
	numeroInvertido := invertesinal(numero)

	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
