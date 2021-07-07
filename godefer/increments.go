package godefer

import (
	"fmt"
)

func DeferIncrements() {
	i := 5

	i++
	defer printIncrements(i)
	i++
	defer printIncrements(i)
	i++
	defer printIncrements(i)
}

func printIncrements(i int) {
	fmt.Println(i)
}
