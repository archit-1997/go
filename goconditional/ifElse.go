package goconditional

//`if`-`else` in Go

import "fmt"

func IfElse() {

	if 10%2 == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}

	//if statement without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//statement preceding conditionals
	//variables declared in the statement available
	//in all the branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
		//parenthesis not requried around conditionals
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
