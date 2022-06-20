package main

import (
	"fmt"
	"sync"
)

// Objective: to use sync.NewCond to suspend a resource until a condition is met
// once the condition is met, release the suspension for the waiting resource to resume

func main() {

	var sharedRsc = map[string]string{}
	var wg sync.WaitGroup = sync.WaitGroup{}
	var m = sync.Mutex{}
	var c = sync.NewCond(&m)

	// 1. Suspend the goroutine operation until sharedRsc is not empty
	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Printf("resource found: %v", sharedRsc)
		c.L.Unlock()

	}()

	// 2. Populate sharedRsc and resume the gorouting operation indicated above
	c.L.Lock()
	sharedRsc["rsc1"] = "pushed 1"
	c.Signal() //Use SIGNAL to releases the wait time for requesting goroutine to resume
	c.L.Unlock()

	wg.Wait()
}
