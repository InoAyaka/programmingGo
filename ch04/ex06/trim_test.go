package main

import "testing"

func TestTrim(t *testing.T) {

	type test struct {
		input []byte
		want  []byte
	}

	tests := []test{
		{
			input: []byte(""),
			want:  []byte(""),
		}, {
			input: []byte("Hello World"),
			want:  []byte("Hello World"),
		}, {
			input: []byte("Hello   World!"),
			want:  []byte("Hello World!"),
		}, {
			input: []byte("Hello  　 世界!"),
			want:  []byte("Hello 世界!"),
		}, {
			input: []byte("  Hello World!   "),
			want:  []byte(" Hello World! "),
		},
	}

	for _, test := range tests {
		r := trim(test.input)
		if len(r) != len(test.want) {
			t.Errorf("trim(%v) want:%v result:%v", string(test.input), string(test.want), string(r))
			continue
		}

		for i, v := range r {
			if v != test.want[i] {
				t.Errorf("trim(%v) want:%v result:%v", string(test.input), string(test.want), string(r))
			}
		}

	}

}
