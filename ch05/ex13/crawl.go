package main

import (
	"Book_gopl/exercise/ch05/ex13/links"
	"fmt"
	"log"
)

func main() {
	breadthFirst(crawl, []string{"http://golang.org"})
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(items, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}

	return list
}
