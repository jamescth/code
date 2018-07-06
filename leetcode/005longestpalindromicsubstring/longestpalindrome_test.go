package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		n   int
		x   string
		ret string
	}{
		{1, "babad", "bab"},
		{2, "cbbd", "bb"},
		{3, "babab", "babab"},
	}

	for _, tc := range tests {
		if ret := longestPalindrome(tc.x); ret != tc.ret {
			t.Errorf("%d failed %s expect %s got %s\n", tc.n, tc.x, tc.ret, ret)
		}
	}
}
