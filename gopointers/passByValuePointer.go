package gopointers

import "fmt"

func PassPointerByValue() {

	x := 1
	addx := &x

	//value stored in the pointer
	fmt.Println("Values before passing")
	fmt.Printf("%p\n", addx)
	fmt.Printf("%p\n", &addx)

	//showing pass by value in pointers
	displayPassByValue(addx)
}

func displayPassByValue(addx *int) {

	//now the value and the address of the pointer
	fmt.Println("Values after passing")
	fmt.Printf("%p\n", addx)
	fmt.Printf("%p\n", &addx)
}
