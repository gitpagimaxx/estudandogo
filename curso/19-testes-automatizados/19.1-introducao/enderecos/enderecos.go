package enderecos

import (
	"strings"
	"testing"
)

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {

	if endereco == "" {
		return "Nenhum Tipo encontrado"
	}

	tiposValidos := []string{"Avenida", "Rua", "Estrada", "Rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if strings.ToLower(tipo) == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}

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
