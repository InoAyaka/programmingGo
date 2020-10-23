package main

import (
	"Book_gopl/exercise/ch05/ex13/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var startDomain string

func main() {
	startURL := "http://golang.org"

	u, err := url.Parse(startURL)
	if err != nil {
		log.Printf("url parse error : %s\n", startURL)
	}
	startDomain = u.Hostname()

	breadthFirst(crawl, []string{startURL})
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if isSameDomain(item) {
					download(item)
					worklist = append(items, f(item)...)
				}
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

func isSameDomain(checkURL string) bool {
	u, err := url.Parse(checkURL)
	if err != nil {
		log.Printf("url parse error : %s\n", checkURL)
		return false
	}

	return strings.HasSuffix(startDomain, u.Hostname())
}

func download(u string) {
	resp, err := http.Get(u)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	/*
		example	"https://golang.org/doc/copyright.html"
		base :	"copyright.html"
		dir  :	"/doc"
	*/
	base := path.Base(resp.Request.URL.Path)
	dir := path.Dir(resp.Request.URL.Path)

	if base == "/" || strings.HasSuffix(dir, base) {
		base = "index.html"
	}

	fileDir := filepath.Join("result", startDomain, dir)
	filePath := filepath.Join(fileDir, base)

	if err := os.MkdirAll(fileDir, 0775); err != nil {
		log.Println(err)
		return
	}

	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Println(err)
	}

}
