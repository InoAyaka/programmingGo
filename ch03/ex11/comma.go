package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {

	if s == "" {
		return s
	}

	//符号記号を取得
	var sign string

	if s[:1] == "-" || s[:1] == "+" {
		sign = s[:1]
		s = s[1:]
	}

	//小数点以下の取得
	var decimal string
	dot := strings.LastIndex(s, ".")

	if dot >= 0 {
		decimal = s[dot:]
		s = s[:dot]
	}

	//3桁未満（小数点を除く）の場合は、そのまま返す
	if len(s) <= 3 {
		return sign + s + decimal
	}

	//3桁以上（小数点を除く）の場合
	var buf bytes.Buffer

	buf.WriteString(sign)

	//,区切りしない頭を取得する
	lead := len(s) % 3

	if lead == 0 {
		lead = 3
	}

	//頭までをバッファに追加
	buf.WriteString(s[:lead])

	//頭以降を3桁単位で繰り返す
	for i := lead; i < len(s); i = i + 3 {
		buf.WriteString("," + s[i:i+3])
	}

	//最後に小数点以下をバッファに追加
	buf.WriteString(decimal)

	return buf.String()
}

func main() {
	fmt.Println(comma(""))
	fmt.Println(comma("1"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12.79008"))
	fmt.Println(comma("123456.79008"))
	fmt.Println(comma("-123456.70786"))
	fmt.Println(comma("+12345.70786"))
	fmt.Println(comma("+0.1234567"))
}
