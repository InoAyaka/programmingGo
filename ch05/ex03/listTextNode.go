package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	f, err := ioutil.ReadFile("/Users/ayaka/go/src/Book_gopl/exercise/ch05/ex03/result.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(f)))
	if err != nil {
		log.Fatal(err)
	}

	lists := make([]html.Node, 0)

	lists = visit(lists, doc)

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", "Type", "Data", "Attr", "FirstChild", "NextSibling")

	for _, list := range lists {
		fmt.Println(strings.Repeat("-", 75))
		fmt.Printf("%v\n%s\n%v\n%v\n%v\n",
			list.Type,
			list.Data,
			list.Attr,
			list.FirstChild,
			list.NextSibling,
		)
	}

}

func visit(lists []html.Node, n *html.Node) []html.Node {

	if n == nil {
		return lists
	}

	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		lists = append(lists, html.Node{
			Type:        n.Type,
			Data:        n.Data,
			Attr:        n.Attr,
			FirstChild:  n.FirstChild,
			NextSibling: n.NextSibling,
		})
	}

	lists = visit(lists, n.FirstChild)
	lists = visit(lists, n.NextSibling)

	return lists
}
