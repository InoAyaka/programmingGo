package main

import (
	"fmt"
)

//p.82 コピー発生する　→　元の値は変更されないのか検証

func main() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Printf("%[1]s %[1]p\n", s)
	fmt.Printf("%[1]s %[1]p\n", b)
	fmt.Printf("%[1]s %[1]p\n", s2)

	s = "cd"

	fmt.Printf("%[1]s %[1]p\n", s)
	fmt.Printf("%[1]s %[1]p\n", b)
	fmt.Printf("%[1]s %[1]p\n", s2)

	b[2] = 'd'
	b[1] = byte("d")

	fmt.Printf("%[1]s %[1]p\n", s)
	fmt.Printf("%[1]s %[1]p\n", b)
	fmt.Printf("%[1]s %[1]p\n", s2)
}
