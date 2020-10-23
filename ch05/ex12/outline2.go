package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	res, err := http.Get("http://gopl.io")
	if err != nil {
		fmt.Fprintf(os.Stderr, "get error : %v\n", err)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "%v\n", res.Status)
	}

	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing error : %v\n", err)
	}

	forEachNode(doc, startElement, endElement)
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
