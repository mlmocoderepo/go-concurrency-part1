package main

import (
	"fmt"
	"time"
)

// function created to test out go routines
func printstring3x(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(1 * time.Millisecond)
	}
}

// the example here does not invoke waitgroups
func main() {
	// 1. direct function call
	printstring3x("function call")

	// 2. go routine function call
	go printstring3x("go routine 1")

	// 3. go routine anonymous function
	go func() {
		printstring3x("go routine 2")
	}()

	// 4. go routine with a function call
	go func(s string) {
		printstring3x(s)
	}("go routine 3")

	// 5. This validates that go routines do not run in order
	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println("test", v)
		}(i)
	}

	// wait for the go routines to end using a timer
	fmt.Println("Waiting for go routine to end.")
	time.Sleep(1 * time.Second)

	fmt.Println("Done")
}
