package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// create a GET request to "https://andcloud.io" with a timeout out 100 milliseconds
	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

	// create a context with timeout of 100 milliseconds
	// pass in the requesting context and duration (100ms)
	cxt, cancel := context.WithTimeout(req.Context(), 3000*time.Millisecond)
	defer cancel()

	// bind the request to the context
	req = req.WithContext(cxt)

	// execute the request initialized above
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Time out!!")
		return
	}

	// defer closing the response
	defer resp.Body.Close()

	// send the response to stdout
	io.Copy(os.Stdout, resp.Body)
}
