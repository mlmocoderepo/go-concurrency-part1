package main

import (
	"fmt"
	"sync"
)

// Objective: to create the instance that while a go routine is invoked several times
// sync.once.do only allows the passed-in function to run once, e.g. for initlialization purposes

func main() {

	var wg sync.WaitGroup = sync.WaitGroup{}

	// 1. Create an instance of sync.Once
	var once sync.Once

	load := func() {
		fmt.Println("This function is only initialized once")
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 2. Call the load function only once using sync.Once.Do's function
			once.Do(load)
		}()
	}

	wg.Wait()
}
