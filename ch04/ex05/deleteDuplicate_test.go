package main

import "testing"

func TestDeleteDuplicate(t *testing.T) {
	var data = []struct {
		s    []string
		want []string
	}{
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"aaa"},
			[]string{"aaa"},
		},
		{
			[]string{"aaa", "aaa"},
			[]string{"aaa"},
		},
		{
			[]string{"aaa", "aaa", "bbb"},
			[]string{"aaa", "bbb"},
		},
		{
			[]string{"aaa", "aaa", "bbb", "ccc", "ccc", "bbb", "bb"},
			[]string{"aaa", "bbb", "ccc", "bbb", "bb"},
		},
	}

	for _, d := range data {
		result := deleteDuplicate(d.s)

		if len(result) != len(d.want) {
			t.Errorf("deleteDuplicate(%s) , want %q , resutl %q", d.s, d.want, result)
		}

		for i := 0; i < len(result)-1; i++ {
			if result[i] != d.want[i] {
				t.Errorf("deleteDuplicate(%s) , want %q , resutl %q", d.s, d.want, result)
			}
		}
	}
}
