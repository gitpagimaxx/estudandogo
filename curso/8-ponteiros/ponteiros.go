package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	fmt.Println("-----------------")

	a := 10
	b := a

	fmt.Println(a, b)

	fmt.Println("-----------------")

	c := 10
	d := &c

	fmt.Println(c, d)
	fmt.Println(c, *d)

	fmt.Println("-----------------")

	e := 10
	f := &e // referenciando, exibe o endereço de memória
	*f = 20

	fmt.Println(e, f)
	fmt.Println(e, *f) // desreferenciando
}