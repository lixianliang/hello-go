package algo02

import (
	"testing"
)

type param2 struct {
	s string
}

type ans2 struct {
	one int
	two string
}

type question2 struct {
	param2
	ans2
}

func Test_pastr(t *testing.T) {
	qs := []question2{
		{
			param2{""},
			ans2{0, ""},
		},
		{
			param2{"abc"},
			ans2{1, "c"},
		},
		{
			param2{"aboiabcdhjj"},
			ans2{3, "bab"},
		},
		{
			param2{"apjljlopasdjpabfcbeaiopipdfasd"},
			ans2{11, "sdpabcbapds"},
		},
	}

	for _, q := range qs {
		ans, param := q.ans2, q.param2
		num, pastr := Palindrom(param.s)
		if num != ans.one || pastr != ans.two {
			t.Errorf("Palindrom(%s) = %d %s not equal one:%d two:%s", param.s, num, pastr, ans.one, ans.two)
		}
	}
}
