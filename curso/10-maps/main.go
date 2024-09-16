package main

import (
	"fmt"
)

func main() {
	// mapas
	// mapas são como dicionários em Python
	// são estruturas de dados que armazenam pares chave-valor
	// a chave é única e o valor pode ser qualquer tipo de dado
	// mapas são uma referência a uma estrutura de dados
	// não precisam ser inicializados
	// nil é o valor zero de um mapa
	// não é possível acessar uma chave que não existe
	// para verificar se uma chave existe, use a vírgula ok
	// delete(mapa, chave) remove um par chave-valor de um mapa
	// mapas não são seguros para concorrência
	// para criar um mapa seguro para concorrência, use sync.Map

	fmt.Println("Mapas")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"último":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
