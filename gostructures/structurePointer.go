package gostructures

import "fmt"

func PointerToStruct() {

	//rohan is pointer of type Employee
	rohan := &Employee{
		firstName: "Rohan",
		lastName:  "Chhabra",
		id:        456,
		age:       23,
		gender:    'M',
	}

	fmt.Println("Hi, I explain about pointer to gostructures")

	//to print the values, we need to dereference
	fmt.Println((*rohan).firstName)

	//we can also use rohan directly without dereferencing, Go handles the rest
	fmt.Println(rohan.firstName)
}
