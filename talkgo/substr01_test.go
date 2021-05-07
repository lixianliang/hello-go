package algo02

import (
	"testing"
)

type param struct {
	s string
	p string
}

type ans struct {
	one int
}

type question struct {
	param
	ans
}

func Test_substr(t *testing.T) {
	qs := []question{
		{
			param{"", "abc"},
			ans{-1},
		},
		{
			param{"abc", ""},
			ans{0},
		},
		{
			param{"aboiabcdhjj", "abc"},
			ans{4},
		},
		{
			param{"abacbabc", "abc"},
			ans{5},
		},
	}

	for _, q := range qs {
		ans, param := q.ans, q.param
		x := bfSubStr(param.s, param.p)
		if x != ans.one {
			t.Errorf("bmSubStr(%s, %s) = %d not equal %d", param.s, param.p, x, ans.one)
		}
	}
}
