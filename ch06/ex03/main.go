package main

import "fmt"

func main() {
	var s IntSet

	s.Add(104)
	s.Add(21)
	s.Add(3)
	s.AddAll()
	s.AddAll(3, 2, 34)

	fmt.Println(&s)
	fmt.Println(s.Has(21))
	fmt.Println(s.Len())

	s.Remove(21)
	fmt.Println(&s)
	fmt.Println(s.Has(21))

	c := s.Copy()
	fmt.Println(c)

	s.Clear()
	fmt.Println(&s)
	fmt.Println(c)

}
