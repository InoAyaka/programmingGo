//dup04の修正
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cntFiles struct {
	cnt   int
	files []string
}

func main() {
	//counts 現れた行の数の集まり
	counts := make(map[string]*cntFiles)

	//files 引数で渡したファイル群
	files := os.Args[1:]

	//引数なしの場合
	if len(files) == 0 {
		//標準入力から読み取った結果を、カウント
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			//ファイル読取失敗時
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				//次のファイルへ移行
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	//結果の出力
	for line, n := range counts {
		//重複したもののみ表示
		if n.cnt > 1 {
			fmt.Printf("%s : %d　（出現ファイル名: %v）\n", line, n.cnt, strings.Join(n.files, " "))
		}
	}
}

//countLines 現れた行のカウントを行う
func countLines(f *os.File, counts map[string]*cntFiles) {
	input := bufio.NewScanner(f)

LOOP:
	for input.Scan() {
		line := input.Text()

		cf, ok := counts[line]
		if ok {
			//既存行の場合
			cf.cnt++
			//ファイル名の追加要否判定
			for _, arg := range cf.files {
				if arg == f.Name() {
					//既にそのファイル名がある場合には、次の行へ移行
					continue LOOP
				}
			}
			//なかった場合には、追加
			cf.files = append(cf.files, f.Name())

		} else {
			//新規行の場合
			counts[line] = &cntFiles{cnt: 1, files: []string{f.Name()}}

		}

	}
}
