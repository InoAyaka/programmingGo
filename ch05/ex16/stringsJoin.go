package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("_", "a", "b", "foo", "hoge", "テスト"))
}

func join(sep string, strs ...string) string {

	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	//接合文字分を含めた文字列の長さを計測
	n := len(sep) * (len(strs) - 1)
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}

	var b strings.Builder
	//文字列の長さ分にcapを拡大
	b.Grow(n)

	b.WriteString(strs[0])
	for _, str := range strs[1:] {
		b.WriteString(sep)
		b.WriteString(str)
	}

	return b.String()

}
