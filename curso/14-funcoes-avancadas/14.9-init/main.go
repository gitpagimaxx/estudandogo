package main

import "fmt"

var variavel int

func init() {
	fmt.Println("Função init foi chamada")
	variavel = 10
}

func main() {

	fmt.Println("Função main foi chamada")
	fmt.Println(variavel)

}
