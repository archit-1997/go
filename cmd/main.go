package main

import (
	"fmt"
	"github.com/archit-1997/go/godefer"
	"github.com/archit-1997/go/goerrors"
	"github.com/archit-1997/go/gofunctions"
	"github.com/archit-1997/go/gostructures"
)

func main() {
	fmt.Println("I'm the main file")

	//gostructures in go
	gostructures.BuildStruct()
	gostructures.BuildAnonymousStruct()
	gostructures.PointerToStruct()
	gostructures.Anonymous()
	gostructures.Nested()

	//use of defer

	//creating, closing and writing to a file
	f := godefer.CreateFile("defer.txt")
	//the below line will be executed as the last statement of func main()
	defer godefer.CloseFile(f)
	godefer.WriteFile(f)

	//order of calls using the defer statement
	godefer.OrderOfCalls()

	//increment operation with defer calls
	godefer.DeferIncrements()

	//incrementing pointers with defer calls
	godefer.DeferPointers()

	//working with errors in go
	goerrors.ErrorImplementation()

	//functions in go
	gofunctions.GoFunction()
	gofunctions.NamedReturn()
	gofunctions.MultipleReturns()

}
