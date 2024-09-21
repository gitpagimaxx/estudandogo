package main

func calculoMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	// retornos nomeados
	// é possível nomear os retornos de uma função
	// isso ajuda a documentar o código
	// é possível retornar os valores sem usar a palavra return
	// isso é chamado de retorno limpo
	// é possível retornar os valores na ordem que foram nomeados
	// é possível retornar os valores sem usar as variáveis de retorno

	soma, subtracao := calculoMatematico(10, 20)
	println(soma, subtracao)
}