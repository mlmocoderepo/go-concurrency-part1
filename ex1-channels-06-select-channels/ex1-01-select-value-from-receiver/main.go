package main

import (
	"fmt"
	"time"
)

// Objective: to use a select to handle goroutines

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(c chan<- string) {
		time.Sleep(1 * time.Second)
		c <- "one"
	}(ch1)

	go func(c chan<- string) {
		time.Sleep(2 * time.Second)
		c <- "two"
	}(ch2)

	// using a multiple on the receivers of channel 1 and channel 2
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
