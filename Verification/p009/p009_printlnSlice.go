package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//strings.Joinを使った方法
	fmt.Println(strings.Join(os.Args[1:], " "))

	//スライスをそのまま渡した方法
	fmt.Println(os.Args[1:])
}
