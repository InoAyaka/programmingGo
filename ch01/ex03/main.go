//非効率な可能性のあるバージョンと、strings.Join()を用いた方法とで、実行時間の差を計測
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/* MEMO：実行結果
テスト文字：Sharp-eyed readers may notice that our year-over-year comparisons don’t quit
結果；
	0.000038s
	0.000003s
*/

func main() {
	//非効率ver.
	inefficiency()
	//strings.Join()　ver.
	strJoin()
}

//非効率ver.
func inefficiency() {
	stT := time.Now()

	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	//実行時間の表示
	fmt.Printf("%2fs \n", time.Since(stT).Seconds())
}

//strings.Join() ver.
func strJoin() {
	stT := time.Now()

	fmt.Println(strings.Join(os.Args, " "))

	//実行時間の表示
	fmt.Printf("%2fs \n", time.Since(stT).Seconds())

}
