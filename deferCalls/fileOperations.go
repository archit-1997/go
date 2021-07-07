package deferCalls

import (
	"fmt"
	"os"
)

func CreateFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func WriteFile(f *os.File) {
	fmt.Println("writing")
	//writing data into the file
	fmt.Fprintln(f, "data")

}

func CloseFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
