package main

import (
	"fmt"
	"github.com/archit-1997/go/deferCalls"
	"github.com/archit-1997/go/structures"
)

func main() {
	fmt.Println("I'm the main file")

	//structures in go
	structures.BuildStruct()
	structures.BuildAnonymousStruct()
	structures.PointerToStruct()
	structures.Anonymous()
	structures.Nested()

	//use of defer

	//creating, closing and writing to a file
	f := deferCalls.CreateFile("defer.txt")
	//the below line will be executed as the last statement of func main()
	defer deferCalls.CloseFile(f)
	deferCalls.WriteFile(f)

	//order of calls using the defer statement
	deferCalls.OrderOfCalls()

	//increment operation with defer calls
	deferCalls.DeferIncrements()

	//incrementing pointers with defer calls
	deferCalls.DeferPointers()
}
