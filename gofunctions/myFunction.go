package gofunctions

import "fmt"

func mult(x int, y int) int {
	return x * y
}

func addmany(a, b, c int) int {
	return a + b + c
}

func GoFunction() {
	fmt.Println("Multiplication fn:", mult(42, 13))
	fmt.Println("Add many fn:", addmany(42, 13, 10))
}
