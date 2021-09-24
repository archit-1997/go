package goloops

import "fmt"

func ForLoop() {

	// classic for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// most basic for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop without condition
	for {
		fmt.Println("loop")
		break
	}
}
