package main

import "testing"

func TestExpend(t *testing.T) {
	f := func(string) string { return "fo fo fo" }
	h := func(string) string { return "hohoho" }

	type test struct {
		input string
		f     func(string) string
		want  string
	}

	tests := []test{
		{"foo $foo", f, "foo fo fo fo"},
		{"foo$foo", f, "foofo fo fo"},
		{"$foo,$foo", f, "fo fo fo,fo fo fo"},
		{"aa $aaa", h, "aa hohoho"},
		{"aa$aaa", h, "aahohoho"},
		{"$aa,$", h, "hohoho,hohoho"},
	}

	for _, test := range tests {

		result := expend(test.input, test.f)

		if test.want != result {
			t.Errorf("%s : got %s , want %s", test.input, result, test.want)
		}
	}
}
