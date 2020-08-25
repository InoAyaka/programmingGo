package main

import (
	"fmt"
)

func main() {
	const (
		//_ = 1 << (iota * 3)
		MB = 1000
		KB = MB * MB
		GB = KB * MB
		TB = GB * MB
		PB = TB * MB
		EB = PB * MB
		ZB = EB * MB
		YB = ZB * MB
	)

	fmt.Println(KB)
	fmt.Println(GB)
	//fmt.Println(YB)
}
