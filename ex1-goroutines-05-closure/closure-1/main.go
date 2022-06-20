package main

import (
	"fmt"
	"sync"
)

// Objective: run the program and check the variable i
// was pinned for access from goroutine, even after
// enclosing function returns

func main() {

	var wg = sync.WaitGroup{}

	// 1. Increment is a function as a variable that accepts a pointer to sync.Waitgroup
	increment := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("Value of i: %v\n", i)
		}()
		fmt.Println("return from function")
		return
	}

	// 2. Since wg is passed as a pointer to the variable function, Wait() for the goroutine
	// to end can be controlled outside the variable function.
	increment(&wg)
	wg.Wait() // 2.1 As Wait() is here, increment function is returned then the goroutine deterministically waits and executes here.
	fmt.Println("Done...")
}
