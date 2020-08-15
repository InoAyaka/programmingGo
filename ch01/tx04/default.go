//p.27 defaultの判定位置は本当に最後かデバッグで見てみる
package main

import "fmt"

func main() {

	x := 0

	switch {
	case x > 0:
		fmt.Println("case1")
	default:
		fmt.Println("defalut")
	case x < 0:
		fmt.Println("case2")
	}

}
