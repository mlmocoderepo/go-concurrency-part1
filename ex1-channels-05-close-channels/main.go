package main

import (
	"fmt"
)

// Objective: create channel owner goroutine which returns a channel and
// writes passed-in data to the channel and closes the channel when done

func main() {

	owner := func(count int) <-chan int {

		ch := make(chan int) //1. create the channel that is to be returned

		go func(val int, c chan<- int) { //2. the goroutine takes in the number of counts to itereate and the sending channel

			defer close(c)             //3. defer the closure of the channel
			for i := 0; i < val; i++ { //3.1 iterate through val as the incrementer
				c <- i //3.2 send each value of the iteration to the sending channel
			}

		}(count, ch)

		return ch
	}

	consumer := func(c <-chan int) {
		for v := range c {
			fmt.Println("Receiver value has:", v)
		}
	}

	fmt.Println("Done receiving...")

	o := owner(10)
	consumer(o)
}
