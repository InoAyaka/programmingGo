package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages("https://golang.org/pkg/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-7s / %-7s\n", "words", "images")
	fmt.Printf("%-7d / %-7d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("getting %s : %s", url, resp.Status)
		return
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		err = fmt.Errorf("parsing HTML : %w", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			words++
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	fw, fi := countWordsAndImages(n.FirstChild)
	words += fw
	images += fi

	nw, ni := countWordsAndImages(n.NextSibling)
	words += nw
	images += ni

	return

}
