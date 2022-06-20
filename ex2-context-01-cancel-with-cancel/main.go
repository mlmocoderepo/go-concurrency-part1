package main

import (
	"context"
	"fmt"
)

// Objective:
// Create a function variable gen that generates integers in a separate goroutine
// and send the value to a returned channel
// The caller of gen will cancel the context once 5 integers are consumed
// so that the internal goroutine started by gen is not leaked.

func main() {

	// gen accepts a context object and returns an <-chan int
	gen := func(cxt context.Context) <-chan int {

		// create ch, a channel that accepts integer
		ch := make(chan int)
		// create n, an integer that is initialized with a value of 1
		n := 1

		// invoke the anonymous goroutine to run
		go func() {

			// defer closing the channel ch
			defer close(ch)

			// create an infinite loop with a mulitplexer
			// listens for a context.Done() to exit the loop
			// otherwise adds 1 to n
			for {
				select {
				case <-cxt.Done():
					return
				case ch <- n:
					n++
				}
			}
		}()

		// returns n
		return ch

	}

	//returns the context and context's cancel channel
	//using cancel sends a context.Done() to stop a process
	cxt, cancel := context.WithCancel(context.Background())

	// call gen and pass the context to the function
	// and returned values from gen is stored in ch
	ch := gen(cxt)

	// use a for loop to receive from channel ch
	// invoke the cancel function trigger the request to send the context.Done()
	for c := range ch {
		fmt.Println(c)
		if c > 5 {
			cancel()
		}
	}

}
