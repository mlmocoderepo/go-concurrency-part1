package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type result struct {
	urls  []string
	url   string
	depth int
	err   error
}

var fetched map[string]bool

func main() {

	fetched = make(map[string]bool)

	t := time.Now()
	crawlUrl("http://andcloud.io", 2)
	fmt.Println("Time took to crawl the site is:", time.Since(t))
}

func crawlUrl(url string, depth int) {

	// 1. make a channel that receives a pointer to result object(s)
	results := make(chan *result)
	count := 0

	// 2. create a go routine that obtains the urls from the given url
	// and store the returned values (urls, url, depth, err) into the pointer of a result object
	// and sent into channel results
	fetch := func(url string, depth int) {
		urls, err := fetchlinks(url)
		results <- &result{urls, url, depth, err}
	}

	// 3. for the first time, crawl the url received
	// and set the fetched url's key/pair value to true
	// print the 1st url that has been crawled
	go fetch(url, depth)
	fetched[url] = true
	fmt.Printf("Found %v\n", url)
	count++

	// 4. run a loop against the depths left
	for fetching := 1; fetching > 0; fetching-- {

		// get the result from the channel
		res := <-results

		// if there is an error, iterate to the next loop
		if res.err != nil {
			continue
		}

		// if the depth of result object is greater than 2
		// if the result's url key/value is false
		// fetch the urls from the url and set the key/value pair to true
		if res.depth > 0 {
			for _, u := range res.urls {
				if !fetched[u] {
					go fetch(u, res.depth-1)
					fetched[u] = true
					fmt.Printf("Found: %v\n", u)
					fetching++
					count++
				}
			}
		}

	}
	fmt.Println(count)
	close(results)
}

func fetchlinks(url string) ([]string, error) {

	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve url %v due to %v\n", url, err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("Response error frm url %v due to %v\n", url, err)
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		return nil, err
	}

	return visit(nil, doc), nil
}

func visit(links []string, node *html.Node) []string {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

// recursive example to crawl
/*
	// 1. if the depth is reduced to zero, stop crawling and end the function
	if depth < 0 {
		return
	}

	// 2. fetch the links from the received url
	urls, err := fetchlinks(url)

	// 3. if there's an error, stop crawling and end the function
	if err != nil {
		return
	}

	// 4. set fetched[url] = true so that the url isn't crawled again
	fetched[url] = true
	fmt.Printf("Found url %v\n", url)

	// 5. recursively check if the urls is found in the fetched[url] == false
	// if it is, continue crawling
	for _, u := range urls {
		if !fetched[u] {
			crawlUrl(u, depth-1)
			fetched[url] = true
		}
	}
*/
