package gopointers

import "fmt"

func PointerImpl() {

	count := 10

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	// Pass the "address of" the variable count.
	increment(&count)

	fmt.Println("After: ", count, &count)
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
func increment(inc *int) {

	// Increment the value that the "pointer points to". (de-referencing)
	*inc++
	fmt.Println("Inc:   ", *inc, &inc, inc)

}
