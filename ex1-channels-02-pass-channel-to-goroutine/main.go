package main

import "fmt"

// Objective: send a channel to an anonymous goroutine to send value(s) into it
// Later, receive the values from the channel to be printed
// TAKE NOTE: to close a sender channel after it it no longer used

func main() {

	// 1. Instantiate a channel variable
	ch := make(chan int)

	// 2. Create a sender go routine
	go func(c chan<- int) {

		defer close(c) //2.2 REMEMBER: gracefully close the channel, otherwise the receiver does not know when it is exited

		for i := 0; i < 3; i++ {
			c <- i
		}

	}(ch) //2.1. Send the channel to the goroutine for it to operate within

	// 3. retrive the values stored in the channel
	for c := range ch {
		fmt.Println(c)
	}
}
