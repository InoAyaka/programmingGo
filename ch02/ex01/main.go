package main

import (
	"Book_gopl/exercise/ch02/ex01/tempconv"
	"fmt"
)

func main() {
	c := tempconv.FToC(212.0)
	fmt.Println(c.String())

	f := tempconv.CToF(c)
	fmt.Println(f.String())

	k := tempconv.FToK(f)
	fmt.Println(k.String())
}
