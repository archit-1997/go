package gostructures

import "fmt"

func BuildAnonymousStruct() {
	dog := struct {
		name   string
		legs   int
		gender byte
	}{
		name:   "tuffy",
		legs:   4,
		gender: 'M',
	}

	fmt.Println("I'm an anonymous structure")
	fmt.Println(dog)
	fmt.Println("My type is :")
	fmt.Printf("%T\n", dog)

}
