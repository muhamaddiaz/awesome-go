package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

func main() {
	diaz := person{
		firstname: "Muhamad",
		lastname: "Diaz",
	}

	fmt.Printf("%v %v", diaz.firstname, diaz.lastname)
}
