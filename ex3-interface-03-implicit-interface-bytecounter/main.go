package main

import (
	"fmt"
	"io"
)

// ByteCounter type (integer)
type ByteCounter int

// Implement a Write method for ByteCounter
// to count the number of bytes written
// In doing so, we are providing a custom writer for the io interface used below
func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return int(*b), nil
}

func main() {

	// printf accepts the custom implementation of the concrete type that has the method signature of Write method from the Writer Inteface
	// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver

	var b ByteCounter
	fmt.Fprintf(&b, "Hello World!!")
	fmt.Println(b)

	// another way of instantiating a ByteCounter variable is by giving it an initialized value
	b2 := ByteCounter(10)
	io.WriteString(&b2, "Hello World!!")
	fmt.Println(b2)
}
