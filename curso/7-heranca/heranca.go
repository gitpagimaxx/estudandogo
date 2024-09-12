package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type student struct {
	person
	grade int
}

func main() {
	fmt.Println("Herança")
	fmt.Println("-----------------")

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
	}

	fmt.Println(p1.firstName, p1.lastName)

	fmt.Println("-----------------")

	student1 := student{
		person: p1,
		grade:  10,
	}

	fmt.Println(student1.firstName, student1.lastName, student1.grade)

	fmt.Println("-----------------")

	var p2 person
	p2.firstName = "Miss"
	p2.lastName = "Moneypenny"

	fmt.Println(p2.firstName, p2.lastName)

	fmt.Println("-----------------")

	p3 := person{firstName: "Dr."}
	fmt.Println(p3.firstName)
}
