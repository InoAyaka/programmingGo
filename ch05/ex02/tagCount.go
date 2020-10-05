package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type tagCounts map[string]int
type item struct {
	name string
	cnt  int
}

func main() {
	f, err := ioutil.ReadFile("/Users/ayaka/go/src/Book_gopl/exercise/ch05/ex02/result.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(f)))
	if err != nil {
		log.Fatal(err)
	}

	m := visit(make(tagCounts), doc)

	for key, val := range m {
		fmt.Printf("%-10s %v\n", key, val)
	}

}

func visit(m tagCounts, n *html.Node) tagCounts {

	if n == nil {
		return m
	}

	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	m = visit(m, n.FirstChild)
	m = visit(m, n.NextSibling)

	return m

}
