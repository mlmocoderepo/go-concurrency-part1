package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// create a http request (GET) for https://andcloud.io
	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

	cxt, cancel := context.WithTimeout(req.Context(), 1000*time.Millisecond)

	req = req.WithContext(cxt)

	defer cancel()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Unable to handle request.")
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}
