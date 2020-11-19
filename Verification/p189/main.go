package main

import (
	"fmt"
)

//p.189 intset

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func main() {
	var s IntSet

	s.Add(104)
	s.Add(21)
	s.Add(3)

	fmt.Println(s.words)
	fmt.Println(s.Has(21))

}
