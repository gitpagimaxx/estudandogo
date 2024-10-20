package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome      string `json:"nome"`
	Idade     int    `json:"idade"`
	Profissao string `json:"profissao"`
}

func main() {
	pessoaJson := `{"idade":25,"nome":"Maria","profissao":"Engenheira"}`

	var pessoa Pessoa

	if erro := json.Unmarshal([]byte(pessoaJson), &pessoa); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(pessoa)

	pessoa2Json := `{"nome":"João","idade":"30","profissao":"Programador"}`

	pessoa2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(pessoa2Json), &pessoa2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(pessoa2)
}
