package gostructures

import "fmt"

func Nested() {

	type Address struct {
		houseNo     int
		street      string
		city        string
		isMetroCity bool
	}

	type Employee struct {
		id                  int
		firstName, lastName string
		address             Address
		age                 int
		isFullTime          bool
	}

	archit := Employee{1234, "Archit", "Singh", Address{343, "Ram Gali Road", "Prayagraj", false}, 23, true}

	fmt.Println("Hi !! I'm a nested structure")
	fmt.Println(archit)
}
