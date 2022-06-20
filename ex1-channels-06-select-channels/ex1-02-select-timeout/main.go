package main

import (
	"fmt"
	"time"
)

// Objective: to implement a timeout to handle goroutines

func main() {
	ch1 := make(chan string)

	go func(c chan<- string) {
		time.Sleep(2 * time.Second)
		c <- "one"
	}(ch1)

	// the second select-case statement handles any timeouts
	// when an expected goroutine ch1 is not receiving its value within the expected time (time.After(1 * time.Second))
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("Time out")
	}

}
