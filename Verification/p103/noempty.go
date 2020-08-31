package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}

	//	fmt.Printf("%q\n", noEmpty(data))
	data = noEmpty(data)
	fmt.Printf("%q\n", data)
}

func noEmpty(strings []string) []string {
	var i int

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}
