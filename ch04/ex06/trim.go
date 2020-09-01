package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello   World!")
	fmt.Println(string(trim(s)))

}

func trim(s []byte) []byte {

	var position int

	var r rune   //検査を行うrune
	var size int //検査を行うruneのバイト数

	//UTF-8は 1byte〜4bytesの可変長データでcode pointを置換するので、lenおよびcapは4にする
	spaceSlice := make([]byte, 4)
	space := ' '
	//' 'をUTF-8へエンコードして、スライスへ書き込み
	spaceSize := utf8.EncodeRune(spaceSlice, space)
	//不要なスライスは破棄
	spaceSlice = spaceSlice[:spaceSize]

	//スペースが続いていることを検知するためのフラグ
	var spaceFlg bool

	//文字によってバイト数は異なるため、sizeを記録してずらしていく
	for i := 0; i < len(s); i += size {
		r, size = utf8.DecodeRune(s[i:]) //TODO : DecodeRune(s[i:]) i:の理由

		if unicode.IsSpace(r) {
			//スペースが連続している場合、次の要素に移る
			if spaceFlg {
				//連続している場合には、1つに集約するため positionの更新などは行わない
				continue
			}

			copy(s[position:], spaceSlice)
			spaceFlg = true

			//元々の文字から、' 'に変更したため追加した' 'のサイズを反映する
			position += spaceSize

		} else {
			copy(s[position:], s[i:i+size])
			position += size

			spaceFlg = false
		}
	}

	return s[:position]

	/*
		//rangeで、UTF-8のデコードを暗黙的に行なっている　→　それを利用
		for i, v := range string(s) {

			if unicode.IsSpace(v) {
				fmt.Println("スペースを検知")
			}

			//スペースが連続しているため、１つのスペース' 'にする
			s[i] = ' '
			copy(s[i+1:], s[i+2:])
			cnt++
		}

		return s[:len(s)-cnt]
	*/

}
