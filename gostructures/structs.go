package gostructures

import "fmt"

type Employee struct {
	firstName, lastName string
	id                  int
	age                 int
	gender              byte
}

func BuildStruct() {

	archit := Employee{
		firstName: "Archit",
		lastName:  "Singh",
		id:        7,
		age:       23,
		gender:    'M',
	}

	fmt.Println("I'm a structure")
	fmt.Println(archit)
}
