package gouserinput

import "fmt"

func UserInput() {
	var first, second string
	fmt.Println("Enter your first name")
	_, err1 := fmt.Scanln(&first)
	if err1 != nil {
		fmt.Println("You haven't properly entered your First Name")
		return
	}

	fmt.Println("Enter your second name")
	_, err2 := fmt.Scanln(&second)
	if err2 != nil {
		fmt.Println("You haven't properly entered your Second Name")
		return
	}

	fmt.Println("Your full name is " + first + " " + second)
}
