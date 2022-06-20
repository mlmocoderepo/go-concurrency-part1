package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// a) fetched is used to check whether the status of the url 'key' has a bool value 'true'
var fetched map[string]bool

func main() {

	fetched = make(map[string]bool)
	now := time.Now()

	CrawlURL("http://andcloud.io", 2)
	fmt.Printf("Time taken to crawl %s\n", time.Since(now))

}

// b) CrawlURL is used to RECURSIVELY crawl through the url / page passed to it
// The maximum level to crawl is dependent on the depth parameter
func CrawlURL(url string, depth int) {

	if depth < 0 {
		return
	}

	urls, err := findLinks(url)
	if err != nil {
		return
	}

	fmt.Printf("Found %s\n", url)
	fetched[url] = true

	for _, u := range urls {
		if fetched[u] != true {
			CrawlURL(u, depth-1)
		}
	}

	return
}

// c) findlinks() gets the url, checks the status code before parsing the response's body
// it then paases the parsed data to the function visit() to iterate through the html node elements
func findLinks(url string) ([]string, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("Unable to read from URL: %s, error: %s\n", url, err)
	}

	doc, err := html.Parse(response.Body)
	response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Unable to parse from URL: %s, error: %s\n", url, err)
	}

	return visit(nil, doc), nil

}

// d) visit() checks the parsed data it receives on whether the node.Type is a html.Elementnode
// with 'a' nodes' as data and loops thorugh the attributes to seive out the href attribute
// and appends the href's value to the the links parameter
// Then, it checks if the the node has siblings before recursively obtaining additional href values from the node's NextSibiling
func visit(links []string, node *html.Node) []string {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// IMPORTANT: The iterative process requires the CHILD to pass the NextSibiling() ==> child=child.NextSibling()
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = visit(links, child)
	}

	return links
}
