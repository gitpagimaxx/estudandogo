package enderecos

import (
	"strings"
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
