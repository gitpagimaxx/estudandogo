// TESTE DE UNIDADE
package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	endereco     string
	tipoEsperado string
}

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TestTipoDeEndereco(t *testing.T) {
	t.Run("Testando um endereço do tipo Rua", func(t *testing.T) {
		endereco := "Rua das Rosas"
		tipoEsperado := "Rua"
		tipoRecebido := TipoDeEndereco(endereco)

		if tipoRecebido != tipoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEsperado, tipoRecebido)
		}
	})

	t.Run("Testando um endereço do tipo Avenida", func(t *testing.T) {
		endereco := "Avenida dos Cravos"
		tipoEsperado := "Avenida"
		tipoRecebido := TipoDeEndereco(endereco)

		if tipoRecebido != tipoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEsperado, tipoRecebido)
		}
	})

	t.Run("Testando um endereço sem tipo", func(t *testing.T) {
		endereco := "Praça das Tulipas"
		tipoEsperado := "Tipo Inválido"
		tipoRecebido := TipoDeEndereco(endereco)

		if tipoRecebido != tipoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEsperado, tipoRecebido)
		}
	})
}

func TestTipoEnderecoComArray(t *testing.T) {
	t.Run("Testando um endereço com tipo válido", func(t *testing.T) {
		enderecos := []string{"Rua ABC", "Avenida ABC", "Estrada ABC", "Rodovia ABC"}
		tiposEsperados := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

		for indice, endereco := range enderecos {
			tipoEsperado := tiposEsperados[indice]
			tipoRecebido := TipoDeEndereco(endereco)

			if tipoRecebido != tipoEsperado {
				t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEsperado, tipoRecebido)
			}
		}
	})

	t.Run("Testando um endereço com tipo inválido", func(t *testing.T) {
		enderecos := []string{"Praça ABC", "Viela ABC", "Estradinha ABC", "R"}
		tipoEsperado := "Tipo Inválido"

		for _, endereco := range enderecos {
			tipoRecebido := TipoDeEndereco(endereco)

			if tipoRecebido != tipoEsperado {
				t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEsperado, tipoRecebido)
			}
		}
	})
}

func TestIgualAoCurso(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Estrada ABC", "Estrada"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça ABC", "Tipo Inválido"},
		{"Viela ABC", "Tipo Inválido"},
		{"Estradinha ABC", "Tipo Inválido"},
		{"R", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoRecebido := TipoDeEndereco(cenario.endereco)

		if tipoRecebido != cenario.tipoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.tipoEsperado, tipoRecebido)
		}
	}
}
