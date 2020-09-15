package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	utflen := [utf8.UTFMax + 1]int{}
	category := make(map[string]int)

	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {

		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		category[getCategory(r)]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\ncategory\tcount\n")
	for cat, n := range category {
		fmt.Printf("%q\t%d\n", cat, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}

func getCategory(r rune) string {

	if unicode.IsLetter(r) {
		return "L"
	}

	if unicode.IsMark(r) {
		return "M"
	}

	if unicode.IsNumber(r) {
		return "N"
	}

	if unicode.IsControl(r) {
		return "C"
	}

	if unicode.IsSpace(r) {
		return "Z"
	}

	if unicode.IsPunct(r) {
		return "P"
	}

	return ""

}
