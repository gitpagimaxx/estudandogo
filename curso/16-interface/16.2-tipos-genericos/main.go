package main

import (
	"fmt"
)

func generico(x interface{}) {
	fmt.Println(x)
}

func main() {

	fmt.Println("Tipos genéricos")

	generico("String")
	generico(1)
	generico(true)

}
