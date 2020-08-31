package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 2))

	s2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(rotate(s2, 7))
}

func rotate(s []int, i int) []int {
	if i >= len(s)-1 || i <= 0 {
		return s
	}

	reverse(s[:i])
	reverse(s[i:])
	reverse(s[:])

	return s

}

func reverse(s []int) {
	l := len(s) - 1

	for i := range s {
		s[i], s[l-i] = s[l-i], s[i]
		if i >= l/2 {
			break
		}
	}
}
