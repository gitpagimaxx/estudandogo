package main

import (
	"bytes"
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
	pessoa := Pessoa{"João", 30, "Programador"}

	pessoaJson, erro := json.Marshal(pessoa)

	if erro != nil {
		fmt.Println("[main] Houve um erro na geração do JSON", erro)
		return
	}

	fmt.Println("A pessoa em JSON: ", string(pessoaJson))

	fmt.Println(bytes.NewBuffer(pessoaJson))

	// Convertendo JSON para struct
	pessoa2 := map[string]string{
		"nome":      "Maria",
		"idade":     "25",
		"profissao": "Engenheira",
	}

	pessoaJson2, erro := json.Marshal(pessoa2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("A pessoa2 em JSON: ", string(pessoaJson2))
	fmt.Println(bytes.NewBuffer(pessoaJson2))

}
