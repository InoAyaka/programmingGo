//echoプログラムを修正して、起動コマンド名も表示する
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
