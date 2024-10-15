package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	fmt.Println("Testes Automatizados Introdução")

	tipoEndereco := enderecos.TipoDeEndereco("Rua ABC")

	fmt.Println(tipoEndereco)
}
