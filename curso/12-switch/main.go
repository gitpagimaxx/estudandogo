package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		//fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	// switch
	// switch é uma estrutura de controle que avalia uma expressão
	// e executa o case que é verdadeiro
	// não é necessário usar break
	// é possível usar mais de um case para uma expressão
	// é possível usar expressões lógicas nos cases
	// o default é executado caso nenhum case seja verdadeiro
	// o default não é obrigatório
	// é possível declarar a variável de comparação dentro do switch
	// a variável declarada no switch só é válida no escopo do switch
	// a variável declarada no switch não pode ser redeclarada no case
	// o switch avalia os cases de cima para baixo
	// o switch interrompe a execução dos cases quando encontra um case verdadeiro
	// o switch não é uma expressão de comparação, como em outras linguagens
	// o switch é uma expressão de controle de fluxo

	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana2(2))
}
