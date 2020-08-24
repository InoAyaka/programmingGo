package main

import (
	"fmt"
)

func isAnagram(s1, s2 string) bool {

	if s1 == s2 {
		return false
	}

	//文字の出現数をカウント
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, v := range s1 {
		m1[v]++
	}

	for _, v := range s2 {
		m2[v]++
	}

	//出現数が一致しているかで判定
	//そもそも文字のバリエーションが一致しない場合は、false
	if len(m1) != len(m2) {
		return false
	}

	//バリエーションも一致した上で、出現数が一致しているかをチェック
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}

func main() {

	s1, s2 := "aiu", "uia"
	fmt.Printf("s1 = %s , s2 = %s : %t\n", s1, s2, isAnagram(s1, s2))

	s1, s2 = "uiia", "aiu"
	fmt.Printf("s1 = %s , s2 = %s : %t\n", s1, s2, isAnagram(s1, s2))

	s1, s2 = "aiu", "aiu"
	fmt.Printf("s1 = %s , s2 = %s : %t\n", s1, s2, isAnagram(s1, s2))

	s1, s2 = "", "aiu"
	fmt.Printf("s1 = %s , s2 = %s : %t\n", s1, s2, isAnagram(s1, s2))

	s1, s2 = "aiu", ""
	fmt.Printf("s1 = %s , s2 = %s : %t\n", s1, s2, isAnagram(s1, s2))

}
