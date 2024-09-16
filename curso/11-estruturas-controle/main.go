package main

import "fmt"

func main() {
	// estruturas de controle
	// if, else, else if

	// if
	fmt.Println("Estruturas de controle")

	if true {
		fmt.Println("Verdadeiro")
	}

	if !false {
		fmt.Println("Falso")
	}

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else {
		fmt.Println("Menor que zero")
	}

}
