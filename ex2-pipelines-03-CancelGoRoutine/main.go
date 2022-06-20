package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Objective: terminate a goroutine using a channel with an empty struct{}
func storeNumInChannel(done <-chan struct{}, in ...int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for _, v := range in {
			select {
			case output <- v:
			case <-done:
				return
			}
		}
	}()
	return output
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for v := range in {
			select {
			case output <- v * v:
			case <-done:
				return
			}
		}
	}()
	return output
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup = sync.WaitGroup{}

	output := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out

}

func main() {

	fmt.Println("**Fan-In, Fan-Out Excercise**")
	fmt.Println("And adding a done channel to terminate the goroutines early.")

	var done = make(chan struct{})

	ch1 := square(done, storeNumInChannel(done, 2, 3))
	ch2 := square(done, storeNumInChannel(done, 2, 3))

	out := merge(done, ch1, ch2)

	//a. after receiving one value
	//b. we close the done channel to terminate the pipeline
	// Note that it is an idiom to pass the first param as a done channel
	fmt.Println(<-out)
	close(done) //terminate the pipeline after printing once

	// use a timer to allow time for the goroutines to terminate
	time.Sleep(10 * time.Millisecond)

	// The checks the number of active goroutines
	// and tell if the above close on done channel has successfully terminate the goroutines
	g := runtime.NumGoroutine()
	fmt.Println("Number of goroutines active: ", g)

}
