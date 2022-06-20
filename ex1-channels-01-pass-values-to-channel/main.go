package main

import "fmt"

// Objective: use a goroutine to send values to a channel
// and later, to retriveve the values from the channels for printing

func main() {
	// 1. Instantiate a channel that accepts an integer
	ch := make(chan int)

	// 2. Create a function as a variable 'multiplier' that accepts two int values
	multiplier := func(a, b int) {
		c := a * b
		ch <- c // 2.1. Send the computed value into the channel 'ch'
	}

	// 3. Run the function variable 'multiplier' as a goroutine
	go multiplier(100, 2)

	// 4. Retrieve the value from the channel <-ch - to be printed out
	fmt.Printf("%v is the value returned from the channel.\n", <-ch)
	fmt.Println("Done...")
}
