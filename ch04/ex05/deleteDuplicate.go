package main

import (
	"fmt"
)

func main() {
	s := []string{"one", "one", "two", "three", "three", "two"}
	s = deleteDuplicate(s)
	fmt.Println(s)
}

func deleteDuplicate(s []string) []string {
	var cnt int

	for i, v := range s {
		if i < len(s)-1-cnt && v == s[i+1] {
			copy(s[i:], s[i+1:])
			cnt++
		}
		fmt.Println(s, cnt)
	}

	return s[:len(s)-cnt]

}
