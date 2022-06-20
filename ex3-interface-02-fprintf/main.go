package main

import (
	"bytes"
	"fmt"
	"os"
)

type BytesCounter int

func (b *BytesCounter) Write(p []byte) (n int, err error) {
	*b = BytesCounter(len(p))
	return int(*b), nil
}

func main() {

	var buf bytes.Buffer
	var bc BytesCounter

	// Fprintf accepts its 1st parameter with objects that has a write method signature that takes in a []bytes slice, and returns and int and error
	// All 3 types here (including custom type BytesCounter) do have a method signature that accepts a []byte slice, and returns an int and error
	// stdOut
	// bytes.Buffer
	// BytesCounter (custom type)

	fmt.Fprintf(os.Stdout, "hello standard output\n") //writes to the standard output
	fmt.Fprintf(&buf, "12345")                        //writes to the buffer as bytes
	fmt.Fprintf(&bc, "123456789")                     //writes to bytecounter which implements a concrete type of Write

	fmt.Println(buf)          //bytecode of 12345 is printed
	fmt.Println(buf.String()) //12345 is printed
	fmt.Println(bc)           //length of bc is printed: 9

}
