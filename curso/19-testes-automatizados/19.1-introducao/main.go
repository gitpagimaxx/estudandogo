package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	fmt.Println("Testes Automatizados Intrudução")

	tipoEndereco := enderecos.TipoDeEndereco("rua ABC")

	fmt.Println(tipoEndereco)
}
