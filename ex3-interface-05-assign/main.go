package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return int(*b), nil
}

func main() {

	// Both os.Stdout and bytes.Buffer are types that satisfy the Writer interface
	// which means that they implement the Write method (or all methods) that Writer (or an interface) requires
	var w io.Writer

	w = os.Stdout                 // contains a writer method that satisfies the Writer interface
	io.WriteString(w, "hello \n") // writes to the standard output

	w = new(bytes.Buffer)              // ALSO contains a writer that satisfies the Writer interface
	io.WriteString(w, "hello again\n") // writes to the the buffer
	fmt.Println(w)

	b := ByteCounter(10)        //ALSO contains a writer that satisfies ... see above
	io.WriteString(&b, "hello") //need to pass as a pointer as the method is passed by pointer
	fmt.Println(b)              //b is printed as the length of the string, as calculated by ByteCounter's Write method

	// t := time.Second // this DOES NOT work as t does not satisfy ... see above
	// fmt.Println(t)   //this would not work
}
