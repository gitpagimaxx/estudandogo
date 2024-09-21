package main

import "fmt"

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	fmt.Println("Recursividade")

	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

	fmt.Println("-----------------")

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
