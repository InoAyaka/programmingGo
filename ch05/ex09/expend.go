package main

import (
	"fmt"
	"regexp"
)

func main() {
	f := func(string) string {
		return "aaa"
	}

	fmt.Println(expend("test$foo$$footest$hoge", f))
}

func expend(s string, f func(string) string) string {
	r := regexp.MustCompile(`\$\w*`)

	return r.ReplaceAllString(s, f(s))

}
