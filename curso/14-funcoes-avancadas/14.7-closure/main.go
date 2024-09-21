package main

import "fmt"

func closure() func() {

	x := 10

	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {

	fmt.Println("Closure")

	funcaoNova := closure()
	funcaoNova()
}
