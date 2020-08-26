package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(diffCount(c1, c2))

}

func diffCount(c1, c2 [32]byte) int {
	var cnt int

	xorBytes := [32]byte{}

	for i, v := range c1 {
		xorBytes[i] = v ^ c2[i]
	}

	for _, b := range xorBytes {
		cnt += int(pc[b])
	}

	return cnt
}
