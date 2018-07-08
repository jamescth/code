package longestpalindromic

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		num      int
		str      string
		expected string
	}{
		{1, "babad", "bab"},
		{2, "cbbd", "bb"},
		{3, "abb", "bb"},
		{4, "abam", "aba"},
		{5, "caba", "aba"},
		{6, "abcbe", "bcb"}}
	for _, test := range tests {
		ret := longestPalindrome(test.str)
		if ret != test.expected {
			t.Errorf("%d failed %s expec %s got %s", test.num, test.str, test.expected, ret)
		}

	}
}

/*
type para struct {
	one string
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "babad",
			},
			a: ans{
				one: "bab",
			},
		},
		question{
			p: para{
				one: "cbbd",
			},
			a: ans{
				one: "bb",
			},
		},
		question{
			p: para{
				one: "abbcccddcccbba",
			},
			a: ans{
				one: "abbcccddcccbba",
			},
		},
		question{
			p: para{
				one: "a",
			},
			a: ans{
				one: "a",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestPalindrome(p.one), "输入:%v", p)
	}
}
*/
