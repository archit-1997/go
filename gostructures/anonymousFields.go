package gostructures

import "fmt"

func Anonymous() {
	type Data struct {
		string
		int
		bool
		byte
	}

	anonymousData := Data{
		"integer",
		4,
		true,
		32,
	}

	fmt.Println("I'm a structure with anonymous fields")
	fmt.Println(anonymousData)
}
