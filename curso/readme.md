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

## Números reais

### Float32 e Float64

`float32` `float64`, deve usar o `.` para os valores reais 

`var numeroReal float32 = 123.45`

