package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file := "/Users/ayaka/go/src/Book_gopl/exercise/ch04/ex09/input.txt"

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	words := wordfreq(f)

	fmt.Printf("word\tcount\n")
	for w, c := range words {
		fmt.Printf("%s\t%d\n", w, c)
	}

	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}

func wordfreq(f *os.File) map[string]int {

	words := make(map[string]int)

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		words[input.Text()]++
	}

	return words
}
