package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	fmt.Println("Structs")
	fmt.Println("-----------------")

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age)

	fmt.Println("-----------------")

	secretAgent1 := secretAgent{
		person:        p1,
		licenseToKill: true,
	}

	fmt.Println(secretAgent1.firstName, secretAgent1.lastName, secretAgent1.age, secretAgent1.licenseToKill)

	fmt.Println("-----------------")

	var p2 person
	p2.firstName = "Miss"
	p2.lastName = "Moneypenny"
	p2.age = 27

	fmt.Println(p2.firstName, p2.lastName, p2.age)

	fmt.Println("-----------------")

	p3 := person{firstName: "Dr."}
	fmt.Println(p3.firstName)
}
