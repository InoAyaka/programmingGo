//echoプログラムを修正して、個々の引数のインデックスと値の組みを１行ごとに表示
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[:] {
		fmt.Println(i, v)
	}

}
