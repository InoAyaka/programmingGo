package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Printf("%d\n", a)

}

func reverse(s *[4]int) {
	l := len(s) - 1

	for i := range s {
		s[i], s[l-i] = s[l-i], s[i]
		if i >= l/2 {
			break
		}
	}
}
