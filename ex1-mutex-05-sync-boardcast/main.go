package main

import (
	"fmt"
	"sync"
)

// Objective: to use sync.NewCond to suspend all requests to a shared resource until a condition is met
// once the objective is met, release the suspension of all requests waiting on the resource

func main() {

	var sharedRsc = map[string]string{}
	var wg sync.WaitGroup = sync.WaitGroup{}
	var m = sync.Mutex{}
	var c = sync.NewCond(&m)

	// 1. Suspend the goroutine operation until sharedRsc is >= 1
	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) < 1 {
			c.Wait()
		}

		fmt.Printf("resource found: %v\n\n", sharedRsc["rsc1"])
		c.L.Unlock()

	}()

	// 2. Suspend the goroutine operation until sharedRsc is >= 2
	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) < 2 {
			c.Wait()
		}

		fmt.Printf("resource found: %v\n\n", sharedRsc["rsc2"])
		c.L.Unlock()

	}()

	// 2. Populate sharedRsc and resume the gorouting operation indicated above
	c.L.Lock()
	sharedRsc["rsc1"] = "pushed 1"
	sharedRsc["rsc2"] = "pushed 2"
	c.Broadcast() //Use BROADCAST to releases the wait time for ALL requesting goroutines to resume
	c.L.Unlock()

	fmt.Printf("resource found: %v\n\n", sharedRsc)
	wg.Wait()
}
