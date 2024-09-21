package main

import "fmt"

type User struct {
	nome  string
	idade uint8
}

func (u User) salvar() {
	fmt.Printf("Salvando os dados de %s no banco de dados\n", u.nome)
}

func (u User) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *User) fazerAniversario() {
	u.idade++
}

func main() {

	fmt.Println("Métodos")

	u := User{nome: "Fulano", idade: 20}
	fmt.Println(u)
	u.salvar()

	u2 := User{"Ciclano", 30}
	fmt.Println(u2)
	u2.salvar()

	maiorDeIdade := u2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	u2.fazerAniversario()
	fmt.Println(u2.idade)
}
