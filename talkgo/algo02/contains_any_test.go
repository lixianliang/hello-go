package algo02

import (
	"testing"
)

type param struct {
	s string
	p string
}

type ans struct {
	one bool
}

type question struct {
	param
	ans
}

func Test_ContainsAny(t *testing.T) {
	qs := []question{
		{
			param{"", "abc"},
			ans{false},
		},
		{
			param{"abc", ""},
			ans{false},
		},
		{
			param{"aboiabcdhjj", "abc"},
			ans{true},
		},
		{
			param{"abacbabc", "abdddeeffc"},
			ans{false},
		},
	}

	for _, q := range qs {
		ans, param := q.ans, q.param
		x := ContainsAny(param.s, param.p)
		if x != ans.one {
			t.Errorf("ContainsAny(%s, %s) = %t not equal %t", param.s, param.p, x, ans.one)
		}
	}
}
