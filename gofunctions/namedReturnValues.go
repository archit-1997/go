package gofunctions

import "fmt"

func add(a, b int) int {
	return a + b
}

func namedReturn(a, b int) (sum int) {
	sum = a + b
	return
}

func NamedReturn() {
	s := add(42, 13)
	fmt.Println("add fn:", s)
	fmt.Println("namedReturn fn:", namedReturn(42, 13))
}
