package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool

type result struct {
	url   string
	urls  []string
	err   error
	depth int
}

func main() {
	fetched = make(map[string]bool)
	now := time.Now()

	CrawlUrl("http://andcloud.io", 2)
	fmt.Printf("Crawling duration: %s\n", time.Since(now))
}

func CrawlUrl(url string, depth int) {

	results := make(chan *result)

	fetch := func(url string, depth int) {
		urls, err := fetchlinks(url)
		result := result{url, urls, err, depth}
		results <- &result
	}

	go fetch(url, depth)
	fetched[url] = true

	defer close(results)

	for fetching := 1; fetching > 0; fetching-- {

		res := <-results
		if res.err != nil {
			continue
		}

		fmt.Printf("Found: %s\n", res.url)

		if res.depth > 0 {
			for _, u := range res.urls {
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth-1)
					fetched[u] = true
				}
			}
		}

	}

}

func fetchlinks(url string) ([]string, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unable to fetch :%s due to %v", url, err)
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse :%s due to %v", url, err)
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

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = visit(links, child)
	}

	return links
}
