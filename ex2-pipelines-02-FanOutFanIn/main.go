package main

import (
	"fmt"
	"sync"
)

// Objective #A: of a pipeline:
// 1. Create a stage to send values to a channel called 'ch'
// 2/ Create a concurrent stage that accepts the values in 'ch' and square the data

// Objective #2: to implement FANOUT stage (multiple square routines) to merge (FANIN) a list of cs into a single channel

// takes in a slice of numbers
// returns a cs that contains the values from the slice
func storeNumInChannel(in ...int) <-chan int {
	output := make(chan int)
	go func() {
		for _, value := range in {
			output <- value
		}
		close(output)
	}()
	return output
}

// takes in a receiver channel
// sqaures the values in the channel and sends it as a returning channel
func square(in <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		for v := range in {
			output <- v * v
		}
		close(output)
	}()
	return output
}

// merge the values from a list of channels to a single channel
func merge(cs ...<-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup = sync.WaitGroup{}

	output := func(c <-chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Q1. Why must wg.Wait() be exectued by a goroutine?
	// Q2. Why can't close be adopted outside of the goroutine; and
	// Q3. why can't defer close(out) be used in the merge function?
	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out

}

func main() {

	fmt.Println("**Pipeline Exercise:**")
	for v := range square(square(storeNumInChannel(2, 3))) {
		fmt.Println(v)
	}

	fmt.Println("**Fan-In, Fan-Out Excercise**")

	// a) Fan out squre stage that run multiple goroutines to thread data from a single channel
	ch1 := square(storeNumInChannel(2, 3))
	ch2 := square(storeNumInChannel(2, 3))

	// b) and produce the output on a single channel
	for data := range merge(ch1, ch2) {
		fmt.Println(data)
	}

}
