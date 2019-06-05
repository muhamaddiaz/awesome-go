package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	posCode int
}

type person struct {
	firstname string
	lastname string
	contactInfo
}

type persons []person

func main() {
	diaz := person{
		firstname: "Muhamad",
		lastname: "Diaz",
		contactInfo: contactInfo{
			email: "muhamaddiaz@outlook.com",
			posCode: 44552,
		},
	}
	anton := person{
		firstname: "Antonio",
		lastname: "Garcia",
		contactInfo: contactInfo{
			email: "antoniogarcia@outlook.com",
			posCode: 44550,
		},
	}
	anton2 := person{
		firstname: "Antonio",
		lastname: "Garcia",
		contactInfo: contactInfo{
			email: "antoniogarcia@outlook.com",
			posCode: 44550,
		},
	}

	people := persons{
		diaz, anton,
	}

	people = append(people, anton2)

	people[0].changeName("Marcos")

	people.printAll()

	//// Structs with shortcut pointer
	// diaz.changeName("Helloworld")
	//
	//// Structs with no shortcut pointer
	//// & is a operation to convert from value to memory address
	// diazPointer := &diaz
	// diazPointer.changeName("Stepen")
	//
	// diaz.print()
}

// * is a operation to convert from memory to value in that address
func (p *person) changeName(newFirstName string) {
	(*p).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("Firstname: %v \n" +
					"Lastname: %v \n" +
					"Email: %v \n" +
					"Kode Pos: %v \n", p.firstname, p.lastname, p.contactInfo.email, p.contactInfo.posCode)
}

func (p persons) printAll() {
	for _, person := range p {
		person.print()
	}
}