package godefer

import (
	"fmt"
)

var p = fmt.Println

func OrderOfCalls() {
	i := 5
	fmt.Println(i)

	defer p(1)
	if true {
		defer p(2)
		p(3)
	}
	defer p(4)
}

func g() {
}
