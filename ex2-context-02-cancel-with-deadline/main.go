package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// set a compute function with a goroutine within to simulate the work - and checks if the deadline is met
	compute := func(cx context.Context, in *data) <-chan data {

		// create an out channel to return the output derived from the processs
		out := make(chan data)

		go func() {
			defer close(out)

			// using the passed in context, check if a deadline has been set
			deadline, ok := cx.Deadline()

			// if deadline has been set, ok == true
			if ok {
				// check if the given deadline has past the given time
				// Subtract the deadline to the time needed to compute: 3 seconds - 2 seconds
				if deadline.Sub(time.Now().Add(2*time.Second)) < 0 {

					// if the time needed is beyond the stipulated duration,
					// end the goroutined without sending Done() to the returning channel
					fmt.Println("Not enough time to run... terminating")
					return
				}
			}

			// simulate the work done that takes 2 second
			time.Sleep(2 * time.Second)

			// work to do is send the 'in' value to the sending channel 'out'
			select {
			case <-cx.Done():
				return
			case out <- *in:
			}
		}()

		// return the out channel
		return out
	}

	//create a variable to provide e deadline of 3 seconds for the program to complete
	deadline := time.Now().Add(3 * time.Second)

	// set the deadline with a new context
	cxt, cancel := context.WithDeadline(context.Background(), deadline)

	// defer the cancel (cxt.Done())
	defer cancel()

	// create the value to be passed in for 'compute'
	in := data{"123"}

	// pass in the context and value to compute
	ch, ok := <-compute(cxt, &in)

	// if cxt.Done() is received from the compute, print 'work complete'
	if ok {
		fmt.Println("work completed", ch)
	}
}
