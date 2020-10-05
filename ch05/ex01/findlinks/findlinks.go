package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	data, err := ioutil.ReadFile("/Users/ayaka/go/src/Book_gopl/exercise/ch05/ex01/result.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink : %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	/*
		for c := n.FirstChild; c != nil; c = n.NextSibling {
			links = visit(links, c)
		}
	*/

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
