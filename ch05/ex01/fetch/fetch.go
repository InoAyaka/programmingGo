package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s : %v\n", url, err)
			os.Exit(1)
		}

		defer resp.Body.Close()

		fmt.Printf("%s", b)
	}
}
