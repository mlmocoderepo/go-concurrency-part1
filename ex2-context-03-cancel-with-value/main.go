package main

import (
	"context"
	"fmt"
)

// create a userIDKey type string
type userIDKey string
type database map[string]bool

var db database = database{
	"jane": true,
}

func main() {

	// create a context
	cxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// send the context and the name of the member for processing
	processStatus(cxt, "jane")
}

func processStatus(cx context.Context, userID string) {

	// create the context with the value to be checked.
	cxtWithValue := context.WithValue(cx, userIDKey("userIDKey"), userID)

	// pass the channel with value to the function that checks the membership status
	status := <-checkMembership(cxtWithValue)

	fmt.Printf("Membership status for %s is %v \n", userID, status)

}

func checkMembership(cx context.Context) <-chan bool {

	// create a channel to return the result
	ch := make(chan bool)

	// call the goroutine to return the value
	go func() {
		defer close(ch)

		// get the Value from the context passed in
		member := cx.Value(userIDKey("userIDKey")).(string)
		ch <- db[member]

	}()
	return ch
}
