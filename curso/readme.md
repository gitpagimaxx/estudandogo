# Anotações Importantes

`go run main.go` rodar o projeto no bash.

`go mod tidy` limpa todo os pacotes não utilizados.

No GO não pode criar uma variável e não usar, isso gerará um erro de compilação, tornando o código mais limpo se nenhuma variável sem ser usado.

## Tipos de dados

### Int
Tem `int` `int8` `int16` `int32` `int64` que podem suportar todos os tipos de números dentro da quantidade de bytes, para não ter que se preocupar com o valor que será inferido é só colocar `int`.

#### Alias 
`int32` = `rune`

`uint8` = `byte`

### Uint

Existe o `uint` que recebe números inteiros apenas positivos

---

### Números reais - float32 e float64

`float32` `float64`, deve usar o `.` para os valores reais, importante saber que não é possível usar o `float` apenas, irá ocorrer erro de compilação então deve ser inferida com 32 ou 64 dependendo da quantidade a ser populada na variável.

`var numeroReal float32 = 123.45`

`var numeroReal float54 = 1230000.45`

`numeroReal := 123.45`

---

### String

Usa-se aspas duplas e aspas simples são para valor char igual ao C#

Quando usar o `char` ao inferir um valor ele vai substituir pelo ASC equivalente tipo B == 66

## Funções

`func calculosMatematicos(numero1 int8, numero2 int8) (int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	return soma, subtracao
}`

Uma função pode retornar mais de um resultado, bastando tipar o retorna como na função acima

E para ignorar um dos resultado você precisa usar o `_` para isso

E para usar basta chamar como no exemplo abaixo:

`resultado_1, resultado_2 := calculosMatematicos(param1, param2)`

## Operadores

Só existe && OR (com pipe) e ! (negar) nem diferente existe

Não existe ternario, tem que usar o if

## Struct

## Herança

Não existe herança por assim dizer, é algo que se aproxima de uma linguagem orientada a objetos

Para não ter que acessar algo como estudante.pessoa.nome ao criar a struct com herança você só adiciona o nome da struct de herança assim você irá acessar a informação com estudante.nome

A struct pode ser criada e não tem problema como varáveis que se criadas devem ser usadas as structs não são assim.

## Ponteiros

Ponteiros são refências de memória, 

## Array e Slices

O slice é sempre criado a partir de um array

## Structs

O GO não tem classe o mais chega perto são as structs

## Defer

Uma coisa interessante em funcções `defer` é que ela executa logo antes de qualquer return de uma função ou seja, como executar algo depois do processamento sem a necessidade de um if por exemplo para dar a resposta. No caso pode ser executada uma função ou um simple fmt.Println, é usado para fechar a conexão com o banco de dados.

## Panic e Recover

O Panic para a execução do programa, e para continuar a vida do programa usa-se o recovery com um if de criação.

`
if r := recover(); r != nil {
	fmt.Println("Execução recuperada com sucesso!")
}
`

## Closure

É uma função que tem uma lembrança de onde ela veio, e irá usar tudo o que tem dentro da função de chamada e não de onde ela foi chamada.

## Funções com ponteiros

E comumente usado quando você precisa fazer uma alteração "global" de uma variável, pois ela vai mudar o valor daquela variável em memória e não uma cópia dela.

## Função init

É uma função que é executada antes da função main, sempre que rodar o sistema a função init, se existir é chamada antes.

## Métodos

## Interface 

É o mesmo modo que usamos em C# mas de um jeito diferente

## Testes

`go test ./... -cover` para verificar todos os testes cobertos

`go test ./... -v` verbose para mmostrar o que está sendo executado

`go test ./... -coverprofile resultado.txt` para gerar um relatório da cobertura

`go tool cover --func=resultado.txt` vai ler o relatório e detalhar 

`go tool cover --html=resultado.txt` exibe um html das linhas que não foram cobertas ficam em vermelho 

## MySQL

go get github.com/go-sql-driver/mysql

