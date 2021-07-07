package gopointers

import "fmt"

func PassReference() {
	a := 5
	changeA(a)
	fmt.Println("After pass as value:", a)

	changeAp(&a)
	fmt.Println("After passing pointer:", a)
}

func changeA(am int) {
	am = 10
}

func changeAp(am *int) {
	*am = 10
}
