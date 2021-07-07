package deferCalls

import (
	"fmt"
)

func DeferPointers() {
	i := 5

	i++
	defer printPointerIncrements(&i)
	i++
	defer printPointerIncrements(&i)
	i++
	defer printPointerIncrements(&i)
}

func printPointerIncrements(i *int) {
	fmt.Println(*i)
}
