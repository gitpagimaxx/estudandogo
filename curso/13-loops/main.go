package main

import (
	"fmt"
)

type usuario struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	fmt.Println("Loops")

	// for
	i := 0
	for i < 10 {
		//time.Sleep(time.Second)
		i++
		//fmt.Println(i)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		// time.Sleep(time.Second)
		// fmt.Println("Incrementando J ", j)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario2 := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	for chave, valor := range usuario2 {
		fmt.Println(chave, valor)
	}

	usuarios := []usuario{
		{nome: "Pedro", sobrenome: "Silva", idade: 30},
		{nome: "Maria", sobrenome: "Silva", idade: 25},
	}

	for _, usuario := range usuarios {
		fmt.Println(usuario)
	}
}
