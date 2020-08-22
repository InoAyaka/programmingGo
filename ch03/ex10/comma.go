package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {

	if len(s) <= 3 {
		return s
	}

	var buf bytes.Buffer

	lead := len(s) % 3

	if lead == 0 {
		lead = 3
	}

	buf.WriteString(s[:lead])

	for i := lead; i < len(s)-lead; i = i + 3 {
		buf.WriteString("," + s[i:i+3])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma(""))
	fmt.Println(comma("1"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234567"))
}
