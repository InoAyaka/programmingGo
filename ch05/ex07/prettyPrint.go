package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("http://www.gopl.io/")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "getting error : %s", err)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing error : %s", err)
	}

	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if n == nil {
		return
	}

	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	/*
		forEachNode(n.FirstChild, pre, post)
		forEachNode(n.NextSibling, pre, post)
	*/

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {

	switch n.Type {
	case html.ElementNode:
		if n.FirstChild == nil {
			return
		}
		fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, getAllAtts(n))
		depth++
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s\n", depth*2, "", n.Data)
	case html.TextNode:
		depth++
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {

	switch n.Type {
	case html.ElementNode:
		depth--
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s />\n", depth*2, "", n.Data, getAllAtts(n))
			return
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	case html.CommentNode:
		fmt.Printf("-->\n")
	case html.TextNode:
		depth--

	}
}

func getAllAtts(n *html.Node) string {
	var atts string
	for _, a := range n.Attr {
		atts += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
	}

	return atts
}
